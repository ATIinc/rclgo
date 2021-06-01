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

#include <px4_msgs/msg/yaw_estimator_status.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/YawEstimatorStatus", &YawEstimatorStatus{})
}

// Do not create instances of this type directly. Always use NewYawEstimatorStatus
// function instead.
type YawEstimatorStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp of the raw data (microseconds)
	YawComposite float32 `yaml:"yaw_composite"`// composite yaw from GSF (rad)
	YawVariance float32 `yaml:"yaw_variance"`// composite yaw variance from GSF (rad^2)
	Yaw [5]float32 `yaml:"yaw"`// yaw estimate for each model in the filter bank (rad)
	InnovVn [5]float32 `yaml:"innov_vn"`// North velocity innovation for each model in the filter bank (m/s)
	InnovVe [5]float32 `yaml:"innov_ve"`// East velocity innovation for each model in the filter bank (m/s)
	Weight [5]float32 `yaml:"weight"`// weighting for each model in the filter bank
}

// NewYawEstimatorStatus creates a new YawEstimatorStatus with default values.
func NewYawEstimatorStatus() *YawEstimatorStatus {
	self := YawEstimatorStatus{}
	self.SetDefaults(nil)
	return &self
}

func (t *YawEstimatorStatus) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *YawEstimatorStatus) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__YawEstimatorStatus())
}
func (t *YawEstimatorStatus) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__YawEstimatorStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__YawEstimatorStatus__create())
}
func (t *YawEstimatorStatus) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__YawEstimatorStatus__destroy((*C.px4_msgs__msg__YawEstimatorStatus)(pointer_to_free))
}
func (t *YawEstimatorStatus) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__YawEstimatorStatus)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.timestamp_sample = C.uint64_t(t.TimestampSample)
	mem.yaw_composite = C.float(t.YawComposite)
	mem.yaw_variance = C.float(t.YawVariance)
	cSlice_yaw := mem.yaw[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_yaw)), t.Yaw[:])
	cSlice_innov_vn := mem.innov_vn[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_innov_vn)), t.InnovVn[:])
	cSlice_innov_ve := mem.innov_ve[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_innov_ve)), t.InnovVe[:])
	cSlice_weight := mem.weight[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_weight)), t.Weight[:])
	return unsafe.Pointer(mem)
}
func (t *YawEstimatorStatus) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__YawEstimatorStatus)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.TimestampSample = uint64(mem.timestamp_sample)
	t.YawComposite = float32(mem.yaw_composite)
	t.YawVariance = float32(mem.yaw_variance)
	cSlice_yaw := mem.yaw[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.Yaw[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_yaw)))
	cSlice_innov_vn := mem.innov_vn[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.InnovVn[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_innov_vn)))
	cSlice_innov_ve := mem.innov_ve[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.InnovVe[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_innov_ve)))
	cSlice_weight := mem.weight[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.Weight[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_weight)))
}
func (t *YawEstimatorStatus) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CYawEstimatorStatus = C.px4_msgs__msg__YawEstimatorStatus
type CYawEstimatorStatus__Sequence = C.px4_msgs__msg__YawEstimatorStatus__Sequence

func YawEstimatorStatus__Sequence_to_Go(goSlice *[]YawEstimatorStatus, cSlice CYawEstimatorStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]YawEstimatorStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__YawEstimatorStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__YawEstimatorStatus * uintptr(i)),
		))
		(*goSlice)[i] = YawEstimatorStatus{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func YawEstimatorStatus__Sequence_to_C(cSlice *CYawEstimatorStatus__Sequence, goSlice []YawEstimatorStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__YawEstimatorStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__YawEstimatorStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__YawEstimatorStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__YawEstimatorStatus * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__YawEstimatorStatus)(v.AsCStruct())
	}
}
func YawEstimatorStatus__Array_to_Go(goSlice []YawEstimatorStatus, cSlice []CYawEstimatorStatus) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func YawEstimatorStatus__Array_to_C(cSlice []CYawEstimatorStatus, goSlice []YawEstimatorStatus) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__YawEstimatorStatus)(goSlice[i].AsCStruct())
	}
}


