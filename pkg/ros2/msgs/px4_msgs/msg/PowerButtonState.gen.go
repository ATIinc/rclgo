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

#include <px4_msgs/msg/power_button_state.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/PowerButtonState", PowerButtonStateTypeSupport)
}
const (
	PowerButtonState_PWR_BUTTON_STATE_IDEL uint8 = 0// Button went up without meeting shutdown button down time (delete event)
	PowerButtonState_PWR_BUTTON_STATE_DOWN uint8 = 1// Button went Down
	PowerButtonState_PWR_BUTTON_STATE_UP uint8 = 2// Button went Up
	PowerButtonState_PWR_BUTTON_STATE_REQUEST_SHUTDOWN uint8 = 3// Button went Up after meeting shutdown button down time
)

// Do not create instances of this type directly. Always use NewPowerButtonState
// function instead.
type PowerButtonState struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Event uint8 `yaml:"event"`// one of PWR_BUTTON_STATE_*
}

// NewPowerButtonState creates a new PowerButtonState with default values.
func NewPowerButtonState() *PowerButtonState {
	self := PowerButtonState{}
	self.SetDefaults()
	return &self
}

func (t *PowerButtonState) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *PowerButtonState) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var PowerButtonStateTypeSupport types.MessageTypeSupport = _PowerButtonStateTypeSupport{}

type _PowerButtonStateTypeSupport struct{}

func (t _PowerButtonStateTypeSupport) New() types.Message {
	return NewPowerButtonState()
}

func (t _PowerButtonStateTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__PowerButtonState
	return (unsafe.Pointer)(C.px4_msgs__msg__PowerButtonState__create())
}

func (t _PowerButtonStateTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__PowerButtonState__destroy((*C.px4_msgs__msg__PowerButtonState)(pointer_to_free))
}

func (t _PowerButtonStateTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*PowerButtonState)
	mem := (*C.px4_msgs__msg__PowerButtonState)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.event = C.uint8_t(m.Event)
}

func (t _PowerButtonStateTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*PowerButtonState)
	mem := (*C.px4_msgs__msg__PowerButtonState)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.Event = uint8(mem.event)
}

func (t _PowerButtonStateTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__PowerButtonState())
}

type CPowerButtonState = C.px4_msgs__msg__PowerButtonState
type CPowerButtonState__Sequence = C.px4_msgs__msg__PowerButtonState__Sequence

func PowerButtonState__Sequence_to_Go(goSlice *[]PowerButtonState, cSlice CPowerButtonState__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PowerButtonState, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__PowerButtonState__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__PowerButtonState * uintptr(i)),
		))
		PowerButtonStateTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func PowerButtonState__Sequence_to_C(cSlice *CPowerButtonState__Sequence, goSlice []PowerButtonState) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__PowerButtonState)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__PowerButtonState * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__PowerButtonState)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__PowerButtonState * uintptr(i)),
		))
		PowerButtonStateTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func PowerButtonState__Array_to_Go(goSlice []PowerButtonState, cSlice []CPowerButtonState) {
	for i := 0; i < len(cSlice); i++ {
		PowerButtonStateTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func PowerButtonState__Array_to_C(cSlice []CPowerButtonState, goSlice []PowerButtonState) {
	for i := 0; i < len(goSlice); i++ {
		PowerButtonStateTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
