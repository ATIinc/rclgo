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

#include <px4_msgs/msg/vehicle_visual_odometry.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/VehicleVisualOdometry", VehicleVisualOdometryTypeSupport)
}
const (
	VehicleVisualOdometry_COVARIANCE_MATRIX_X_VARIANCE uint8 = 0// Covariance matrix index constants
	VehicleVisualOdometry_COVARIANCE_MATRIX_Y_VARIANCE uint8 = 6// Covariance matrix index constants
	VehicleVisualOdometry_COVARIANCE_MATRIX_Z_VARIANCE uint8 = 11// Covariance matrix index constants
	VehicleVisualOdometry_COVARIANCE_MATRIX_ROLL_VARIANCE uint8 = 15// Covariance matrix index constants
	VehicleVisualOdometry_COVARIANCE_MATRIX_PITCH_VARIANCE uint8 = 18// Covariance matrix index constants
	VehicleVisualOdometry_COVARIANCE_MATRIX_YAW_VARIANCE uint8 = 20// Covariance matrix index constants
	VehicleVisualOdometry_COVARIANCE_MATRIX_VX_VARIANCE uint8 = 0// Covariance matrix index constants
	VehicleVisualOdometry_COVARIANCE_MATRIX_VY_VARIANCE uint8 = 6// Covariance matrix index constants
	VehicleVisualOdometry_COVARIANCE_MATRIX_VZ_VARIANCE uint8 = 11// Covariance matrix index constants
	VehicleVisualOdometry_COVARIANCE_MATRIX_ROLLRATE_VARIANCE uint8 = 15// Covariance matrix index constants
	VehicleVisualOdometry_COVARIANCE_MATRIX_PITCHRATE_VARIANCE uint8 = 18// Covariance matrix index constants
	VehicleVisualOdometry_COVARIANCE_MATRIX_YAWRATE_VARIANCE uint8 = 20// Covariance matrix index constants
	VehicleVisualOdometry_LOCAL_FRAME_NED uint8 = 0// NED earth-fixed frame. Position and linear velocity frame of reference constants
	VehicleVisualOdometry_LOCAL_FRAME_FRD uint8 = 1// FRD earth-fixed frame, arbitrary heading reference. Position and linear velocity frame of reference constants
	VehicleVisualOdometry_LOCAL_FRAME_OTHER uint8 = 2// Not aligned with the std frames of reference. Position and linear velocity frame of reference constants
	VehicleVisualOdometry_BODY_FRAME_FRD uint8 = 3// FRD body-fixed frame. Position and linear velocity frame of reference constants
)

// Do not create instances of this type directly. Always use NewVehicleVisualOdometry
// function instead.
type VehicleVisualOdometry struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds). Vehicle odometry data. Fits ROS REP 147 for aerial vehicles
	TimestampSample uint64 `yaml:"timestamp_sample"`
	LocalFrame uint8 `yaml:"local_frame"`// Position and linear velocity local frame of reference
	X float32 `yaml:"x"`// North position. Position in meters. Frame of reference defined by local_frame. NaN if invalid/unknown
	Y float32 `yaml:"y"`// East position. Position in meters. Frame of reference defined by local_frame. NaN if invalid/unknown
	Z float32 `yaml:"z"`// Down position. Position in meters. Frame of reference defined by local_frame. NaN if invalid/unknown
	Q [4]float32 `yaml:"q"`// Quaternion rotation from FRD body frame to refernce frame. Orientation quaternion. First value NaN if invalid/unknown
	QOffset [4]float32 `yaml:"q_offset"`// Quaternion rotation from odometry reference frame to navigation frame. Orientation quaternion. First value NaN if invalid/unknown
	PoseCovariance [21]float32 `yaml:"pose_covariance"`// Row-major representation of 6x6 pose cross-covariance matrix URT.NED earth-fixed frame.Order: x, y, z, rotation about X axis, rotation about Y axis, rotation about Z axisIf position covariance invalid/unknown, first cell is NaNIf orientation covariance invalid/unknown, 16th cell is NaN
	VelocityFrame uint8 `yaml:"velocity_frame"`// Reference frame of the velocity data
	Vx float32 `yaml:"vx"`// North velocity. Velocity in meters/sec. Frame of reference defined by velocity_frame variable. NaN if invalid/unknown
	Vy float32 `yaml:"vy"`// East velocity. Velocity in meters/sec. Frame of reference defined by velocity_frame variable. NaN if invalid/unknown
	Vz float32 `yaml:"vz"`// Down velocity. Velocity in meters/sec. Frame of reference defined by velocity_frame variable. NaN if invalid/unknown
	Rollspeed float32 `yaml:"rollspeed"`// Angular velocity about X body axis. Angular rate in body-fixed frame (rad/s). NaN if invalid/unknown
	Pitchspeed float32 `yaml:"pitchspeed"`// Angular velocity about Y body axis. Angular rate in body-fixed frame (rad/s). NaN if invalid/unknown
	Yawspeed float32 `yaml:"yawspeed"`// Angular velocity about Z body axis. Angular rate in body-fixed frame (rad/s). NaN if invalid/unknown
	VelocityCovariance [21]float32 `yaml:"velocity_covariance"`// Row-major representation of 6x6 velocity cross-covariance matrix URT.Linear velocity in NED earth-fixed frame. Angular velocity in body-fixed frame.Order: vx, vy, vz, rotation rate about X axis, rotation rate about Y axis, rotation rate about Z axisIf linear velocity covariance invalid/unknown, first cell is NaNIf angular velocity covariance invalid/unknown, 16th cell is NaN
}

