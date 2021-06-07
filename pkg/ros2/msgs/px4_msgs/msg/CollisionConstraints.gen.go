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
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/collision_constraints.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/CollisionConstraints", CollisionConstraintsTypeSupport)
}

// Do not create instances of this type directly. Always use NewCollisionConstraints
// function instead.
type CollisionConstraints struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	OriginalSetpoint [2]float32 `yaml:"original_setpoint"`// velocities demanded
	AdaptedSetpoint [2]float32 `yaml:"adapted_setpoint"`// velocities allowed
}

// NewCollisionConstraints creates a new CollisionConstraints with default values.
func NewCollisionConstraints() *CollisionConstraints {
	self := CollisionConstraints{}
	self.SetDefaults()
	return &self
}

func (t *CollisionConstraints) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *CollisionConstraints) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var CollisionConstraintsTypeSupport types.MessageTypeSupport = _CollisionConstraintsTypeSupport{}

type _CollisionConstraintsTypeSupport struct{}

func (t _CollisionConstraintsTypeSupport) New() types.Message {
	return NewCollisionConstraints()
}

func (t _CollisionConstraintsTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__CollisionConstraints
	return (unsafe.Pointer)(C.px4_msgs__msg__CollisionConstraints__create())
}

func (t _CollisionConstraintsTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__CollisionConstraints__destroy((*C.px4_msgs__msg__CollisionConstraints)(pointer_to_free))
}

func (t _CollisionConstraintsTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*CollisionConstraints)
	mem := (*C.px4_msgs__msg__CollisionConstraints)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	cSlice_original_setpoint := mem.original_setpoint[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_original_setpoint)), m.OriginalSetpoint[:])
	cSlice_adapted_setpoint := mem.adapted_setpoint[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_adapted_setpoint)), m.AdaptedSetpoint[:])
}

func (t _CollisionConstraintsTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*CollisionConstraints)
	mem := (*C.px4_msgs__msg__CollisionConstraints)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	cSlice_original_setpoint := mem.original_setpoint[:]
	rosidl_runtime_c.Float32__Array_to_Go(m.OriginalSetpoint[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_original_setpoint)))
	cSlice_adapted_setpoint := mem.adapted_setpoint[:]
	rosidl_runtime_c.Float32__Array_to_Go(m.AdaptedSetpoint[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_adapted_setpoint)))
}

func (t _CollisionConstraintsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__CollisionConstraints())
}

type CCollisionConstraints = C.px4_msgs__msg__CollisionConstraints
type CCollisionConstraints__Sequence = C.px4_msgs__msg__CollisionConstraints__Sequence

func CollisionConstraints__Sequence_to_Go(goSlice *[]CollisionConstraints, cSlice CCollisionConstraints__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]CollisionConstraints, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__CollisionConstraints__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__CollisionConstraints * uintptr(i)),
		))
		CollisionConstraintsTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func CollisionConstraints__Sequence_to_C(cSlice *CCollisionConstraints__Sequence, goSlice []CollisionConstraints) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__CollisionConstraints)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__CollisionConstraints * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__CollisionConstraints)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__CollisionConstraints * uintptr(i)),
		))
		CollisionConstraintsTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func CollisionConstraints__Array_to_Go(goSlice []CollisionConstraints, cSlice []CCollisionConstraints) {
	for i := 0; i < len(cSlice); i++ {
		CollisionConstraintsTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func CollisionConstraints__Array_to_C(cSlice []CCollisionConstraints, goSlice []CollisionConstraints) {
	for i := 0; i < len(goSlice); i++ {
		CollisionConstraintsTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
