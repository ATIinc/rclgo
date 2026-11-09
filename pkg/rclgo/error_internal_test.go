package rclgo

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrorStringIsNotStale(t *testing.T) {
	// rcutils error state is per thread; pin the test so both errors are built
	// on the thread whose state the first failure set.
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	rclCtx := newTestContext(t)

	_, err := rclCtx.NewNode("invalid node name!", "")
	require.Error(t, err)
	assert.NotContains(t, err.Error(), "error not set", "the rcl failure set no message")

	// A code that sets no message, like a take that found nothing, must not
	// report the previous failure's text.
	second := errorsCastC(1, "second failure")
	assert.Contains(t, second.Error(), "second failure error not set")
	assert.NotContains(t, second.Error(), "node name")
}
