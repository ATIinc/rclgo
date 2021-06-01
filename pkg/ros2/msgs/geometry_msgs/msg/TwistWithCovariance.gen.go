/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package geometry_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/twist_with_covariance.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("geometry_msgs/TwistWithCovariance", &TwistWithCovariance{})
}

// Do not create instances of this type directly. Always use NewTwistWithCovariance
// function instead.
type TwistWithCovariance struct {
	Twist Twist `yaml:"twist"`
	Covariance [36]float64 `yaml:"covariance"`// Row-major representation of the 6x6 covariance matrixThe orientation parameters use a fixed-axis representation.In order, the parameters are:(x, y, z, rotation about X axis, rotation about Y axis, rotation about Z axis)
}

// NewTwistWithCovariance creates a new TwistWithCovariance with default values.
func NewTwistWithCovariance() *TwistWithCovariance {
	self := TwistWithCovariance{}
	self.SetDefaults(nil)
	return &self
}

func (t *TwistWithCovariance) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Twist.SetDefaults(nil)
	
	return t
}

func (t *TwistWithCovariance) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__TwistWithCovariance())
}
func (t *TwistWithCovariance) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__TwistWithCovariance
	return (unsafe.Pointer)(C.geometry_msgs__msg__TwistWithCovariance__create())
}
func (t *TwistWithCovariance) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__TwistWithCovariance__destroy((*C.geometry_msgs__msg__TwistWithCovariance)(pointer_to_free))
}
func (t *TwistWithCovariance) AsCStruct() unsafe.Pointer {
	mem := (*C.geometry_msgs__msg__TwistWithCovariance)(t.PrepareMemory())
	mem.twist = *(*C.geometry_msgs__msg__Twist)(t.Twist.AsCStruct())
	cSlice_covariance := mem.covariance[:]
	rosidl_runtime_c.Float64__Array_to_C(*(*[]rosidl_runtime_c.CFloat64)(unsafe.Pointer(&cSlice_covariance)), t.Covariance[:])
	return unsafe.Pointer(mem)
}
func (t *TwistWithCovariance) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.geometry_msgs__msg__TwistWithCovariance)(ros2_message_buffer)
	t.Twist.AsGoStruct(unsafe.Pointer(&mem.twist))
	cSlice_covariance := mem.covariance[:]
	rosidl_runtime_c.Float64__Array_to_Go(t.Covariance[:], *(*[]rosidl_runtime_c.CFloat64)(unsafe.Pointer(&cSlice_covariance)))
}
func (t *TwistWithCovariance) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CTwistWithCovariance = C.geometry_msgs__msg__TwistWithCovariance
type CTwistWithCovariance__Sequence = C.geometry_msgs__msg__TwistWithCovariance__Sequence

func TwistWithCovariance__Sequence_to_Go(goSlice *[]TwistWithCovariance, cSlice CTwistWithCovariance__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TwistWithCovariance, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__TwistWithCovariance__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__TwistWithCovariance * uintptr(i)),
		))
		(*goSlice)[i] = TwistWithCovariance{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func TwistWithCovariance__Sequence_to_C(cSlice *CTwistWithCovariance__Sequence, goSlice []TwistWithCovariance) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__TwistWithCovariance)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__TwistWithCovariance * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__TwistWithCovariance)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__TwistWithCovariance * uintptr(i)),
		))
		*cIdx = *(*C.geometry_msgs__msg__TwistWithCovariance)(v.AsCStruct())
	}
}
func TwistWithCovariance__Array_to_Go(goSlice []TwistWithCovariance, cSlice []CTwistWithCovariance) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func TwistWithCovariance__Array_to_C(cSlice []CTwistWithCovariance, goSlice []TwistWithCovariance) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.geometry_msgs__msg__TwistWithCovariance)(goSlice[i].AsCStruct())
	}
}


