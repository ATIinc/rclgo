/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package lifecycle_msgs_srv

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -llifecycle_msgs__rosidl_typesupport_c -llifecycle_msgs__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <lifecycle_msgs/srv/get_available_transitions.h>
*/
import "C"

import (
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	"github.com/tiiuae/rclgo/pkg/ros2/types"

	"unsafe"
)

func init() {
	ros2_type_dispatcher.RegisterROS2ServiceTypeNameAlias("lifecycle_msgs/GetAvailableTransitions", GetAvailableTransitionsTypeSupport)
}

type _GetAvailableTransitionsTypeSupport struct {}

func (s _GetAvailableTransitionsTypeSupport) Request() types.MessageTypeSupport {
	return GetAvailableTransitions_RequestTypeSupport
}

func (s _GetAvailableTransitionsTypeSupport) Response() types.MessageTypeSupport {
	return GetAvailableTransitions_ResponseTypeSupport
}

func (s _GetAvailableTransitionsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__lifecycle_msgs__srv__GetAvailableTransitions())
}

// Modifying this variable is undefined behavior.
var GetAvailableTransitionsTypeSupport types.ServiceTypeSupport = _GetAvailableTransitionsTypeSupport{}
