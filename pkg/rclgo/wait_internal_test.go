/*
This file is part of rclgo

Licensed under the Apache License, Version 2.0 (the "License");
    http://www.apache.org/licenses/LICENSE-2.0
*/

package rclgo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
