/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package fog_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lfog_msgs__rosidl_typesupport_c -lfog_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <fog_msgs/srv/vec4.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("fog_msgs/Vec4_Response", Vec4_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewVec4_Response
// function instead.
type Vec4_Response struct {
	Success bool `yaml:"success"`
	Message string `yaml:"message"`
}

// NewVec4_Response creates a new Vec4_Response with default values.
func NewVec4_Response() *Vec4_Response {
	self := Vec4_Response{}
	self.SetDefaults()
	return &self
}

func (t *Vec4_Response) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *Vec4_Response) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var Vec4_ResponseTypeSupport types.MessageTypeSupport = _Vec4_ResponseTypeSupport{}

type _Vec4_ResponseTypeSupport struct{}

func (t _Vec4_ResponseTypeSupport) New() types.Message {
	return NewVec4_Response()
}

func (t _Vec4_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.fog_msgs__srv__Vec4_Response
	return (unsafe.Pointer)(C.fog_msgs__srv__Vec4_Response__create())
}

func (t _Vec4_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.fog_msgs__srv__Vec4_Response__destroy((*C.fog_msgs__srv__Vec4_Response)(pointer_to_free))
}

func (t _Vec4_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Vec4_Response)
	mem := (*C.fog_msgs__srv__Vec4_Response)(dst)
	mem.success = C.bool(m.Success)
	rosidl_runtime_c.StringAsCStruct(unsafe.Pointer(&mem.message), m.Message)
}

func (t _Vec4_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Vec4_Response)
	mem := (*C.fog_msgs__srv__Vec4_Response)(ros2_message_buffer)
	m.Success = bool(mem.success)
	rosidl_runtime_c.StringAsGoStruct(&m.Message, unsafe.Pointer(&mem.message))
}

func (t _Vec4_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__fog_msgs__srv__Vec4_Response())
}

type CVec4_Response = C.fog_msgs__srv__Vec4_Response
type CVec4_Response__Sequence = C.fog_msgs__srv__Vec4_Response__Sequence

func Vec4_Response__Sequence_to_Go(goSlice *[]Vec4_Response, cSlice CVec4_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Vec4_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.fog_msgs__srv__Vec4_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_fog_msgs__srv__Vec4_Response * uintptr(i)),
		))
		Vec4_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Vec4_Response__Sequence_to_C(cSlice *CVec4_Response__Sequence, goSlice []Vec4_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.fog_msgs__srv__Vec4_Response)(C.malloc((C.size_t)(C.sizeof_struct_fog_msgs__srv__Vec4_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.fog_msgs__srv__Vec4_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_fog_msgs__srv__Vec4_Response * uintptr(i)),
		))
		Vec4_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Vec4_Response__Array_to_Go(goSlice []Vec4_Response, cSlice []CVec4_Response) {
	for i := 0; i < len(cSlice); i++ {
		Vec4_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Vec4_Response__Array_to_C(cSlice []CVec4_Response, goSlice []Vec4_Response) {
	for i := 0; i < len(goSlice); i++ {
		Vec4_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
