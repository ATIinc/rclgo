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
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lexample_interfaces__rosidl_typesupport_c -lexample_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/msg/float32.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("example_interfaces/Float32", Float32TypeSupport)
}

// Do not create instances of this type directly. Always use NewFloat32
// function instead.
type Float32 struct {
	Data float32 `yaml:"data"`// This is an example message of using a primitive datatype, float32.If you want to test with this that's fine, but if you are deployingit into a system you should create a semantically meaningful message type.If you want to embed it in another message, use the primitive data type instead.
}

// NewFloat32 creates a new Float32 with default values.
func NewFloat32() *Float32 {
	self := Float32{}
	self.SetDefaults()
	return &self
}

func (t *Float32) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *Float32) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var Float32TypeSupport types.MessageTypeSupport = _Float32TypeSupport{}

type _Float32TypeSupport struct{}

func (t _Float32TypeSupport) New() types.Message {
	return NewFloat32()
}

func (t _Float32TypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__Float32
	return (unsafe.Pointer)(C.example_interfaces__msg__Float32__create())
}

func (t _Float32TypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__Float32__destroy((*C.example_interfaces__msg__Float32)(pointer_to_free))
}

func (t _Float32TypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Float32)
	mem := (*C.example_interfaces__msg__Float32)(dst)
	mem.data = C.float(m.Data)
}

func (t _Float32TypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Float32)
	mem := (*C.example_interfaces__msg__Float32)(ros2_message_buffer)
	m.Data = float32(mem.data)
}

func (t _Float32TypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__Float32())
}

type CFloat32 = C.example_interfaces__msg__Float32
type CFloat32__Sequence = C.example_interfaces__msg__Float32__Sequence

func Float32__Sequence_to_Go(goSlice *[]Float32, cSlice CFloat32__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Float32, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__msg__Float32__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__Float32 * uintptr(i)),
		))
		Float32TypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Float32__Sequence_to_C(cSlice *CFloat32__Sequence, goSlice []Float32) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__Float32)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__msg__Float32 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__msg__Float32)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__Float32 * uintptr(i)),
		))
		Float32TypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Float32__Array_to_Go(goSlice []Float32, cSlice []CFloat32) {
	for i := 0; i < len(cSlice); i++ {
		Float32TypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Float32__Array_to_C(cSlice []CFloat32, goSlice []Float32) {
	for i := 0; i < len(goSlice); i++ {
		Float32TypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
