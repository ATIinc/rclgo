/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package example_interfaces_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/msg/float32.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/Float32", Float32TypeSupport)
}

// Do not create instances of this type directly. Always use NewFloat32
// function instead.
type Float32 struct {
	Data float32 `yaml:"data"`// This is an example message of using a primitive datatype, float32.If you want to test with this that's fine, but if you are deployingit into a system you should create a semantically meaningful message type.If you want to embed it in another message, use the primitive data type instead.
}

// NewFloat32 creates a new Float32 with default values.
func NewFloat32() *Float32 {
	self := Float32{}
	self.SetDefaults()
	return &self
}

func (t *Float32) Clone() *Float32 {
	c := &Float32{}
	c.Data = t.Data
	return c
}

func (t *Float32) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Float32) SetDefaults() {
	t.Data = 0
}

// Float32Publisher wraps rclgo.Publisher to provide type safe helper
// functions
type Float32Publisher struct {
	*rclgo.Publisher
}

// NewFloat32Publisher creates and returns a new publisher for the
// Float32
func NewFloat32Publisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*Float32Publisher, error) {
	pub, err := node.NewPublisher(topic_name, Float32TypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Float32Publisher{pub}, nil
}

func (p *Float32Publisher) Publish(msg *Float32) error {
	return p.Publisher.Publish(msg)
}

// Float32Subscription wraps rclgo.Subscription to provide type safe helper
// functions
type Float32Subscription struct {
	*rclgo.Subscription
}

// Float32SubscriptionCallback type is used to provide a subscription
// handler function for a Float32Subscription.
type Float32SubscriptionCallback func(msg *Float32, info *rclgo.RmwMessageInfo, err error)

// NewFloat32Subscription creates and returns a new subscription for the
// Float32
func NewFloat32Subscription(node *rclgo.Node, topic_name string, subscriptionCallback Float32SubscriptionCallback) (*Float32Subscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Float32
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, Float32TypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &Float32Subscription{sub}, nil
}

func (s *Float32Subscription) TakeMessage(out *Float32) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneFloat32Slice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneFloat32Slice(dst, src []Float32) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Float32TypeSupport types.MessageTypeSupport = _Float32TypeSupport{}

type _Float32TypeSupport struct{}

func (t _Float32TypeSupport) New() types.Message {
	return NewFloat32()
}

func (t _Float32TypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__Float32
	return (unsafe.Pointer)(C.example_interfaces__msg__Float32__create())
}

func (t _Float32TypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__Float32__destroy((*C.example_interfaces__msg__Float32)(pointer_to_free))
}

func (t _Float32TypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Float32)
	mem := (*C.example_interfaces__msg__Float32)(dst)
	mem.data = C.float(m.Data)
}

func (t _Float32TypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Float32)
	mem := (*C.example_interfaces__msg__Float32)(ros2_message_buffer)
	m.Data = float32(mem.data)
}

func (t _Float32TypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__Float32())
}

type CFloat32 = C.example_interfaces__msg__Float32
type CFloat32__Sequence = C.example_interfaces__msg__Float32__Sequence

func Float32__Sequence_to_Go(goSlice *[]Float32, cSlice CFloat32__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Float32, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__msg__Float32__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__Float32 * uintptr(i)),
		))
		Float32TypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Float32__Sequence_to_C(cSlice *CFloat32__Sequence, goSlice []Float32) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__Float32)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__msg__Float32 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__msg__Float32)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__Float32 * uintptr(i)),
		))
		Float32TypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Float32__Array_to_Go(goSlice []Float32, cSlice []CFloat32) {
	for i := 0; i < len(cSlice); i++ {
		Float32TypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Float32__Array_to_C(cSlice []CFloat32, goSlice []Float32) {
	for i := 0; i < len(goSlice); i++ {
		Float32TypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
