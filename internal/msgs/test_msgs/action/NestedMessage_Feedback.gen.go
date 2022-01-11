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
	builtin_interfaces_msg "github.com/tiiuae/rclgo/internal/msgs/builtin_interfaces/msg"
	test_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/test_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/action/nested_message.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/NestedMessage_Feedback", NestedMessage_FeedbackTypeSupport)
}

// Do not create instances of this type directly. Always use NewNestedMessage_Feedback
// function instead.
type NestedMessage_Feedback struct {
	NestedFieldNoPkg test_msgs_msg.Builtins `yaml:"nested_field_no_pkg"`// goal definitionresult definitionfeedback
	NestedField test_msgs_msg.BasicTypes `yaml:"nested_field"`// goal definitionresult definitionfeedback
	NestedDifferentPkg builtin_interfaces_msg.Time `yaml:"nested_different_pkg"`// goal definitionresult definitionfeedback
}

// NewNestedMessage_Feedback creates a new NestedMessage_Feedback with default values.
func NewNestedMessage_Feedback() *NestedMessage_Feedback {
	self := NestedMessage_Feedback{}
	self.SetDefaults()
	return &self
}

func (t *NestedMessage_Feedback) Clone() *NestedMessage_Feedback {
	c := &NestedMessage_Feedback{}
	c.NestedFieldNoPkg = *t.NestedFieldNoPkg.Clone()
	c.NestedField = *t.NestedField.Clone()
	c.NestedDifferentPkg = *t.NestedDifferentPkg.Clone()
	return c
}

func (t *NestedMessage_Feedback) CloneMsg() types.Message {
	return t.Clone()
}

func (t *NestedMessage_Feedback) SetDefaults() {
	t.NestedFieldNoPkg.SetDefaults()
	t.NestedField.SetDefaults()
	t.NestedDifferentPkg.SetDefaults()
}

// NestedMessage_FeedbackPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type NestedMessage_FeedbackPublisher struct {
	*rclgo.Publisher
}

// NewNestedMessage_FeedbackPublisher creates and returns a new publisher for the
// NestedMessage_Feedback
func NewNestedMessage_FeedbackPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*NestedMessage_FeedbackPublisher, error) {
	pub, err := node.NewPublisher(topic_name, NestedMessage_FeedbackTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &NestedMessage_FeedbackPublisher{pub}, nil
}

func (p *NestedMessage_FeedbackPublisher) Publish(msg *NestedMessage_Feedback) error {
	return p.Publisher.Publish(msg)
}

// NestedMessage_FeedbackSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type NestedMessage_FeedbackSubscription struct {
	*rclgo.Subscription
}

// NestedMessage_FeedbackSubscriptionCallback type is used to provide a subscription
// handler function for a NestedMessage_FeedbackSubscription.
type NestedMessage_FeedbackSubscriptionCallback func(msg *NestedMessage_Feedback, info *rclgo.RmwMessageInfo, err error)

// NewNestedMessage_FeedbackSubscription creates and returns a new subscription for the
// NestedMessage_Feedback
func NewNestedMessage_FeedbackSubscription(node *rclgo.Node, topic_name string, subscriptionCallback NestedMessage_FeedbackSubscriptionCallback) (*NestedMessage_FeedbackSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg NestedMessage_Feedback
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, NestedMessage_FeedbackTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &NestedMessage_FeedbackSubscription{sub}, nil
}

func (s *NestedMessage_FeedbackSubscription) TakeMessage(out *NestedMessage_Feedback) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneNestedMessage_FeedbackSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneNestedMessage_FeedbackSlice(dst, src []NestedMessage_Feedback) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var NestedMessage_FeedbackTypeSupport types.MessageTypeSupport = _NestedMessage_FeedbackTypeSupport{}

type _NestedMessage_FeedbackTypeSupport struct{}

func (t _NestedMessage_FeedbackTypeSupport) New() types.Message {
	return NewNestedMessage_Feedback()
}

func (t _NestedMessage_FeedbackTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__action__NestedMessage_Feedback
	return (unsafe.Pointer)(C.test_msgs__action__NestedMessage_Feedback__create())
}

func (t _NestedMessage_FeedbackTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__action__NestedMessage_Feedback__destroy((*C.test_msgs__action__NestedMessage_Feedback)(pointer_to_free))
}

func (t _NestedMessage_FeedbackTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*NestedMessage_Feedback)
	mem := (*C.test_msgs__action__NestedMessage_Feedback)(dst)
	test_msgs_msg.BuiltinsTypeSupport.AsCStruct(unsafe.Pointer(&mem.nested_field_no_pkg), &m.NestedFieldNoPkg)
	test_msgs_msg.BasicTypesTypeSupport.AsCStruct(unsafe.Pointer(&mem.nested_field), &m.NestedField)
	builtin_interfaces_msg.TimeTypeSupport.AsCStruct(unsafe.Pointer(&mem.nested_different_pkg), &m.NestedDifferentPkg)
}

func (t _NestedMessage_FeedbackTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*NestedMessage_Feedback)
	mem := (*C.test_msgs__action__NestedMessage_Feedback)(ros2_message_buffer)
	test_msgs_msg.BuiltinsTypeSupport.AsGoStruct(&m.NestedFieldNoPkg, unsafe.Pointer(&mem.nested_field_no_pkg))
	test_msgs_msg.BasicTypesTypeSupport.AsGoStruct(&m.NestedField, unsafe.Pointer(&mem.nested_field))
	builtin_interfaces_msg.TimeTypeSupport.AsGoStruct(&m.NestedDifferentPkg, unsafe.Pointer(&mem.nested_different_pkg))
}

func (t _NestedMessage_FeedbackTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__action__NestedMessage_Feedback())
}

type CNestedMessage_Feedback = C.test_msgs__action__NestedMessage_Feedback
type CNestedMessage_Feedback__Sequence = C.test_msgs__action__NestedMessage_Feedback__Sequence

func NestedMessage_Feedback__Sequence_to_Go(goSlice *[]NestedMessage_Feedback, cSlice CNestedMessage_Feedback__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]NestedMessage_Feedback, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		NestedMessage_FeedbackTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func NestedMessage_Feedback__Sequence_to_C(cSlice *CNestedMessage_Feedback__Sequence, goSlice []NestedMessage_Feedback) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.test_msgs__action__NestedMessage_Feedback)(C.malloc(C.sizeof_struct_test_msgs__action__NestedMessage_Feedback * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		NestedMessage_FeedbackTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func NestedMessage_Feedback__Array_to_Go(goSlice []NestedMessage_Feedback, cSlice []CNestedMessage_Feedback) {
	for i := 0; i < len(cSlice); i++ {
		NestedMessage_FeedbackTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func NestedMessage_Feedback__Array_to_C(cSlice []CNestedMessage_Feedback, goSlice []NestedMessage_Feedback) {
	for i := 0; i < len(goSlice); i++ {
		NestedMessage_FeedbackTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
