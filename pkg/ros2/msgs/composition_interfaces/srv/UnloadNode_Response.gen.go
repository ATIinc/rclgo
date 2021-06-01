/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package composition_interfaces_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lcomposition_interfaces__rosidl_typesupport_c -lcomposition_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <composition_interfaces/srv/unload_node.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("composition_interfaces/UnloadNode_Response", &UnloadNode_Response{})
}

// Do not create instances of this type directly. Always use NewUnloadNode_Response
// function instead.
type UnloadNode_Response struct {
	Success bool `yaml:"success"`// Container specific unique id of a loaded node.True if the node existed and was unloaded.
	ErrorMessage rosidl_runtime_c.String `yaml:"error_message"`// Human readable error message if success is false, else empty string.
}

// NewUnloadNode_Response creates a new UnloadNode_Response with default values.
func NewUnloadNode_Response() *UnloadNode_Response {
	self := UnloadNode_Response{}
	self.SetDefaults(nil)
	return &self
}

func (t *UnloadNode_Response) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.ErrorMessage.SetDefaults("")
	
	return t
}

func (t *UnloadNode_Response) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__composition_interfaces__srv__UnloadNode_Response())
}
func (t *UnloadNode_Response) PrepareMemory() unsafe.Pointer { //returns *C.composition_interfaces__srv__UnloadNode_Response
	return (unsafe.Pointer)(C.composition_interfaces__srv__UnloadNode_Response__create())
}
func (t *UnloadNode_Response) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.composition_interfaces__srv__UnloadNode_Response__destroy((*C.composition_interfaces__srv__UnloadNode_Response)(pointer_to_free))
}
func (t *UnloadNode_Response) AsCStruct() unsafe.Pointer {
	mem := (*C.composition_interfaces__srv__UnloadNode_Response)(t.PrepareMemory())
	mem.success = C.bool(t.Success)
	mem.error_message = *(*C.rosidl_runtime_c__String)(t.ErrorMessage.AsCStruct())
	return unsafe.Pointer(mem)
}
func (t *UnloadNode_Response) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.composition_interfaces__srv__UnloadNode_Response)(ros2_message_buffer)
	t.Success = bool(mem.success)
	t.ErrorMessage.AsGoStruct(unsafe.Pointer(&mem.error_message))
}
func (t *UnloadNode_Response) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CUnloadNode_Response = C.composition_interfaces__srv__UnloadNode_Response
type CUnloadNode_Response__Sequence = C.composition_interfaces__srv__UnloadNode_Response__Sequence

func UnloadNode_Response__Sequence_to_Go(goSlice *[]UnloadNode_Response, cSlice CUnloadNode_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UnloadNode_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.composition_interfaces__srv__UnloadNode_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_composition_interfaces__srv__UnloadNode_Response * uintptr(i)),
		))
		(*goSlice)[i] = UnloadNode_Response{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func UnloadNode_Response__Sequence_to_C(cSlice *CUnloadNode_Response__Sequence, goSlice []UnloadNode_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.composition_interfaces__srv__UnloadNode_Response)(C.malloc((C.size_t)(C.sizeof_struct_composition_interfaces__srv__UnloadNode_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.composition_interfaces__srv__UnloadNode_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_composition_interfaces__srv__UnloadNode_Response * uintptr(i)),
		))
		*cIdx = *(*C.composition_interfaces__srv__UnloadNode_Response)(v.AsCStruct())
	}
}
func UnloadNode_Response__Array_to_Go(goSlice []UnloadNode_Response, cSlice []CUnloadNode_Response) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func UnloadNode_Response__Array_to_C(cSlice []CUnloadNode_Response, goSlice []UnloadNode_Response) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.composition_interfaces__srv__UnloadNode_Response)(goSlice[i].AsCStruct())
	}
}


