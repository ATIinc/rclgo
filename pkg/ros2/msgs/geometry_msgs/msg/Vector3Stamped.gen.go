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

	"github.com/tiiuae/rclgo/pkg/ros2/types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	std_msgs_msg "github.com/tiiuae/rclgo/pkg/ros2/msgs/std_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/vector3_stamped.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("geometry_msgs/Vector3Stamped", Vector3StampedTypeSupport)
}

// Do not create instances of this type directly. Always use NewVector3Stamped
// function instead.
type Vector3Stamped struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Vector Vector3 `yaml:"vector"`
}

// NewVector3Stamped creates a new Vector3Stamped with default values.
func NewVector3Stamped() *Vector3Stamped {
	self := Vector3Stamped{}
	self.SetDefaults()
	return &self
}

func (t *Vector3Stamped) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *Vector3Stamped) SetDefaults() {
	t.Header.SetDefaults()
	t.Vector.SetDefaults()
	
}

// Modifying this variable is undefined behavior.
var Vector3StampedTypeSupport types.MessageTypeSupport = _Vector3StampedTypeSupport{}

type _Vector3StampedTypeSupport struct{}

func (t _Vector3StampedTypeSupport) New() types.Message {
	return NewVector3Stamped()
}

func (t _Vector3StampedTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__Vector3Stamped
	return (unsafe.Pointer)(C.geometry_msgs__msg__Vector3Stamped__create())
}

func (t _Vector3StampedTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__Vector3Stamped__destroy((*C.geometry_msgs__msg__Vector3Stamped)(pointer_to_free))
}

func (t _Vector3StampedTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Vector3Stamped)
	mem := (*C.geometry_msgs__msg__Vector3Stamped)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	Vector3TypeSupport.AsCStruct(unsafe.Pointer(&mem.vector), &m.Vector)
}

func (t _Vector3StampedTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Vector3Stamped)
	mem := (*C.geometry_msgs__msg__Vector3Stamped)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	Vector3TypeSupport.AsGoStruct(&m.Vector, unsafe.Pointer(&mem.vector))
}

func (t _Vector3StampedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__Vector3Stamped())
}

type CVector3Stamped = C.geometry_msgs__msg__Vector3Stamped
type CVector3Stamped__Sequence = C.geometry_msgs__msg__Vector3Stamped__Sequence

func Vector3Stamped__Sequence_to_Go(goSlice *[]Vector3Stamped, cSlice CVector3Stamped__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Vector3Stamped, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__Vector3Stamped__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Vector3Stamped * uintptr(i)),
		))
		Vector3StampedTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Vector3Stamped__Sequence_to_C(cSlice *CVector3Stamped__Sequence, goSlice []Vector3Stamped) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__Vector3Stamped)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__Vector3Stamped * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__Vector3Stamped)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Vector3Stamped * uintptr(i)),
		))
		Vector3StampedTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Vector3Stamped__Array_to_Go(goSlice []Vector3Stamped, cSlice []CVector3Stamped) {
	for i := 0; i < len(cSlice); i++ {
		Vector3StampedTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Vector3Stamped__Array_to_C(cSlice []CVector3Stamped, goSlice []Vector3Stamped) {
	for i := 0; i < len(goSlice); i++ {
		Vector3StampedTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
