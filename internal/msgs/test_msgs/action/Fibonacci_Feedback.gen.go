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

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/action/fibonacci.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/Fibonacci_Feedback", Fibonacci_FeedbackTypeSupport)
}

// Do not create instances of this type directly. Always use NewFibonacci_Feedback
// function instead.
type Fibonacci_Feedback struct {
	Sequence []int32 `yaml:"sequence"`// goal definitionresult definitionfeedback
}

// NewFibonacci_Feedback creates a new Fibonacci_Feedback with default values.
func NewFibonacci_Feedback() *Fibonacci_Feedback {
	self := Fibonacci_Feedback{}
	self.SetDefaults()
	return &self
}

func (t *Fibonacci_Feedback) Clone() *Fibonacci_Feedback {
	c := &Fibonacci_Feedback{}
	if t.Sequence != nil {
		c.Sequence = make([]int32, len(t.Sequence))
		copy(c.Sequence, t.Sequence)
	}
	return c
}

func (t *Fibonacci_Feedback) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Fibonacci_Feedback) SetDefaults() {
	t.Sequence = nil
}

// Fibonacci_FeedbackPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type Fibonacci_FeedbackPublisher struct {
	*rclgo.Publisher
}

// NewFibonacci_FeedbackPublisher creates and returns a new publisher for the
// Fibonacci_Feedback
func NewFibonacci_FeedbackPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*Fibonacci_FeedbackPublisher, error) {
	pub, err := node.NewPublisher(topic_name, Fibonacci_FeedbackTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Fibonacci_FeedbackPublisher{pub}, nil
}

func (p *Fibonacci_FeedbackPublisher) Publish(msg *Fibonacci_Feedback) error {
	return p.Publisher.Publish(msg)
}

// Fibonacci_FeedbackSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type Fibonacci_FeedbackSubscription struct {
	*rclgo.Subscription
}

// Fibonacci_FeedbackSubscriptionCallback type is used to provide a subscription
// handler function for a Fibonacci_FeedbackSubscription.
type Fibonacci_FeedbackSubscriptionCallback func(msg *Fibonacci_Feedback, info *rclgo.RmwMessageInfo, err error)

// NewFibonacci_FeedbackSubscription creates and returns a new subscription for the
// Fibonacci_Feedback
func NewFibonacci_FeedbackSubscription(node *rclgo.Node, topic_name string, subscriptionCallback Fibonacci_FeedbackSubscriptionCallback) (*Fibonacci_FeedbackSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Fibonacci_Feedback
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, Fibonacci_FeedbackTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &Fibonacci_FeedbackSubscription{sub}, nil
}

func (s *Fibonacci_FeedbackSubscription) TakeMessage(out *Fibonacci_Feedback) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneFibonacci_FeedbackSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneFibonacci_FeedbackSlice(dst, src []Fibonacci_Feedback) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Fibonacci_FeedbackTypeSupport types.MessageTypeSupport = _Fibonacci_FeedbackTypeSupport{}

type _Fibonacci_FeedbackTypeSupport struct{}

func (t _Fibonacci_FeedbackTypeSupport) New() types.Message {
	return NewFibonacci_Feedback()
}

func (t _Fibonacci_FeedbackTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__action__Fibonacci_Feedback
	return (unsafe.Pointer)(C.test_msgs__action__Fibonacci_Feedback__create())
}

func (t _Fibonacci_FeedbackTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__action__Fibonacci_Feedback__destroy((*C.test_msgs__action__Fibonacci_Feedback)(pointer_to_free))
}

func (t _Fibonacci_FeedbackTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Fibonacci_Feedback)
	mem := (*C.test_msgs__action__Fibonacci_Feedback)(dst)
	primitives.Int32__Sequence_to_C((*primitives.CInt32__Sequence)(unsafe.Pointer(&mem.sequence)), m.Sequence)
}

func (t _Fibonacci_FeedbackTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Fibonacci_Feedback)
	mem := (*C.test_msgs__action__Fibonacci_Feedback)(ros2_message_buffer)
	primitives.Int32__Sequence_to_Go(&m.Sequence, *(*primitives.CInt32__Sequence)(unsafe.Pointer(&mem.sequence)))
}

func (t _Fibonacci_FeedbackTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__action__Fibonacci_Feedback())
}

type CFibonacci_Feedback = C.test_msgs__action__Fibonacci_Feedback
type CFibonacci_Feedback__Sequence = C.test_msgs__action__Fibonacci_Feedback__Sequence

func Fibonacci_Feedback__Sequence_to_Go(goSlice *[]Fibonacci_Feedback, cSlice CFibonacci_Feedback__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Fibonacci_Feedback, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.test_msgs__action__Fibonacci_Feedback__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__action__Fibonacci_Feedback * uintptr(i)),
		))
		Fibonacci_FeedbackTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Fibonacci_Feedback__Sequence_to_C(cSlice *CFibonacci_Feedback__Sequence, goSlice []Fibonacci_Feedback) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.test_msgs__action__Fibonacci_Feedback)(C.malloc((C.size_t)(C.sizeof_struct_test_msgs__action__Fibonacci_Feedback * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.test_msgs__action__Fibonacci_Feedback)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__action__Fibonacci_Feedback * uintptr(i)),
		))
		Fibonacci_FeedbackTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Fibonacci_Feedback__Array_to_Go(goSlice []Fibonacci_Feedback, cSlice []CFibonacci_Feedback) {
	for i := 0; i < len(cSlice); i++ {
		Fibonacci_FeedbackTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Fibonacci_Feedback__Array_to_C(cSlice []CFibonacci_Feedback, goSlice []Fibonacci_Feedback) {
	for i := 0; i < len(goSlice); i++ {
		Fibonacci_FeedbackTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
