/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package test_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ltest_msgs__rosidl_typesupport_c -ltest_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/msg/nested.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("test_msgs/Nested", &Nested{})
}

// Do not create instances of this type directly. Always use NewNested
// function instead.
type Nested struct {
	BasicTypesValue BasicTypes `yaml:"basic_types_value"`
}

// NewNested creates a new Nested with default values.
func NewNested() *Nested {
	self := Nested{}
	self.SetDefaults(nil)
	return &self
}

func (t *Nested) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.BasicTypesValue.SetDefaults(nil)
	
	return t
}

func (t *Nested) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__msg__Nested())
}
func (t *Nested) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__msg__Nested
	return (unsafe.Pointer)(C.test_msgs__msg__Nested__create())
}
func (t *Nested) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__msg__Nested__destroy((*C.test_msgs__msg__Nested)(pointer_to_free))
}
func (t *Nested) AsCStruct() unsafe.Pointer {
	mem := (*C.test_msgs__msg__Nested)(t.PrepareMemory())
	mem.basic_types_value = *(*C.test_msgs__msg__BasicTypes)(t.BasicTypesValue.AsCStruct())
	return unsafe.Pointer(mem)
}
func (t *Nested) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.test_msgs__msg__Nested)(ros2_message_buffer)
	t.BasicTypesValue.AsGoStruct(unsafe.Pointer(&mem.basic_types_value))
}
func (t *Nested) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CNested = C.test_msgs__msg__Nested
type CNested__Sequence = C.test_msgs__msg__Nested__Sequence

func Nested__Sequence_to_Go(goSlice *[]Nested, cSlice CNested__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Nested, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.test_msgs__msg__Nested__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__msg__Nested * uintptr(i)),
		))
		(*goSlice)[i] = Nested{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func Nested__Sequence_to_C(cSlice *CNested__Sequence, goSlice []Nested) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.test_msgs__msg__Nested)(C.malloc((C.size_t)(C.sizeof_struct_test_msgs__msg__Nested * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.test_msgs__msg__Nested)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__msg__Nested * uintptr(i)),
		))
		*cIdx = *(*C.test_msgs__msg__Nested)(v.AsCStruct())
	}
}
func Nested__Array_to_Go(goSlice []Nested, cSlice []CNested) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func Nested__Array_to_C(cSlice []CNested, goSlice []Nested) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.test_msgs__msg__Nested)(goSlice[i].AsCStruct())
	}
}


