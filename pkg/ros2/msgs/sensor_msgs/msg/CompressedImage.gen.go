/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package sensor_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	std_msgs_msg "github.com/tiiuae/rclgo/pkg/ros2/msgs/std_msgs/msg"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lsensor_msgs__rosidl_typesupport_c -lsensor_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/compressed_image.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("sensor_msgs/CompressedImage", &CompressedImage{})
}

// Do not create instances of this type directly. Always use NewCompressedImage
// function instead.
type CompressedImage struct {
	Header std_msgs_msg.Header `yaml:"header"`// Header timestamp should be acquisition time of image
	Format rosidl_runtime_c.String `yaml:"format"`// Specifies the format of the data
	Data []uint8 `yaml:"data"`// Compressed image buffer
}

// NewCompressedImage creates a new CompressedImage with default values.
func NewCompressedImage() *CompressedImage {
	self := CompressedImage{}
	self.SetDefaults(nil)
	return &self
}

func (t *CompressedImage) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Header.SetDefaults(nil)
	t.Format.SetDefaults("")
	
	return t
}

func (t *CompressedImage) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__CompressedImage())
}
func (t *CompressedImage) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__CompressedImage
	return (unsafe.Pointer)(C.sensor_msgs__msg__CompressedImage__create())
}
func (t *CompressedImage) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__CompressedImage__destroy((*C.sensor_msgs__msg__CompressedImage)(pointer_to_free))
}
func (t *CompressedImage) AsCStruct() unsafe.Pointer {
	mem := (*C.sensor_msgs__msg__CompressedImage)(t.PrepareMemory())
	mem.header = *(*C.std_msgs__msg__Header)(t.Header.AsCStruct())
	mem.format = *(*C.rosidl_runtime_c__String)(t.Format.AsCStruct())
	rosidl_runtime_c.Uint8__Sequence_to_C((*rosidl_runtime_c.CUint8__Sequence)(unsafe.Pointer(&mem.data)), t.Data)
	return unsafe.Pointer(mem)
}
func (t *CompressedImage) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.sensor_msgs__msg__CompressedImage)(ros2_message_buffer)
	t.Header.AsGoStruct(unsafe.Pointer(&mem.header))
	t.Format.AsGoStruct(unsafe.Pointer(&mem.format))
	rosidl_runtime_c.Uint8__Sequence_to_Go(&t.Data, *(*rosidl_runtime_c.CUint8__Sequence)(unsafe.Pointer(&mem.data)))
}
func (t *CompressedImage) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CCompressedImage = C.sensor_msgs__msg__CompressedImage
type CCompressedImage__Sequence = C.sensor_msgs__msg__CompressedImage__Sequence

func CompressedImage__Sequence_to_Go(goSlice *[]CompressedImage, cSlice CCompressedImage__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]CompressedImage, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.sensor_msgs__msg__CompressedImage__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__CompressedImage * uintptr(i)),
		))
		(*goSlice)[i] = CompressedImage{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func CompressedImage__Sequence_to_C(cSlice *CCompressedImage__Sequence, goSlice []CompressedImage) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__CompressedImage)(C.malloc((C.size_t)(C.sizeof_struct_sensor_msgs__msg__CompressedImage * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.sensor_msgs__msg__CompressedImage)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__CompressedImage * uintptr(i)),
		))
		*cIdx = *(*C.sensor_msgs__msg__CompressedImage)(v.AsCStruct())
	}
}
func CompressedImage__Array_to_Go(goSlice []CompressedImage, cSlice []CCompressedImage) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func CompressedImage__Array_to_C(cSlice []CCompressedImage, goSlice []CompressedImage) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.sensor_msgs__msg__CompressedImage)(goSlice[i].AsCStruct())
	}
}


