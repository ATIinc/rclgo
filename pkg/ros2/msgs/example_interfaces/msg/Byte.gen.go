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

#include <example_interfaces/msg/byte.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("example_interfaces/Byte", &Byte{})
}

// Do not create instances of this type directly. Always use NewByte
// function instead.
type Byte struct {
	Data byte `yaml:"data"`// This is an example message of using a primitive datatype, byte.If you want to test with this that's fine, but if you are deployingit into a system you should create a semantically meaningful message type.If you want to embed it in another message, use the primitive data type instead.
}

// NewByte creates a new Byte with default values.
func NewByte() *Byte {
	self := Byte{}
	self.SetDefaults(nil)
	return &self
}

func (t *Byte) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *Byte) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__Byte())
}
func (t *Byte) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__Byte
	return (unsafe.Pointer)(C.example_interfaces__msg__Byte__create())
}
func (t *Byte) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__Byte__destroy((*C.example_interfaces__msg__Byte)(pointer_to_free))
}
func (t *Byte) AsCStruct() unsafe.Pointer {
	mem := (*C.example_interfaces__msg__Byte)(t.PrepareMemory())
	mem.data = C.uint8_t(t.Data)
	return unsafe.Pointer(mem)
}
func (t *Byte) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.example_interfaces__msg__Byte)(ros2_message_buffer)
	t.Data = byte(mem.data)
}
func (t *Byte) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CByte = C.example_interfaces__msg__Byte
type CByte__Sequence = C.example_interfaces__msg__Byte__Sequence

func Byte__Sequence_to_Go(goSlice *[]Byte, cSlice CByte__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Byte, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__msg__Byte__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__Byte * uintptr(i)),
		))
		(*goSlice)[i] = Byte{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func Byte__Sequence_to_C(cSlice *CByte__Sequence, goSlice []Byte) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__Byte)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__msg__Byte * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__msg__Byte)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__Byte * uintptr(i)),
		))
		*cIdx = *(*C.example_interfaces__msg__Byte)(v.AsCStruct())
	}
}
func Byte__Array_to_Go(goSlice []Byte, cSlice []CByte) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func Byte__Array_to_C(cSlice []CByte, goSlice []Byte) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.example_interfaces__msg__Byte)(goSlice[i].AsCStruct())
	}
}


