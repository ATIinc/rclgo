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
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/point32.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("geometry_msgs/Point32", &Point32{})
}

// Do not create instances of this type directly. Always use NewPoint32
// function instead.
type Point32 struct {
	X float32 `yaml:"x"`
	Y float32 `yaml:"y"`
	Z float32 `yaml:"z"`
}

// NewPoint32 creates a new Point32 with default values.
func NewPoint32() *Point32 {
	self := Point32{}
	self.SetDefaults(nil)
	return &self
}

func (t *Point32) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *Point32) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__Point32())
}
func (t *Point32) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__Point32
	return (unsafe.Pointer)(C.geometry_msgs__msg__Point32__create())
}
func (t *Point32) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__Point32__destroy((*C.geometry_msgs__msg__Point32)(pointer_to_free))
}
func (t *Point32) AsCStruct() unsafe.Pointer {
	mem := (*C.geometry_msgs__msg__Point32)(t.PrepareMemory())
	mem.x = C.float(t.X)
	mem.y = C.float(t.Y)
	mem.z = C.float(t.Z)
	return unsafe.Pointer(mem)
}
func (t *Point32) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.geometry_msgs__msg__Point32)(ros2_message_buffer)
	t.X = float32(mem.x)
	t.Y = float32(mem.y)
	t.Z = float32(mem.z)
}
func (t *Point32) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CPoint32 = C.geometry_msgs__msg__Point32
type CPoint32__Sequence = C.geometry_msgs__msg__Point32__Sequence

func Point32__Sequence_to_Go(goSlice *[]Point32, cSlice CPoint32__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Point32, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__Point32__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Point32 * uintptr(i)),
		))
		(*goSlice)[i] = Point32{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func Point32__Sequence_to_C(cSlice *CPoint32__Sequence, goSlice []Point32) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__Point32)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__Point32 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__Point32)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Point32 * uintptr(i)),
		))
		*cIdx = *(*C.geometry_msgs__msg__Point32)(v.AsCStruct())
	}
}
func Point32__Array_to_Go(goSlice []Point32, cSlice []CPoint32) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func Point32__Array_to_C(cSlice []CPoint32, goSlice []Point32) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.geometry_msgs__msg__Point32)(goSlice[i].AsCStruct())
	}
}


