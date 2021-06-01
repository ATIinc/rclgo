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
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/vehicle_angular_velocity.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/VehicleAngularVelocity", &VehicleAngularVelocity{})
}

// Do not create instances of this type directly. Always use NewVehicleAngularVelocity
// function instead.
type VehicleAngularVelocity struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// timestamp of the data sample on which this message is based (microseconds)
	Xyz [3]float32 `yaml:"xyz"`// Bias corrected angular velocity about the FRD body frame XYZ-axis in rad/s
}

// NewVehicleAngularVelocity creates a new VehicleAngularVelocity with default values.
func NewVehicleAngularVelocity() *VehicleAngularVelocity {
	self := VehicleAngularVelocity{}
	self.SetDefaults(nil)
	return &self
}

func (t *VehicleAngularVelocity) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *VehicleAngularVelocity) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleAngularVelocity())
}
func (t *VehicleAngularVelocity) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleAngularVelocity
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleAngularVelocity__create())
}
func (t *VehicleAngularVelocity) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleAngularVelocity__destroy((*C.px4_msgs__msg__VehicleAngularVelocity)(pointer_to_free))
}
func (t *VehicleAngularVelocity) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__VehicleAngularVelocity)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.timestamp_sample = C.uint64_t(t.TimestampSample)
	cSlice_xyz := mem.xyz[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_xyz)), t.Xyz[:])
	return unsafe.Pointer(mem)
}
func (t *VehicleAngularVelocity) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__VehicleAngularVelocity)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.TimestampSample = uint64(mem.timestamp_sample)
	cSlice_xyz := mem.xyz[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.Xyz[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_xyz)))
}
func (t *VehicleAngularVelocity) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CVehicleAngularVelocity = C.px4_msgs__msg__VehicleAngularVelocity
type CVehicleAngularVelocity__Sequence = C.px4_msgs__msg__VehicleAngularVelocity__Sequence

func VehicleAngularVelocity__Sequence_to_Go(goSlice *[]VehicleAngularVelocity, cSlice CVehicleAngularVelocity__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleAngularVelocity, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleAngularVelocity__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleAngularVelocity * uintptr(i)),
		))
		(*goSlice)[i] = VehicleAngularVelocity{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func VehicleAngularVelocity__Sequence_to_C(cSlice *CVehicleAngularVelocity__Sequence, goSlice []VehicleAngularVelocity) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleAngularVelocity)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleAngularVelocity * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleAngularVelocity)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleAngularVelocity * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__VehicleAngularVelocity)(v.AsCStruct())
	}
}
func VehicleAngularVelocity__Array_to_Go(goSlice []VehicleAngularVelocity, cSlice []CVehicleAngularVelocity) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleAngularVelocity__Array_to_C(cSlice []CVehicleAngularVelocity, goSlice []VehicleAngularVelocity) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__VehicleAngularVelocity)(goSlice[i].AsCStruct())
	}
}


