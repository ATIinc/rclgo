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

#include <px4_msgs/msg/rate_ctrl_status.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/RateCtrlStatus", RateCtrlStatusTypeSupport)
}

// Do not create instances of this type directly. Always use NewRateCtrlStatus
// function instead.
type RateCtrlStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	RollspeedInteg float32 `yaml:"rollspeed_integ"`// rate controller integrator status
	PitchspeedInteg float32 `yaml:"pitchspeed_integ"`// rate controller integrator status
	YawspeedInteg float32 `yaml:"yawspeed_integ"`// rate controller integrator status
	AdditionalInteg1 float32 `yaml:"additional_integ1"`// FW: wheel rate integrator (optional). rate controller integrator status
}

// NewRateCtrlStatus creates a new RateCtrlStatus with default values.
func NewRateCtrlStatus() *RateCtrlStatus {
	self := RateCtrlStatus{}
	self.SetDefaults()
	return &self
}

func (t *RateCtrlStatus) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *RateCtrlStatus) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var RateCtrlStatusTypeSupport types.MessageTypeSupport = _RateCtrlStatusTypeSupport{}

type _RateCtrlStatusTypeSupport struct{}

func (t _RateCtrlStatusTypeSupport) New() types.Message {
	return NewRateCtrlStatus()
}

func (t _RateCtrlStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__RateCtrlStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__RateCtrlStatus__create())
}

func (t _RateCtrlStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__RateCtrlStatus__destroy((*C.px4_msgs__msg__RateCtrlStatus)(pointer_to_free))
}

func (t _RateCtrlStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*RateCtrlStatus)
	mem := (*C.px4_msgs__msg__RateCtrlStatus)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.rollspeed_integ = C.float(m.RollspeedInteg)
	mem.pitchspeed_integ = C.float(m.PitchspeedInteg)
	mem.yawspeed_integ = C.float(m.YawspeedInteg)
	mem.additional_integ1 = C.float(m.AdditionalInteg1)
}

func (t _RateCtrlStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*RateCtrlStatus)
	mem := (*C.px4_msgs__msg__RateCtrlStatus)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.RollspeedInteg = float32(mem.rollspeed_integ)
	m.PitchspeedInteg = float32(mem.pitchspeed_integ)
	m.YawspeedInteg = float32(mem.yawspeed_integ)
	m.AdditionalInteg1 = float32(mem.additional_integ1)
}

func (t _RateCtrlStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__RateCtrlStatus())
}

type CRateCtrlStatus = C.px4_msgs__msg__RateCtrlStatus
type CRateCtrlStatus__Sequence = C.px4_msgs__msg__RateCtrlStatus__Sequence

func RateCtrlStatus__Sequence_to_Go(goSlice *[]RateCtrlStatus, cSlice CRateCtrlStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]RateCtrlStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__RateCtrlStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__RateCtrlStatus * uintptr(i)),
		))
		RateCtrlStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func RateCtrlStatus__Sequence_to_C(cSlice *CRateCtrlStatus__Sequence, goSlice []RateCtrlStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__RateCtrlStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__RateCtrlStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__RateCtrlStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__RateCtrlStatus * uintptr(i)),
		))
		RateCtrlStatusTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func RateCtrlStatus__Array_to_Go(goSlice []RateCtrlStatus, cSlice []CRateCtrlStatus) {
	for i := 0; i < len(cSlice); i++ {
		RateCtrlStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func RateCtrlStatus__Array_to_C(cSlice []CRateCtrlStatus, goSlice []RateCtrlStatus) {
	for i := 0; i < len(goSlice); i++ {
		RateCtrlStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
