/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package sensor_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/std_msgs/msg"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/compressed_image.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/CompressedImage", CompressedImageTypeSupport)
}

// Do not create instances of this type directly. Always use NewCompressedImage
// function instead.
type CompressedImage struct {
	Header std_msgs_msg.Header `yaml:"header"`// Header timestamp should be acquisition time of image
	Format string `yaml:"format"`// Specifies the format of the data
	Data []uint8 `yaml:"data"`// Compressed image buffer
}

// NewCompressedImage creates a new CompressedImage with default values.
func NewCompressedImage() *CompressedImage {
	self := CompressedImage{}
	self.SetDefaults()
	return &self
}

func (t *CompressedImage) Clone() *CompressedImage {
	c := &CompressedImage{}
	c.Header = *t.Header.Clone()
	c.Format = t.Format
	if t.Data != nil {
		c.Data = make([]uint8, len(t.Data))
		copy(c.Data, t.Data)
	}
	return c
}

func (t *CompressedImage) CloneMsg() types.Message {
	return t.Clone()
}

func (t *CompressedImage) SetDefaults() {
	t.Header.SetDefaults()
	t.Format = ""
	t.Data = nil
}

// CompressedImagePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type CompressedImagePublisher struct {
	*rclgo.Publisher
}

// NewCompressedImagePublisher creates and returns a new publisher for the
// CompressedImage
func NewCompressedImagePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*CompressedImagePublisher, error) {
	pub, err := node.NewPublisher(topic_name, CompressedImageTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &CompressedImagePublisher{pub}, nil
}

func (p *CompressedImagePublisher) Publish(msg *CompressedImage) error {
	return p.Publisher.Publish(msg)
}

// CompressedImageSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type CompressedImageSubscription struct {
	*rclgo.Subscription
}

// CompressedImageSubscriptionCallback type is used to provide a subscription
// handler function for a CompressedImageSubscription.
type CompressedImageSubscriptionCallback func(msg *CompressedImage, info *rclgo.RmwMessageInfo, err error)

// NewCompressedImageSubscription creates and returns a new subscription for the
// CompressedImage
func NewCompressedImageSubscription(node *rclgo.Node, topic_name string, subscriptionCallback CompressedImageSubscriptionCallback) (*CompressedImageSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg CompressedImage
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, CompressedImageTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &CompressedImageSubscription{sub}, nil
}

func (s *CompressedImageSubscription) TakeMessage(out *CompressedImage) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}


// CloneCompressedImageSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneCompressedImageSlice(dst, src []CompressedImage) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var CompressedImageTypeSupport types.MessageTypeSupport = _CompressedImageTypeSupport{}

type _CompressedImageTypeSupport struct{}

func (t _CompressedImageTypeSupport) New() types.Message {
	return NewCompressedImage()
}

func (t _CompressedImageTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__CompressedImage
	return (unsafe.Pointer)(C.sensor_msgs__msg__CompressedImage__create())
}

func (t _CompressedImageTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__CompressedImage__destroy((*C.sensor_msgs__msg__CompressedImage)(pointer_to_free))
}

func (t _CompressedImageTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*CompressedImage)
	mem := (*C.sensor_msgs__msg__CompressedImage)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.format), m.Format)
	primitives.Uint8__Sequence_to_C((*primitives.CUint8__Sequence)(unsafe.Pointer(&mem.data)), m.Data)
}

func (t _CompressedImageTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*CompressedImage)
	mem := (*C.sensor_msgs__msg__CompressedImage)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	primitives.StringAsGoStruct(&m.Format, unsafe.Pointer(&mem.format))
	primitives.Uint8__Sequence_to_Go(&m.Data, *(*primitives.CUint8__Sequence)(unsafe.Pointer(&mem.data)))
}

func (t _CompressedImageTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__CompressedImage())
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
		CompressedImageTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
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
		CompressedImageTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func CompressedImage__Array_to_Go(goSlice []CompressedImage, cSlice []CCompressedImage) {
	for i := 0; i < len(cSlice); i++ {
		CompressedImageTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func CompressedImage__Array_to_C(cSlice []CCompressedImage, goSlice []CompressedImage) {
	for i := 0; i < len(goSlice); i++ {
		CompressedImageTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
