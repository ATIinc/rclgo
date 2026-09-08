package rclgo_test

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey" //nolint:revive

	test_msgs_action "github.com/ATIinc/rclgo/internal/msgs/test_msgs/action"
	"github.com/ATIinc/rclgo/pkg/rclgo"
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
