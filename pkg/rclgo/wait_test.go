/*
This file is part of rclgo

Licensed under the Apache License, Version 2.0 (the "License");
    http://www.apache.org/licenses/LICENSE-2.0
*/

package rclgo_test

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	std_msgs "github.com/ATIinc/rclgo/internal/msgs/std_msgs/msg"
	"github.com/ATIinc/rclgo/pkg/rclgo"
)

// runIdleWaitSet runs an otherwise-empty wait set until ctx is done and
// returns how long Run took to return and its error.
func runIdleWaitSet(t *testing.T, waitTimeout, ctxTimeout time.Duration) (time.Duration, error) {
	t.Helper()
	rclCtx, err := newDefaultRCLContext()
	require.NoError(t, err)
	t.Cleanup(func() { rclCtx.Close() })
	ws, err := rclCtx.NewWaitSet()
	require.NoError(t, err)
	ws.WaitTimeout = waitTimeout

	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	started := time.Now()
	runErr := ws.Run(ctx)
	return time.Since(started), runErr
}

func TestWaitSetRunLoopsAcrossTimeouts(t *testing.T) {
	// A 10 ms wait timeout against a 300 ms context: Run must keep looping
	// through RCL_RET_TIMEOUT returns and still exit on cancellation.
	took, err := runIdleWaitSet(t, 10*time.Millisecond, 300*time.Millisecond)
	require.ErrorIs(t, err, context.DeadlineExceeded)
	require.Less(t, took, 2*time.Second, "Run did not return promptly after cancellation")
}

func TestWaitSetRunInfiniteTimeoutStillCancels(t *testing.T) {
	// Negative WaitTimeout restores the historical infinite rcl_wait; the
	// cancel guard condition is what wakes it.
	took, err := runIdleWaitSet(t, -1, 200*time.Millisecond)
	require.ErrorIs(t, err, context.DeadlineExceeded)
	require.Less(t, took, 2*time.Second, "Run did not return promptly after cancellation")
}

func TestWaitSetRunNoSpuriousLostWakeups(t *testing.T) {
	// An idle wait set produces timeouts only; none of them may be counted as
	// a recovered lost wakeup.
	rclCtx, err := newDefaultRCLContext()
	require.NoError(t, err)
	t.Cleanup(func() { rclCtx.Close() })
	ws, err := rclCtx.NewWaitSet()
	require.NoError(t, err)
	ws.WaitTimeout = 5 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	_ = ws.Run(ctx)
	require.Zero(t, ws.LostWakeupsRecovered())
}

// TestWaitSetIdleIterationCost measures the CPU cost of the empty timeout
// iterations the finite WaitTimeout introduces, so the default period can be
// chosen against a number rather than a guess. It only reports.
func TestWaitSetIdleIterationCost(t *testing.T) {
	rclCtx, err := newDefaultRCLContext()
	require.NoError(t, err)
	t.Cleanup(func() { rclCtx.Close() })
	ws, err := rclCtx.NewWaitSet()
	require.NoError(t, err)
	const (
		period   = time.Millisecond
		duration = time.Second
	)
	ws.WaitTimeout = period

	var before, after syscall.Rusage
	require.NoError(t, syscall.Getrusage(syscall.RUSAGE_SELF, &before))
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	_ = ws.Run(ctx)
	require.NoError(t, syscall.Getrusage(syscall.RUSAGE_SELF, &after))

	cpu := time.Duration(after.Utime.Nano()+after.Stime.Nano()) -
		time.Duration(before.Utime.Nano()+before.Stime.Nano())
	iterations := int(duration / period)
	t.Logf("~%d idle iterations cost %v CPU (~%v each); at DefaultWaitTimeout=%v that is ~%.3f%% of one core per wait set",
		iterations, cpu, cpu/time.Duration(iterations), rclgo.DefaultWaitTimeout,
		100*float64(cpu)/float64(iterations)/float64(rclgo.DefaultWaitTimeout))
}

// startStringPublisher publishes on topic from its own context every
// millisecond until the test ends, so late subscriptions still see traffic.
// The returned function stops publishing and closes the context early.
func startStringPublisher(t *testing.T, topic string) (stop func()) {
	t.Helper()
	pubCtx, err := newDefaultRCLContext()
	require.NoError(t, err)
	node, err := pubCtx.NewNode("ws_test_pub", "")
	require.NoError(t, err)
	pub, err := node.NewPublisher(topic, std_msgs.StringTypeSupport, nil)
	require.NoError(t, err)
	done := make(chan struct{})
	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		ticker := time.NewTicker(time.Millisecond)
		defer ticker.Stop()
		msg := &std_msgs.String{Data: "hello"}
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				_ = pub.Publish(msg)
			}
		}
	}()
	var once sync.Once
	stop = func() {
		once.Do(func() {
			close(done)
			<-stopped
			pubCtx.Close()
		})
	}
	t.Cleanup(stop)
	return stop
}

