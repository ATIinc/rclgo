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

#include <px4_msgs/msg/orb_test_medium.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/OrbTestMedium", &OrbTestMedium{})
}

// Do not create instances of this type directly. Always use NewOrbTestMedium
// function instead.
type OrbTestMedium struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Val int32 `yaml:"val"`
	Junk [64]uint8 `yaml:"junk"`
}

// NewOrbTestMedium creates a new OrbTestMedium with default values.
func NewOrbTestMedium() *OrbTestMedium {
	self := OrbTestMedium{}
	self.SetDefaults(nil)
	return &self
}

func (t *OrbTestMedium) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *OrbTestMedium) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__OrbTestMedium())
}
func (t *OrbTestMedium) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__OrbTestMedium
	return (unsafe.Pointer)(C.px4_msgs__msg__OrbTestMedium__create())
}
func (t *OrbTestMedium) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__OrbTestMedium__destroy((*C.px4_msgs__msg__OrbTestMedium)(pointer_to_free))
}
func (t *OrbTestMedium) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__OrbTestMedium)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.val = C.int32_t(t.Val)
	cSlice_junk := mem.junk[:]
	rosidl_runtime_c.Uint8__Array_to_C(*(*[]rosidl_runtime_c.CUint8)(unsafe.Pointer(&cSlice_junk)), t.Junk[:])
	return unsafe.Pointer(mem)
}
func (t *OrbTestMedium) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__OrbTestMedium)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.Val = int32(mem.val)
	cSlice_junk := mem.junk[:]
	rosidl_runtime_c.Uint8__Array_to_Go(t.Junk[:], *(*[]rosidl_runtime_c.CUint8)(unsafe.Pointer(&cSlice_junk)))
}
func (t *OrbTestMedium) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type COrbTestMedium = C.px4_msgs__msg__OrbTestMedium
type COrbTestMedium__Sequence = C.px4_msgs__msg__OrbTestMedium__Sequence

func OrbTestMedium__Sequence_to_Go(goSlice *[]OrbTestMedium, cSlice COrbTestMedium__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]OrbTestMedium, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__OrbTestMedium__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__OrbTestMedium * uintptr(i)),
		))
		(*goSlice)[i] = OrbTestMedium{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func OrbTestMedium__Sequence_to_C(cSlice *COrbTestMedium__Sequence, goSlice []OrbTestMedium) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__OrbTestMedium)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__OrbTestMedium * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__OrbTestMedium)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__OrbTestMedium * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__OrbTestMedium)(v.AsCStruct())
	}
}
func OrbTestMedium__Array_to_Go(goSlice []OrbTestMedium, cSlice []COrbTestMedium) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func OrbTestMedium__Array_to_C(cSlice []COrbTestMedium, goSlice []OrbTestMedium) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__OrbTestMedium)(goSlice[i].AsCStruct())
	}
}


