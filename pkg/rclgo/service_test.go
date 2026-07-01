package rclgo_test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey" //nolint:revive

	example_interfaces_srv "github.com/ATIinc/rclgo/internal/msgs/example_interfaces/srv"
	"github.com/ATIinc/rclgo/pkg/rclgo"
	"github.com/ATIinc/rclgo/pkg/rclgo/types"
)

func TestServiceAndClient(t *testing.T) {
	type testSendResult struct {
		req  types.Message
		resp types.Message
		info *rclgo.ServiceInfo
		err  error
		sum  int64
	}
	var (
		serviceCtx, clientCtx *rclgo.Context
		client                *rclgo.Client
		err                   error

		spinCtx, cancelSpin = context.WithCancel(context.Background())
		spinErrs            = make(chan error, 2)

		requestReceivedChan = make(
			chan *example_interfaces_srv.AddTwoInts_Request,
			1,
		)
		responseSentErrChan = make(chan error, 1)

		randGen = rand.NewSource(42)

		qosProfile = rclgo.NewDefaultServiceQosProfile()
	)
	qosProfile.History = rclgo.HistoryKeepAll
	sendReq := func(a, b int64) *testSendResult {
		req := example_interfaces_srv.NewAddTwoInts_Request()
		req.A = a
		req.B = b
		result := testSendResult{req: req, sum: a + b}
		result.resp, result.info, result.err = client.Send(spinCtx, req)
		return &result
	}
	defer func() {
		cancelSpin()
		if serviceCtx != nil {
			serviceCtx.Close()
		}
		if clientCtx != nil {
			clientCtx.Close()
		}
	}()
	Convey("Scenario: Client calls a service", t, func() {
		Convey("Create a service", func() {
			serviceCtx, err = newDefaultRCLContext()
			So(err, ShouldBeNil)
			node, err := serviceCtx.NewNode("service_node", "/test")
			So(err, ShouldBeNil)
			_, err = node.NewService(
				"add",
				example_interfaces_srv.AddTwoIntsTypeSupport,
				&rclgo.ServiceOptions{Qos: qosProfile},
				func(_ *rclgo.ServiceInfo, rm types.Message, srs rclgo.ServiceResponseSender) {
					req := rm.(*example_interfaces_srv.AddTwoInts_Request)
					requestReceivedChan <- req
					resp := example_interfaces_srv.NewAddTwoInts_Response()
					resp.Sum = req.A + req.B
					responseSentErrChan <- srs.SendResponse(resp)
				},
			)
			So(err, ShouldBeNil)
			go func() { spinErrs <- serviceCtx.Spin(spinCtx) }()
		})
		Convey("Create a client", func() {
			clientCtx, err = newDefaultRCLContext()
			So(err, ShouldBeNil)
			node, err := clientCtx.NewNode("client_node", "/test")
			So(err, ShouldBeNil)
			client, err = node.NewClient(
				"add",
				example_interfaces_srv.AddTwoIntsTypeSupport,
				&rclgo.ClientOptions{Qos: qosProfile},
			)
			So(err, ShouldBeNil)
			go func() { spinErrs <- clientCtx.Spin(spinCtx) }()
		})
		Convey("The client sends a request", func() {
			time.Sleep(200 * time.Millisecond)
			var result *testSendResult
			timeOut(2000, func() { result = sendReq(3, -7) }, "Sending request")

			So(result.err, ShouldBeNil)
			So(result.info, ShouldNotBeNil)
			So(
				result.resp.(*example_interfaces_srv.AddTwoInts_Response).Sum,
				ShouldEqual,
				-4,
			)

			So(<-requestReceivedChan, ShouldResemble, result.req)
			So(<-responseSentErrChan, ShouldBeNil)
		})
		Convey("The client sends many requests in quick succession", func() {
			const reqCount = 100
			testResults := make(chan *testSendResult, reqCount)
			requestReceivedChan = make(
				chan *example_interfaces_srv.AddTwoInts_Request,
				reqCount,
			)
			responseSentErrChan = make(chan error, reqCount)
			for range reqCount {
				a, b := randGen.Int63(), randGen.Int63()
				go func() { testResults <- sendReq(a, b) }()
			}
			for range reqCount {
				res := <-testResults
				So(res, ShouldNotBeNil)
				So(res.err, ShouldBeNil)
				So(res.info, ShouldNotBeNil)
				So(
					res.resp.(*example_interfaces_srv.AddTwoInts_Response).Sum,
					ShouldEqual,
					res.sum,
				)
			}
		})
		Convey("The service and client are stopped", func() {
			cancelSpin()
			So(<-spinErrs, shouldContainError, context.Canceled)
			So(<-spinErrs, shouldContainError, context.Canceled)
		})
		Convey("The service context is closed without errors", func() {
			timeOut(2000, func() {
				err = serviceCtx.Close()
			}, "Service context is closing")
			So(err, ShouldBeNil)
		})
		Convey("The client context is closed without errors", func() {
			timeOut(2000, func() {
				err = clientCtx.Close()
			}, "Client context is closing")
			So(err, ShouldBeNil)
		})
	})
}

