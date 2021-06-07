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

#include <px4_msgs/msg/vehicle_attitude.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/VehicleAttitude", VehicleAttitudeTypeSupport)
}

// Do not create instances of this type directly. Always use NewVehicleAttitude
// function instead.
type VehicleAttitude struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp of the raw data (microseconds)
	Q [4]float32 `yaml:"q"`// Quaternion rotation from the FRD body frame to the NED earth frame
	DeltaQReset [4]float32 `yaml:"delta_q_reset"`// Amount by which quaternion has changed during last reset
	QuatResetCounter uint8 `yaml:"quat_reset_counter"`// Quaternion reset counter
}

// NewVehicleAttitude creates a new VehicleAttitude with default values.
func NewVehicleAttitude() *VehicleAttitude {
	self := VehicleAttitude{}
	self.SetDefaults()
	return &self
}

func (t *VehicleAttitude) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *VehicleAttitude) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var VehicleAttitudeTypeSupport types.MessageTypeSupport = _VehicleAttitudeTypeSupport{}

type _VehicleAttitudeTypeSupport struct{}

func (t _VehicleAttitudeTypeSupport) New() types.Message {
	return NewVehicleAttitude()
}

func (t _VehicleAttitudeTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleAttitude
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleAttitude__create())
}

func (t _VehicleAttitudeTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleAttitude__destroy((*C.px4_msgs__msg__VehicleAttitude)(pointer_to_free))
}

func (t _VehicleAttitudeTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*VehicleAttitude)
	mem := (*C.px4_msgs__msg__VehicleAttitude)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	cSlice_q := mem.q[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_q)), m.Q[:])
	cSlice_delta_q_reset := mem.delta_q_reset[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_delta_q_reset)), m.DeltaQReset[:])
	mem.quat_reset_counter = C.uint8_t(m.QuatResetCounter)
}

func (t _VehicleAttitudeTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*VehicleAttitude)
	mem := (*C.px4_msgs__msg__VehicleAttitude)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	cSlice_q := mem.q[:]
	rosidl_runtime_c.Float32__Array_to_Go(m.Q[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_q)))
	cSlice_delta_q_reset := mem.delta_q_reset[:]
	rosidl_runtime_c.Float32__Array_to_Go(m.DeltaQReset[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_delta_q_reset)))
	m.QuatResetCounter = uint8(mem.quat_reset_counter)
}

func (t _VehicleAttitudeTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleAttitude())
}

type CVehicleAttitude = C.px4_msgs__msg__VehicleAttitude
type CVehicleAttitude__Sequence = C.px4_msgs__msg__VehicleAttitude__Sequence

func VehicleAttitude__Sequence_to_Go(goSlice *[]VehicleAttitude, cSlice CVehicleAttitude__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleAttitude, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleAttitude__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleAttitude * uintptr(i)),
		))
		VehicleAttitudeTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func VehicleAttitude__Sequence_to_C(cSlice *CVehicleAttitude__Sequence, goSlice []VehicleAttitude) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleAttitude)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleAttitude * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleAttitude)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleAttitude * uintptr(i)),
		))
		VehicleAttitudeTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func VehicleAttitude__Array_to_Go(goSlice []VehicleAttitude, cSlice []CVehicleAttitude) {
	for i := 0; i < len(cSlice); i++ {
		VehicleAttitudeTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleAttitude__Array_to_C(cSlice []CVehicleAttitude, goSlice []VehicleAttitude) {
	for i := 0; i < len(goSlice); i++ {
		VehicleAttitudeTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