// startWaitSet runs ws on its own goroutine and returns a function that stops
// it and returns Run's error.
func startWaitSet(t *testing.T, ws *rclgo.WaitSet) func() error {
	t.Helper()
	ctx, cancel := context.WithCancel(context.Background())
	runErr := make(chan error, 1)
	go func() { runErr <- ws.Run(ctx) }()
	var once sync.Once
	var err error
	stop := func() error {
		once.Do(func() {
			cancel()
			select {
			case err = <-runErr:
			case <-time.After(5 * time.Second):
				err = errors.New("Run did not return after cancellation")
			}
		})
		return err
	}
	t.Cleanup(func() { _ = stop() })
	return stop
}

// newSignalingSubscription creates a subscription that closes the returned
// channel on its first message.
func newSignalingSubscription(node *rclgo.Node, topic string) (*rclgo.Subscription, <-chan struct{}, error) {
	received := make(chan struct{})
	var once sync.Once
	sub, err := node.NewSubscription(topic, std_msgs.StringTypeSupport, nil, func(s *rclgo.Subscription) {
		if _, _, err := s.TakeSerializedMessage(); err == nil {
			once.Do(func() { close(received) })
		}
	})
	return sub, received, err
}

func waitReceived(t *testing.T, received <-chan struct{}, what string) {
	t.Helper()
	select {
	case <-received:
	case <-time.After(10 * time.Second):
		t.Fatalf("%s received no message", what)
	}
}

func TestWaitSetConcurrentAddSubscriptions(t *testing.T) {
	const (
		topic    = "/ws_concurrent_add"
		subCount = 40
	)
	stopPublisher := startStringPublisher(t, topic)
	rclCtx, err := newDefaultRCLContext()
	require.NoError(t, err)
	t.Cleanup(func() {
		// rmw_zenoh can deadlock destroying a subscription that is still
		// receiving, so stop the traffic first.
		stopPublisher()
		rclCtx.Close()
	})
	node, err := rclCtx.NewNode("ws_concurrent_add", "")
	require.NoError(t, err)
	ws, err := rclCtx.NewWaitSet()
	require.NoError(t, err)

	// One subscription is present when Run starts so messages already flow
	// while the others are added.
	first, firstReceived, err := newSignalingSubscription(node, topic)
	require.NoError(t, err)
	ws.AddSubscriptions(first)
	stop := startWaitSet(t, ws)
	waitReceived(t, firstReceived, "initial subscription")

	receivedChans := make(chan (<-chan struct{}), subCount)
	addErr := make(chan error, 1)
	go func() {
		defer close(receivedChans)
		for range subCount {
			sub, received, err := newSignalingSubscription(node, topic)
			if err != nil {
				addErr <- err
				return
			}
			ws.AddSubscriptions(sub)
			receivedChans <- received
		}
	}()
	i := 0
	for received := range receivedChans {
		waitReceived(t, received, fmt.Sprintf("subscription %d", i))
		i++
	}
	select {
	case err := <-addErr:
		require.NoError(t, err)
	default:
	}
	require.Equal(t, subCount, i)
	require.ErrorIs(t, stop(), context.Canceled)
}

func TestWaitSetCall(t *testing.T) {
	const topic = "/ws_call"
	startStringPublisher(t, topic)
	rclCtx, err := newDefaultRCLContext()
	require.NoError(t, err)
	t.Cleanup(func() { rclCtx.Close() })
	node, err := rclCtx.NewNode("ws_call", "")
	require.NoError(t, err)
	ws, err := rclCtx.NewWaitSet()
	require.NoError(t, err)
	errSentinel := errors.New("sentinel")

	// Before Run, Call runs inline and must not block a later Run on the same
	// goroutine.
	var beforeReceived <-chan struct{}
	err = ws.Call(func() error {
		sub, received, err := newSignalingSubscription(node, topic)
		if err != nil {
			return err
		}
		beforeReceived = received
		ws.AddSubscriptions(sub)
		return errSentinel
	})
	require.ErrorIs(t, err, errSentinel)

	stop := startWaitSet(t, ws)
	waitReceived(t, beforeReceived, "subscription added before Run")

	var duringReceived <-chan struct{}
	err = ws.Call(func() error {
		sub, received, err := newSignalingSubscription(node, topic)
		if err != nil {
			return err
		}
		duringReceived = received
		ws.AddSubscriptions(sub)
		return errSentinel
	})
	require.ErrorIs(t, err, errSentinel)
	waitReceived(t, duringReceived, "subscription added by Call during Run")

	require.ErrorIs(t, stop(), context.Canceled)

	// After Run exits, Call runs inline again.
	ran := false
	require.NoError(t, ws.Call(func() error { ran = true; return nil }))
	require.True(t, ran)

	require.NoError(t, ws.Close())
	ran = false
	require.ErrorIs(t, ws.Call(func() error { ran = true; return nil }), rclgo.ErrWaitSetClosed)
	require.False(t, ran)
	require.ErrorIs(t, ws.Run(context.Background()), rclgo.ErrWaitSetClosed)
}

