/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package test_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ltest_msgs__rosidl_typesupport_c -ltest_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/srv/basic_types.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("test_msgs/BasicTypes_Response", &BasicTypes_Response{})
}

// Do not create instances of this type directly. Always use NewBasicTypes_Response
// function instead.
type BasicTypes_Response struct {
	BoolValue bool `yaml:"bool_value"`
	ByteValue byte `yaml:"byte_value"`
	CharValue byte `yaml:"char_value"`
	Float32Value float32 `yaml:"float32_value"`
	Float64Value float64 `yaml:"float64_value"`
	Int8Value int8 `yaml:"int8_value"`
	Uint8Value uint8 `yaml:"uint8_value"`
	Int16Value int16 `yaml:"int16_value"`
	Uint16Value uint16 `yaml:"uint16_value"`
	Int32Value int32 `yaml:"int32_value"`
	Uint32Value uint32 `yaml:"uint32_value"`
	Int64Value int64 `yaml:"int64_value"`
	Uint64Value uint64 `yaml:"uint64_value"`
	StringValue rosidl_runtime_c.String `yaml:"string_value"`
}

// NewBasicTypes_Response creates a new BasicTypes_Response with default values.
func NewBasicTypes_Response() *BasicTypes_Response {
	self := BasicTypes_Response{}
	self.SetDefaults(nil)
	return &self
}

func (t *BasicTypes_Response) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.StringValue.SetDefaults("")
	
	return t
}

func (t *BasicTypes_Response) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__srv__BasicTypes_Response())
}
func (t *BasicTypes_Response) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__srv__BasicTypes_Response
	return (unsafe.Pointer)(C.test_msgs__srv__BasicTypes_Response__create())
}
func (t *BasicTypes_Response) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__srv__BasicTypes_Response__destroy((*C.test_msgs__srv__BasicTypes_Response)(pointer_to_free))
}
func (t *BasicTypes_Response) AsCStruct() unsafe.Pointer {
	mem := (*C.test_msgs__srv__BasicTypes_Response)(t.PrepareMemory())
	mem.bool_value = C.bool(t.BoolValue)
	mem.byte_value = C.uint8_t(t.ByteValue)
	mem.char_value = C.uchar(t.CharValue)
	mem.float32_value = C.float(t.Float32Value)
	mem.float64_value = C.double(t.Float64Value)
	mem.int8_value = C.int8_t(t.Int8Value)
	mem.uint8_value = C.uint8_t(t.Uint8Value)
	mem.int16_value = C.int16_t(t.Int16Value)
	mem.uint16_value = C.uint16_t(t.Uint16Value)
	mem.int32_value = C.int32_t(t.Int32Value)
	mem.uint32_value = C.uint32_t(t.Uint32Value)
	mem.int64_value = C.int64_t(t.Int64Value)
	mem.uint64_value = C.uint64_t(t.Uint64Value)
	mem.string_value = *(*C.rosidl_runtime_c__String)(t.StringValue.AsCStruct())
	return unsafe.Pointer(mem)
}
func (t *BasicTypes_Response) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.test_msgs__srv__BasicTypes_Response)(ros2_message_buffer)
	t.BoolValue = bool(mem.bool_value)
	t.ByteValue = byte(mem.byte_value)
	t.CharValue = byte(mem.char_value)
	t.Float32Value = float32(mem.float32_value)
	t.Float64Value = float64(mem.float64_value)
	t.Int8Value = int8(mem.int8_value)
	t.Uint8Value = uint8(mem.uint8_value)
	t.Int16Value = int16(mem.int16_value)
	t.Uint16Value = uint16(mem.uint16_value)
	t.Int32Value = int32(mem.int32_value)
	t.Uint32Value = uint32(mem.uint32_value)
	t.Int64Value = int64(mem.int64_value)
	t.Uint64Value = uint64(mem.uint64_value)
	t.StringValue.AsGoStruct(unsafe.Pointer(&mem.string_value))
}
func (t *BasicTypes_Response) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CBasicTypes_Response = C.test_msgs__srv__BasicTypes_Response
type CBasicTypes_Response__Sequence = C.test_msgs__srv__BasicTypes_Response__Sequence

func BasicTypes_Response__Sequence_to_Go(goSlice *[]BasicTypes_Response, cSlice CBasicTypes_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]BasicTypes_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.test_msgs__srv__BasicTypes_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__srv__BasicTypes_Response * uintptr(i)),
		))
		(*goSlice)[i] = BasicTypes_Response{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func BasicTypes_Response__Sequence_to_C(cSlice *CBasicTypes_Response__Sequence, goSlice []BasicTypes_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.test_msgs__srv__BasicTypes_Response)(C.malloc((C.size_t)(C.sizeof_struct_test_msgs__srv__BasicTypes_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.test_msgs__srv__BasicTypes_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__srv__BasicTypes_Response * uintptr(i)),
		))
		*cIdx = *(*C.test_msgs__srv__BasicTypes_Response)(v.AsCStruct())
	}
}
func BasicTypes_Response__Array_to_Go(goSlice []BasicTypes_Response, cSlice []CBasicTypes_Response) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func BasicTypes_Response__Array_to_C(cSlice []CBasicTypes_Response, goSlice []BasicTypes_Response) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.test_msgs__srv__BasicTypes_Response)(goSlice[i].AsCStruct())
	}
}


