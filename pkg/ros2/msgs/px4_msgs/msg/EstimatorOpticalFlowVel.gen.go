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

#include <px4_msgs/msg/estimator_optical_flow_vel.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/EstimatorOpticalFlowVel", &EstimatorOpticalFlowVel{})
}

// Do not create instances of this type directly. Always use NewEstimatorOpticalFlowVel
// function instead.
type EstimatorOpticalFlowVel struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp of the raw data (microseconds)
	VelBody [2]float32 `yaml:"vel_body"`// velocity obtained from gyro-compensated and distance-scaled optical flow raw measurements in body frame(m/s)
	VelNe [2]float32 `yaml:"vel_ne"`// same as vel_body but in local frame (m/s)
	FlowUncompensatedIntegral [2]float32 `yaml:"flow_uncompensated_integral"`// integrated optical flow measurement (rad)
	FlowCompensatedIntegral [2]float32 `yaml:"flow_compensated_integral"`// integrated optical flow measurement compensated for angular motion (rad)
	GyroRateIntegral [3]float32 `yaml:"gyro_rate_integral"`// gyro measurement integrated to flow rate and synchronized with flow measurements (rad)
}

// NewEstimatorOpticalFlowVel creates a new EstimatorOpticalFlowVel with default values.
func NewEstimatorOpticalFlowVel() *EstimatorOpticalFlowVel {
	self := EstimatorOpticalFlowVel{}
	self.SetDefaults(nil)
	return &self
}

func (t *EstimatorOpticalFlowVel) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *EstimatorOpticalFlowVel) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__EstimatorOpticalFlowVel())
}
func (t *EstimatorOpticalFlowVel) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__EstimatorOpticalFlowVel
	return (unsafe.Pointer)(C.px4_msgs__msg__EstimatorOpticalFlowVel__create())
}
func (t *EstimatorOpticalFlowVel) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__EstimatorOpticalFlowVel__destroy((*C.px4_msgs__msg__EstimatorOpticalFlowVel)(pointer_to_free))
}
func (t *EstimatorOpticalFlowVel) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__EstimatorOpticalFlowVel)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.timestamp_sample = C.uint64_t(t.TimestampSample)
	cSlice_vel_body := mem.vel_body[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_vel_body)), t.VelBody[:])
	cSlice_vel_ne := mem.vel_ne[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_vel_ne)), t.VelNe[:])
	cSlice_flow_uncompensated_integral := mem.flow_uncompensated_integral[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_flow_uncompensated_integral)), t.FlowUncompensatedIntegral[:])
	cSlice_flow_compensated_integral := mem.flow_compensated_integral[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_flow_compensated_integral)), t.FlowCompensatedIntegral[:])
	cSlice_gyro_rate_integral := mem.gyro_rate_integral[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_gyro_rate_integral)), t.GyroRateIntegral[:])
	return unsafe.Pointer(mem)
}
func (t *EstimatorOpticalFlowVel) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__EstimatorOpticalFlowVel)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.TimestampSample = uint64(mem.timestamp_sample)
	cSlice_vel_body := mem.vel_body[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.VelBody[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_vel_body)))
	cSlice_vel_ne := mem.vel_ne[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.VelNe[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_vel_ne)))
	cSlice_flow_uncompensated_integral := mem.flow_uncompensated_integral[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.FlowUncompensatedIntegral[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_flow_uncompensated_integral)))
	cSlice_flow_compensated_integral := mem.flow_compensated_integral[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.FlowCompensatedIntegral[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_flow_compensated_integral)))
	cSlice_gyro_rate_integral := mem.gyro_rate_integral[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.GyroRateIntegral[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_gyro_rate_integral)))
}
func (t *EstimatorOpticalFlowVel) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CEstimatorOpticalFlowVel = C.px4_msgs__msg__EstimatorOpticalFlowVel
type CEstimatorOpticalFlowVel__Sequence = C.px4_msgs__msg__EstimatorOpticalFlowVel__Sequence

func EstimatorOpticalFlowVel__Sequence_to_Go(goSlice *[]EstimatorOpticalFlowVel, cSlice CEstimatorOpticalFlowVel__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]EstimatorOpticalFlowVel, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__EstimatorOpticalFlowVel__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorOpticalFlowVel * uintptr(i)),
		))
		(*goSlice)[i] = EstimatorOpticalFlowVel{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func EstimatorOpticalFlowVel__Sequence_to_C(cSlice *CEstimatorOpticalFlowVel__Sequence, goSlice []EstimatorOpticalFlowVel) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__EstimatorOpticalFlowVel)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__EstimatorOpticalFlowVel * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__EstimatorOpticalFlowVel)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorOpticalFlowVel * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__EstimatorOpticalFlowVel)(v.AsCStruct())
	}
}
func EstimatorOpticalFlowVel__Array_to_Go(goSlice []EstimatorOpticalFlowVel, cSlice []CEstimatorOpticalFlowVel) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func EstimatorOpticalFlowVel__Array_to_C(cSlice []CEstimatorOpticalFlowVel, goSlice []EstimatorOpticalFlowVel) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__EstimatorOpticalFlowVel)(goSlice[i].AsCStruct())
	}
}


