package rclgo_test

import (
	"context"
	"crypto/rand"
	"sync/atomic"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey" //nolint:revive

	test_msgs_action "github.com/ATIinc/rclgo/internal/msgs/test_msgs/action"
	"github.com/ATIinc/rclgo/pkg/rclgo"
	"github.com/ATIinc/rclgo/pkg/rclgo/types"
)

// TestWatchFeedbackChannelSignalsAfterDrain checks the contract callers rely on
// to tear down handler state: the channel returned by WatchFeedback delivers
// its value only after the handler has been unregistered and any in-flight
// invocation has returned, and no invocation starts after that.
func TestWatchFeedbackChannelSignalsAfterDrain(t *testing.T) {
	var (
		ctx, cancel = context.WithCancel(context.Background())
		rclctx      *rclgo.Context
		serverNode  *rclgo.Node
		clientNode  *rclgo.Node
		client      *rclgo.ActionClient
		err         error
		spinErr     = make(chan error, 2)
		action      = &fibonacciAction{continueChan: make(chan struct{}, 100)}
	)
	defer func() {
		cancel()
		if rclctx != nil {
			rclctx.Close()
		}
	}()
	Convey("Scenario: WatchFeedback's channel means the handler is done", t, func() {
		Convey("Create a spinning server and client", func() {
			rclctx, err = newDefaultRCLContext()
			So(err, ShouldBeNil)
			serverNode, err = rclctx.NewNode("watch_teardown_server", "actions_test")
			So(err, ShouldBeNil)
			_, err = serverNode.NewActionServer("watch_teardown", action, actionServerOpts)
			So(err, ShouldBeNil)
			clientNode, err = rclctx.NewNode("watch_teardown_client", "actions_test")
			So(err, ShouldBeNil)
			client, err = clientNode.NewActionClient(
				"watch_teardown",
				test_msgs_action.FibonacciTypeSupport,
				actionClientOpts,
			)
			So(err, ShouldBeNil)
			go func() { spinErr <- serverNode.Spin(ctx) }()
			go func() { spinErr <- clientNode.Spin(ctx) }()
			avail, err := waitForAvailability(client, true)
			So(err, ShouldBeNil)
			So(avail, ShouldBeTrue)
		})
		Convey("The channel waits for an in-flight handler and nothing runs after it", func() {
			goal := test_msgs_action.NewFibonacci_Goal()
			goal.Order = 10
			req := test_msgs_action.NewFibonacci_SendGoal_Request()
			req.SetGoalDescription(goal)
			goalID := req.GetGoalID()
			_, err = rand.Read(goalID[:])
			So(err, ShouldBeNil)

			entered := make(chan struct{}, 1)
			release := make(chan struct{})
			var calls atomic.Int32
			var afterSignal atomic.Int32
			var signaled atomic.Bool
			watchCtx, stopWatch := context.WithCancel(ctx)
			errc := client.WatchFeedback(watchCtx, goalID, func(_ context.Context, _ types.Message) {
				if signaled.Load() {
					afterSignal.Add(1)
				}
				if calls.Add(1) == 1 {
					entered <- struct{}{}
					<-release // hold the first invocation open
				}
			})

			// the server sends one feedback per continue token, so let exactly one through
			resp, err := client.SendGoalRequest(ctx, req)
			So(err, ShouldBeNil)
			So(resp.(*test_msgs_action.Fibonacci_SendGoal_Response).Accepted, ShouldBeTrue)
			action.continueChan <- struct{}{}
			timeOut(5000, func() { <-entered }, "Waiting for the first feedback invocation")

			// cancel the watch while the handler is blocked: the channel must not fire yet
			stopWatch()
			select {
			case <-errc:
				t.Fatal("channel signaled while the handler was still running")
			case <-time.After(100 * time.Millisecond):
			}

			// release the handler: now the channel fires, and later feedback must not reach the handler
			close(release)
			timeOut(5000, func() { <-errc }, "Waiting for the watch channel after the handler returned")
			signaled.Store(true)
			for range int(goal.Order) {
				action.continueChan <- struct{}{}
			}
			result, err := client.GetResult(ctx, goalID)
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
			So(afterSignal.Load(), ShouldEqual, 0)
		})
		Convey("Resources are released properly", func() {
			cancel()
			timeOut(1000, func() {
				<-spinErr
				<-spinErr
			}, "Waiting for spinning to stop")
			So(rclctx.Close(), ShouldBeNil)
		})
	})
}
