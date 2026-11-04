package rclgo_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	std_msgs "github.com/ATIinc/rclgo/internal/msgs/std_msgs/msg"
	"github.com/ATIinc/rclgo/pkg/rclgo"
)

// BenchmarkWaitSetInitEntities measures the per-iteration refill of a wait set
// holding as many subscriptions as a whole-graph recorder.
func BenchmarkWaitSetInitEntities(b *testing.B) {
	const subscriptionCount = 700
	rclCtx, err := rclgo.NewContext(0, nil)
	require.NoError(b, err)
	b.Cleanup(func() { rclCtx.Close() })
	node, err := rclCtx.NewNode("bench_init_entities", "")
	require.NoError(b, err)
	ws, err := rclCtx.NewWaitSet()
	require.NoError(b, err)
	for i := range subscriptionCount {
		sub, err := node.NewSubscription(fmt.Sprintf("/bench_init_entities/t%d", i), std_msgs.EmptyTypeSupport, nil, nil)
		require.NoError(b, err)
		ws.AddSubscriptions(sub)
	}
	require.NoError(b, ws.Testing_initEntities())
	for b.Loop() {
		if err := ws.Testing_initEntities(); err != nil {
			b.Fatal(err)
		}
	}
}
