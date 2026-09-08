package rclgo_test

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey" //nolint:revive

	test_msgs_action "github.com/ATIinc/rclgo/internal/msgs/test_msgs/action"
	"github.com/ATIinc/rclgo/pkg/rclgo"
	"github.com/ATIinc/rclgo/pkg/rclgo/types"
)

// waitForAvailability polls client.IsServerAvailable until it reports want or
// the timeout elapses, returning the last observed value and error.
func waitForAvailability(client *rclgo.ActionClient, want bool, timeout time.Duration) (bool, error) {
	deadline := time.Now().Add(timeout)
	for {
		avail, err := client.IsServerAvailable()
		if err != nil || avail == want || time.Now().After(deadline) {
			return avail, err
		}
		time.Sleep(20 * time.Millisecond)
	}
}

func TestActionClientIsServerAvailable(t *testing.T) {
	var (
		rclctx     *rclgo.Context
		clientNode *rclgo.Node
		serverNode *rclgo.Node
		client     *rclgo.ActionClient
		server     *rclgo.ActionServer
		err        error
	)
	defer func() {
		if rclctx != nil {
			rclctx.Close()
		}
	}()
	Convey("Scenario: ActionClient.IsServerAvailable tracks the presence of an ActionServer", t, func() {
		Convey("Create an ActionClient with no server", func() {
			rclctx, err = newDefaultRCLContext()
			So(err, ShouldBeNil)
			clientNode, err = rclctx.NewNode("availability_client", "actions_test")
			So(err, ShouldBeNil)
			client, err = clientNode.NewActionClient(
				"availability",
				test_msgs_action.FibonacciTypeSupport,
				actionClientOpts,
			)
			So(err, ShouldBeNil)
		})
		Convey("The server is reported unavailable before it exists", func() {
			avail, err := client.IsServerAvailable()
			So(err, ShouldBeNil)
			So(avail, ShouldBeFalse)
		})
		Convey("The server is reported available once it has been created and discovered", func() {
			serverNode, err = rclctx.NewNode("availability_server", "actions_test")
			So(err, ShouldBeNil)
			_, action := newWaitAction()
			server, err = serverNode.NewActionServer("availability", action, actionServerOpts)
			So(err, ShouldBeNil)
			avail, err := waitForAvailability(client, true, 10*time.Second)
			So(err, ShouldBeNil)
			So(avail, ShouldBeTrue)
		})
		Convey("The server is reported unavailable again after it is closed", func() {
			So(server.Close(), ShouldBeNil)
			avail, err := waitForAvailability(client, false, 10*time.Second)
			So(err, ShouldBeNil)
			So(avail, ShouldBeFalse)
		})
		Convey("A closed client reports an error", func() {
			So(client.Close(), ShouldBeNil)
			_, err := client.IsServerAvailable()
			So(err, ShouldNotBeNil)
		})
		Convey("Resources are released properly", func() {
			So(rclctx.Close(), ShouldBeNil)
		})
	})
}

// TestActionClientIsServerAvailableWhileSpinning exercises the availability
// probe from many goroutines while the client's node is spinning and goals
// are executing, i.e. while the wait set concurrently registers and takes from
// the same rcl client handle. Run under -race: the probe, the wait-set
// registration and the take path must all serialize on the client's handle
// lock.
func TestActionClientIsServerAvailableWhileSpinning(t *testing.T) {
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
	Convey("Scenario: IsServerAvailable is safe to call concurrently while spinning", t, func() {
		Convey("Create a spinning server and client", func() {
			rclctx, err = newDefaultRCLContext()
			So(err, ShouldBeNil)
			serverNode, err = rclctx.NewNode("availability_spin_server", "actions_test")
			So(err, ShouldBeNil)
			_, err = serverNode.NewActionServer("availability_spin", action, actionServerOpts)
			So(err, ShouldBeNil)
			clientNode, err = rclctx.NewNode("availability_spin_client", "actions_test")
			So(err, ShouldBeNil)
			client, err = clientNode.NewActionClient(
				"availability_spin",
				test_msgs_action.FibonacciTypeSupport,
				actionClientOpts,
			)
			So(err, ShouldBeNil)
			go func() { spinErr <- serverNode.Spin(ctx) }()
			go func() { spinErr <- clientNode.Spin(ctx) }()
			avail, err := waitForAvailability(client, true, 10*time.Second)
			So(err, ShouldBeNil)
			So(avail, ShouldBeTrue)
		})
		Convey("Concurrent probes and goals complete without error", func() {
			const probers = 8
			const probesEach = 200
			var wg sync.WaitGroup
			probeErrs := make(chan error, probers*probesEach)
			for range probers {
				wg.Go(func() {
					for range probesEach {
						avail, err := client.IsServerAvailable()
						if err != nil {
							probeErrs <- err
							return
						}
						if !avail {
							probeErrs <- errors.New("server reported unavailable while spinning")
							return
						}
					}
				})
			}
			// goals in flight keep the take path busy on the same client handle
			for range 5 {
				goal := test_msgs_action.NewFibonacci_Goal()
				goal.Order = 10
				for range int(goal.Order) {
					action.continueChan <- struct{}{}
				}
				result, _, err := client.WatchGoal(ctx, goal, func(context.Context, types.Message) {})
				So(err, ShouldBeNil)
				So(result, ShouldNotBeNil)
			}
			wg.Wait()
			close(probeErrs)
			for err := range probeErrs {
				So(err, ShouldBeNil)
			}
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
