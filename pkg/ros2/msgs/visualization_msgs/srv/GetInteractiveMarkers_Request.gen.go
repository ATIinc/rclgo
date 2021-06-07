/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package visualization_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lvisualization_msgs__rosidl_typesupport_c -lvisualization_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <visualization_msgs/srv/get_interactive_markers.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("visualization_msgs/GetInteractiveMarkers_Request", GetInteractiveMarkers_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewGetInteractiveMarkers_Request
// function instead.
type GetInteractiveMarkers_Request struct {
}

// NewGetInteractiveMarkers_Request creates a new GetInteractiveMarkers_Request with default values.
func NewGetInteractiveMarkers_Request() *GetInteractiveMarkers_Request {
	self := GetInteractiveMarkers_Request{}
	self.SetDefaults()
	return &self
}

func (t *GetInteractiveMarkers_Request) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *GetInteractiveMarkers_Request) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var GetInteractiveMarkers_RequestTypeSupport types.MessageTypeSupport = _GetInteractiveMarkers_RequestTypeSupport{}

type _GetInteractiveMarkers_RequestTypeSupport struct{}

func (t _GetInteractiveMarkers_RequestTypeSupport) New() types.Message {
	return NewGetInteractiveMarkers_Request()
}

func (t _GetInteractiveMarkers_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.visualization_msgs__srv__GetInteractiveMarkers_Request
	return (unsafe.Pointer)(C.visualization_msgs__srv__GetInteractiveMarkers_Request__create())
}

func (t _GetInteractiveMarkers_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.visualization_msgs__srv__GetInteractiveMarkers_Request__destroy((*C.visualization_msgs__srv__GetInteractiveMarkers_Request)(pointer_to_free))
}

func (t _GetInteractiveMarkers_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _GetInteractiveMarkers_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _GetInteractiveMarkers_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__visualization_msgs__srv__GetInteractiveMarkers_Request())
}

type CGetInteractiveMarkers_Request = C.visualization_msgs__srv__GetInteractiveMarkers_Request
type CGetInteractiveMarkers_Request__Sequence = C.visualization_msgs__srv__GetInteractiveMarkers_Request__Sequence

func GetInteractiveMarkers_Request__Sequence_to_Go(goSlice *[]GetInteractiveMarkers_Request, cSlice CGetInteractiveMarkers_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GetInteractiveMarkers_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.visualization_msgs__srv__GetInteractiveMarkers_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__srv__GetInteractiveMarkers_Request * uintptr(i)),
		))
		GetInteractiveMarkers_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GetInteractiveMarkers_Request__Sequence_to_C(cSlice *CGetInteractiveMarkers_Request__Sequence, goSlice []GetInteractiveMarkers_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.visualization_msgs__srv__GetInteractiveMarkers_Request)(C.malloc((C.size_t)(C.sizeof_struct_visualization_msgs__srv__GetInteractiveMarkers_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.visualization_msgs__srv__GetInteractiveMarkers_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__srv__GetInteractiveMarkers_Request * uintptr(i)),
		))
		GetInteractiveMarkers_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GetInteractiveMarkers_Request__Array_to_Go(goSlice []GetInteractiveMarkers_Request, cSlice []CGetInteractiveMarkers_Request) {
	for i := 0; i < len(cSlice); i++ {
		GetInteractiveMarkers_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GetInteractiveMarkers_Request__Array_to_C(cSlice []CGetInteractiveMarkers_Request, goSlice []GetInteractiveMarkers_Request) {
	for i := 0; i < len(goSlice); i++ {
		GetInteractiveMarkers_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
