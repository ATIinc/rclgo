/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package std_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <std_msgs/msg/u_int16.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("std_msgs/UInt16", UInt16TypeSupport)
}

// Do not create instances of this type directly. Always use NewUInt16
// function instead.
type UInt16 struct {
	Data uint16 `yaml:"data"`
}

// NewUInt16 creates a new UInt16 with default values.
func NewUInt16() *UInt16 {
	self := UInt16{}
	self.SetDefaults()
	return &self
}

func (t *UInt16) Clone() *UInt16 {
	c := &UInt16{}
	c.Data = t.Data
	return c
}

func (t *UInt16) CloneMsg() types.Message {
	return t.Clone()
}

func (t *UInt16) SetDefaults() {
	t.Data = 0
}

// UInt16Publisher wraps rclgo.Publisher to provide type safe helper
// functions
type UInt16Publisher struct {
	*rclgo.Publisher
}

// NewUInt16Publisher creates and returns a new publisher for the
// UInt16
func NewUInt16Publisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*UInt16Publisher, error) {
	pub, err := node.NewPublisher(topic_name, UInt16TypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &UInt16Publisher{pub}, nil
}

func (p *UInt16Publisher) Publish(msg *UInt16) error {
	return p.Publisher.Publish(msg)
}

// UInt16Subscription wraps rclgo.Subscription to provide type safe helper
// functions
type UInt16Subscription struct {
	*rclgo.Subscription
}

// UInt16SubscriptionCallback type is used to provide a subscription
// handler function for a UInt16Subscription.
type UInt16SubscriptionCallback func(msg *UInt16, info *rclgo.RmwMessageInfo, err error)

// NewUInt16Subscription creates and returns a new subscription for the
// UInt16
func NewUInt16Subscription(node *rclgo.Node, topic_name string, subscriptionCallback UInt16SubscriptionCallback) (*UInt16Subscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg UInt16
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, UInt16TypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &UInt16Subscription{sub}, nil
}

func (s *UInt16Subscription) TakeMessage(out *UInt16) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}


// CloneUInt16Slice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneUInt16Slice(dst, src []UInt16) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var UInt16TypeSupport types.MessageTypeSupport = _UInt16TypeSupport{}

type _UInt16TypeSupport struct{}

func (t _UInt16TypeSupport) New() types.Message {
	return NewUInt16()
}

func (t _UInt16TypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__UInt16
	return (unsafe.Pointer)(C.std_msgs__msg__UInt16__create())
}

func (t _UInt16TypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__UInt16__destroy((*C.std_msgs__msg__UInt16)(pointer_to_free))
}

func (t _UInt16TypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*UInt16)
	mem := (*C.std_msgs__msg__UInt16)(dst)
	mem.data = C.uint16_t(m.Data)
}

func (t _UInt16TypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*UInt16)
	mem := (*C.std_msgs__msg__UInt16)(ros2_message_buffer)
	m.Data = uint16(mem.data)
}

func (t _UInt16TypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__UInt16())
}

type CUInt16 = C.std_msgs__msg__UInt16
type CUInt16__Sequence = C.std_msgs__msg__UInt16__Sequence

func UInt16__Sequence_to_Go(goSlice *[]UInt16, cSlice CUInt16__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UInt16, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.std_msgs__msg__UInt16__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__UInt16 * uintptr(i)),
		))
		UInt16TypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func UInt16__Sequence_to_C(cSlice *CUInt16__Sequence, goSlice []UInt16) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__UInt16)(C.malloc((C.size_t)(C.sizeof_struct_std_msgs__msg__UInt16 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.std_msgs__msg__UInt16)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__UInt16 * uintptr(i)),
		))
		UInt16TypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func UInt16__Array_to_Go(goSlice []UInt16, cSlice []CUInt16) {
	for i := 0; i < len(cSlice); i++ {
		UInt16TypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func UInt16__Array_to_C(cSlice []CUInt16, goSlice []UInt16) {
	for i := 0; i < len(goSlice); i++ {
		UInt16TypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
