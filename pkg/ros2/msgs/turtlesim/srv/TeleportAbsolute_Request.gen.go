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
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lturtlesim__rosidl_typesupport_c -lturtlesim__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <turtlesim/srv/teleport_absolute.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("turtlesim/TeleportAbsolute_Request", &TeleportAbsolute_Request{})
}

// Do not create instances of this type directly. Always use NewTeleportAbsolute_Request
// function instead.
type TeleportAbsolute_Request struct {
	X float32 `yaml:"x"`
	Y float32 `yaml:"y"`
	Theta float32 `yaml:"theta"`
}

// NewTeleportAbsolute_Request creates a new TeleportAbsolute_Request with default values.
func NewTeleportAbsolute_Request() *TeleportAbsolute_Request {
	self := TeleportAbsolute_Request{}
	self.SetDefaults(nil)
	return &self
}

func (t *TeleportAbsolute_Request) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *TeleportAbsolute_Request) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__turtlesim__srv__TeleportAbsolute_Request())
}
func (t *TeleportAbsolute_Request) PrepareMemory() unsafe.Pointer { //returns *C.turtlesim__srv__TeleportAbsolute_Request
	return (unsafe.Pointer)(C.turtlesim__srv__TeleportAbsolute_Request__create())
}
func (t *TeleportAbsolute_Request) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.turtlesim__srv__TeleportAbsolute_Request__destroy((*C.turtlesim__srv__TeleportAbsolute_Request)(pointer_to_free))
}
func (t *TeleportAbsolute_Request) AsCStruct() unsafe.Pointer {
	mem := (*C.turtlesim__srv__TeleportAbsolute_Request)(t.PrepareMemory())
	mem.x = C.float(t.X)
	mem.y = C.float(t.Y)
	mem.theta = C.float(t.Theta)
	return unsafe.Pointer(mem)
}
func (t *TeleportAbsolute_Request) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.turtlesim__srv__TeleportAbsolute_Request)(ros2_message_buffer)
	t.X = float32(mem.x)
	t.Y = float32(mem.y)
	t.Theta = float32(mem.theta)
}
func (t *TeleportAbsolute_Request) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CTeleportAbsolute_Request = C.turtlesim__srv__TeleportAbsolute_Request
type CTeleportAbsolute_Request__Sequence = C.turtlesim__srv__TeleportAbsolute_Request__Sequence

func TeleportAbsolute_Request__Sequence_to_Go(goSlice *[]TeleportAbsolute_Request, cSlice CTeleportAbsolute_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TeleportAbsolute_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.turtlesim__srv__TeleportAbsolute_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_turtlesim__srv__TeleportAbsolute_Request * uintptr(i)),
		))
		(*goSlice)[i] = TeleportAbsolute_Request{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func TeleportAbsolute_Request__Sequence_to_C(cSlice *CTeleportAbsolute_Request__Sequence, goSlice []TeleportAbsolute_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.turtlesim__srv__TeleportAbsolute_Request)(C.malloc((C.size_t)(C.sizeof_struct_turtlesim__srv__TeleportAbsolute_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.turtlesim__srv__TeleportAbsolute_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_turtlesim__srv__TeleportAbsolute_Request * uintptr(i)),
		))
		*cIdx = *(*C.turtlesim__srv__TeleportAbsolute_Request)(v.AsCStruct())
	}
}
func TeleportAbsolute_Request__Array_to_Go(goSlice []TeleportAbsolute_Request, cSlice []CTeleportAbsolute_Request) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func TeleportAbsolute_Request__Array_to_C(cSlice []CTeleportAbsolute_Request, goSlice []TeleportAbsolute_Request) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.turtlesim__srv__TeleportAbsolute_Request)(goSlice[i].AsCStruct())
	}
}


