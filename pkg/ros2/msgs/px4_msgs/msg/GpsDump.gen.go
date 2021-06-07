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

#include <px4_msgs/msg/gps_dump.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/GpsDump", GpsDumpTypeSupport)
}
const (
	GpsDump_ORB_QUEUE_LENGTH uint8 = 8
)

// Do not create instances of this type directly. Always use NewGpsDump
// function instead.
type GpsDump struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Instance uint8 `yaml:"instance"`// Instance of GNSS reciever
	Len uint8 `yaml:"len"`// length of data, MSB bit set = message to the gps device,
	Data [79]uint8 `yaml:"data"`// data to write to the log. clear = message from the device
}

// NewGpsDump creates a new GpsDump with default values.
func NewGpsDump() *GpsDump {
	self := GpsDump{}
	self.SetDefaults()
	return &self
}

func (t *GpsDump) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *GpsDump) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var GpsDumpTypeSupport types.MessageTypeSupport = _GpsDumpTypeSupport{}

type _GpsDumpTypeSupport struct{}

func (t _GpsDumpTypeSupport) New() types.Message {
	return NewGpsDump()
}

func (t _GpsDumpTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__GpsDump
	return (unsafe.Pointer)(C.px4_msgs__msg__GpsDump__create())
}

func (t _GpsDumpTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__GpsDump__destroy((*C.px4_msgs__msg__GpsDump)(pointer_to_free))
}

func (t _GpsDumpTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GpsDump)
	mem := (*C.px4_msgs__msg__GpsDump)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.instance = C.uint8_t(m.Instance)
	mem.len = C.uint8_t(m.Len)
	cSlice_data := mem.data[:]
	rosidl_runtime_c.Uint8__Array_to_C(*(*[]rosidl_runtime_c.CUint8)(unsafe.Pointer(&cSlice_data)), m.Data[:])
}

func (t _GpsDumpTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GpsDump)
	mem := (*C.px4_msgs__msg__GpsDump)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.Instance = uint8(mem.instance)
	m.Len = uint8(mem.len)
	cSlice_data := mem.data[:]
	rosidl_runtime_c.Uint8__Array_to_Go(m.Data[:], *(*[]rosidl_runtime_c.CUint8)(unsafe.Pointer(&cSlice_data)))
}

func (t _GpsDumpTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__GpsDump())
}

type CGpsDump = C.px4_msgs__msg__GpsDump
type CGpsDump__Sequence = C.px4_msgs__msg__GpsDump__Sequence

func GpsDump__Sequence_to_Go(goSlice *[]GpsDump, cSlice CGpsDump__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GpsDump, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__GpsDump__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__GpsDump * uintptr(i)),
		))
		GpsDumpTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GpsDump__Sequence_to_C(cSlice *CGpsDump__Sequence, goSlice []GpsDump) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__GpsDump)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__GpsDump * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__GpsDump)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__GpsDump * uintptr(i)),
		))
		GpsDumpTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GpsDump__Array_to_Go(goSlice []GpsDump, cSlice []CGpsDump) {
	for i := 0; i < len(cSlice); i++ {
		GpsDumpTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GpsDump__Array_to_C(cSlice []CGpsDump, goSlice []GpsDump) {
	for i := 0; i < len(goSlice); i++ {
		GpsDumpTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
