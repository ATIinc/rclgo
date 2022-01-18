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
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/msg/float64_multi_array.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/Float64MultiArray", Float64MultiArrayTypeSupport)
}

// Do not create instances of this type directly. Always use NewFloat64MultiArray
// function instead.
type Float64MultiArray struct {
	Layout MultiArrayLayout `yaml:"layout"`// specification of data layout
	Data []float64 `yaml:"data"`// array of data
}

// NewFloat64MultiArray creates a new Float64MultiArray with default values.
func NewFloat64MultiArray() *Float64MultiArray {
	self := Float64MultiArray{}
	self.SetDefaults()
	return &self
}

func (t *Float64MultiArray) Clone() *Float64MultiArray {
	c := &Float64MultiArray{}
	c.Layout = *t.Layout.Clone()
	if t.Data != nil {
		c.Data = make([]float64, len(t.Data))
		copy(c.Data, t.Data)
	}
	return c
}

func (t *Float64MultiArray) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Float64MultiArray) SetDefaults() {
	t.Layout.SetDefaults()
	t.Data = nil
}

// Float64MultiArrayPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type Float64MultiArrayPublisher struct {
	*rclgo.Publisher
}

// NewFloat64MultiArrayPublisher creates and returns a new publisher for the
// Float64MultiArray
func NewFloat64MultiArrayPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*Float64MultiArrayPublisher, error) {
	pub, err := node.NewPublisher(topic_name, Float64MultiArrayTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Float64MultiArrayPublisher{pub}, nil
}

func (p *Float64MultiArrayPublisher) Publish(msg *Float64MultiArray) error {
	return p.Publisher.Publish(msg)
}

// Float64MultiArraySubscription wraps rclgo.Subscription to provide type safe helper
// functions
type Float64MultiArraySubscription struct {
	*rclgo.Subscription
}

// Float64MultiArraySubscriptionCallback type is used to provide a subscription
// handler function for a Float64MultiArraySubscription.
type Float64MultiArraySubscriptionCallback func(msg *Float64MultiArray, info *rclgo.RmwMessageInfo, err error)

// NewFloat64MultiArraySubscription creates and returns a new subscription for the
// Float64MultiArray
func NewFloat64MultiArraySubscription(node *rclgo.Node, topic_name string, subscriptionCallback Float64MultiArraySubscriptionCallback) (*Float64MultiArraySubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Float64MultiArray
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, Float64MultiArrayTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &Float64MultiArraySubscription{sub}, nil
}

func (s *Float64MultiArraySubscription) TakeMessage(out *Float64MultiArray) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneFloat64MultiArraySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneFloat64MultiArraySlice(dst, src []Float64MultiArray) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Float64MultiArrayTypeSupport types.MessageTypeSupport = _Float64MultiArrayTypeSupport{}

type _Float64MultiArrayTypeSupport struct{}

func (t _Float64MultiArrayTypeSupport) New() types.Message {
	return NewFloat64MultiArray()
}

func (t _Float64MultiArrayTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__Float64MultiArray
	return (unsafe.Pointer)(C.example_interfaces__msg__Float64MultiArray__create())
}

func (t _Float64MultiArrayTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__Float64MultiArray__destroy((*C.example_interfaces__msg__Float64MultiArray)(pointer_to_free))
}

func (t _Float64MultiArrayTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Float64MultiArray)
	mem := (*C.example_interfaces__msg__Float64MultiArray)(dst)
	MultiArrayLayoutTypeSupport.AsCStruct(unsafe.Pointer(&mem.layout), &m.Layout)
	primitives.Float64__Sequence_to_C((*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.data)), m.Data)
}

func (t _Float64MultiArrayTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Float64MultiArray)
	mem := (*C.example_interfaces__msg__Float64MultiArray)(ros2_message_buffer)
	MultiArrayLayoutTypeSupport.AsGoStruct(&m.Layout, unsafe.Pointer(&mem.layout))
	primitives.Float64__Sequence_to_Go(&m.Data, *(*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.data)))
}

func (t _Float64MultiArrayTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__Float64MultiArray())
}

type CFloat64MultiArray = C.example_interfaces__msg__Float64MultiArray
type CFloat64MultiArray__Sequence = C.example_interfaces__msg__Float64MultiArray__Sequence

func Float64MultiArray__Sequence_to_Go(goSlice *[]Float64MultiArray, cSlice CFloat64MultiArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Float64MultiArray, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__msg__Float64MultiArray__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__Float64MultiArray * uintptr(i)),
		))
		Float64MultiArrayTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Float64MultiArray__Sequence_to_C(cSlice *CFloat64MultiArray__Sequence, goSlice []Float64MultiArray) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__Float64MultiArray)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__msg__Float64MultiArray * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__msg__Float64MultiArray)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__Float64MultiArray * uintptr(i)),
		))
		Float64MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Float64MultiArray__Array_to_Go(goSlice []Float64MultiArray, cSlice []CFloat64MultiArray) {
	for i := 0; i < len(cSlice); i++ {
		Float64MultiArrayTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Float64MultiArray__Array_to_C(cSlice []CFloat64MultiArray, goSlice []Float64MultiArray) {
	for i := 0; i < len(goSlice); i++ {
		Float64MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
