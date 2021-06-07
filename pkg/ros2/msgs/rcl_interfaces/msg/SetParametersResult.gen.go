/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package rcl_interfaces_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lrcl_interfaces__rosidl_typesupport_c -lrcl_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <rcl_interfaces/msg/set_parameters_result.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("rcl_interfaces/SetParametersResult", SetParametersResultTypeSupport)
}

// Do not create instances of this type directly. Always use NewSetParametersResult
// function instead.
type SetParametersResult struct {
	Successful bool `yaml:"successful"`// A true value of the same index indicates that the parameter was setsuccessfully. A false value indicates the change was rejected.
	Reason string `yaml:"reason"`// Reason why the setting was either successful or a failure. This should only beused for logging and user interfaces.
}

// NewSetParametersResult creates a new SetParametersResult with default values.
func NewSetParametersResult() *SetParametersResult {
	self := SetParametersResult{}
	self.SetDefaults()
	return &self
}

func (t *SetParametersResult) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *SetParametersResult) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var SetParametersResultTypeSupport types.MessageTypeSupport = _SetParametersResultTypeSupport{}

type _SetParametersResultTypeSupport struct{}

func (t _SetParametersResultTypeSupport) New() types.Message {
	return NewSetParametersResult()
}

func (t _SetParametersResultTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.rcl_interfaces__msg__SetParametersResult
	return (unsafe.Pointer)(C.rcl_interfaces__msg__SetParametersResult__create())
}

func (t _SetParametersResultTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rcl_interfaces__msg__SetParametersResult__destroy((*C.rcl_interfaces__msg__SetParametersResult)(pointer_to_free))
}

func (t _SetParametersResultTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SetParametersResult)
	mem := (*C.rcl_interfaces__msg__SetParametersResult)(dst)
	mem.successful = C.bool(m.Successful)
	rosidl_runtime_c.StringAsCStruct(unsafe.Pointer(&mem.reason), m.Reason)
}

func (t _SetParametersResultTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SetParametersResult)
	mem := (*C.rcl_interfaces__msg__SetParametersResult)(ros2_message_buffer)
	m.Successful = bool(mem.successful)
	rosidl_runtime_c.StringAsGoStruct(&m.Reason, unsafe.Pointer(&mem.reason))
}

func (t _SetParametersResultTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rcl_interfaces__msg__SetParametersResult())
}

type CSetParametersResult = C.rcl_interfaces__msg__SetParametersResult
type CSetParametersResult__Sequence = C.rcl_interfaces__msg__SetParametersResult__Sequence

func SetParametersResult__Sequence_to_Go(goSlice *[]SetParametersResult, cSlice CSetParametersResult__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SetParametersResult, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.rcl_interfaces__msg__SetParametersResult__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__msg__SetParametersResult * uintptr(i)),
		))
		SetParametersResultTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func SetParametersResult__Sequence_to_C(cSlice *CSetParametersResult__Sequence, goSlice []SetParametersResult) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.rcl_interfaces__msg__SetParametersResult)(C.malloc((C.size_t)(C.sizeof_struct_rcl_interfaces__msg__SetParametersResult * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.rcl_interfaces__msg__SetParametersResult)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__msg__SetParametersResult * uintptr(i)),
		))
		SetParametersResultTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func SetParametersResult__Array_to_Go(goSlice []SetParametersResult, cSlice []CSetParametersResult) {
	for i := 0; i < len(cSlice); i++ {
		SetParametersResultTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SetParametersResult__Array_to_C(cSlice []CSetParametersResult, goSlice []SetParametersResult) {
	for i := 0; i < len(goSlice); i++ {
		SetParametersResultTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
