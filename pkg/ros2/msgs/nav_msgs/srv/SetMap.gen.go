/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package nav_msgs_srv

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lnav_msgs__rosidl_typesupport_c -lnav_msgs__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <nav_msgs/srv/set_map.h>
*/
import "C"

import (
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	"github.com/tiiuae/rclgo/pkg/ros2/types"

	"unsafe"
)

func init() {
	ros2_type_dispatcher.RegisterROS2ServiceTypeNameAlias("nav_msgs/SetMap", SetMapTypeSupport)
}

type _SetMapTypeSupport struct {}

func (s _SetMapTypeSupport) Request() types.MessageTypeSupport {
	return SetMap_RequestTypeSupport
}

func (s _SetMapTypeSupport) Response() types.MessageTypeSupport {
	return SetMap_ResponseTypeSupport
}

func (s _SetMapTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__nav_msgs__srv__SetMap())
}

// Modifying this variable is undefined behavior.
var SetMapTypeSupport types.ServiceTypeSupport = _SetMapTypeSupport{}
