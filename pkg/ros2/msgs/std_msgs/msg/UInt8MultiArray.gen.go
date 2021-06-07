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
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <std_msgs/msg/u_int8_multi_array.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("std_msgs/UInt8MultiArray", UInt8MultiArrayTypeSupport)
}

// Do not create instances of this type directly. Always use NewUInt8MultiArray
// function instead.
type UInt8MultiArray struct {
	Layout MultiArrayLayout `yaml:"layout"`// specification of data layout
	Data []uint8 `yaml:"data"`// array of data
}

// NewUInt8MultiArray creates a new UInt8MultiArray with default values.
func NewUInt8MultiArray() *UInt8MultiArray {
	self := UInt8MultiArray{}
	self.SetDefaults()
	return &self
}

func (t *UInt8MultiArray) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *UInt8MultiArray) SetDefaults() {
	t.Layout.SetDefaults()
	
}

// Modifying this variable is undefined behavior.
var UInt8MultiArrayTypeSupport types.MessageTypeSupport = _UInt8MultiArrayTypeSupport{}

type _UInt8MultiArrayTypeSupport struct{}

func (t _UInt8MultiArrayTypeSupport) New() types.Message {
	return NewUInt8MultiArray()
}

func (t _UInt8MultiArrayTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__UInt8MultiArray
	return (unsafe.Pointer)(C.std_msgs__msg__UInt8MultiArray__create())
}

func (t _UInt8MultiArrayTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__UInt8MultiArray__destroy((*C.std_msgs__msg__UInt8MultiArray)(pointer_to_free))
}

func (t _UInt8MultiArrayTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*UInt8MultiArray)
	mem := (*C.std_msgs__msg__UInt8MultiArray)(dst)
	MultiArrayLayoutTypeSupport.AsCStruct(unsafe.Pointer(&mem.layout), &m.Layout)
	rosidl_runtime_c.Uint8__Sequence_to_C((*rosidl_runtime_c.CUint8__Sequence)(unsafe.Pointer(&mem.data)), m.Data)
}

func (t _UInt8MultiArrayTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*UInt8MultiArray)
	mem := (*C.std_msgs__msg__UInt8MultiArray)(ros2_message_buffer)
	MultiArrayLayoutTypeSupport.AsGoStruct(&m.Layout, unsafe.Pointer(&mem.layout))
	rosidl_runtime_c.Uint8__Sequence_to_Go(&m.Data, *(*rosidl_runtime_c.CUint8__Sequence)(unsafe.Pointer(&mem.data)))
}

func (t _UInt8MultiArrayTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__UInt8MultiArray())
}

type CUInt8MultiArray = C.std_msgs__msg__UInt8MultiArray
type CUInt8MultiArray__Sequence = C.std_msgs__msg__UInt8MultiArray__Sequence

func UInt8MultiArray__Sequence_to_Go(goSlice *[]UInt8MultiArray, cSlice CUInt8MultiArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UInt8MultiArray, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.std_msgs__msg__UInt8MultiArray__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__UInt8MultiArray * uintptr(i)),
		))
		UInt8MultiArrayTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func UInt8MultiArray__Sequence_to_C(cSlice *CUInt8MultiArray__Sequence, goSlice []UInt8MultiArray) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__UInt8MultiArray)(C.malloc((C.size_t)(C.sizeof_struct_std_msgs__msg__UInt8MultiArray * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.std_msgs__msg__UInt8MultiArray)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__UInt8MultiArray * uintptr(i)),
		))
		UInt8MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func UInt8MultiArray__Array_to_Go(goSlice []UInt8MultiArray, cSlice []CUInt8MultiArray) {
	for i := 0; i < len(cSlice); i++ {
		UInt8MultiArrayTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func UInt8MultiArray__Array_to_C(cSlice []CUInt8MultiArray, goSlice []UInt8MultiArray) {
	for i := 0; i < len(goSlice); i++ {
		UInt8MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
