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

#include <px4_msgs/msg/debug_vect.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/DebugVect", &DebugVect{})
}

// Do not create instances of this type directly. Always use NewDebugVect
// function instead.
type DebugVect struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Name [10]byte `yaml:"name"`// max. 10 characters as key / name
	X float32 `yaml:"x"`// x value
	Y float32 `yaml:"y"`// y value
	Z float32 `yaml:"z"`// z value
}

// NewDebugVect creates a new DebugVect with default values.
func NewDebugVect() *DebugVect {
	self := DebugVect{}
	self.SetDefaults(nil)
	return &self
}

func (t *DebugVect) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *DebugVect) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__DebugVect())
}
func (t *DebugVect) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__DebugVect
	return (unsafe.Pointer)(C.px4_msgs__msg__DebugVect__create())
}
func (t *DebugVect) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__DebugVect__destroy((*C.px4_msgs__msg__DebugVect)(pointer_to_free))
}
func (t *DebugVect) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__DebugVect)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	cSlice_name := mem.name[:]
	rosidl_runtime_c.Char__Array_to_C(*(*[]rosidl_runtime_c.CChar)(unsafe.Pointer(&cSlice_name)), t.Name[:])
	mem.x = C.float(t.X)
	mem.y = C.float(t.Y)
	mem.z = C.float(t.Z)
	return unsafe.Pointer(mem)
}
func (t *DebugVect) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__DebugVect)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	cSlice_name := mem.name[:]
	rosidl_runtime_c.Char__Array_to_Go(t.Name[:], *(*[]rosidl_runtime_c.CChar)(unsafe.Pointer(&cSlice_name)))
	t.X = float32(mem.x)
	t.Y = float32(mem.y)
	t.Z = float32(mem.z)
}
func (t *DebugVect) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CDebugVect = C.px4_msgs__msg__DebugVect
type CDebugVect__Sequence = C.px4_msgs__msg__DebugVect__Sequence

func DebugVect__Sequence_to_Go(goSlice *[]DebugVect, cSlice CDebugVect__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]DebugVect, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__DebugVect__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__DebugVect * uintptr(i)),
		))
		(*goSlice)[i] = DebugVect{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func DebugVect__Sequence_to_C(cSlice *CDebugVect__Sequence, goSlice []DebugVect) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__DebugVect)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__DebugVect * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__DebugVect)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__DebugVect * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__DebugVect)(v.AsCStruct())
	}
}
func DebugVect__Array_to_Go(goSlice []DebugVect, cSlice []CDebugVect) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func DebugVect__Array_to_C(cSlice []CDebugVect, goSlice []DebugVect) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__DebugVect)(goSlice[i].AsCStruct())
	}
}


