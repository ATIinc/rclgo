/*
This file is part of rclgo

Licensed under the Apache License, Version 2.0 (the "License");
    http://www.apache.org/licenses/LICENSE-2.0
*/

package rclgo_test

import (
	"context"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

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
