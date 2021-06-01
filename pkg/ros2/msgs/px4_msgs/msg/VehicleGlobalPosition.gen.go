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
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/vehicle_global_position.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/VehicleGlobalPosition", &VehicleGlobalPosition{})
}

// Do not create instances of this type directly. Always use NewVehicleGlobalPosition
// function instead.
type VehicleGlobalPosition struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp of the raw data (microseconds)
	Lat float64 `yaml:"lat"`// Latitude, (degrees)
	Lon float64 `yaml:"lon"`// Longitude, (degrees)
	Alt float32 `yaml:"alt"`// Altitude AMSL, (meters)
	AltEllipsoid float32 `yaml:"alt_ellipsoid"`// Altitude above ellipsoid, (meters)
	DeltaAlt float32 `yaml:"delta_alt"`// Reset delta for altitude
	LatLonResetCounter uint8 `yaml:"lat_lon_reset_counter"`// Counter for reset events on horizontal position coordinates
	AltResetCounter uint8 `yaml:"alt_reset_counter"`// Counter for reset events on altitude
	Eph float32 `yaml:"eph"`// Standard deviation of horizontal position error, (metres)
	Epv float32 `yaml:"epv"`// Standard deviation of vertical position error, (metres)
	TerrainAlt float32 `yaml:"terrain_alt"`// Terrain altitude WGS84, (metres)
	TerrainAltValid bool `yaml:"terrain_alt_valid"`// Terrain altitude estimate is valid
	DeadReckoning bool `yaml:"dead_reckoning"`// True if this position is estimated through dead-reckoning
}

// NewVehicleGlobalPosition creates a new VehicleGlobalPosition with default values.
func NewVehicleGlobalPosition() *VehicleGlobalPosition {
	self := VehicleGlobalPosition{}
	self.SetDefaults(nil)
	return &self
}

func (t *VehicleGlobalPosition) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *VehicleGlobalPosition) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleGlobalPosition())
}
func (t *VehicleGlobalPosition) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleGlobalPosition
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleGlobalPosition__create())
}
func (t *VehicleGlobalPosition) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleGlobalPosition__destroy((*C.px4_msgs__msg__VehicleGlobalPosition)(pointer_to_free))
}
func (t *VehicleGlobalPosition) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__VehicleGlobalPosition)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.timestamp_sample = C.uint64_t(t.TimestampSample)
	mem.lat = C.double(t.Lat)
	mem.lon = C.double(t.Lon)
	mem.alt = C.float(t.Alt)
	mem.alt_ellipsoid = C.float(t.AltEllipsoid)
	mem.delta_alt = C.float(t.DeltaAlt)
	mem.lat_lon_reset_counter = C.uint8_t(t.LatLonResetCounter)
	mem.alt_reset_counter = C.uint8_t(t.AltResetCounter)
	mem.eph = C.float(t.Eph)
	mem.epv = C.float(t.Epv)
	mem.terrain_alt = C.float(t.TerrainAlt)
	mem.terrain_alt_valid = C.bool(t.TerrainAltValid)
	mem.dead_reckoning = C.bool(t.DeadReckoning)
	return unsafe.Pointer(mem)
}
func (t *VehicleGlobalPosition) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__VehicleGlobalPosition)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.TimestampSample = uint64(mem.timestamp_sample)
	t.Lat = float64(mem.lat)
	t.Lon = float64(mem.lon)
	t.Alt = float32(mem.alt)
	t.AltEllipsoid = float32(mem.alt_ellipsoid)
	t.DeltaAlt = float32(mem.delta_alt)
	t.LatLonResetCounter = uint8(mem.lat_lon_reset_counter)
	t.AltResetCounter = uint8(mem.alt_reset_counter)
	t.Eph = float32(mem.eph)
	t.Epv = float32(mem.epv)
	t.TerrainAlt = float32(mem.terrain_alt)
	t.TerrainAltValid = bool(mem.terrain_alt_valid)
	t.DeadReckoning = bool(mem.dead_reckoning)
}
func (t *VehicleGlobalPosition) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CVehicleGlobalPosition = C.px4_msgs__msg__VehicleGlobalPosition
type CVehicleGlobalPosition__Sequence = C.px4_msgs__msg__VehicleGlobalPosition__Sequence

func VehicleGlobalPosition__Sequence_to_Go(goSlice *[]VehicleGlobalPosition, cSlice CVehicleGlobalPosition__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleGlobalPosition, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleGlobalPosition__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleGlobalPosition * uintptr(i)),
		))
		(*goSlice)[i] = VehicleGlobalPosition{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func VehicleGlobalPosition__Sequence_to_C(cSlice *CVehicleGlobalPosition__Sequence, goSlice []VehicleGlobalPosition) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleGlobalPosition)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleGlobalPosition * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleGlobalPosition)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleGlobalPosition * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__VehicleGlobalPosition)(v.AsCStruct())
	}
}
func VehicleGlobalPosition__Array_to_Go(goSlice []VehicleGlobalPosition, cSlice []CVehicleGlobalPosition) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleGlobalPosition__Array_to_C(cSlice []CVehicleGlobalPosition, goSlice []VehicleGlobalPosition) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__VehicleGlobalPosition)(goSlice[i].AsCStruct())
	}
}


