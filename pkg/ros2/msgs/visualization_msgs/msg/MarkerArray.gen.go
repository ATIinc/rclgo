/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package visualization_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lvisualization_msgs__rosidl_typesupport_c -lvisualization_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <visualization_msgs/msg/marker_array.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("visualization_msgs/MarkerArray", &MarkerArray{})
}

// Do not create instances of this type directly. Always use NewMarkerArray
// function instead.
type MarkerArray struct {
	Markers []Marker `yaml:"markers"`
}

// NewMarkerArray creates a new MarkerArray with default values.
func NewMarkerArray() *MarkerArray {
	self := MarkerArray{}
	self.SetDefaults(nil)
	return &self
}

func (t *MarkerArray) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *MarkerArray) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__visualization_msgs__msg__MarkerArray())
}
func (t *MarkerArray) PrepareMemory() unsafe.Pointer { //returns *C.visualization_msgs__msg__MarkerArray
	return (unsafe.Pointer)(C.visualization_msgs__msg__MarkerArray__create())
}
func (t *MarkerArray) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.visualization_msgs__msg__MarkerArray__destroy((*C.visualization_msgs__msg__MarkerArray)(pointer_to_free))
}
func (t *MarkerArray) AsCStruct() unsafe.Pointer {
	mem := (*C.visualization_msgs__msg__MarkerArray)(t.PrepareMemory())
	Marker__Sequence_to_C(&mem.markers, t.Markers)
	return unsafe.Pointer(mem)
}
func (t *MarkerArray) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.visualization_msgs__msg__MarkerArray)(ros2_message_buffer)
	Marker__Sequence_to_Go(&t.Markers, mem.markers)
}
func (t *MarkerArray) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CMarkerArray = C.visualization_msgs__msg__MarkerArray
type CMarkerArray__Sequence = C.visualization_msgs__msg__MarkerArray__Sequence

func MarkerArray__Sequence_to_Go(goSlice *[]MarkerArray, cSlice CMarkerArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MarkerArray, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.visualization_msgs__msg__MarkerArray__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__MarkerArray * uintptr(i)),
		))
		(*goSlice)[i] = MarkerArray{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func MarkerArray__Sequence_to_C(cSlice *CMarkerArray__Sequence, goSlice []MarkerArray) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.visualization_msgs__msg__MarkerArray)(C.malloc((C.size_t)(C.sizeof_struct_visualization_msgs__msg__MarkerArray * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.visualization_msgs__msg__MarkerArray)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__MarkerArray * uintptr(i)),
		))
		*cIdx = *(*C.visualization_msgs__msg__MarkerArray)(v.AsCStruct())
	}
}
func MarkerArray__Array_to_Go(goSlice []MarkerArray, cSlice []CMarkerArray) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func MarkerArray__Array_to_C(cSlice []CMarkerArray, goSlice []MarkerArray) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.visualization_msgs__msg__MarkerArray)(goSlice[i].AsCStruct())
	}
}


