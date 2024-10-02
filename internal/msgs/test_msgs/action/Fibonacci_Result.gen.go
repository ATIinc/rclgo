/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package test_msgs_action
import (
	"unsafe"

	"github.com/ATIinc/rclgo/pkg/rclgo"
	"github.com/ATIinc/rclgo/pkg/rclgo/types"
	"github.com/ATIinc/rclgo/pkg/rclgo/typemap"
	primitives "github.com/ATIinc/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/action/fibonacci.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/Fibonacci_Result", Fibonacci_ResultTypeSupport)
	typemap.RegisterMessage("test_msgs/action/Fibonacci_Result", Fibonacci_ResultTypeSupport)
}

type Fibonacci_Result struct {
	Sequence []int32 `yaml:"sequence"`// result definition
}

// NewFibonacci_Result creates a new Fibonacci_Result with default values.
func NewFibonacci_Result() *Fibonacci_Result {
	self := Fibonacci_Result{}
	self.SetDefaults()
	return &self
}

func (t *Fibonacci_Result) Clone() *Fibonacci_Result {
	c := &Fibonacci_Result{}
	if t.Sequence != nil {
		c.Sequence = make([]int32, len(t.Sequence))
		copy(c.Sequence, t.Sequence)
	}
	return c
}

func (t *Fibonacci_Result) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Fibonacci_Result) SetDefaults() {
	t.Sequence = nil
}

func (t *Fibonacci_Result) GetTypeSupport() types.MessageTypeSupport {
	return Fibonacci_ResultTypeSupport
}

// Fibonacci_ResultPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type Fibonacci_ResultPublisher struct {
	*rclgo.Publisher
}

// NewFibonacci_ResultPublisher creates and returns a new publisher for the
// Fibonacci_Result
func NewFibonacci_ResultPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*Fibonacci_ResultPublisher, error) {
	pub, err := node.NewPublisher(topic_name, Fibonacci_ResultTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Fibonacci_ResultPublisher{pub}, nil
}

func (p *Fibonacci_ResultPublisher) Publish(msg *Fibonacci_Result) error {
	return p.Publisher.Publish(msg)
}

// Fibonacci_ResultSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type Fibonacci_ResultSubscription struct {
	*rclgo.Subscription
}

// Fibonacci_ResultSubscriptionCallback type is used to provide a subscription
// handler function for a Fibonacci_ResultSubscription.
type Fibonacci_ResultSubscriptionCallback func(msg *Fibonacci_Result, info *rclgo.MessageInfo, err error)

// NewFibonacci_ResultSubscription creates and returns a new subscription for the
// Fibonacci_Result
func NewFibonacci_ResultSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback Fibonacci_ResultSubscriptionCallback) (*Fibonacci_ResultSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Fibonacci_Result
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, Fibonacci_ResultTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &Fibonacci_ResultSubscription{sub}, nil
}

func (s *Fibonacci_ResultSubscription) TakeMessage(out *Fibonacci_Result) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneFibonacci_ResultSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneFibonacci_ResultSlice(dst, src []Fibonacci_Result) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Fibonacci_ResultTypeSupport types.MessageTypeSupport = _Fibonacci_ResultTypeSupport{}

type _Fibonacci_ResultTypeSupport struct{}

func (t _Fibonacci_ResultTypeSupport) New() types.Message {
	return NewFibonacci_Result()
}

func (t _Fibonacci_ResultTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__action__Fibonacci_Result
	return (unsafe.Pointer)(C.test_msgs__action__Fibonacci_Result__create())
}

func (t _Fibonacci_ResultTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__action__Fibonacci_Result__destroy((*C.test_msgs__action__Fibonacci_Result)(pointer_to_free))
}

func (t _Fibonacci_ResultTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Fibonacci_Result)
	mem := (*C.test_msgs__action__Fibonacci_Result)(dst)
	primitives.Int32__Sequence_to_C((*primitives.CInt32__Sequence)(unsafe.Pointer(&mem.sequence)), m.Sequence)
}

func (t _Fibonacci_ResultTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Fibonacci_Result)
	mem := (*C.test_msgs__action__Fibonacci_Result)(ros2_message_buffer)
	primitives.Int32__Sequence_to_Go(&m.Sequence, *(*primitives.CInt32__Sequence)(unsafe.Pointer(&mem.sequence)))
}

func (t _Fibonacci_ResultTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__action__Fibonacci_Result())
}

type CFibonacci_Result = C.test_msgs__action__Fibonacci_Result
type CFibonacci_Result__Sequence = C.test_msgs__action__Fibonacci_Result__Sequence

func Fibonacci_Result__Sequence_to_Go(goSlice *[]Fibonacci_Result, cSlice CFibonacci_Result__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Fibonacci_Result, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		Fibonacci_ResultTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Fibonacci_Result__Sequence_to_C(cSlice *CFibonacci_Result__Sequence, goSlice []Fibonacci_Result) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.test_msgs__action__Fibonacci_Result)(C.malloc(C.sizeof_struct_test_msgs__action__Fibonacci_Result * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		Fibonacci_ResultTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Fibonacci_Result__Array_to_Go(goSlice []Fibonacci_Result, cSlice []CFibonacci_Result) {
	for i := 0; i < len(cSlice); i++ {
		Fibonacci_ResultTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Fibonacci_Result__Array_to_C(cSlice []CFibonacci_Result, goSlice []Fibonacci_Result) {
	for i := 0; i < len(goSlice); i++ {
		Fibonacci_ResultTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
