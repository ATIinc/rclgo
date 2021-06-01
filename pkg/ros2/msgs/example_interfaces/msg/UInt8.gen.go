/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package example_interfaces_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lexample_interfaces__rosidl_typesupport_c -lexample_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/msg/u_int8.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("example_interfaces/UInt8", &UInt8{})
}

// Do not create instances of this type directly. Always use NewUInt8
// function instead.
type UInt8 struct {
	Data uint8 `yaml:"data"`// This is an example message of using a primitive datatype, uint8.If you want to test with this that's fine, but if you are deployingit into a system you should create a semantically meaningful message type.If you want to embed it in another message, use the primitive data type instead.
}

// NewUInt8 creates a new UInt8 with default values.
func NewUInt8() *UInt8 {
	self := UInt8{}
	self.SetDefaults(nil)
	return &self
}

func (t *UInt8) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *UInt8) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__UInt8())
}
func (t *UInt8) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__UInt8
	return (unsafe.Pointer)(C.example_interfaces__msg__UInt8__create())
}
func (t *UInt8) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__UInt8__destroy((*C.example_interfaces__msg__UInt8)(pointer_to_free))
}
func (t *UInt8) AsCStruct() unsafe.Pointer {
	mem := (*C.example_interfaces__msg__UInt8)(t.PrepareMemory())
	mem.data = C.uint8_t(t.Data)
	return unsafe.Pointer(mem)
}
func (t *UInt8) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.example_interfaces__msg__UInt8)(ros2_message_buffer)
	t.Data = uint8(mem.data)
}
func (t *UInt8) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CUInt8 = C.example_interfaces__msg__UInt8
type CUInt8__Sequence = C.example_interfaces__msg__UInt8__Sequence

func UInt8__Sequence_to_Go(goSlice *[]UInt8, cSlice CUInt8__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UInt8, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__msg__UInt8__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__UInt8 * uintptr(i)),
		))
		(*goSlice)[i] = UInt8{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func UInt8__Sequence_to_C(cSlice *CUInt8__Sequence, goSlice []UInt8) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__UInt8)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__msg__UInt8 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__msg__UInt8)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__UInt8 * uintptr(i)),
		))
		*cIdx = *(*C.example_interfaces__msg__UInt8)(v.AsCStruct())
	}
}
func UInt8__Array_to_Go(goSlice []UInt8, cSlice []CUInt8) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func UInt8__Array_to_C(cSlice []CUInt8, goSlice []UInt8) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.example_interfaces__msg__UInt8)(goSlice[i].AsCStruct())
	}
}


