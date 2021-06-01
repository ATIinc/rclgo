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

#include <composition_interfaces/srv/list_nodes.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("composition_interfaces/ListNodes_Response", &ListNodes_Response{})
}

// Do not create instances of this type directly. Always use NewListNodes_Response
// function instead.
type ListNodes_Response struct {
	FullNodeNames []rosidl_runtime_c.String `yaml:"full_node_names"`// List of full node names including namespace.
	UniqueIds []uint64 `yaml:"unique_ids"`// List of full node names including namespace.corresponding unique ids (must have same length as full_node_names).
}

// NewListNodes_Response creates a new ListNodes_Response with default values.
func NewListNodes_Response() *ListNodes_Response {
	self := ListNodes_Response{}
	self.SetDefaults(nil)
	return &self
}

func (t *ListNodes_Response) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *ListNodes_Response) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__composition_interfaces__srv__ListNodes_Response())
}
func (t *ListNodes_Response) PrepareMemory() unsafe.Pointer { //returns *C.composition_interfaces__srv__ListNodes_Response
	return (unsafe.Pointer)(C.composition_interfaces__srv__ListNodes_Response__create())
}
func (t *ListNodes_Response) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.composition_interfaces__srv__ListNodes_Response__destroy((*C.composition_interfaces__srv__ListNodes_Response)(pointer_to_free))
}
func (t *ListNodes_Response) AsCStruct() unsafe.Pointer {
	mem := (*C.composition_interfaces__srv__ListNodes_Response)(t.PrepareMemory())
	rosidl_runtime_c.String__Sequence_to_C((*rosidl_runtime_c.CString__Sequence)(unsafe.Pointer(&mem.full_node_names)), t.FullNodeNames)
	rosidl_runtime_c.Uint64__Sequence_to_C((*rosidl_runtime_c.CUint64__Sequence)(unsafe.Pointer(&mem.unique_ids)), t.UniqueIds)
	return unsafe.Pointer(mem)
}
func (t *ListNodes_Response) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.composition_interfaces__srv__ListNodes_Response)(ros2_message_buffer)
	rosidl_runtime_c.String__Sequence_to_Go(&t.FullNodeNames, *(*rosidl_runtime_c.CString__Sequence)(unsafe.Pointer(&mem.full_node_names)))
	rosidl_runtime_c.Uint64__Sequence_to_Go(&t.UniqueIds, *(*rosidl_runtime_c.CUint64__Sequence)(unsafe.Pointer(&mem.unique_ids)))
}
func (t *ListNodes_Response) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CListNodes_Response = C.composition_interfaces__srv__ListNodes_Response
type CListNodes_Response__Sequence = C.composition_interfaces__srv__ListNodes_Response__Sequence

func ListNodes_Response__Sequence_to_Go(goSlice *[]ListNodes_Response, cSlice CListNodes_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ListNodes_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.composition_interfaces__srv__ListNodes_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_composition_interfaces__srv__ListNodes_Response * uintptr(i)),
		))
		(*goSlice)[i] = ListNodes_Response{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func ListNodes_Response__Sequence_to_C(cSlice *CListNodes_Response__Sequence, goSlice []ListNodes_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.composition_interfaces__srv__ListNodes_Response)(C.malloc((C.size_t)(C.sizeof_struct_composition_interfaces__srv__ListNodes_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.composition_interfaces__srv__ListNodes_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_composition_interfaces__srv__ListNodes_Response * uintptr(i)),
		))
		*cIdx = *(*C.composition_interfaces__srv__ListNodes_Response)(v.AsCStruct())
	}
}
func ListNodes_Response__Array_to_Go(goSlice []ListNodes_Response, cSlice []CListNodes_Response) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func ListNodes_Response__Array_to_C(cSlice []CListNodes_Response, goSlice []ListNodes_Response) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.composition_interfaces__srv__ListNodes_Response)(goSlice[i].AsCStruct())
	}
}


