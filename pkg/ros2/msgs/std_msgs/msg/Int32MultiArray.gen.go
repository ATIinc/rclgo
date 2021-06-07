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

#include <std_msgs/msg/int32_multi_array.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("std_msgs/Int32MultiArray", Int32MultiArrayTypeSupport)
}

// Do not create instances of this type directly. Always use NewInt32MultiArray
// function instead.
type Int32MultiArray struct {
	Layout MultiArrayLayout `yaml:"layout"`// specification of data layout
	Data []int32 `yaml:"data"`// array of data
}

// NewInt32MultiArray creates a new Int32MultiArray with default values.
func NewInt32MultiArray() *Int32MultiArray {
	self := Int32MultiArray{}
	self.SetDefaults()
	return &self
}

func (t *Int32MultiArray) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *Int32MultiArray) SetDefaults() {
	t.Layout.SetDefaults()
	
}

// Modifying this variable is undefined behavior.
var Int32MultiArrayTypeSupport types.MessageTypeSupport = _Int32MultiArrayTypeSupport{}

type _Int32MultiArrayTypeSupport struct{}

func (t _Int32MultiArrayTypeSupport) New() types.Message {
	return NewInt32MultiArray()
}

func (t _Int32MultiArrayTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__Int32MultiArray
	return (unsafe.Pointer)(C.std_msgs__msg__Int32MultiArray__create())
}

func (t _Int32MultiArrayTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__Int32MultiArray__destroy((*C.std_msgs__msg__Int32MultiArray)(pointer_to_free))
}

func (t _Int32MultiArrayTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Int32MultiArray)
	mem := (*C.std_msgs__msg__Int32MultiArray)(dst)
	MultiArrayLayoutTypeSupport.AsCStruct(unsafe.Pointer(&mem.layout), &m.Layout)
	rosidl_runtime_c.Int32__Sequence_to_C((*rosidl_runtime_c.CInt32__Sequence)(unsafe.Pointer(&mem.data)), m.Data)
}

func (t _Int32MultiArrayTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Int32MultiArray)
	mem := (*C.std_msgs__msg__Int32MultiArray)(ros2_message_buffer)
	MultiArrayLayoutTypeSupport.AsGoStruct(&m.Layout, unsafe.Pointer(&mem.layout))
	rosidl_runtime_c.Int32__Sequence_to_Go(&m.Data, *(*rosidl_runtime_c.CInt32__Sequence)(unsafe.Pointer(&mem.data)))
}

func (t _Int32MultiArrayTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__Int32MultiArray())
}

type CInt32MultiArray = C.std_msgs__msg__Int32MultiArray
type CInt32MultiArray__Sequence = C.std_msgs__msg__Int32MultiArray__Sequence

func Int32MultiArray__Sequence_to_Go(goSlice *[]Int32MultiArray, cSlice CInt32MultiArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Int32MultiArray, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.std_msgs__msg__Int32MultiArray__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__Int32MultiArray * uintptr(i)),
		))
		Int32MultiArrayTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Int32MultiArray__Sequence_to_C(cSlice *CInt32MultiArray__Sequence, goSlice []Int32MultiArray) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__Int32MultiArray)(C.malloc((C.size_t)(C.sizeof_struct_std_msgs__msg__Int32MultiArray * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.std_msgs__msg__Int32MultiArray)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__Int32MultiArray * uintptr(i)),
		))
		Int32MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Int32MultiArray__Array_to_Go(goSlice []Int32MultiArray, cSlice []CInt32MultiArray) {
	for i := 0; i < len(cSlice); i++ {
		Int32MultiArrayTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Int32MultiArray__Array_to_C(cSlice []CInt32MultiArray, goSlice []Int32MultiArray) {
	for i := 0; i < len(goSlice); i++ {
		Int32MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
