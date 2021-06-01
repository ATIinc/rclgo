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

#include <px4_msgs/msg/vehicle_rates_setpoint.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/VehicleRatesSetpoint", &VehicleRatesSetpoint{})
}

// Do not create instances of this type directly. Always use NewVehicleRatesSetpoint
// function instead.
type VehicleRatesSetpoint struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Roll float32 `yaml:"roll"`// body angular rates in NED frame
	Pitch float32 `yaml:"pitch"`// body angular rates in NED frame
	Yaw float32 `yaml:"yaw"`// body angular rates in NED frame
	ThrustBody [3]float32 `yaml:"thrust_body"`// Normalized thrust command in body NED frame [-1,1]. For clarification: For multicopters thrust_body[0] and thrust[1] are usually 0 and thrust[2] is the negative throttle demand.For fixed wings thrust_x is the throttle demand and thrust_y, thrust_z will usually be zero.
}

// NewVehicleRatesSetpoint creates a new VehicleRatesSetpoint with default values.
func NewVehicleRatesSetpoint() *VehicleRatesSetpoint {
	self := VehicleRatesSetpoint{}
	self.SetDefaults(nil)
	return &self
}

func (t *VehicleRatesSetpoint) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *VehicleRatesSetpoint) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleRatesSetpoint())
}
func (t *VehicleRatesSetpoint) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleRatesSetpoint
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleRatesSetpoint__create())
}
func (t *VehicleRatesSetpoint) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleRatesSetpoint__destroy((*C.px4_msgs__msg__VehicleRatesSetpoint)(pointer_to_free))
}
func (t *VehicleRatesSetpoint) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__VehicleRatesSetpoint)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.roll = C.float(t.Roll)
	mem.pitch = C.float(t.Pitch)
	mem.yaw = C.float(t.Yaw)
	cSlice_thrust_body := mem.thrust_body[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_thrust_body)), t.ThrustBody[:])
	return unsafe.Pointer(mem)
}
func (t *VehicleRatesSetpoint) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__VehicleRatesSetpoint)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.Roll = float32(mem.roll)
	t.Pitch = float32(mem.pitch)
	t.Yaw = float32(mem.yaw)
	cSlice_thrust_body := mem.thrust_body[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.ThrustBody[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_thrust_body)))
}
func (t *VehicleRatesSetpoint) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CVehicleRatesSetpoint = C.px4_msgs__msg__VehicleRatesSetpoint
type CVehicleRatesSetpoint__Sequence = C.px4_msgs__msg__VehicleRatesSetpoint__Sequence

func VehicleRatesSetpoint__Sequence_to_Go(goSlice *[]VehicleRatesSetpoint, cSlice CVehicleRatesSetpoint__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleRatesSetpoint, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleRatesSetpoint__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleRatesSetpoint * uintptr(i)),
		))
		(*goSlice)[i] = VehicleRatesSetpoint{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func VehicleRatesSetpoint__Sequence_to_C(cSlice *CVehicleRatesSetpoint__Sequence, goSlice []VehicleRatesSetpoint) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleRatesSetpoint)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleRatesSetpoint * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleRatesSetpoint)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleRatesSetpoint * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__VehicleRatesSetpoint)(v.AsCStruct())
	}
}
func VehicleRatesSetpoint__Array_to_Go(goSlice []VehicleRatesSetpoint, cSlice []CVehicleRatesSetpoint) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleRatesSetpoint__Array_to_C(cSlice []CVehicleRatesSetpoint, goSlice []VehicleRatesSetpoint) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__VehicleRatesSetpoint)(goSlice[i].AsCStruct())
	}
}


