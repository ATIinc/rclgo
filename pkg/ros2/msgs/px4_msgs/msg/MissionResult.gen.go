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

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/mission_result.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/MissionResult", &MissionResult{})
}
const (
	MissionResult_MISSION_EXECUTION_MODE_NORMAL uint8 = 0// Execute the mission according to the planned items
	MissionResult_MISSION_EXECUTION_MODE_REVERSE uint8 = 1// Execute the mission in reverse order, ignoring commands and converting all waypoints to normal ones
	MissionResult_MISSION_EXECUTION_MODE_FAST_FORWARD uint8 = 2// Execute the mission as fast as possible, for example converting loiter waypoints to normal ones
)

// Do not create instances of this type directly. Always use NewMissionResult
// function instead.
type MissionResult struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	InstanceCount uint32 `yaml:"instance_count"`// Instance count of this mission. Increments monotonically whenever the mission is modified
	SeqReached int32 `yaml:"seq_reached"`// Sequence of the mission item which has been reached, default -1
	SeqCurrent uint16 `yaml:"seq_current"`// Sequence of the current mission item
	SeqTotal uint16 `yaml:"seq_total"`// Total number of mission items
	Valid bool `yaml:"valid"`// true if mission is valid
	Warning bool `yaml:"warning"`// true if mission is valid, but has potentially problematic items leading to safety warnings
	Finished bool `yaml:"finished"`// true if mission has been completed
	Failure bool `yaml:"failure"`// true if the mission cannot continue or be completed for some reason
	StayInFailsafe bool `yaml:"stay_in_failsafe"`// true if the commander should not switch out of the failsafe mode
	FlightTermination bool `yaml:"flight_termination"`// true if the navigator demands a flight termination from the commander app
	ItemDoJumpChanged bool `yaml:"item_do_jump_changed"`// true if the number of do jumps remaining has changed
	ItemChangedIndex uint16 `yaml:"item_changed_index"`// indicate which item has changed
	ItemDoJumpRemaining uint16 `yaml:"item_do_jump_remaining"`// set to the number of do jumps remaining for that item
	ExecutionMode uint8 `yaml:"execution_mode"`// indicates the mode in which the mission is executed
}

// NewMissionResult creates a new MissionResult with default values.
func NewMissionResult() *MissionResult {
	self := MissionResult{}
	self.SetDefaults(nil)
	return &self
}

func (t *MissionResult) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *MissionResult) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__MissionResult())
}
func (t *MissionResult) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__MissionResult
	return (unsafe.Pointer)(C.px4_msgs__msg__MissionResult__create())
}
func (t *MissionResult) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__MissionResult__destroy((*C.px4_msgs__msg__MissionResult)(pointer_to_free))
}
func (t *MissionResult) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__MissionResult)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.instance_count = C.uint32_t(t.InstanceCount)
	mem.seq_reached = C.int32_t(t.SeqReached)
	mem.seq_current = C.uint16_t(t.SeqCurrent)
	mem.seq_total = C.uint16_t(t.SeqTotal)
	mem.valid = C.bool(t.Valid)
	mem.warning = C.bool(t.Warning)
	mem.finished = C.bool(t.Finished)
	mem.failure = C.bool(t.Failure)
	mem.stay_in_failsafe = C.bool(t.StayInFailsafe)
	mem.flight_termination = C.bool(t.FlightTermination)
	mem.item_do_jump_changed = C.bool(t.ItemDoJumpChanged)
	mem.item_changed_index = C.uint16_t(t.ItemChangedIndex)
	mem.item_do_jump_remaining = C.uint16_t(t.ItemDoJumpRemaining)
	mem.execution_mode = C.uint8_t(t.ExecutionMode)
	return unsafe.Pointer(mem)
}
func (t *MissionResult) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__MissionResult)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.InstanceCount = uint32(mem.instance_count)
	t.SeqReached = int32(mem.seq_reached)
	t.SeqCurrent = uint16(mem.seq_current)
	t.SeqTotal = uint16(mem.seq_total)
	t.Valid = bool(mem.valid)
	t.Warning = bool(mem.warning)
	t.Finished = bool(mem.finished)
	t.Failure = bool(mem.failure)
	t.StayInFailsafe = bool(mem.stay_in_failsafe)
	t.FlightTermination = bool(mem.flight_termination)
	t.ItemDoJumpChanged = bool(mem.item_do_jump_changed)
	t.ItemChangedIndex = uint16(mem.item_changed_index)
	t.ItemDoJumpRemaining = uint16(mem.item_do_jump_remaining)
	t.ExecutionMode = uint8(mem.execution_mode)
}
func (t *MissionResult) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CMissionResult = C.px4_msgs__msg__MissionResult
type CMissionResult__Sequence = C.px4_msgs__msg__MissionResult__Sequence

func MissionResult__Sequence_to_Go(goSlice *[]MissionResult, cSlice CMissionResult__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MissionResult, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__MissionResult__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__MissionResult * uintptr(i)),
		))
		(*goSlice)[i] = MissionResult{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func MissionResult__Sequence_to_C(cSlice *CMissionResult__Sequence, goSlice []MissionResult) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__MissionResult)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__MissionResult * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__MissionResult)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__MissionResult * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__MissionResult)(v.AsCStruct())
	}
}
func MissionResult__Array_to_Go(goSlice []MissionResult, cSlice []CMissionResult) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func MissionResult__Array_to_C(cSlice []CMissionResult, goSlice []MissionResult) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__MissionResult)(goSlice[i].AsCStruct())
	}
}


