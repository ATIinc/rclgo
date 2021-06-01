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

#include <px4_msgs/msg/sensor_correction.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/SensorCorrection", &SensorCorrection{})
}

// Do not create instances of this type directly. Always use NewSensorCorrection
// function instead.
type SensorCorrection struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	GyroDeviceIds [4]uint32 `yaml:"gyro_device_ids"`// Corrections for gyro angular rate outputs where corrected_rate = raw_rate * gyro_scale + gyro_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	GyroOffset0 [3]float32 `yaml:"gyro_offset_0"`// gyro 0 XYZ offsets in the sensor frame in rad/s. Corrections for gyro angular rate outputs where corrected_rate = raw_rate * gyro_scale + gyro_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	GyroOffset1 [3]float32 `yaml:"gyro_offset_1"`// gyro 1 XYZ offsets in the sensor frame in rad/s. Corrections for gyro angular rate outputs where corrected_rate = raw_rate * gyro_scale + gyro_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	GyroOffset2 [3]float32 `yaml:"gyro_offset_2"`// gyro 2 XYZ offsets in the sensor frame in rad/s. Corrections for gyro angular rate outputs where corrected_rate = raw_rate * gyro_scale + gyro_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	GyroOffset3 [3]float32 `yaml:"gyro_offset_3"`// gyro 3 XYZ offsets in the sensor frame in rad/s. Corrections for gyro angular rate outputs where corrected_rate = raw_rate * gyro_scale + gyro_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	AccelDeviceIds [4]uint32 `yaml:"accel_device_ids"`// Corrections for acceleromter acceleration outputs where corrected_accel = raw_accel * accel_scale + accel_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	AccelOffset0 [3]float32 `yaml:"accel_offset_0"`// accelerometer 0 offsets in the FRD board frame XYZ-axis in m/s^s. Corrections for acceleromter acceleration outputs where corrected_accel = raw_accel * accel_scale + accel_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	AccelOffset1 [3]float32 `yaml:"accel_offset_1"`// accelerometer 1 offsets in the FRD board frame XYZ-axis in m/s^s. Corrections for acceleromter acceleration outputs where corrected_accel = raw_accel * accel_scale + accel_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	AccelOffset2 [3]float32 `yaml:"accel_offset_2"`// accelerometer 2 offsets in the FRD board frame XYZ-axis in m/s^s. Corrections for acceleromter acceleration outputs where corrected_accel = raw_accel * accel_scale + accel_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	AccelOffset3 [3]float32 `yaml:"accel_offset_3"`// accelerometer 3 offsets in the FRD board frame XYZ-axis in m/s^s. Corrections for acceleromter acceleration outputs where corrected_accel = raw_accel * accel_scale + accel_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	BaroDeviceIds [4]uint32 `yaml:"baro_device_ids"`// Corrections for barometric pressure outputs where corrected_pressure = raw_pressure * pressure_scale + pressure_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	BaroOffset0 float32 `yaml:"baro_offset_0"`// barometric pressure 0 offsets in the sensor frame in Pascals. Corrections for barometric pressure outputs where corrected_pressure = raw_pressure * pressure_scale + pressure_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	BaroOffset1 float32 `yaml:"baro_offset_1"`// barometric pressure 1 offsets in the sensor frame in Pascals. Corrections for barometric pressure outputs where corrected_pressure = raw_pressure * pressure_scale + pressure_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	BaroOffset2 float32 `yaml:"baro_offset_2"`// barometric pressure 2 offsets in the sensor frame in Pascals. Corrections for barometric pressure outputs where corrected_pressure = raw_pressure * pressure_scale + pressure_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	BaroOffset3 float32 `yaml:"baro_offset_3"`// barometric pressure 3 offsets in the sensor frame in Pascals. Corrections for barometric pressure outputs where corrected_pressure = raw_pressure * pressure_scale + pressure_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
}

// NewSensorCorrection creates a new SensorCorrection with default values.
func NewSensorCorrection() *SensorCorrection {
	self := SensorCorrection{}
	self.SetDefaults(nil)
	return &self
}

