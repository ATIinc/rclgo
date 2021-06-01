/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package example_interfaces_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lexample_interfaces__rosidl_typesupport_c -lexample_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/srv/trigger.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("example_interfaces/Trigger_Response", &Trigger_Response{})
}

// Do not create instances of this type directly. Always use NewTrigger_Response
// function instead.
type Trigger_Response struct {
	Success bool `yaml:"success"`// indicate successful run of triggered service
	Message rosidl_runtime_c.String `yaml:"message"`// informational, e.g. for error messages.
}

// NewTrigger_Response creates a new Trigger_Response with default values.
func NewTrigger_Response() *Trigger_Response {
	self := Trigger_Response{}
	self.SetDefaults(nil)
	return &self
}

func (t *Trigger_Response) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Message.SetDefaults("")
	
	return t
}

func (t *Trigger_Response) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__srv__Trigger_Response())
}
func (t *Trigger_Response) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__srv__Trigger_Response
	return (unsafe.Pointer)(C.example_interfaces__srv__Trigger_Response__create())
}
func (t *Trigger_Response) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__srv__Trigger_Response__destroy((*C.example_interfaces__srv__Trigger_Response)(pointer_to_free))
}
func (t *Trigger_Response) AsCStruct() unsafe.Pointer {
	mem := (*C.example_interfaces__srv__Trigger_Response)(t.PrepareMemory())
	mem.success = C.bool(t.Success)
	mem.message = *(*C.rosidl_runtime_c__String)(t.Message.AsCStruct())
	return unsafe.Pointer(mem)
}
func (t *Trigger_Response) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.example_interfaces__srv__Trigger_Response)(ros2_message_buffer)
	t.Success = bool(mem.success)
	t.Message.AsGoStruct(unsafe.Pointer(&mem.message))
}
func (t *Trigger_Response) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CTrigger_Response = C.example_interfaces__srv__Trigger_Response
type CTrigger_Response__Sequence = C.example_interfaces__srv__Trigger_Response__Sequence

func Trigger_Response__Sequence_to_Go(goSlice *[]Trigger_Response, cSlice CTrigger_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Trigger_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__srv__Trigger_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__srv__Trigger_Response * uintptr(i)),
		))
		(*goSlice)[i] = Trigger_Response{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func Trigger_Response__Sequence_to_C(cSlice *CTrigger_Response__Sequence, goSlice []Trigger_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__srv__Trigger_Response)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__srv__Trigger_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__srv__Trigger_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__srv__Trigger_Response * uintptr(i)),
		))
		*cIdx = *(*C.example_interfaces__srv__Trigger_Response)(v.AsCStruct())
	}
}
func Trigger_Response__Array_to_Go(goSlice []Trigger_Response, cSlice []CTrigger_Response) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func Trigger_Response__Array_to_C(cSlice []CTrigger_Response, goSlice []Trigger_Response) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.example_interfaces__srv__Trigger_Response)(goSlice[i].AsCStruct())
	}
}


