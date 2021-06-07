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

#include <px4_msgs/msg/logger_status.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/LoggerStatus", LoggerStatusTypeSupport)
}
const (
	LoggerStatus_LOGGER_TYPE_FULL uint8 = 0// Normal, full size log
	LoggerStatus_LOGGER_TYPE_MISSION uint8 = 1// reduced mission log (e.g. for geotagging)
	LoggerStatus_BACKEND_FILE uint8 = 1
	LoggerStatus_BACKEND_MAVLINK uint8 = 2
	LoggerStatus_BACKEND_ALL uint8 = 3
)

// Do not create instances of this type directly. Always use NewLoggerStatus
// function instead.
type LoggerStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Type uint8 `yaml:"type"`
	Backend uint8 `yaml:"backend"`
	TotalWrittenKb float32 `yaml:"total_written_kb"`// total written to log in kiloBytes
	WriteRateKbS float32 `yaml:"write_rate_kb_s"`// write rate in kiloBytes/s
	Dropouts uint32 `yaml:"dropouts"`// number of failed buffer writes due to buffer overflow
	MessageGaps uint32 `yaml:"message_gaps"`// messages misssed
	BufferUsedBytes uint32 `yaml:"buffer_used_bytes"`// current buffer fill in Bytes
	BufferSizeBytes uint32 `yaml:"buffer_size_bytes"`// total buffer size in Bytes
	NumMessages uint8 `yaml:"num_messages"`
}

// NewLoggerStatus creates a new LoggerStatus with default values.
func NewLoggerStatus() *LoggerStatus {
	self := LoggerStatus{}
	self.SetDefaults()
	return &self
}

func (t *LoggerStatus) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *LoggerStatus) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var LoggerStatusTypeSupport types.MessageTypeSupport = _LoggerStatusTypeSupport{}

type _LoggerStatusTypeSupport struct{}

func (t _LoggerStatusTypeSupport) New() types.Message {
	return NewLoggerStatus()
}

func (t _LoggerStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__LoggerStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__LoggerStatus__create())
}

func (t _LoggerStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__LoggerStatus__destroy((*C.px4_msgs__msg__LoggerStatus)(pointer_to_free))
}

func (t _LoggerStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*LoggerStatus)
	mem := (*C.px4_msgs__msg__LoggerStatus)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem._type = C.uint8_t(m.Type)
	mem.backend = C.uint8_t(m.Backend)
	mem.total_written_kb = C.float(m.TotalWrittenKb)
	mem.write_rate_kb_s = C.float(m.WriteRateKbS)
	mem.dropouts = C.uint32_t(m.Dropouts)
	mem.message_gaps = C.uint32_t(m.MessageGaps)
	mem.buffer_used_bytes = C.uint32_t(m.BufferUsedBytes)
	mem.buffer_size_bytes = C.uint32_t(m.BufferSizeBytes)
	mem.num_messages = C.uint8_t(m.NumMessages)
}

func (t _LoggerStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*LoggerStatus)
	mem := (*C.px4_msgs__msg__LoggerStatus)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.Type = uint8(mem._type)
	m.Backend = uint8(mem.backend)
	m.TotalWrittenKb = float32(mem.total_written_kb)
	m.WriteRateKbS = float32(mem.write_rate_kb_s)
	m.Dropouts = uint32(mem.dropouts)
	m.MessageGaps = uint32(mem.message_gaps)
	m.BufferUsedBytes = uint32(mem.buffer_used_bytes)
	m.BufferSizeBytes = uint32(mem.buffer_size_bytes)
	m.NumMessages = uint8(mem.num_messages)
}

func (t _LoggerStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__LoggerStatus())
}

type CLoggerStatus = C.px4_msgs__msg__LoggerStatus
type CLoggerStatus__Sequence = C.px4_msgs__msg__LoggerStatus__Sequence

func LoggerStatus__Sequence_to_Go(goSlice *[]LoggerStatus, cSlice CLoggerStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]LoggerStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__LoggerStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__LoggerStatus * uintptr(i)),
		))
		LoggerStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func LoggerStatus__Sequence_to_C(cSlice *CLoggerStatus__Sequence, goSlice []LoggerStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__LoggerStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__LoggerStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__LoggerStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__LoggerStatus * uintptr(i)),
		))
		LoggerStatusTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func LoggerStatus__Array_to_Go(goSlice []LoggerStatus, cSlice []CLoggerStatus) {
	for i := 0; i < len(cSlice); i++ {
		LoggerStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func LoggerStatus__Array_to_C(cSlice []CLoggerStatus, goSlice []LoggerStatus) {
	for i := 0; i < len(goSlice); i++ {
		LoggerStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