func (t *SensorCorrection) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *SensorCorrection) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__SensorCorrection())
}
func (t *SensorCorrection) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__SensorCorrection
	return (unsafe.Pointer)(C.px4_msgs__msg__SensorCorrection__create())
}
func (t *SensorCorrection) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__SensorCorrection__destroy((*C.px4_msgs__msg__SensorCorrection)(pointer_to_free))
}
func (t *SensorCorrection) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__SensorCorrection)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	cSlice_gyro_device_ids := mem.gyro_device_ids[:]
	rosidl_runtime_c.Uint32__Array_to_C(*(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_gyro_device_ids)), t.GyroDeviceIds[:])
	cSlice_gyro_offset_0 := mem.gyro_offset_0[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_gyro_offset_0)), t.GyroOffset0[:])
	cSlice_gyro_offset_1 := mem.gyro_offset_1[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_gyro_offset_1)), t.GyroOffset1[:])
	cSlice_gyro_offset_2 := mem.gyro_offset_2[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_gyro_offset_2)), t.GyroOffset2[:])
	cSlice_gyro_offset_3 := mem.gyro_offset_3[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_gyro_offset_3)), t.GyroOffset3[:])
	cSlice_accel_device_ids := mem.accel_device_ids[:]
	rosidl_runtime_c.Uint32__Array_to_C(*(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_accel_device_ids)), t.AccelDeviceIds[:])
	cSlice_accel_offset_0 := mem.accel_offset_0[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_accel_offset_0)), t.AccelOffset0[:])
	cSlice_accel_offset_1 := mem.accel_offset_1[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_accel_offset_1)), t.AccelOffset1[:])
	cSlice_accel_offset_2 := mem.accel_offset_2[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_accel_offset_2)), t.AccelOffset2[:])
	cSlice_accel_offset_3 := mem.accel_offset_3[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_accel_offset_3)), t.AccelOffset3[:])
	cSlice_baro_device_ids := mem.baro_device_ids[:]
	rosidl_runtime_c.Uint32__Array_to_C(*(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_baro_device_ids)), t.BaroDeviceIds[:])
	mem.baro_offset_0 = C.float(t.BaroOffset0)
	mem.baro_offset_1 = C.float(t.BaroOffset1)
	mem.baro_offset_2 = C.float(t.BaroOffset2)
	mem.baro_offset_3 = C.float(t.BaroOffset3)
	return unsafe.Pointer(mem)
}
func (t *SensorCorrection) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__SensorCorrection)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	cSlice_gyro_device_ids := mem.gyro_device_ids[:]
	rosidl_runtime_c.Uint32__Array_to_Go(t.GyroDeviceIds[:], *(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_gyro_device_ids)))
	cSlice_gyro_offset_0 := mem.gyro_offset_0[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.GyroOffset0[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_gyro_offset_0)))
	cSlice_gyro_offset_1 := mem.gyro_offset_1[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.GyroOffset1[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_gyro_offset_1)))
	cSlice_gyro_offset_2 := mem.gyro_offset_2[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.GyroOffset2[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_gyro_offset_2)))
	cSlice_gyro_offset_3 := mem.gyro_offset_3[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.GyroOffset3[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_gyro_offset_3)))
	cSlice_accel_device_ids := mem.accel_device_ids[:]
	rosidl_runtime_c.Uint32__Array_to_Go(t.AccelDeviceIds[:], *(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_accel_device_ids)))
	cSlice_accel_offset_0 := mem.accel_offset_0[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.AccelOffset0[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_accel_offset_0)))
	cSlice_accel_offset_1 := mem.accel_offset_1[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.AccelOffset1[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_accel_offset_1)))
	cSlice_accel_offset_2 := mem.accel_offset_2[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.AccelOffset2[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_accel_offset_2)))
	cSlice_accel_offset_3 := mem.accel_offset_3[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.AccelOffset3[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_accel_offset_3)))
	cSlice_baro_device_ids := mem.baro_device_ids[:]
	rosidl_runtime_c.Uint32__Array_to_Go(t.BaroDeviceIds[:], *(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_baro_device_ids)))
	t.BaroOffset0 = float32(mem.baro_offset_0)
	t.BaroOffset1 = float32(mem.baro_offset_1)
	t.BaroOffset2 = float32(mem.baro_offset_2)
	t.BaroOffset3 = float32(mem.baro_offset_3)
}
func (t *SensorCorrection) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CSensorCorrection = C.px4_msgs__msg__SensorCorrection
type CSensorCorrection__Sequence = C.px4_msgs__msg__SensorCorrection__Sequence

func SensorCorrection__Sequence_to_Go(goSlice *[]SensorCorrection, cSlice CSensorCorrection__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SensorCorrection, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__SensorCorrection__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__SensorCorrection * uintptr(i)),
		))
		(*goSlice)[i] = SensorCorrection{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func SensorCorrection__Sequence_to_C(cSlice *CSensorCorrection__Sequence, goSlice []SensorCorrection) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__SensorCorrection)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__SensorCorrection * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__SensorCorrection)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__SensorCorrection * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__SensorCorrection)(v.AsCStruct())
	}
}
func SensorCorrection__Array_to_Go(goSlice []SensorCorrection, cSlice []CSensorCorrection) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func SensorCorrection__Array_to_C(cSlice []CSensorCorrection, goSlice []SensorCorrection) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__SensorCorrection)(goSlice[i].AsCStruct())
	}
}


