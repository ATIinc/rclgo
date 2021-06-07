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
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/twist.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("geometry_msgs/Twist", TwistTypeSupport)
}

// Do not create instances of this type directly. Always use NewTwist
// function instead.
type Twist struct {
	Linear Vector3 `yaml:"linear"`
	Angular Vector3 `yaml:"angular"`
}

// NewTwist creates a new Twist with default values.
func NewTwist() *Twist {
	self := Twist{}
	self.SetDefaults()
	return &self
}

func (t *Twist) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *Twist) SetDefaults() {
	t.Linear.SetDefaults()
	t.Angular.SetDefaults()
	
}

// Modifying this variable is undefined behavior.
var TwistTypeSupport types.MessageTypeSupport = _TwistTypeSupport{}

type _TwistTypeSupport struct{}

func (t _TwistTypeSupport) New() types.Message {
	return NewTwist()
}

func (t _TwistTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__Twist
	return (unsafe.Pointer)(C.geometry_msgs__msg__Twist__create())
}

func (t _TwistTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__Twist__destroy((*C.geometry_msgs__msg__Twist)(pointer_to_free))
}

func (t _TwistTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Twist)
	mem := (*C.geometry_msgs__msg__Twist)(dst)
	Vector3TypeSupport.AsCStruct(unsafe.Pointer(&mem.linear), &m.Linear)
	Vector3TypeSupport.AsCStruct(unsafe.Pointer(&mem.angular), &m.Angular)
}

func (t _TwistTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Twist)
	mem := (*C.geometry_msgs__msg__Twist)(ros2_message_buffer)
	Vector3TypeSupport.AsGoStruct(&m.Linear, unsafe.Pointer(&mem.linear))
	Vector3TypeSupport.AsGoStruct(&m.Angular, unsafe.Pointer(&mem.angular))
}

func (t _TwistTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__Twist())
}

type CTwist = C.geometry_msgs__msg__Twist
type CTwist__Sequence = C.geometry_msgs__msg__Twist__Sequence

func Twist__Sequence_to_Go(goSlice *[]Twist, cSlice CTwist__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Twist, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__Twist__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Twist * uintptr(i)),
		))
		TwistTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Twist__Sequence_to_C(cSlice *CTwist__Sequence, goSlice []Twist) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__Twist)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__Twist * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__Twist)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Twist * uintptr(i)),
		))
		TwistTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Twist__Array_to_Go(goSlice []Twist, cSlice []CTwist) {
	for i := 0; i < len(cSlice); i++ {
		TwistTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Twist__Array_to_C(cSlice []CTwist, goSlice []Twist) {
	for i := 0; i < len(goSlice); i++ {
		TwistTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
