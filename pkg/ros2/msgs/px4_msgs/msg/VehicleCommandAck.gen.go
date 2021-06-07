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
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/vehicle_command_ack.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/VehicleCommandAck", VehicleCommandAckTypeSupport)
}
const (
	VehicleCommandAck_VEHICLE_RESULT_ACCEPTED uint8 = 0
	VehicleCommandAck_VEHICLE_RESULT_TEMPORARILY_REJECTED uint8 = 1
	VehicleCommandAck_VEHICLE_RESULT_DENIED uint8 = 2
	VehicleCommandAck_VEHICLE_RESULT_UNSUPPORTED uint8 = 3
	VehicleCommandAck_VEHICLE_RESULT_FAILED uint8 = 4
	VehicleCommandAck_VEHICLE_RESULT_IN_PROGRESS uint8 = 5
	VehicleCommandAck_ARM_AUTH_DENIED_REASON_GENERIC uint16 = 0
	VehicleCommandAck_ARM_AUTH_DENIED_REASON_NONE uint16 = 1
	VehicleCommandAck_ARM_AUTH_DENIED_REASON_INVALID_WAYPOINT uint16 = 2
	VehicleCommandAck_ARM_AUTH_DENIED_REASON_TIMEOUT uint16 = 3
	VehicleCommandAck_ARM_AUTH_DENIED_REASON_AIRSPACE_IN_USE uint16 = 4
	VehicleCommandAck_ARM_AUTH_DENIED_REASON_BAD_WEATHER uint16 = 5
	VehicleCommandAck_ORB_QUEUE_LENGTH uint8 = 4
)

// Do not create instances of this type directly. Always use NewVehicleCommandAck
// function instead.
type VehicleCommandAck struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Command uint32 `yaml:"command"`
	Result uint8 `yaml:"result"`
	FromExternal bool `yaml:"from_external"`
	ResultParam1 uint8 `yaml:"result_param1"`
	ResultParam2 int32 `yaml:"result_param2"`
	TargetSystem uint8 `yaml:"target_system"`
	TargetComponent uint8 `yaml:"target_component"`
}

// NewVehicleCommandAck creates a new VehicleCommandAck with default values.
func NewVehicleCommandAck() *VehicleCommandAck {
	self := VehicleCommandAck{}
	self.SetDefaults()
	return &self
}

func (t *VehicleCommandAck) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *VehicleCommandAck) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var VehicleCommandAckTypeSupport types.MessageTypeSupport = _VehicleCommandAckTypeSupport{}

type _VehicleCommandAckTypeSupport struct{}

func (t _VehicleCommandAckTypeSupport) New() types.Message {
	return NewVehicleCommandAck()
}

func (t _VehicleCommandAckTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleCommandAck
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleCommandAck__create())
}

func (t _VehicleCommandAckTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleCommandAck__destroy((*C.px4_msgs__msg__VehicleCommandAck)(pointer_to_free))
}

func (t _VehicleCommandAckTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*VehicleCommandAck)
	mem := (*C.px4_msgs__msg__VehicleCommandAck)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.command = C.uint32_t(m.Command)
	mem.result = C.uint8_t(m.Result)
	mem.from_external = C.bool(m.FromExternal)
	mem.result_param1 = C.uint8_t(m.ResultParam1)
	mem.result_param2 = C.int32_t(m.ResultParam2)
	mem.target_system = C.uint8_t(m.TargetSystem)
	mem.target_component = C.uint8_t(m.TargetComponent)
}

func (t _VehicleCommandAckTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*VehicleCommandAck)
	mem := (*C.px4_msgs__msg__VehicleCommandAck)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.Command = uint32(mem.command)
	m.Result = uint8(mem.result)
	m.FromExternal = bool(mem.from_external)
	m.ResultParam1 = uint8(mem.result_param1)
	m.ResultParam2 = int32(mem.result_param2)
	m.TargetSystem = uint8(mem.target_system)
	m.TargetComponent = uint8(mem.target_component)
}

func (t _VehicleCommandAckTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleCommandAck())
}

type CVehicleCommandAck = C.px4_msgs__msg__VehicleCommandAck
type CVehicleCommandAck__Sequence = C.px4_msgs__msg__VehicleCommandAck__Sequence

func VehicleCommandAck__Sequence_to_Go(goSlice *[]VehicleCommandAck, cSlice CVehicleCommandAck__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleCommandAck, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleCommandAck__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleCommandAck * uintptr(i)),
		))
		VehicleCommandAckTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func VehicleCommandAck__Sequence_to_C(cSlice *CVehicleCommandAck__Sequence, goSlice []VehicleCommandAck) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleCommandAck)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleCommandAck * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleCommandAck)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleCommandAck * uintptr(i)),
		))
		VehicleCommandAckTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func VehicleCommandAck__Array_to_Go(goSlice []VehicleCommandAck, cSlice []CVehicleCommandAck) {
	for i := 0; i < len(cSlice); i++ {
		VehicleCommandAckTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleCommandAck__Array_to_C(cSlice []CVehicleCommandAck, goSlice []VehicleCommandAck) {
	for i := 0; i < len(goSlice); i++ {
		VehicleCommandAckTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
