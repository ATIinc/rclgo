/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package map_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	map_msgs_msg "github.com/tiiuae/rclgo/pkg/ros2/msgs/map_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lmap_msgs__rosidl_typesupport_c -lmap_msgs__rosidl_generator_c
#cgo LDFLAGS: -lmap_msgs__rosidl_typesupport_c -lmap_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <map_msgs/srv/set_map_projections.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("map_msgs/SetMapProjections_Response", &SetMapProjections_Response{})
}

// Do not create instances of this type directly. Always use NewSetMapProjections_Response
// function instead.
type SetMapProjections_Response struct {
	ProjectedMapsInfo []map_msgs_msg.ProjectedMapInfo `yaml:"projected_maps_info"`
}

// NewSetMapProjections_Response creates a new SetMapProjections_Response with default values.
func NewSetMapProjections_Response() *SetMapProjections_Response {
	self := SetMapProjections_Response{}
	self.SetDefaults(nil)
	return &self
}

func (t *SetMapProjections_Response) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *SetMapProjections_Response) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__map_msgs__srv__SetMapProjections_Response())
}
func (t *SetMapProjections_Response) PrepareMemory() unsafe.Pointer { //returns *C.map_msgs__srv__SetMapProjections_Response
	return (unsafe.Pointer)(C.map_msgs__srv__SetMapProjections_Response__create())
}
func (t *SetMapProjections_Response) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.map_msgs__srv__SetMapProjections_Response__destroy((*C.map_msgs__srv__SetMapProjections_Response)(pointer_to_free))
}
func (t *SetMapProjections_Response) AsCStruct() unsafe.Pointer {
	mem := (*C.map_msgs__srv__SetMapProjections_Response)(t.PrepareMemory())
	map_msgs_msg.ProjectedMapInfo__Sequence_to_C((*map_msgs_msg.CProjectedMapInfo__Sequence)(unsafe.Pointer(&mem.projected_maps_info)), t.ProjectedMapsInfo)
	return unsafe.Pointer(mem)
}
func (t *SetMapProjections_Response) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.map_msgs__srv__SetMapProjections_Response)(ros2_message_buffer)
	map_msgs_msg.ProjectedMapInfo__Sequence_to_Go(&t.ProjectedMapsInfo, *(*map_msgs_msg.CProjectedMapInfo__Sequence)(unsafe.Pointer(&mem.projected_maps_info)))
}
func (t *SetMapProjections_Response) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CSetMapProjections_Response = C.map_msgs__srv__SetMapProjections_Response
type CSetMapProjections_Response__Sequence = C.map_msgs__srv__SetMapProjections_Response__Sequence

func SetMapProjections_Response__Sequence_to_Go(goSlice *[]SetMapProjections_Response, cSlice CSetMapProjections_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SetMapProjections_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.map_msgs__srv__SetMapProjections_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_map_msgs__srv__SetMapProjections_Response * uintptr(i)),
		))
		(*goSlice)[i] = SetMapProjections_Response{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func SetMapProjections_Response__Sequence_to_C(cSlice *CSetMapProjections_Response__Sequence, goSlice []SetMapProjections_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.map_msgs__srv__SetMapProjections_Response)(C.malloc((C.size_t)(C.sizeof_struct_map_msgs__srv__SetMapProjections_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.map_msgs__srv__SetMapProjections_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_map_msgs__srv__SetMapProjections_Response * uintptr(i)),
		))
		*cIdx = *(*C.map_msgs__srv__SetMapProjections_Response)(v.AsCStruct())
	}
}
func SetMapProjections_Response__Array_to_Go(goSlice []SetMapProjections_Response, cSlice []CSetMapProjections_Response) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func SetMapProjections_Response__Array_to_C(cSlice []CSetMapProjections_Response, goSlice []SetMapProjections_Response) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.map_msgs__srv__SetMapProjections_Response)(goSlice[i].AsCStruct())
	}
}


