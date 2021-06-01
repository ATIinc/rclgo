/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package tf2_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ltf2_msgs__rosidl_typesupport_c -ltf2_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <tf2_msgs/srv/frame_graph.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("tf2_msgs/FrameGraph_Request", &FrameGraph_Request{})
}

// Do not create instances of this type directly. Always use NewFrameGraph_Request
// function instead.
type FrameGraph_Request struct {
}

// NewFrameGraph_Request creates a new FrameGraph_Request with default values.
func NewFrameGraph_Request() *FrameGraph_Request {
	self := FrameGraph_Request{}
	self.SetDefaults(nil)
	return &self
}

func (t *FrameGraph_Request) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *FrameGraph_Request) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__tf2_msgs__srv__FrameGraph_Request())
}
func (t *FrameGraph_Request) PrepareMemory() unsafe.Pointer { //returns *C.tf2_msgs__srv__FrameGraph_Request
	return (unsafe.Pointer)(C.tf2_msgs__srv__FrameGraph_Request__create())
}
func (t *FrameGraph_Request) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.tf2_msgs__srv__FrameGraph_Request__destroy((*C.tf2_msgs__srv__FrameGraph_Request)(pointer_to_free))
}
func (t *FrameGraph_Request) AsCStruct() unsafe.Pointer {
	mem := (*C.tf2_msgs__srv__FrameGraph_Request)(t.PrepareMemory())
	return unsafe.Pointer(mem)
}
func (t *FrameGraph_Request) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	
}
func (t *FrameGraph_Request) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CFrameGraph_Request = C.tf2_msgs__srv__FrameGraph_Request
type CFrameGraph_Request__Sequence = C.tf2_msgs__srv__FrameGraph_Request__Sequence

func FrameGraph_Request__Sequence_to_Go(goSlice *[]FrameGraph_Request, cSlice CFrameGraph_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]FrameGraph_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.tf2_msgs__srv__FrameGraph_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_tf2_msgs__srv__FrameGraph_Request * uintptr(i)),
		))
		(*goSlice)[i] = FrameGraph_Request{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func FrameGraph_Request__Sequence_to_C(cSlice *CFrameGraph_Request__Sequence, goSlice []FrameGraph_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.tf2_msgs__srv__FrameGraph_Request)(C.malloc((C.size_t)(C.sizeof_struct_tf2_msgs__srv__FrameGraph_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.tf2_msgs__srv__FrameGraph_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_tf2_msgs__srv__FrameGraph_Request * uintptr(i)),
		))
		*cIdx = *(*C.tf2_msgs__srv__FrameGraph_Request)(v.AsCStruct())
	}
}
func FrameGraph_Request__Array_to_Go(goSlice []FrameGraph_Request, cSlice []CFrameGraph_Request) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func FrameGraph_Request__Array_to_C(cSlice []CFrameGraph_Request, goSlice []FrameGraph_Request) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.tf2_msgs__srv__FrameGraph_Request)(goSlice[i].AsCStruct())
	}
}


