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

#include <px4_msgs/msg/wind.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/Wind", WindTypeSupport)
}

// Do not create instances of this type directly. Always use NewWind
// function instead.
type Wind struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp of the raw data (microseconds)
	WindspeedNorth float32 `yaml:"windspeed_north"`// Wind component in north / X direction (m/sec)
	WindspeedEast float32 `yaml:"windspeed_east"`// Wind component in east / Y direction (m/sec)
	VarianceNorth float32 `yaml:"variance_north"`// Wind estimate error variance in north / X direction (m/sec)**2 - set to zero (no uncertainty) if not estimated
	VarianceEast float32 `yaml:"variance_east"`// Wind estimate error variance in east / Y direction (m/sec)**2 - set to zero (no uncertainty) if not estimated
	TasInnov float32 `yaml:"tas_innov"`// True airspeed innovation
	TasInnovVar float32 `yaml:"tas_innov_var"`// True airspeed innovation variance
	BetaInnov float32 `yaml:"beta_innov"`// Sideslip measurement innovation
	BetaInnovVar float32 `yaml:"beta_innov_var"`// Sideslip measurement innovation variance
}

// NewWind creates a new Wind with default values.
func NewWind() *Wind {
	self := Wind{}
	self.SetDefaults()
	return &self
}

func (t *Wind) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *Wind) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var WindTypeSupport types.MessageTypeSupport = _WindTypeSupport{}

type _WindTypeSupport struct{}

func (t _WindTypeSupport) New() types.Message {
	return NewWind()
}

func (t _WindTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__Wind
	return (unsafe.Pointer)(C.px4_msgs__msg__Wind__create())
}

func (t _WindTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__Wind__destroy((*C.px4_msgs__msg__Wind)(pointer_to_free))
}

func (t _WindTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Wind)
	mem := (*C.px4_msgs__msg__Wind)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	mem.windspeed_north = C.float(m.WindspeedNorth)
	mem.windspeed_east = C.float(m.WindspeedEast)
	mem.variance_north = C.float(m.VarianceNorth)
	mem.variance_east = C.float(m.VarianceEast)
	mem.tas_innov = C.float(m.TasInnov)
	mem.tas_innov_var = C.float(m.TasInnovVar)
	mem.beta_innov = C.float(m.BetaInnov)
	mem.beta_innov_var = C.float(m.BetaInnovVar)
}

func (t _WindTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Wind)
	mem := (*C.px4_msgs__msg__Wind)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	m.WindspeedNorth = float32(mem.windspeed_north)
	m.WindspeedEast = float32(mem.windspeed_east)
	m.VarianceNorth = float32(mem.variance_north)
	m.VarianceEast = float32(mem.variance_east)
	m.TasInnov = float32(mem.tas_innov)
	m.TasInnovVar = float32(mem.tas_innov_var)
	m.BetaInnov = float32(mem.beta_innov)
	m.BetaInnovVar = float32(mem.beta_innov_var)
}

func (t _WindTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__Wind())
}

type CWind = C.px4_msgs__msg__Wind
type CWind__Sequence = C.px4_msgs__msg__Wind__Sequence

func Wind__Sequence_to_Go(goSlice *[]Wind, cSlice CWind__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Wind, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__Wind__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__Wind * uintptr(i)),
		))
		WindTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Wind__Sequence_to_C(cSlice *CWind__Sequence, goSlice []Wind) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__Wind)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__Wind * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__Wind)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__Wind * uintptr(i)),
		))
		WindTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Wind__Array_to_Go(goSlice []Wind, cSlice []CWind) {
	for i := 0; i < len(cSlice); i++ {
		WindTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Wind__Array_to_C(cSlice []CWind, goSlice []Wind) {
	for i := 0; i < len(goSlice); i++ {
		WindTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
