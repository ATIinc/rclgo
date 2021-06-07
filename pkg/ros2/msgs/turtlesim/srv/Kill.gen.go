/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package turtlesim_srv

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lturtlesim__rosidl_typesupport_c -lturtlesim__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <turtlesim/srv/kill.h>
*/
import "C"

import (
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	"github.com/tiiuae/rclgo/pkg/ros2/types"

	"unsafe"
)

func init() {
	ros2_type_dispatcher.RegisterROS2ServiceTypeNameAlias("turtlesim/Kill", KillTypeSupport)
}

type _KillTypeSupport struct {}

func (s _KillTypeSupport) Request() types.MessageTypeSupport {
	return Kill_RequestTypeSupport
}

func (s _KillTypeSupport) Response() types.MessageTypeSupport {
	return Kill_ResponseTypeSupport
}

func (s _KillTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__turtlesim__srv__Kill())
}

// Modifying this variable is undefined behavior.
var KillTypeSupport types.ServiceTypeSupport = _KillTypeSupport{}
