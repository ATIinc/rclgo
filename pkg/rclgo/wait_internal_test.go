/*
This file is part of rclgo

Licensed under the Apache License, Version 2.0 (the "License");
    http://www.apache.org/licenses/LICENSE-2.0
*/

package rclgo

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWaitTimeoutNanos(t *testing.T) {
	tests := []struct {
		name    string
		timeout time.Duration
		want    int64
	}{
		{"zero uses the package default", 0, DefaultWaitTimeout.Nanoseconds()},
		{"negative waits forever", -1, -1},
		{"very negative waits forever", -time.Hour, -1},
		{"positive is passed through", 250 * time.Millisecond, (250 * time.Millisecond).Nanoseconds()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ws WaitSet
			ws.WaitTimeout = tt.timeout
			assert.Equal(t, tt.want, ws.waitTimeoutNanos())
		})
	}
}

func TestRecoveredLostWakeup(t *testing.T) {
	tests := []struct {
		name         string
		prevTimedOut bool
		elapsed      time.Duration
		want         bool
	}{
		{"immediate ready after a timeout is a recovery", true, 20 * time.Microsecond, true},
		{"just under the threshold counts", true, LostWakeupThreshold - time.Nanosecond, true},
		{"at the threshold does not count", true, LostWakeupThreshold, false},
		{"a message arriving later in the wait is normal traffic", true, 30 * time.Millisecond, false},
		{"immediate ready after a ready iteration is a busy wait set", false, 20 * time.Microsecond, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, recoveredLostWakeup(tt.prevTimedOut, tt.elapsed))
		})
	}
}

func (w *WaitSet) queueLen() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return len(w.queue)
}

func newTestContext(t *testing.T) *Context {
	t.Helper()
	ctx, err := NewContext(0, nil)
	require.NoError(t, err)
	t.Cleanup(func() { ctx.Close() })
	return ctx
}

// runInBackground runs ws until the returned function cancels it, which then
// returns Run's error.
func runInBackground(t *testing.T, ws *WaitSet) func() error {
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

func TestWaitSetSkipsConflictingAddWhileRunning(t *testing.T) {
	rclCtx := newTestContext(t)
	reserved := func(timer *Timer) bool { return (*atomic.Bool)(&timer.waitable).Load() }
	var a, b, c, aInFlight, aOverlaps atomic.Int64
	timerA, err := rclCtx.NewTimer(time.Millisecond, func(*Timer) {
		if aInFlight.Add(1) > 1 {
			aOverlaps.Add(1)
		}
		a.Add(1)
		time.Sleep(100 * time.Microsecond)
		aInFlight.Add(-1)
	})
	require.NoError(t, err)
	timerB, err := rclCtx.NewTimer(time.Millisecond, func(*Timer) { b.Add(1) })
	require.NoError(t, err)
	timerC, err := rclCtx.NewTimer(time.Millisecond, func(*Timer) { c.Add(1) })
	require.NoError(t, err)

	ws1, err := rclCtx.NewWaitSet()
	require.NoError(t, err)
	ws1.AddTimers(timerA)
	stop1 := runInBackground(t, ws1)
	ws2, err := rclCtx.NewWaitSet()
	require.NoError(t, err)
	ws2.AddTimers(timerB)
	stop2 := runInBackground(t, ws2)
	require.Eventually(t, func() bool { return a.Load() > 0 && b.Load() > 0 }, 5*time.Second, time.Millisecond)

	// timerA is held by ws1, so ws2 must skip it but still take timerC.
	ws2.AddTimers(timerA, timerC)
	require.NoError(t, ws2.Call(func() error {
		assert.Equal(t, []*Timer{timerB, timerC}, ws2.timers)
		return nil
	}))
	// A Call queued during the same drain runs before the cache rebuild.
	require.Eventually(t, func() bool {
		var cached uint64
		_ = ws2.Call(func() error { cached = uint64(ws2.handles.timers); return nil })
		return cached == 2
	}, 5*time.Second, time.Millisecond, "handle cache never held exactly timerB and timerC")
	aBefore, bBefore := a.Load(), b.Load()
	require.Eventually(t, func() bool {
		return c.Load() > 0 && a.Load() > aBefore && b.Load() > bBefore
	}, 5*time.Second, time.Millisecond, "timers stopped firing after the skipped add")
	assert.Zero(t, aOverlaps.Load(), "timerA dispatched by two wait sets at once")
	assert.True(t, reserved(timerC), "timerC not reserved while ws2 runs")

	require.ErrorIs(t, stop2(), context.Canceled)
	assert.True(t, reserved(timerA), "ws2 released timerA, which ws1 holds")
	assert.False(t, reserved(timerB))
	assert.False(t, reserved(timerC))
	require.ErrorIs(t, stop1(), context.Canceled)
	assert.False(t, reserved(timerA))
}

func TestWaitSetExitRunsPendingCalls(t *testing.T) {
	// A Call queued after Run's last drain must still run when Run exits.
	// A panicking callback forces the exit while the Call is queued.
	rclCtx := newTestContext(t)
	ws, err := rclCtx.NewWaitSet()
	require.NoError(t, err)
	entered := make(chan struct{})
	release := make(chan struct{})
	var once sync.Once
	timer, err := rclCtx.NewTimer(time.Millisecond, func(*Timer) {
		once.Do(func() { close(entered) })
		<-release
		panic("stop Run")
	})
	require.NoError(t, err)
	ws.AddTimers(timer)

	runPanic := make(chan any, 1)
	go func() {
		defer func() { runPanic <- recover() }()
		_ = ws.Run(context.Background())
	}()
	<-entered

	errSentinel := errors.New("sentinel")
	callErr := make(chan error, 1)
	go func() { callErr <- ws.Call(func() error { return errSentinel }) }()
	require.Eventually(t, func() bool { return ws.queueLen() == 1 }, 5*time.Second, time.Millisecond)
	close(release)

	assert.Equal(t, "stop Run", <-runPanic)
	select {
	case err := <-callErr:
		require.ErrorIs(t, err, errSentinel)
	case <-time.After(5 * time.Second):
		t.Fatal("Call queued before Run exited never returned")
	}
}

func TestWaitSetHandleCacheDropsClosedEntities(t *testing.T) {
	// The wait set caches entity handles in C memory; closing an entity must
	// invalidate the cache so the freed handle is never passed to rcl again.
	rclCtx := newTestContext(t)
	ws, err := rclCtx.NewWaitSet()
	require.NoError(t, err)
	timer, err := rclCtx.NewTimer(time.Hour, func(*Timer) {})
	require.NoError(t, err)
	ws.AddTimers(timer)
	require.NoError(t, ws.initEntities())
	require.Equal(t, unsafe.Pointer(timer.rcl_timer_t), unsafe.Slice(ws.handles.handles, 1)[0])

	require.NoError(t, timer.Close())
	err = ws.initEntities()
	require.ErrorContains(t, err, "rcl_wait_set_add_timer() failed")
	require.ErrorContains(t, err, "timers=1")
	assert.Nil(t, unsafe.Slice(ws.handles.handles, 1)[0])
}
