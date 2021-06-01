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
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lexample_interfaces__rosidl_typesupport_c -lexample_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/msg/multi_array_dimension.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("example_interfaces/MultiArrayDimension", &MultiArrayDimension{})
}

// Do not create instances of this type directly. Always use NewMultiArrayDimension
// function instead.
type MultiArrayDimension struct {
	Label rosidl_runtime_c.String `yaml:"label"`// label of given dimension
	Size uint32 `yaml:"size"`// size of given dimension (in type units)
	Stride uint32 `yaml:"stride"`// stride of given dimension
}

// NewMultiArrayDimension creates a new MultiArrayDimension with default values.
func NewMultiArrayDimension() *MultiArrayDimension {
	self := MultiArrayDimension{}
	self.SetDefaults(nil)
	return &self
}

func (t *MultiArrayDimension) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Label.SetDefaults("")
	
	return t
}

func (t *MultiArrayDimension) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__MultiArrayDimension())
}
func (t *MultiArrayDimension) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__MultiArrayDimension
	return (unsafe.Pointer)(C.example_interfaces__msg__MultiArrayDimension__create())
}
func (t *MultiArrayDimension) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__MultiArrayDimension__destroy((*C.example_interfaces__msg__MultiArrayDimension)(pointer_to_free))
}
func (t *MultiArrayDimension) AsCStruct() unsafe.Pointer {
	mem := (*C.example_interfaces__msg__MultiArrayDimension)(t.PrepareMemory())
	mem.label = *(*C.rosidl_runtime_c__String)(t.Label.AsCStruct())
	mem.size = C.uint32_t(t.Size)
	mem.stride = C.uint32_t(t.Stride)
	return unsafe.Pointer(mem)
}
func (t *MultiArrayDimension) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.example_interfaces__msg__MultiArrayDimension)(ros2_message_buffer)
	t.Label.AsGoStruct(unsafe.Pointer(&mem.label))
	t.Size = uint32(mem.size)
	t.Stride = uint32(mem.stride)
}
func (t *MultiArrayDimension) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CMultiArrayDimension = C.example_interfaces__msg__MultiArrayDimension
type CMultiArrayDimension__Sequence = C.example_interfaces__msg__MultiArrayDimension__Sequence

func MultiArrayDimension__Sequence_to_Go(goSlice *[]MultiArrayDimension, cSlice CMultiArrayDimension__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MultiArrayDimension, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__msg__MultiArrayDimension__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__MultiArrayDimension * uintptr(i)),
		))
		(*goSlice)[i] = MultiArrayDimension{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func MultiArrayDimension__Sequence_to_C(cSlice *CMultiArrayDimension__Sequence, goSlice []MultiArrayDimension) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__MultiArrayDimension)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__msg__MultiArrayDimension * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__msg__MultiArrayDimension)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__MultiArrayDimension * uintptr(i)),
		))
		*cIdx = *(*C.example_interfaces__msg__MultiArrayDimension)(v.AsCStruct())
	}
}
func MultiArrayDimension__Array_to_Go(goSlice []MultiArrayDimension, cSlice []CMultiArrayDimension) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func MultiArrayDimension__Array_to_C(cSlice []CMultiArrayDimension, goSlice []MultiArrayDimension) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.example_interfaces__msg__MultiArrayDimension)(goSlice[i].AsCStruct())
	}
}