func TestWaitSetCallsDoNotOutliveRun(t *testing.T) {
	// Calls racing with Run's exit must all return, whether Run executed them
	// while running, while exiting, or they ran inline afterwards.
	rclCtx, err := newDefaultRCLContext()
	require.NoError(t, err)
	t.Cleanup(func() { rclCtx.Close() })
	ws, err := rclCtx.NewWaitSet()
	require.NoError(t, err)
	ws.WaitTimeout = -1
	stop := startWaitSet(t, ws)

	const callers = 8
	var executed atomic.Int64
	var wg sync.WaitGroup
	callerCtx, stopCallers := context.WithCancel(context.Background())
	for range callers {
		wg.Go(func() {
			for callerCtx.Err() == nil {
				_ = ws.Call(func() error { executed.Add(1); return nil })
			}
		})
	}
	time.Sleep(50 * time.Millisecond)
	require.ErrorIs(t, stop(), context.Canceled)
	time.Sleep(10 * time.Millisecond)
	stopCallers()

	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("a Call blocked after Run exited")
	}
	require.Positive(t, executed.Load())
}

func TestSpinConcurrentWithResourceCreation(t *testing.T) {
	// Run with -race: Spin collects the context's resources while another
	// goroutine adds timers to it.
	rclCtx, err := rclgo.NewContext(0, nil)
	require.NoError(t, err)
	t.Cleanup(func() { rclCtx.Close() })
	canceled, cancel := context.WithCancel(context.Background())
	cancel()

	const timers = 200
	var created atomic.Int32
	var wg sync.WaitGroup
	wg.Go(func() {
		for range timers {
			if _, err := rclCtx.NewTimer(time.Hour, func(*rclgo.Timer) {}); err != nil {
				t.Error(err)
				return
			}
			created.Add(1)
		}
	})
	spins := 0
	for spins == 0 || created.Load() < timers {
		require.ErrorIs(t, rclCtx.Spin(canceled), context.Canceled)
		spins++
	}
	wg.Wait()
	t.Logf("%d spins overlapped %d timer creations", spins, timers)
}

func TestWaitSetSharedSubscriptionRunsInOneWaitSetAtATime(t *testing.T) {
	const topic = "/ws_shared_sub"
	stopPublisher := startStringPublisher(t, topic)
	rclCtx, err := newDefaultRCLContext()
	require.NoError(t, err)
	t.Cleanup(func() {
		// rmw_zenoh can deadlock destroying a subscription that is still
		// receiving, so stop the traffic first.
		stopPublisher()
		rclCtx.Close()
	})
	node, err := rclCtx.NewNode("ws_shared_sub", "")
	require.NoError(t, err)
	var received atomic.Int64
	sub, err := node.NewSubscription(topic, std_msgs.StringTypeSupport, nil, func(s *rclgo.Subscription) {
		if _, _, err := s.TakeSerializedMessage(); err == nil {
			received.Add(1)
		}
	})
	require.NoError(t, err)
	receivesMore := func(what string) {
		t.Helper()
		before := received.Load()
		require.Eventually(t, func() bool { return received.Load() > before }, 10*time.Second, time.Millisecond, what)
	}

	ws1, err := rclCtx.NewWaitSet()
	require.NoError(t, err)
	ws1.AddSubscriptions(sub)
	ws2, err := rclCtx.NewWaitSet()
	require.NoError(t, err)
	ws2.AddSubscriptions(sub)

	stop1 := startWaitSet(t, ws1)
	receivesMore("first wait set")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	inUse, ok := errors.AsType[*rclgo.EntityInUseError](ws2.Run(ctx))
	require.True(t, ok, "second Run did not report the shared subscription")
	assert.Equal(t, "subscription", inUse.Kind)
	assert.Equal(t, topic, inUse.Name)
	receivesMore("first wait set after the conflict")

	require.ErrorIs(t, stop1(), context.Canceled)
	stop2 := startWaitSet(t, ws2)
	receivesMore("second wait set after the first stopped")
	require.ErrorIs(t, stop2(), context.Canceled)
}

func TestSpinFailsOnEntityInRunningWaitSet(t *testing.T) {
	rclCtx, err := rclgo.NewContext(0, nil)
	require.NoError(t, err)
	t.Cleanup(func() { rclCtx.Close() })
	var fired atomic.Int64
	timer, err := rclCtx.NewTimer(time.Millisecond, func(*rclgo.Timer) { fired.Add(1) })
	require.NoError(t, err)
	ws, err := rclCtx.NewWaitSet()
	require.NoError(t, err)
	ws.AddTimers(timer)
	stop := startWaitSet(t, ws)
	require.Eventually(t, func() bool { return fired.Load() > 0 }, 5*time.Second, time.Millisecond, "wait set never ran")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	inUse, ok := errors.AsType[*rclgo.EntityInUseError](rclCtx.Spin(ctx))
	require.True(t, ok, "Spin did not report the timer held by the running wait set")
	assert.Equal(t, "timer", inUse.Kind)

	require.ErrorIs(t, stop(), context.Canceled)
	canceled, cancelNow := context.WithCancel(context.Background())
	cancelNow()
	require.ErrorIs(t, rclCtx.Spin(canceled), context.Canceled)
}
