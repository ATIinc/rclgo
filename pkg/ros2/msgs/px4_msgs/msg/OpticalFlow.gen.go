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

#include <px4_msgs/msg/optical_flow.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/OpticalFlow", OpticalFlowTypeSupport)
}
const (
	OpticalFlow_MODE_UNKNOWN uint8 = 0
	OpticalFlow_MODE_BRIGHT uint8 = 1
	OpticalFlow_MODE_LOWLIGHT uint8 = 2
	OpticalFlow_MODE_SUPER_LOWLIGHT uint8 = 3
)

// Do not create instances of this type directly. Always use NewOpticalFlow
// function instead.
type OpticalFlow struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	SensorId uint8 `yaml:"sensor_id"`// id of the sensor emitting the flow value
	PixelFlowXIntegral float32 `yaml:"pixel_flow_x_integral"`// accumulated optical flow in radians where a positive value is produced by a RH rotation about the X body axis
	PixelFlowYIntegral float32 `yaml:"pixel_flow_y_integral"`// accumulated optical flow in radians where a positive value is produced by a RH rotation about the Y body axis
	GyroXRateIntegral float32 `yaml:"gyro_x_rate_integral"`// accumulated gyro value in radians where a positive value is produced by a RH rotation about the X body axis. Set to NaN if flow sensor does not have 3-axis gyro data.
	GyroYRateIntegral float32 `yaml:"gyro_y_rate_integral"`// accumulated gyro value in radians where a positive value is produced by a RH rotation about the Y body axis. Set to NaN if flow sensor does not have 3-axis gyro data.
	GyroZRateIntegral float32 `yaml:"gyro_z_rate_integral"`// accumulated gyro value in radians where a positive value is produced by a RH rotation about the Z body axis. Set to NaN if flow sensor does not have 3-axis gyro data.
	GroundDistanceM float32 `yaml:"ground_distance_m"`// Altitude / distance to ground in meters
	IntegrationTimespan uint32 `yaml:"integration_timespan"`// accumulation timespan in microseconds
	TimeSinceLastSonarUpdate uint32 `yaml:"time_since_last_sonar_update"`// time since last sonar update in microseconds
	FrameCountSinceLastReadout uint16 `yaml:"frame_count_since_last_readout"`// number of accumulated frames in timespan
	GyroTemperature int16 `yaml:"gyro_temperature"`// Temperature * 100 in centi-degrees Celsius
	Quality uint8 `yaml:"quality"`// Average of quality of accumulated frames, 0: bad quality, 255: maximum quality
	MaxFlowRate float32 `yaml:"max_flow_rate"`// Magnitude of maximum angular which the optical flow sensor can measure reliably
	MinGroundDistance float32 `yaml:"min_ground_distance"`// Minimum distance from ground at which the optical flow sensor operates reliably
	MaxGroundDistance float32 `yaml:"max_ground_distance"`// Maximum distance from ground at which the optical flow sensor operates reliably
	Mode uint8 `yaml:"mode"`
}

// NewOpticalFlow creates a new OpticalFlow with default values.
func NewOpticalFlow() *OpticalFlow {
	self := OpticalFlow{}
	self.SetDefaults()
	return &self
}

func (t *OpticalFlow) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *OpticalFlow) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var OpticalFlowTypeSupport types.MessageTypeSupport = _OpticalFlowTypeSupport{}

type _OpticalFlowTypeSupport struct{}

func (t _OpticalFlowTypeSupport) New() types.Message {
	return NewOpticalFlow()
}

func (t _OpticalFlowTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__OpticalFlow
	return (unsafe.Pointer)(C.px4_msgs__msg__OpticalFlow__create())
}

func (t _OpticalFlowTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__OpticalFlow__destroy((*C.px4_msgs__msg__OpticalFlow)(pointer_to_free))
}

func (t _OpticalFlowTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*OpticalFlow)
	mem := (*C.px4_msgs__msg__OpticalFlow)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.sensor_id = C.uint8_t(m.SensorId)
	mem.pixel_flow_x_integral = C.float(m.PixelFlowXIntegral)
	mem.pixel_flow_y_integral = C.float(m.PixelFlowYIntegral)
	mem.gyro_x_rate_integral = C.float(m.GyroXRateIntegral)
	mem.gyro_y_rate_integral = C.float(m.GyroYRateIntegral)
	mem.gyro_z_rate_integral = C.float(m.GyroZRateIntegral)
	mem.ground_distance_m = C.float(m.GroundDistanceM)
	mem.integration_timespan = C.uint32_t(m.IntegrationTimespan)
	mem.time_since_last_sonar_update = C.uint32_t(m.TimeSinceLastSonarUpdate)
	mem.frame_count_since_last_readout = C.uint16_t(m.FrameCountSinceLastReadout)
	mem.gyro_temperature = C.int16_t(m.GyroTemperature)
	mem.quality = C.uint8_t(m.Quality)
	mem.max_flow_rate = C.float(m.MaxFlowRate)
	mem.min_ground_distance = C.float(m.MinGroundDistance)
	mem.max_ground_distance = C.float(m.MaxGroundDistance)
	mem.mode = C.uint8_t(m.Mode)
}

func (t _OpticalFlowTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*OpticalFlow)
	mem := (*C.px4_msgs__msg__OpticalFlow)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.SensorId = uint8(mem.sensor_id)
	m.PixelFlowXIntegral = float32(mem.pixel_flow_x_integral)
	m.PixelFlowYIntegral = float32(mem.pixel_flow_y_integral)
	m.GyroXRateIntegral = float32(mem.gyro_x_rate_integral)
	m.GyroYRateIntegral = float32(mem.gyro_y_rate_integral)
	m.GyroZRateIntegral = float32(mem.gyro_z_rate_integral)
	m.GroundDistanceM = float32(mem.ground_distance_m)
	m.IntegrationTimespan = uint32(mem.integration_timespan)
	m.TimeSinceLastSonarUpdate = uint32(mem.time_since_last_sonar_update)
	m.FrameCountSinceLastReadout = uint16(mem.frame_count_since_last_readout)
	m.GyroTemperature = int16(mem.gyro_temperature)
	m.Quality = uint8(mem.quality)
	m.MaxFlowRate = float32(mem.max_flow_rate)
	m.MinGroundDistance = float32(mem.min_ground_distance)
	m.MaxGroundDistance = float32(mem.max_ground_distance)
	m.Mode = uint8(mem.mode)
}

func (t _OpticalFlowTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__OpticalFlow())
}

type COpticalFlow = C.px4_msgs__msg__OpticalFlow
type COpticalFlow__Sequence = C.px4_msgs__msg__OpticalFlow__Sequence

func OpticalFlow__Sequence_to_Go(goSlice *[]OpticalFlow, cSlice COpticalFlow__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]OpticalFlow, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__OpticalFlow__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__OpticalFlow * uintptr(i)),
		))
		OpticalFlowTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func OpticalFlow__Sequence_to_C(cSlice *COpticalFlow__Sequence, goSlice []OpticalFlow) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__OpticalFlow)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__OpticalFlow * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__OpticalFlow)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__OpticalFlow * uintptr(i)),
		))
		OpticalFlowTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func OpticalFlow__Array_to_Go(goSlice []OpticalFlow, cSlice []COpticalFlow) {
	for i := 0; i < len(cSlice); i++ {
		OpticalFlowTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func OpticalFlow__Array_to_C(cSlice []COpticalFlow, goSlice []OpticalFlow) {
	for i := 0; i < len(goSlice); i++ {
		OpticalFlowTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
