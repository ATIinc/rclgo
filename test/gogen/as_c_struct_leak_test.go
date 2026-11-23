package test

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	test_msgs "github.com/ATIinc/rclgo/internal/msgs/test_msgs/msg"
	"github.com/ATIinc/rclgo/pkg/rclgo/types"
)

// PrepareMemory runs the rosidl __init, which allocates a buffer for every string
// and the IDL default of every sequence. AsCStruct must free those before it
// replaces them, or every publish, request, response and goal leaks.
func TestAsCStructDoesNotLeak(t *testing.T) {
	const (
		warmup     = 1_000
		iterations = 100_000
		// A leak of even one 1-byte string per conversion is 32 bytes of heap
		// per iteration, about 3 MiB over the loop.
		maxGrowth = 64 << 10
	)
	cases := []struct {
		name string
		ts   types.MessageTypeSupport
		msg  types.Message
	}{
		{"Strings zero", test_msgs.StringsTypeSupport, &test_msgs.Strings{}},
		{"Strings defaults", test_msgs.StringsTypeSupport, test_msgs.NewStrings()},
		{"WStrings zero", test_msgs.WStringsTypeSupport, &test_msgs.WStrings{}},
		{"WStrings fixture", test_msgs.WStringsTypeSupport, Fixture_Go_test_msgs__WStrings()},
		{"Arrays zero", test_msgs.ArraysTypeSupport, &test_msgs.Arrays{}},
		{"Arrays fixture", test_msgs.ArraysTypeSupport, Fixture_Go_test_msgs__Arrays()},
		{"UnboundedSequences zero", test_msgs.UnboundedSequencesTypeSupport, &test_msgs.UnboundedSequences{}},
		{"UnboundedSequences fixture", test_msgs.UnboundedSequencesTypeSupport, Fixture_Go_test_msgs__UnboundedSequences()},
		{"BoundedSequences zero", test_msgs.BoundedSequencesTypeSupport, &test_msgs.BoundedSequences{}},
		{"BoundedSequences fixture", test_msgs.BoundedSequencesTypeSupport, Fixture_Go_test_msgs__BoundedSequences()},
		{"MultiNested fixture", test_msgs.MultiNestedTypeSupport, Fixture_Go_test_msgs__MultiNested()},
	}

	convert := func(ts types.MessageTypeSupport, msg types.Message, n int) {
		for range n {
			buf := ts.PrepareMemory()
			ts.AsCStruct(buf, msg)
			ts.ReleaseMemory(buf)
		}
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// glibc keeps per-thread arenas and caches; stay on one thread so
			// the warmup primes the same ones the measured loop uses.
			runtime.LockOSThread()
			defer runtime.UnlockOSThread()
			convert(c.ts, c.msg, warmup)
			before := cHeapInUse()
			convert(c.ts, c.msg, iterations)
			growth := cHeapInUse() - before
			t.Logf("C heap growth over %d conversions: %d bytes", iterations, growth)
			assert.Less(t, growth, int64(maxGrowth))
		})
	}
}
