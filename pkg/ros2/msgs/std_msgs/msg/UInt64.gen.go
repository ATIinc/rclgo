/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package std_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <std_msgs/msg/u_int64.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("std_msgs/UInt64", &UInt64{})
}

// Do not create instances of this type directly. Always use NewUInt64
// function instead.
type UInt64 struct {
	Data uint64 `yaml:"data"`
}

// NewUInt64 creates a new UInt64 with default values.
func NewUInt64() *UInt64 {
	self := UInt64{}
	self.SetDefaults(nil)
	return &self
}

func (t *UInt64) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *UInt64) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__UInt64())
}
func (t *UInt64) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__UInt64
	return (unsafe.Pointer)(C.std_msgs__msg__UInt64__create())
}
func (t *UInt64) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__UInt64__destroy((*C.std_msgs__msg__UInt64)(pointer_to_free))
}
func (t *UInt64) AsCStruct() unsafe.Pointer {
	mem := (*C.std_msgs__msg__UInt64)(t.PrepareMemory())
	mem.data = C.uint64_t(t.Data)
	return unsafe.Pointer(mem)
}
func (t *UInt64) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.std_msgs__msg__UInt64)(ros2_message_buffer)
	t.Data = uint64(mem.data)
}
func (t *UInt64) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CUInt64 = C.std_msgs__msg__UInt64
type CUInt64__Sequence = C.std_msgs__msg__UInt64__Sequence

func UInt64__Sequence_to_Go(goSlice *[]UInt64, cSlice CUInt64__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UInt64, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.std_msgs__msg__UInt64__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__UInt64 * uintptr(i)),
		))
		(*goSlice)[i] = UInt64{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func UInt64__Sequence_to_C(cSlice *CUInt64__Sequence, goSlice []UInt64) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__UInt64)(C.malloc((C.size_t)(C.sizeof_struct_std_msgs__msg__UInt64 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.std_msgs__msg__UInt64)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__UInt64 * uintptr(i)),
		))
		*cIdx = *(*C.std_msgs__msg__UInt64)(v.AsCStruct())
	}
}
func UInt64__Array_to_Go(goSlice []UInt64, cSlice []CUInt64) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func UInt64__Array_to_C(cSlice []CUInt64, goSlice []UInt64) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.std_msgs__msg__UInt64)(goSlice[i].AsCStruct())
	}
}


