package rclgo_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	std_msgs "github.com/ATIinc/rclgo/internal/msgs/std_msgs/msg"
	"github.com/ATIinc/rclgo/pkg/rclgo"
)

// newMatchedPair creates a publisher and a callback-less subscription on topic
// in separate contexts and waits until they match.
func newMatchedPair(t *testing.T, topic string) (*rclgo.Publisher, *rclgo.Subscription) {
	t.Helper()
	pubCtx, pub, err := newContextWithPublisher(nil, "take_pub", "", topic, std_msgs.ColorRGBATypeSupport)
	require.NoError(t, err)
	t.Cleanup(func() { pubCtx.Close() })
	subCtx, err := newDefaultRCLContext()
	require.NoError(t, err)
	t.Cleanup(func() { subCtx.Close() })
	node, err := subCtx.NewNode("take_sub", "")
	require.NoError(t, err)
	opts := rclgo.NewDefaultSubscriptionOptions()
	opts.Qos = reliableQos
	sub, err := node.NewSubscription(topic, std_msgs.ColorRGBATypeSupport, opts, nil)
	require.NoError(t, err)
	require.Eventually(t, func() bool {
		n, err := pub.GetSubscriptionCount()
		return err == nil && n == 1
	}, 5*time.Second, 10*time.Millisecond)
	return pub, sub
}

// takeEventually retries a take until a message has arrived.
func takeEventually(t *testing.T, take func() error) {
	t.Helper()
	var err error
	require.Eventually(t, func() bool {
		err = take()
		return err == nil
	}, 5*time.Second, time.Millisecond, "last take error: %v", err)
}

func TestTakeSerializedMessageInto(t *testing.T) {
	const topic = "/take_serialized_into"
	pub, sub := newMatchedPair(t, topic)
	msg := &std_msgs.ColorRGBA{R: 1, G: 2, B: 3, A: 4}
	want, err := rclgo.Serialize(msg)
	require.NoError(t, err)

	t.Run("round trip", func(t *testing.T) {
		require.NoError(t, pub.Publish(msg))
		var got []byte
		var info rclgo.MessageInfo
		takeEventually(t, func() (err error) {
			got, err = sub.TakeSerializedMessageInto(nil, &info)
			return err
		})
		assert.Equal(t, want, got)
		back, err := rclgo.Deserialize(got, std_msgs.ColorRGBATypeSupport)
		require.NoError(t, err)
		assert.Equal(t, msg, back)
		assert.False(t, info.ReceivedTimestamp.IsZero())

		pubs, err := sub.Node().GetPublishersInfoByTopic(topic, false)
		require.NoError(t, err)
		require.Len(t, pubs, 1)
		assert.NotEqual(t, rclgo.GID{}, info.PublisherGID)
		assert.Equal(t, pubs[0].EndpointGID, info.PublisherGID)
		t.Logf("publication seq %d, reception seq %d", info.PublicationSequenceNumber, info.ReceptionSequenceNumber)

		// TakeMessage reports the same publisher.
		require.NoError(t, pub.Publish(msg))
		var typed std_msgs.ColorRGBA
		var typedInfo *rclgo.MessageInfo
		takeEventually(t, func() (err error) {
			typedInfo, err = sub.TakeMessage(&typed)
			return err
		})
		assert.Equal(t, pubs[0].EndpointGID, typedInfo.PublisherGID)
		assert.Greater(t, typedInfo.ReceptionSequenceNumber, info.ReceptionSequenceNumber)
	})

	t.Run("reuses dst", func(t *testing.T) {
		require.NoError(t, pub.Publish(msg))
		dst := make([]byte, 0, 256)
		var got []byte
		takeEventually(t, func() (err error) {
			got, err = sub.TakeSerializedMessageInto(dst, nil)
			return err
		})
		assert.Equal(t, want, got)
		assert.Same(t, &dst[:1][0], &got[0])
	})

	t.Run("no allocations in steady state", func(t *testing.T) {
		const runs = 100
		// AllocsPerRun makes one warm-up call in addition to runs.
		for range runs + 1 {
			require.NoError(t, pub.Publish(msg))
		}
		dst := make([]byte, 0, 256)
		takeEventually(t, func() (err error) {
			dst, err = sub.TakeSerializedMessageInto(dst, nil)
			return err
		})
		time.Sleep(200 * time.Millisecond) // let the rest of the burst arrive
		var info rclgo.MessageInfo
		failures := 0
		allocs := testing.AllocsPerRun(runs-1, func() {
			var err error
			if dst, err = sub.TakeSerializedMessageInto(dst, &info); err != nil {
				failures++
			}
		})
		require.Zero(t, failures, "takes failed; not all messages arrived")
		assert.Zero(t, allocs)
	})

	t.Run("no message available", func(t *testing.T) {
		var takeFailed *rclgo.SubscriptionTakeFailed
		_, err := sub.TakeSerializedMessageInto(nil, nil)
		require.ErrorAs(t, err, &takeFailed)
		var typed std_msgs.ColorRGBA
		_, err = sub.TakeMessage(&typed)
		require.ErrorAs(t, err, &takeFailed)
		// Callers poll for this result, so it must not allocate.
		allocs := testing.AllocsPerRun(100, func() {
			_, _ = sub.TakeSerializedMessageInto(nil, nil)
		})
		assert.Zero(t, allocs)
	})

	t.Run("take after close", func(t *testing.T) {
		require.NoError(t, sub.Close())
		_, err := sub.TakeSerializedMessageInto(nil, nil)
		var invalid *rclgo.SubscriptionInvalid
		require.ErrorAs(t, err, &invalid)
		_, _, err = sub.TakeSerializedMessage()
		require.ErrorAs(t, err, &invalid)
	})
}
