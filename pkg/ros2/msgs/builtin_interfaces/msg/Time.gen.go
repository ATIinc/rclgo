/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package builtin_interfaces_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lbuiltin_interfaces__rosidl_typesupport_c -lbuiltin_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <builtin_interfaces/msg/time.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("builtin_interfaces/Time", TimeTypeSupport)
}

// Do not create instances of this type directly. Always use NewTime
// function instead.
type Time struct {
	Sec int32 `yaml:"sec"`// The seconds component, valid over all int32 values.
	Nanosec uint32 `yaml:"nanosec"`// The nanoseconds component, valid in the range [0, 10e9).
}

// NewTime creates a new Time with default values.
func NewTime() *Time {
	self := Time{}
	self.SetDefaults()
	return &self
}

func (t *Time) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *Time) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var TimeTypeSupport types.MessageTypeSupport = _TimeTypeSupport{}

type _TimeTypeSupport struct{}

func (t _TimeTypeSupport) New() types.Message {
	return NewTime()
}

func (t _TimeTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.builtin_interfaces__msg__Time
	return (unsafe.Pointer)(C.builtin_interfaces__msg__Time__create())
}

func (t _TimeTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.builtin_interfaces__msg__Time__destroy((*C.builtin_interfaces__msg__Time)(pointer_to_free))
}

func (t _TimeTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Time)
	mem := (*C.builtin_interfaces__msg__Time)(dst)
	mem.sec = C.int32_t(m.Sec)
	mem.nanosec = C.uint32_t(m.Nanosec)
}

func (t _TimeTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Time)
	mem := (*C.builtin_interfaces__msg__Time)(ros2_message_buffer)
	m.Sec = int32(mem.sec)
	m.Nanosec = uint32(mem.nanosec)
}

func (t _TimeTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__builtin_interfaces__msg__Time())
}

type CTime = C.builtin_interfaces__msg__Time
type CTime__Sequence = C.builtin_interfaces__msg__Time__Sequence

func Time__Sequence_to_Go(goSlice *[]Time, cSlice CTime__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Time, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.builtin_interfaces__msg__Time__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_builtin_interfaces__msg__Time * uintptr(i)),
		))
		TimeTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Time__Sequence_to_C(cSlice *CTime__Sequence, goSlice []Time) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.builtin_interfaces__msg__Time)(C.malloc((C.size_t)(C.sizeof_struct_builtin_interfaces__msg__Time * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.builtin_interfaces__msg__Time)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_builtin_interfaces__msg__Time * uintptr(i)),
		))
		TimeTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Time__Array_to_Go(goSlice []Time, cSlice []CTime) {
	for i := 0; i < len(cSlice); i++ {
		TimeTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Time__Array_to_C(cSlice []CTime, goSlice []Time) {
	for i := 0; i < len(goSlice); i++ {
		TimeTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