func TestServiceIntrospection(t *testing.T) {
	var (
		rclCtx *rclgo.Context
		err    error
	)
	defer func() {
		if rclCtx != nil {
			rclCtx.Close()
		}
	}()
	Convey("Scenario: configuring service introspection", t, func() {
		rclCtx, err = newDefaultRCLContext()
		So(err, ShouldBeNil)
		node, err := rclCtx.NewNode("introspection_node", "/test")
		So(err, ShouldBeNil)

		Convey("A service defaults to introspection off", func() {
			service, err := node.NewService(
				"introspect_add",
				example_interfaces_srv.AddTwoIntsTypeSupport,
				nil,
				func(_ *rclgo.ServiceInfo, _ types.Message, _ rclgo.ServiceResponseSender) {},
			)
			So(err, ShouldBeNil)
			So(service.IntrospectionState(), ShouldEqual, rclgo.ServiceIntrospectionOff)

			Convey("Its introspection state can be changed at runtime", func() {
				So(service.ConfigureIntrospection(rclgo.ServiceIntrospectionContents), ShouldBeNil)
				So(service.IntrospectionState(), ShouldEqual, rclgo.ServiceIntrospectionContents)

				So(service.ConfigureIntrospection(rclgo.ServiceIntrospectionMetadata), ShouldBeNil)
				So(service.IntrospectionState(), ShouldEqual, rclgo.ServiceIntrospectionMetadata)

				So(service.ConfigureIntrospection(rclgo.ServiceIntrospectionOff), ShouldBeNil)
				So(service.IntrospectionState(), ShouldEqual, rclgo.ServiceIntrospectionOff)
			})
		})

		Convey("A client can enable introspection via options", func() {
			client, err := node.NewClient(
				"introspect_add",
				example_interfaces_srv.AddTwoIntsTypeSupport,
				&rclgo.ClientOptions{
					Qos:           rclgo.NewDefaultServiceQosProfile(),
					Introspection: rclgo.ServiceIntrospectionMetadata,
				},
			)
			So(err, ShouldBeNil)
			So(client.IntrospectionState(), ShouldEqual, rclgo.ServiceIntrospectionMetadata)

			So(client.ConfigureIntrospection(rclgo.ServiceIntrospectionOff), ShouldBeNil)
			So(client.IntrospectionState(), ShouldEqual, rclgo.ServiceIntrospectionOff)
		})
	})
}

func TestServiceIntrospectionStateZeroValue(t *testing.T) {
	Convey("The zero value of ServiceIntrospectionState is ServiceIntrospectionOff", t, func() {
		// Guards against the (unlikely) possibility that the underlying RCL
		// enum value RCL_SERVICE_INTROSPECTION_OFF is not zero. Callers rely on
		// the zero value being "off" so that the default ServiceOptions /
		// ClientOptions disable introspection.
		var state rclgo.ServiceIntrospectionState
		So(state, ShouldEqual, rclgo.ServiceIntrospectionOff)
	})
}
