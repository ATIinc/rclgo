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

#include <example_interfaces/msg/u_int32_multi_array.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("example_interfaces/UInt32MultiArray", &UInt32MultiArray{})
}

// Do not create instances of this type directly. Always use NewUInt32MultiArray
// function instead.
type UInt32MultiArray struct {
	Layout MultiArrayLayout `yaml:"layout"`// specification of data layout
	Data []uint32 `yaml:"data"`// array of data
}

// NewUInt32MultiArray creates a new UInt32MultiArray with default values.
func NewUInt32MultiArray() *UInt32MultiArray {
	self := UInt32MultiArray{}
	self.SetDefaults(nil)
	return &self
}

func (t *UInt32MultiArray) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Layout.SetDefaults(nil)
	
	return t
}

func (t *UInt32MultiArray) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__UInt32MultiArray())
}
func (t *UInt32MultiArray) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__UInt32MultiArray
	return (unsafe.Pointer)(C.example_interfaces__msg__UInt32MultiArray__create())
}
func (t *UInt32MultiArray) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__UInt32MultiArray__destroy((*C.example_interfaces__msg__UInt32MultiArray)(pointer_to_free))
}
func (t *UInt32MultiArray) AsCStruct() unsafe.Pointer {
	mem := (*C.example_interfaces__msg__UInt32MultiArray)(t.PrepareMemory())
	mem.layout = *(*C.example_interfaces__msg__MultiArrayLayout)(t.Layout.AsCStruct())
	rosidl_runtime_c.Uint32__Sequence_to_C((*rosidl_runtime_c.CUint32__Sequence)(unsafe.Pointer(&mem.data)), t.Data)
	return unsafe.Pointer(mem)
}
func (t *UInt32MultiArray) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.example_interfaces__msg__UInt32MultiArray)(ros2_message_buffer)
	t.Layout.AsGoStruct(unsafe.Pointer(&mem.layout))
	rosidl_runtime_c.Uint32__Sequence_to_Go(&t.Data, *(*rosidl_runtime_c.CUint32__Sequence)(unsafe.Pointer(&mem.data)))
}
func (t *UInt32MultiArray) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CUInt32MultiArray = C.example_interfaces__msg__UInt32MultiArray
type CUInt32MultiArray__Sequence = C.example_interfaces__msg__UInt32MultiArray__Sequence

func UInt32MultiArray__Sequence_to_Go(goSlice *[]UInt32MultiArray, cSlice CUInt32MultiArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UInt32MultiArray, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__msg__UInt32MultiArray__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__UInt32MultiArray * uintptr(i)),
		))
		(*goSlice)[i] = UInt32MultiArray{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func UInt32MultiArray__Sequence_to_C(cSlice *CUInt32MultiArray__Sequence, goSlice []UInt32MultiArray) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__UInt32MultiArray)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__msg__UInt32MultiArray * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__msg__UInt32MultiArray)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__UInt32MultiArray * uintptr(i)),
		))
		*cIdx = *(*C.example_interfaces__msg__UInt32MultiArray)(v.AsCStruct())
	}
}
func UInt32MultiArray__Array_to_Go(goSlice []UInt32MultiArray, cSlice []CUInt32MultiArray) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func UInt32MultiArray__Array_to_C(cSlice []CUInt32MultiArray, goSlice []UInt32MultiArray) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.example_interfaces__msg__UInt32MultiArray)(goSlice[i].AsCStruct())
	}
}