// NewVehicleVisualOdometry creates a new VehicleVisualOdometry with default values.
func NewVehicleVisualOdometry() *VehicleVisualOdometry {
	self := VehicleVisualOdometry{}
	self.SetDefaults()
	return &self
}

func (t *VehicleVisualOdometry) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *VehicleVisualOdometry) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var VehicleVisualOdometryTypeSupport types.MessageTypeSupport = _VehicleVisualOdometryTypeSupport{}

type _VehicleVisualOdometryTypeSupport struct{}

func (t _VehicleVisualOdometryTypeSupport) New() types.Message {
	return NewVehicleVisualOdometry()
}

func (t _VehicleVisualOdometryTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleVisualOdometry
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleVisualOdometry__create())
}

func (t _VehicleVisualOdometryTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleVisualOdometry__destroy((*C.px4_msgs__msg__VehicleVisualOdometry)(pointer_to_free))
}

func (t _VehicleVisualOdometryTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*VehicleVisualOdometry)
	mem := (*C.px4_msgs__msg__VehicleVisualOdometry)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	mem.local_frame = C.uint8_t(m.LocalFrame)
	mem.x = C.float(m.X)
	mem.y = C.float(m.Y)
	mem.z = C.float(m.Z)
	cSlice_q := mem.q[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_q)), m.Q[:])
	cSlice_q_offset := mem.q_offset[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_q_offset)), m.QOffset[:])
	cSlice_pose_covariance := mem.pose_covariance[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_pose_covariance)), m.PoseCovariance[:])
	mem.velocity_frame = C.uint8_t(m.VelocityFrame)
	mem.vx = C.float(m.Vx)
	mem.vy = C.float(m.Vy)
	mem.vz = C.float(m.Vz)
	mem.rollspeed = C.float(m.Rollspeed)
	mem.pitchspeed = C.float(m.Pitchspeed)
	mem.yawspeed = C.float(m.Yawspeed)
	cSlice_velocity_covariance := mem.velocity_covariance[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_velocity_covariance)), m.VelocityCovariance[:])
}

func (t _VehicleVisualOdometryTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*VehicleVisualOdometry)
	mem := (*C.px4_msgs__msg__VehicleVisualOdometry)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	m.LocalFrame = uint8(mem.local_frame)
	m.X = float32(mem.x)
	m.Y = float32(mem.y)
	m.Z = float32(mem.z)
	cSlice_q := mem.q[:]
	rosidl_runtime_c.Float32__Array_to_Go(m.Q[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_q)))
	cSlice_q_offset := mem.q_offset[:]
	rosidl_runtime_c.Float32__Array_to_Go(m.QOffset[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_q_offset)))
	cSlice_pose_covariance := mem.pose_covariance[:]
	rosidl_runtime_c.Float32__Array_to_Go(m.PoseCovariance[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_pose_covariance)))
	m.VelocityFrame = uint8(mem.velocity_frame)
	m.Vx = float32(mem.vx)
	m.Vy = float32(mem.vy)
	m.Vz = float32(mem.vz)
	m.Rollspeed = float32(mem.rollspeed)
	m.Pitchspeed = float32(mem.pitchspeed)
	m.Yawspeed = float32(mem.yawspeed)
	cSlice_velocity_covariance := mem.velocity_covariance[:]
	rosidl_runtime_c.Float32__Array_to_Go(m.VelocityCovariance[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_velocity_covariance)))
}

func (t _VehicleVisualOdometryTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleVisualOdometry())
}

type CVehicleVisualOdometry = C.px4_msgs__msg__VehicleVisualOdometry
type CVehicleVisualOdometry__Sequence = C.px4_msgs__msg__VehicleVisualOdometry__Sequence

func VehicleVisualOdometry__Sequence_to_Go(goSlice *[]VehicleVisualOdometry, cSlice CVehicleVisualOdometry__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleVisualOdometry, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleVisualOdometry__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleVisualOdometry * uintptr(i)),
		))
		VehicleVisualOdometryTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func VehicleVisualOdometry__Sequence_to_C(cSlice *CVehicleVisualOdometry__Sequence, goSlice []VehicleVisualOdometry) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleVisualOdometry)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleVisualOdometry * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleVisualOdometry)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleVisualOdometry * uintptr(i)),
		))
		VehicleVisualOdometryTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func VehicleVisualOdometry__Array_to_Go(goSlice []VehicleVisualOdometry, cSlice []CVehicleVisualOdometry) {
	for i := 0; i < len(cSlice); i++ {
		VehicleVisualOdometryTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleVisualOdometry__Array_to_C(cSlice []CVehicleVisualOdometry, goSlice []VehicleVisualOdometry) {
	for i := 0; i < len(goSlice); i++ {
		VehicleVisualOdometryTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
