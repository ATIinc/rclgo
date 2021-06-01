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

#include <geometry_msgs/msg/accel.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("geometry_msgs/Accel", &Accel{})
}

// Do not create instances of this type directly. Always use NewAccel
// function instead.
type Accel struct {
	Linear Vector3 `yaml:"linear"`// This expresses acceleration in free space broken into its linear and angular parts.
	Angular Vector3 `yaml:"angular"`// This expresses acceleration in free space broken into its linear and angular parts.
}

// NewAccel creates a new Accel with default values.
func NewAccel() *Accel {
	self := Accel{}
	self.SetDefaults(nil)
	return &self
}

func (t *Accel) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Linear.SetDefaults(nil)
	t.Angular.SetDefaults(nil)
	
	return t
}

func (t *Accel) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__Accel())
}
func (t *Accel) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__Accel
	return (unsafe.Pointer)(C.geometry_msgs__msg__Accel__create())
}
func (t *Accel) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__Accel__destroy((*C.geometry_msgs__msg__Accel)(pointer_to_free))
}
func (t *Accel) AsCStruct() unsafe.Pointer {
	mem := (*C.geometry_msgs__msg__Accel)(t.PrepareMemory())
	mem.linear = *(*C.geometry_msgs__msg__Vector3)(t.Linear.AsCStruct())
	mem.angular = *(*C.geometry_msgs__msg__Vector3)(t.Angular.AsCStruct())
	return unsafe.Pointer(mem)
}
func (t *Accel) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.geometry_msgs__msg__Accel)(ros2_message_buffer)
	t.Linear.AsGoStruct(unsafe.Pointer(&mem.linear))
	t.Angular.AsGoStruct(unsafe.Pointer(&mem.angular))
}
func (t *Accel) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CAccel = C.geometry_msgs__msg__Accel
type CAccel__Sequence = C.geometry_msgs__msg__Accel__Sequence

func Accel__Sequence_to_Go(goSlice *[]Accel, cSlice CAccel__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Accel, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__Accel__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Accel * uintptr(i)),
		))
		(*goSlice)[i] = Accel{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func Accel__Sequence_to_C(cSlice *CAccel__Sequence, goSlice []Accel) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__Accel)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__Accel * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__Accel)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Accel * uintptr(i)),
		))
		*cIdx = *(*C.geometry_msgs__msg__Accel)(v.AsCStruct())
	}
}
func Accel__Array_to_Go(goSlice []Accel, cSlice []CAccel) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func Accel__Array_to_C(cSlice []CAccel, goSlice []Accel) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.geometry_msgs__msg__Accel)(goSlice[i].AsCStruct())
	}
}


