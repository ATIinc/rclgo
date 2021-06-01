/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
    http://www.apache.org/licenses/LICENSE-2.0
*/

package ros2

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	std_msgs "github.com/tiiuae/rclgo/pkg/ros2/msgs/std_msgs/msg"
	std_srvs_srv "github.com/tiiuae/rclgo/pkg/ros2/msgs/std_srvs/srv"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
)

func TestContextClose(t *testing.T) {
	var context *Context
	defer func() {
		if context != nil {
			context.Close()
		}
	}()
	SetDefaultFailureMode(FailureContinues)
	Convey("Scenario: Does Context handle close resources correctly", t, func() {
		Convey("Given a context with resources", func() {
			var err error
			context, err = newDefaultRCLContext()
			So(err, ShouldBeNil)
			_, err = context.NewNode("node1", "/test/context_close")
			So(err, ShouldBeNil)
			node2, err := context.NewNode("node2", "/test/context_close")
			So(err, ShouldBeNil)
			_, err = context.NewWaitSet(time.Second)
			So(err, ShouldBeNil)
			_, err = node2.NewClient("client2", std_srvs_srv.Empty, nil)
			So(err, ShouldBeNil)
			_, err = context.NewNode("node3", "/test/context_close")
			So(err, ShouldBeNil)
			_, err = node2.NewPublisher("/test_topic", &std_msgs.String{})
			So(err, ShouldBeNil)
			_, err = node2.NewClient("client1", std_srvs_srv.Empty, nil)
			So(err, ShouldBeNil)
			_, err = node2.NewSubscription("/test_topic", &std_msgs.ColorRGBA{}, func(s *Subscription) {})
			So(err, ShouldBeNil)
			_, err = node2.NewPublisher("/test_topic2", &std_msgs.ColorRGBA{})
			So(err, ShouldBeNil)
			_, err = node2.NewService("service1", std_srvs_srv.Empty, nil, func(rsi *RmwServiceInfo, rm ros2types.ROS2Msg, srs ServiceResponseSender) {})
			So(err, ShouldBeNil)
			_, err = node2.NewService("service2", std_srvs_srv.Empty, nil, func(rsi *RmwServiceInfo, rm ros2types.ROS2Msg, srs ServiceResponseSender) {})
			So(err, ShouldBeNil)
		})
		Convey("When the context is closed the first time, no errors occur", func() {
			So(context.Close(), ShouldBeNil)
		})
		Convey("When the context is closed the second time, an error occurs", func() {
			So(context.Close(), ShouldNotBeNil)
		})
	})
}
