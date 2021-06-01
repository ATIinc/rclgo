/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package diagnostic_msgs_srv

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ldiagnostic_msgs__rosidl_typesupport_c -ldiagnostic_msgs__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <diagnostic_msgs/srv/add_diagnostics.h>
*/
import "C"

import (
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"

	"unsafe"
)

func init() {
	ros2_type_dispatcher.RegisterROS2ServiceTypeNameAlias("diagnostic_msgs/AddDiagnostics", AddDiagnostics)
}

type _AddDiagnostics struct {
	req,resp ros2types.ROS2Msg
}

func (s *_AddDiagnostics) Request() ros2types.ROS2Msg {
	return s.req
}

func (s *_AddDiagnostics) Response() ros2types.ROS2Msg {
	return s.resp
}

func (s *_AddDiagnostics) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__diagnostic_msgs__srv__AddDiagnostics())
}

// Modifying this variable is undefined behavior.
var AddDiagnostics ros2types.Service = &_AddDiagnostics{
	req: &AddDiagnostics_Request{},
	resp: &AddDiagnostics_Response{},
}
