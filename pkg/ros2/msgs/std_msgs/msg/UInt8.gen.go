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

	"github.com/tiiuae/rclgo/pkg/ros2/types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <std_msgs/msg/u_int8.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("std_msgs/UInt8", UInt8TypeSupport)
}

// Do not create instances of this type directly. Always use NewUInt8
// function instead.
type UInt8 struct {
	Data uint8 `yaml:"data"`
}

// NewUInt8 creates a new UInt8 with default values.
func NewUInt8() *UInt8 {
	self := UInt8{}
	self.SetDefaults()
	return &self
}

func (t *UInt8) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *UInt8) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var UInt8TypeSupport types.MessageTypeSupport = _UInt8TypeSupport{}

type _UInt8TypeSupport struct{}

func (t _UInt8TypeSupport) New() types.Message {
	return NewUInt8()
}

func (t _UInt8TypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__UInt8
	return (unsafe.Pointer)(C.std_msgs__msg__UInt8__create())
}

func (t _UInt8TypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__UInt8__destroy((*C.std_msgs__msg__UInt8)(pointer_to_free))
}

func (t _UInt8TypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*UInt8)
	mem := (*C.std_msgs__msg__UInt8)(dst)
	mem.data = C.uint8_t(m.Data)
}

func (t _UInt8TypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*UInt8)
	mem := (*C.std_msgs__msg__UInt8)(ros2_message_buffer)
	m.Data = uint8(mem.data)
}

func (t _UInt8TypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__UInt8())
}

type CUInt8 = C.std_msgs__msg__UInt8
type CUInt8__Sequence = C.std_msgs__msg__UInt8__Sequence

func UInt8__Sequence_to_Go(goSlice *[]UInt8, cSlice CUInt8__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UInt8, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.std_msgs__msg__UInt8__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__UInt8 * uintptr(i)),
		))
		UInt8TypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func UInt8__Sequence_to_C(cSlice *CUInt8__Sequence, goSlice []UInt8) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__UInt8)(C.malloc((C.size_t)(C.sizeof_struct_std_msgs__msg__UInt8 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.std_msgs__msg__UInt8)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__UInt8 * uintptr(i)),
		))
		UInt8TypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func UInt8__Array_to_Go(goSlice []UInt8, cSlice []CUInt8) {
	for i := 0; i < len(cSlice); i++ {
		UInt8TypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func UInt8__Array_to_C(cSlice []CUInt8, goSlice []UInt8) {
	for i := 0; i < len(goSlice); i++ {
		UInt8TypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
