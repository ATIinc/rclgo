/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package px4_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/uavcan_parameter_request.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/UavcanParameterRequest", UavcanParameterRequestTypeSupport)
}
const (
	UavcanParameterRequest_MESSAGE_TYPE_PARAM_REQUEST_READ uint8 = 20// MAVLINK_MSG_ID_PARAM_REQUEST_READ
	UavcanParameterRequest_MESSAGE_TYPE_PARAM_REQUEST_LIST uint8 = 21// MAVLINK_MSG_ID_PARAM_REQUEST_LIST
	UavcanParameterRequest_MESSAGE_TYPE_PARAM_SET uint8 = 23// MAVLINK_MSG_ID_PARAM_SET
	UavcanParameterRequest_NODE_ID_ALL uint8 = 0// MAV_COMP_ID_ALL
	UavcanParameterRequest_PARAM_TYPE_UINT8 uint8 = 1// MAV_PARAM_TYPE_UINT8
	UavcanParameterRequest_PARAM_TYPE_INT64 uint8 = 8// MAV_PARAM_TYPE_INT64
	UavcanParameterRequest_PARAM_TYPE_REAL32 uint8 = 9// MAV_PARAM_TYPE_REAL32
	UavcanParameterRequest_ORB_QUEUE_LENGTH uint8 = 4
)

// Do not create instances of this type directly. Always use NewUavcanParameterRequest
// function instead.
type UavcanParameterRequest struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds). UAVCAN-MAVLink parameter bridge request type
	MessageType uint8 `yaml:"message_type"`// MAVLink message type: PARAM_REQUEST_READ, PARAM_REQUEST_LIST, PARAM_SET
	NodeId uint8 `yaml:"node_id"`// UAVCAN node ID mapped from MAVLink component ID
	ParamId [17]byte `yaml:"param_id"`// MAVLink/UAVCAN parameter name
	ParamIndex int16 `yaml:"param_index"`// -1 if the param_id field should be used as identifier
	ParamType uint8 `yaml:"param_type"`// MAVLink parameter type
	IntValue int64 `yaml:"int_value"`// current value if param_type is int-like
	RealValue float32 `yaml:"real_value"`// current value if param_type is float-like
}

// NewUavcanParameterRequest creates a new UavcanParameterRequest with default values.
func NewUavcanParameterRequest() *UavcanParameterRequest {
	self := UavcanParameterRequest{}
	self.SetDefaults()
	return &self
}

func (t *UavcanParameterRequest) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *UavcanParameterRequest) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var UavcanParameterRequestTypeSupport types.MessageTypeSupport = _UavcanParameterRequestTypeSupport{}

type _UavcanParameterRequestTypeSupport struct{}

func (t _UavcanParameterRequestTypeSupport) New() types.Message {
	return NewUavcanParameterRequest()
}

func (t _UavcanParameterRequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__UavcanParameterRequest
	return (unsafe.Pointer)(C.px4_msgs__msg__UavcanParameterRequest__create())
}

func (t _UavcanParameterRequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__UavcanParameterRequest__destroy((*C.px4_msgs__msg__UavcanParameterRequest)(pointer_to_free))
}

func (t _UavcanParameterRequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*UavcanParameterRequest)
	mem := (*C.px4_msgs__msg__UavcanParameterRequest)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.message_type = C.uint8_t(m.MessageType)
	mem.node_id = C.uint8_t(m.NodeId)
	cSlice_param_id := mem.param_id[:]
	rosidl_runtime_c.Char__Array_to_C(*(*[]rosidl_runtime_c.CChar)(unsafe.Pointer(&cSlice_param_id)), m.ParamId[:])
	mem.param_index = C.int16_t(m.ParamIndex)
	mem.param_type = C.uint8_t(m.ParamType)
	mem.int_value = C.int64_t(m.IntValue)
	mem.real_value = C.float(m.RealValue)
}

func (t _UavcanParameterRequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*UavcanParameterRequest)
	mem := (*C.px4_msgs__msg__UavcanParameterRequest)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.MessageType = uint8(mem.message_type)
	m.NodeId = uint8(mem.node_id)
	cSlice_param_id := mem.param_id[:]
	rosidl_runtime_c.Char__Array_to_Go(m.ParamId[:], *(*[]rosidl_runtime_c.CChar)(unsafe.Pointer(&cSlice_param_id)))
	m.ParamIndex = int16(mem.param_index)
	m.ParamType = uint8(mem.param_type)
	m.IntValue = int64(mem.int_value)
	m.RealValue = float32(mem.real_value)
}

func (t _UavcanParameterRequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__UavcanParameterRequest())
}

type CUavcanParameterRequest = C.px4_msgs__msg__UavcanParameterRequest
type CUavcanParameterRequest__Sequence = C.px4_msgs__msg__UavcanParameterRequest__Sequence

func UavcanParameterRequest__Sequence_to_Go(goSlice *[]UavcanParameterRequest, cSlice CUavcanParameterRequest__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UavcanParameterRequest, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__UavcanParameterRequest__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__UavcanParameterRequest * uintptr(i)),
		))
		UavcanParameterRequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func UavcanParameterRequest__Sequence_to_C(cSlice *CUavcanParameterRequest__Sequence, goSlice []UavcanParameterRequest) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__UavcanParameterRequest)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__UavcanParameterRequest * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__UavcanParameterRequest)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__UavcanParameterRequest * uintptr(i)),
		))
		UavcanParameterRequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func UavcanParameterRequest__Array_to_Go(goSlice []UavcanParameterRequest, cSlice []CUavcanParameterRequest) {
	for i := 0; i < len(cSlice); i++ {
		UavcanParameterRequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func UavcanParameterRequest__Array_to_C(cSlice []CUavcanParameterRequest, goSlice []UavcanParameterRequest) {
	for i := 0; i < len(goSlice); i++ {
		UavcanParameterRequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
