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
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <std_msgs/msg/int8_multi_array.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("std_msgs/Int8MultiArray", Int8MultiArrayTypeSupport)
}

// Do not create instances of this type directly. Always use NewInt8MultiArray
// function instead.
type Int8MultiArray struct {
	Layout MultiArrayLayout `yaml:"layout"`// specification of data layout
	Data []int8 `yaml:"data"`// array of data
}

// NewInt8MultiArray creates a new Int8MultiArray with default values.
func NewInt8MultiArray() *Int8MultiArray {
	self := Int8MultiArray{}
	self.SetDefaults()
	return &self
}

func (t *Int8MultiArray) Clone() *Int8MultiArray {
	c := &Int8MultiArray{}
	c.Layout = *t.Layout.Clone()
	if t.Data != nil {
		c.Data = make([]int8, len(t.Data))
		copy(c.Data, t.Data)
	}
	return c
}

func (t *Int8MultiArray) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Int8MultiArray) SetDefaults() {
	t.Layout.SetDefaults()
	t.Data = nil
}

// Int8MultiArrayPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type Int8MultiArrayPublisher struct {
	*rclgo.Publisher
}

// NewInt8MultiArrayPublisher creates and returns a new publisher for the
// Int8MultiArray
func NewInt8MultiArrayPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*Int8MultiArrayPublisher, error) {
	pub, err := node.NewPublisher(topic_name, Int8MultiArrayTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Int8MultiArrayPublisher{pub}, nil
}

func (p *Int8MultiArrayPublisher) Publish(msg *Int8MultiArray) error {
	return p.Publisher.Publish(msg)
}

// Int8MultiArraySubscription wraps rclgo.Subscription to provide type safe helper
// functions
type Int8MultiArraySubscription struct {
	*rclgo.Subscription
}

// Int8MultiArraySubscriptionCallback type is used to provide a subscription
// handler function for a Int8MultiArraySubscription.
type Int8MultiArraySubscriptionCallback func(msg *Int8MultiArray, info *rclgo.RmwMessageInfo, err error)

// NewInt8MultiArraySubscription creates and returns a new subscription for the
// Int8MultiArray
func NewInt8MultiArraySubscription(node *rclgo.Node, topic_name string, subscriptionCallback Int8MultiArraySubscriptionCallback) (*Int8MultiArraySubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Int8MultiArray
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, Int8MultiArrayTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &Int8MultiArraySubscription{sub}, nil
}

func (s *Int8MultiArraySubscription) TakeMessage(out *Int8MultiArray) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}


// CloneInt8MultiArraySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneInt8MultiArraySlice(dst, src []Int8MultiArray) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Int8MultiArrayTypeSupport types.MessageTypeSupport = _Int8MultiArrayTypeSupport{}

type _Int8MultiArrayTypeSupport struct{}

func (t _Int8MultiArrayTypeSupport) New() types.Message {
	return NewInt8MultiArray()
}

func (t _Int8MultiArrayTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__Int8MultiArray
	return (unsafe.Pointer)(C.std_msgs__msg__Int8MultiArray__create())
}

func (t _Int8MultiArrayTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__Int8MultiArray__destroy((*C.std_msgs__msg__Int8MultiArray)(pointer_to_free))
}

func (t _Int8MultiArrayTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Int8MultiArray)
	mem := (*C.std_msgs__msg__Int8MultiArray)(dst)
	MultiArrayLayoutTypeSupport.AsCStruct(unsafe.Pointer(&mem.layout), &m.Layout)
	primitives.Int8__Sequence_to_C((*primitives.CInt8__Sequence)(unsafe.Pointer(&mem.data)), m.Data)
}

func (t _Int8MultiArrayTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Int8MultiArray)
	mem := (*C.std_msgs__msg__Int8MultiArray)(ros2_message_buffer)
	MultiArrayLayoutTypeSupport.AsGoStruct(&m.Layout, unsafe.Pointer(&mem.layout))
	primitives.Int8__Sequence_to_Go(&m.Data, *(*primitives.CInt8__Sequence)(unsafe.Pointer(&mem.data)))
}

func (t _Int8MultiArrayTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__Int8MultiArray())
}

type CInt8MultiArray = C.std_msgs__msg__Int8MultiArray
type CInt8MultiArray__Sequence = C.std_msgs__msg__Int8MultiArray__Sequence

func Int8MultiArray__Sequence_to_Go(goSlice *[]Int8MultiArray, cSlice CInt8MultiArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Int8MultiArray, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.std_msgs__msg__Int8MultiArray__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__Int8MultiArray * uintptr(i)),
		))
		Int8MultiArrayTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Int8MultiArray__Sequence_to_C(cSlice *CInt8MultiArray__Sequence, goSlice []Int8MultiArray) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__Int8MultiArray)(C.malloc((C.size_t)(C.sizeof_struct_std_msgs__msg__Int8MultiArray * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.std_msgs__msg__Int8MultiArray)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__Int8MultiArray * uintptr(i)),
		))
		Int8MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Int8MultiArray__Array_to_Go(goSlice []Int8MultiArray, cSlice []CInt8MultiArray) {
	for i := 0; i < len(cSlice); i++ {
		Int8MultiArrayTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Int8MultiArray__Array_to_C(cSlice []CInt8MultiArray, goSlice []Int8MultiArray) {
	for i := 0; i < len(goSlice); i++ {
		Int8MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
