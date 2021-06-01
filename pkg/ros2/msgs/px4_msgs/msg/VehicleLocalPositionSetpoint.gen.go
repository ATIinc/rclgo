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

#include <px4_msgs/msg/vehicle_local_position_setpoint.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/VehicleLocalPositionSetpoint", &VehicleLocalPositionSetpoint{})
}

// Do not create instances of this type directly. Always use NewVehicleLocalPositionSetpoint
// function instead.
type VehicleLocalPositionSetpoint struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	X float32 `yaml:"x"`// in meters NED
	Y float32 `yaml:"y"`// in meters NED
	Z float32 `yaml:"z"`// in meters NED
	Yaw float32 `yaml:"yaw"`// in radians NED -PI..+PI
	Yawspeed float32 `yaml:"yawspeed"`// in radians/sec
	Vx float32 `yaml:"vx"`// in meters/sec
	Vy float32 `yaml:"vy"`// in meters/sec
	Vz float32 `yaml:"vz"`// in meters/sec
	Acceleration [3]float32 `yaml:"acceleration"`// in meters/sec^2
	Jerk [3]float32 `yaml:"jerk"`// in meters/sec^3
	Thrust [3]float32 `yaml:"thrust"`// normalized thrust vector in NED
}

// NewVehicleLocalPositionSetpoint creates a new VehicleLocalPositionSetpoint with default values.
func NewVehicleLocalPositionSetpoint() *VehicleLocalPositionSetpoint {
	self := VehicleLocalPositionSetpoint{}
	self.SetDefaults(nil)
	return &self
}

func (t *VehicleLocalPositionSetpoint) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *VehicleLocalPositionSetpoint) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleLocalPositionSetpoint())
}
func (t *VehicleLocalPositionSetpoint) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleLocalPositionSetpoint
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleLocalPositionSetpoint__create())
}
func (t *VehicleLocalPositionSetpoint) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleLocalPositionSetpoint__destroy((*C.px4_msgs__msg__VehicleLocalPositionSetpoint)(pointer_to_free))
}
func (t *VehicleLocalPositionSetpoint) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__VehicleLocalPositionSetpoint)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.x = C.float(t.X)
	mem.y = C.float(t.Y)
	mem.z = C.float(t.Z)
	mem.yaw = C.float(t.Yaw)
	mem.yawspeed = C.float(t.Yawspeed)
	mem.vx = C.float(t.Vx)
	mem.vy = C.float(t.Vy)
	mem.vz = C.float(t.Vz)
	cSlice_acceleration := mem.acceleration[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_acceleration)), t.Acceleration[:])
	cSlice_jerk := mem.jerk[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_jerk)), t.Jerk[:])
	cSlice_thrust := mem.thrust[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_thrust)), t.Thrust[:])
	return unsafe.Pointer(mem)
}
func (t *VehicleLocalPositionSetpoint) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__VehicleLocalPositionSetpoint)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.X = float32(mem.x)
	t.Y = float32(mem.y)
	t.Z = float32(mem.z)
	t.Yaw = float32(mem.yaw)
	t.Yawspeed = float32(mem.yawspeed)
	t.Vx = float32(mem.vx)
	t.Vy = float32(mem.vy)
	t.Vz = float32(mem.vz)
	cSlice_acceleration := mem.acceleration[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.Acceleration[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_acceleration)))
	cSlice_jerk := mem.jerk[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.Jerk[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_jerk)))
	cSlice_thrust := mem.thrust[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.Thrust[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_thrust)))
}
func (t *VehicleLocalPositionSetpoint) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CVehicleLocalPositionSetpoint = C.px4_msgs__msg__VehicleLocalPositionSetpoint
type CVehicleLocalPositionSetpoint__Sequence = C.px4_msgs__msg__VehicleLocalPositionSetpoint__Sequence

func VehicleLocalPositionSetpoint__Sequence_to_Go(goSlice *[]VehicleLocalPositionSetpoint, cSlice CVehicleLocalPositionSetpoint__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleLocalPositionSetpoint, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleLocalPositionSetpoint__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleLocalPositionSetpoint * uintptr(i)),
		))
		(*goSlice)[i] = VehicleLocalPositionSetpoint{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func VehicleLocalPositionSetpoint__Sequence_to_C(cSlice *CVehicleLocalPositionSetpoint__Sequence, goSlice []VehicleLocalPositionSetpoint) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleLocalPositionSetpoint)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleLocalPositionSetpoint * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleLocalPositionSetpoint)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleLocalPositionSetpoint * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__VehicleLocalPositionSetpoint)(v.AsCStruct())
	}
}
func VehicleLocalPositionSetpoint__Array_to_Go(goSlice []VehicleLocalPositionSetpoint, cSlice []CVehicleLocalPositionSetpoint) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleLocalPositionSetpoint__Array_to_C(cSlice []CVehicleLocalPositionSetpoint, goSlice []VehicleLocalPositionSetpoint) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__VehicleLocalPositionSetpoint)(goSlice[i].AsCStruct())
	}
}


