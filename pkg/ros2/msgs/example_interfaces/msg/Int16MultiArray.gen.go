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

	"github.com/tiiuae/rclgo/pkg/ros2/types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lexample_interfaces__rosidl_typesupport_c -lexample_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/msg/int16_multi_array.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("example_interfaces/Int16MultiArray", Int16MultiArrayTypeSupport)
}

// Do not create instances of this type directly. Always use NewInt16MultiArray
// function instead.
type Int16MultiArray struct {
	Layout MultiArrayLayout `yaml:"layout"`// specification of data layout
	Data []int16 `yaml:"data"`// array of data
}

// NewInt16MultiArray creates a new Int16MultiArray with default values.
func NewInt16MultiArray() *Int16MultiArray {
	self := Int16MultiArray{}
	self.SetDefaults()
	return &self
}

func (t *Int16MultiArray) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *Int16MultiArray) SetDefaults() {
	t.Layout.SetDefaults()
	
}

// Modifying this variable is undefined behavior.
var Int16MultiArrayTypeSupport types.MessageTypeSupport = _Int16MultiArrayTypeSupport{}

type _Int16MultiArrayTypeSupport struct{}

func (t _Int16MultiArrayTypeSupport) New() types.Message {
	return NewInt16MultiArray()
}

func (t _Int16MultiArrayTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__Int16MultiArray
	return (unsafe.Pointer)(C.example_interfaces__msg__Int16MultiArray__create())
}

func (t _Int16MultiArrayTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__Int16MultiArray__destroy((*C.example_interfaces__msg__Int16MultiArray)(pointer_to_free))
}

func (t _Int16MultiArrayTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Int16MultiArray)
	mem := (*C.example_interfaces__msg__Int16MultiArray)(dst)
	MultiArrayLayoutTypeSupport.AsCStruct(unsafe.Pointer(&mem.layout), &m.Layout)
	rosidl_runtime_c.Int16__Sequence_to_C((*rosidl_runtime_c.CInt16__Sequence)(unsafe.Pointer(&mem.data)), m.Data)
}

func (t _Int16MultiArrayTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Int16MultiArray)
	mem := (*C.example_interfaces__msg__Int16MultiArray)(ros2_message_buffer)
	MultiArrayLayoutTypeSupport.AsGoStruct(&m.Layout, unsafe.Pointer(&mem.layout))
	rosidl_runtime_c.Int16__Sequence_to_Go(&m.Data, *(*rosidl_runtime_c.CInt16__Sequence)(unsafe.Pointer(&mem.data)))
}

func (t _Int16MultiArrayTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__Int16MultiArray())
}

type CInt16MultiArray = C.example_interfaces__msg__Int16MultiArray
type CInt16MultiArray__Sequence = C.example_interfaces__msg__Int16MultiArray__Sequence

func Int16MultiArray__Sequence_to_Go(goSlice *[]Int16MultiArray, cSlice CInt16MultiArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Int16MultiArray, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__msg__Int16MultiArray__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__Int16MultiArray * uintptr(i)),
		))
		Int16MultiArrayTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Int16MultiArray__Sequence_to_C(cSlice *CInt16MultiArray__Sequence, goSlice []Int16MultiArray) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__Int16MultiArray)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__msg__Int16MultiArray * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__msg__Int16MultiArray)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__Int16MultiArray * uintptr(i)),
		))
		Int16MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Int16MultiArray__Array_to_Go(goSlice []Int16MultiArray, cSlice []CInt16MultiArray) {
	for i := 0; i < len(cSlice); i++ {
		Int16MultiArrayTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Int16MultiArray__Array_to_C(cSlice []CInt16MultiArray, goSlice []Int16MultiArray) {
	for i := 0; i < len(goSlice); i++ {
		Int16MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
