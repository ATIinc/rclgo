/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package pcl_msgs_srv

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpcl_msgs__rosidl_typesupport_c -lpcl_msgs__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <pcl_msgs/srv/update_filename.h>
*/
import "C"

import (
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"

	"unsafe"
)

func init() {
	ros2_type_dispatcher.RegisterROS2ServiceTypeNameAlias("pcl_msgs/UpdateFilename", UpdateFilename)
}

type _UpdateFilename struct {
	req,resp ros2types.ROS2Msg
}

func (s *_UpdateFilename) Request() ros2types.ROS2Msg {
	return s.req
}

func (s *_UpdateFilename) Response() ros2types.ROS2Msg {
	return s.resp
}

func (s *_UpdateFilename) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__pcl_msgs__srv__UpdateFilename())
}

// Modifying this variable is undefined behavior.
var UpdateFilename ros2types.Service = &_UpdateFilename{
	req: &UpdateFilename_Request{},
	resp: &UpdateFilename_Response{},
}
