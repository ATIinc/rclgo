/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package diagnostic_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	diagnostic_msgs_msg "github.com/tiiuae/rclgo/pkg/ros2/msgs/diagnostic_msgs/msg"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ldiagnostic_msgs__rosidl_typesupport_c -ldiagnostic_msgs__rosidl_generator_c
#cgo LDFLAGS: -ldiagnostic_msgs__rosidl_typesupport_c -ldiagnostic_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <diagnostic_msgs/srv/self_test.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("diagnostic_msgs/SelfTest_Response", &SelfTest_Response{})
}

// Do not create instances of this type directly. Always use NewSelfTest_Response
// function instead.
type SelfTest_Response struct {
	Id rosidl_runtime_c.String `yaml:"id"`
	Passed byte `yaml:"passed"`
	Status []diagnostic_msgs_msg.DiagnosticStatus `yaml:"status"`
}

// NewSelfTest_Response creates a new SelfTest_Response with default values.
func NewSelfTest_Response() *SelfTest_Response {
	self := SelfTest_Response{}
	self.SetDefaults(nil)
	return &self
}

func (t *SelfTest_Response) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Id.SetDefaults("")
	
	return t
}

func (t *SelfTest_Response) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__diagnostic_msgs__srv__SelfTest_Response())
}
func (t *SelfTest_Response) PrepareMemory() unsafe.Pointer { //returns *C.diagnostic_msgs__srv__SelfTest_Response
	return (unsafe.Pointer)(C.diagnostic_msgs__srv__SelfTest_Response__create())
}
func (t *SelfTest_Response) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.diagnostic_msgs__srv__SelfTest_Response__destroy((*C.diagnostic_msgs__srv__SelfTest_Response)(pointer_to_free))
}
func (t *SelfTest_Response) AsCStruct() unsafe.Pointer {
	mem := (*C.diagnostic_msgs__srv__SelfTest_Response)(t.PrepareMemory())
	mem.id = *(*C.rosidl_runtime_c__String)(t.Id.AsCStruct())
	mem.passed = C.uint8_t(t.Passed)
	diagnostic_msgs_msg.DiagnosticStatus__Sequence_to_C((*diagnostic_msgs_msg.CDiagnosticStatus__Sequence)(unsafe.Pointer(&mem.status)), t.Status)
	return unsafe.Pointer(mem)
}
func (t *SelfTest_Response) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.diagnostic_msgs__srv__SelfTest_Response)(ros2_message_buffer)
	t.Id.AsGoStruct(unsafe.Pointer(&mem.id))
	t.Passed = byte(mem.passed)
	diagnostic_msgs_msg.DiagnosticStatus__Sequence_to_Go(&t.Status, *(*diagnostic_msgs_msg.CDiagnosticStatus__Sequence)(unsafe.Pointer(&mem.status)))
}
func (t *SelfTest_Response) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CSelfTest_Response = C.diagnostic_msgs__srv__SelfTest_Response
type CSelfTest_Response__Sequence = C.diagnostic_msgs__srv__SelfTest_Response__Sequence

func SelfTest_Response__Sequence_to_Go(goSlice *[]SelfTest_Response, cSlice CSelfTest_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SelfTest_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.diagnostic_msgs__srv__SelfTest_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_diagnostic_msgs__srv__SelfTest_Response * uintptr(i)),
		))
		(*goSlice)[i] = SelfTest_Response{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func SelfTest_Response__Sequence_to_C(cSlice *CSelfTest_Response__Sequence, goSlice []SelfTest_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.diagnostic_msgs__srv__SelfTest_Response)(C.malloc((C.size_t)(C.sizeof_struct_diagnostic_msgs__srv__SelfTest_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.diagnostic_msgs__srv__SelfTest_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_diagnostic_msgs__srv__SelfTest_Response * uintptr(i)),
		))
		*cIdx = *(*C.diagnostic_msgs__srv__SelfTest_Response)(v.AsCStruct())
	}
}
func SelfTest_Response__Array_to_Go(goSlice []SelfTest_Response, cSlice []CSelfTest_Response) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func SelfTest_Response__Array_to_C(cSlice []CSelfTest_Response, goSlice []SelfTest_Response) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.diagnostic_msgs__srv__SelfTest_Response)(goSlice[i].AsCStruct())
	}
}


