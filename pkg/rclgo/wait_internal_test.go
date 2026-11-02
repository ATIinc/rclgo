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

func TestWaitSetReservesEntitiesAddedWhileRunning(t *testing.T) {
	rclCtx := newTestContext(t)
	ws, err := rclCtx.NewWaitSet()
	require.NoError(t, err)
	fired := make(chan struct{}, 1)
	timer, err := rclCtx.NewTimer(time.Millisecond, func(*Timer) {
		select {
		case fired <- struct{}{}:
		default:
		}
	})
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	runErr := make(chan error, 1)
	go func() { runErr <- ws.Run(ctx) }()
	ws.AddTimers(timer)
	select {
	case <-fired:
	case <-time.After(5 * time.Second):
		t.Fatal("timer added while running never fired")
	}
	assert.True(t, (*atomic.Bool)(&timer.waitable).Load(), "timer not reserved while running")
	cancel()
	require.ErrorIs(t, <-runErr, context.Canceled)
	assert.False(t, (*atomic.Bool)(&timer.waitable).Load(), "timer still reserved after Run returned")
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
