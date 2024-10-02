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
	builtin_interfaces_msg "github.com/ATIinc/rclgo/internal/msgs/builtin_interfaces/msg"
	test_msgs_msg "github.com/ATIinc/rclgo/internal/msgs/test_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/action/nested_message.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/NestedMessage_Result", NestedMessage_ResultTypeSupport)
	typemap.RegisterMessage("test_msgs/action/NestedMessage_Result", NestedMessage_ResultTypeSupport)
}

type NestedMessage_Result struct {
	NestedFieldNoPkg test_msgs_msg.Builtins `yaml:"nested_field_no_pkg"`// result definition
	NestedField test_msgs_msg.BasicTypes `yaml:"nested_field"`
	NestedDifferentPkg builtin_interfaces_msg.Time `yaml:"nested_different_pkg"`
}

// NewNestedMessage_Result creates a new NestedMessage_Result with default values.
func NewNestedMessage_Result() *NestedMessage_Result {
	self := NestedMessage_Result{}
	self.SetDefaults()
	return &self
}

func (t *NestedMessage_Result) Clone() *NestedMessage_Result {
	c := &NestedMessage_Result{}
	c.NestedFieldNoPkg = *t.NestedFieldNoPkg.Clone()
	c.NestedField = *t.NestedField.Clone()
	c.NestedDifferentPkg = *t.NestedDifferentPkg.Clone()
	return c
}

func (t *NestedMessage_Result) CloneMsg() types.Message {
	return t.Clone()
}

func (t *NestedMessage_Result) SetDefaults() {
	t.NestedFieldNoPkg.SetDefaults()
	t.NestedField.SetDefaults()
	t.NestedDifferentPkg.SetDefaults()
}

func (t *NestedMessage_Result) GetTypeSupport() types.MessageTypeSupport {
	return NestedMessage_ResultTypeSupport
}

// NestedMessage_ResultPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type NestedMessage_ResultPublisher struct {
	*rclgo.Publisher
}

// NewNestedMessage_ResultPublisher creates and returns a new publisher for the
// NestedMessage_Result
func NewNestedMessage_ResultPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*NestedMessage_ResultPublisher, error) {
	pub, err := node.NewPublisher(topic_name, NestedMessage_ResultTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &NestedMessage_ResultPublisher{pub}, nil
}

func (p *NestedMessage_ResultPublisher) Publish(msg *NestedMessage_Result) error {
	return p.Publisher.Publish(msg)
}

// NestedMessage_ResultSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type NestedMessage_ResultSubscription struct {
	*rclgo.Subscription
}

// NestedMessage_ResultSubscriptionCallback type is used to provide a subscription
// handler function for a NestedMessage_ResultSubscription.
type NestedMessage_ResultSubscriptionCallback func(msg *NestedMessage_Result, info *rclgo.MessageInfo, err error)

// NewNestedMessage_ResultSubscription creates and returns a new subscription for the
// NestedMessage_Result
func NewNestedMessage_ResultSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback NestedMessage_ResultSubscriptionCallback) (*NestedMessage_ResultSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg NestedMessage_Result
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, NestedMessage_ResultTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &NestedMessage_ResultSubscription{sub}, nil
}

func (s *NestedMessage_ResultSubscription) TakeMessage(out *NestedMessage_Result) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneNestedMessage_ResultSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneNestedMessage_ResultSlice(dst, src []NestedMessage_Result) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var NestedMessage_ResultTypeSupport types.MessageTypeSupport = _NestedMessage_ResultTypeSupport{}

type _NestedMessage_ResultTypeSupport struct{}

func (t _NestedMessage_ResultTypeSupport) New() types.Message {
	return NewNestedMessage_Result()
}

func (t _NestedMessage_ResultTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__action__NestedMessage_Result
	return (unsafe.Pointer)(C.test_msgs__action__NestedMessage_Result__create())
}

func (t _NestedMessage_ResultTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__action__NestedMessage_Result__destroy((*C.test_msgs__action__NestedMessage_Result)(pointer_to_free))
}

func (t _NestedMessage_ResultTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*NestedMessage_Result)
	mem := (*C.test_msgs__action__NestedMessage_Result)(dst)
	test_msgs_msg.BuiltinsTypeSupport.AsCStruct(unsafe.Pointer(&mem.nested_field_no_pkg), &m.NestedFieldNoPkg)
	test_msgs_msg.BasicTypesTypeSupport.AsCStruct(unsafe.Pointer(&mem.nested_field), &m.NestedField)
	builtin_interfaces_msg.TimeTypeSupport.AsCStruct(unsafe.Pointer(&mem.nested_different_pkg), &m.NestedDifferentPkg)
}

func (t _NestedMessage_ResultTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*NestedMessage_Result)
	mem := (*C.test_msgs__action__NestedMessage_Result)(ros2_message_buffer)
	test_msgs_msg.BuiltinsTypeSupport.AsGoStruct(&m.NestedFieldNoPkg, unsafe.Pointer(&mem.nested_field_no_pkg))
	test_msgs_msg.BasicTypesTypeSupport.AsGoStruct(&m.NestedField, unsafe.Pointer(&mem.nested_field))
	builtin_interfaces_msg.TimeTypeSupport.AsGoStruct(&m.NestedDifferentPkg, unsafe.Pointer(&mem.nested_different_pkg))
}

func (t _NestedMessage_ResultTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__action__NestedMessage_Result())
}

type CNestedMessage_Result = C.test_msgs__action__NestedMessage_Result
type CNestedMessage_Result__Sequence = C.test_msgs__action__NestedMessage_Result__Sequence

func NestedMessage_Result__Sequence_to_Go(goSlice *[]NestedMessage_Result, cSlice CNestedMessage_Result__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]NestedMessage_Result, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		NestedMessage_ResultTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func NestedMessage_Result__Sequence_to_C(cSlice *CNestedMessage_Result__Sequence, goSlice []NestedMessage_Result) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.test_msgs__action__NestedMessage_Result)(C.malloc(C.sizeof_struct_test_msgs__action__NestedMessage_Result * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		NestedMessage_ResultTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func NestedMessage_Result__Array_to_Go(goSlice []NestedMessage_Result, cSlice []CNestedMessage_Result) {
	for i := 0; i < len(cSlice); i++ {
		NestedMessage_ResultTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func NestedMessage_Result__Array_to_C(cSlice []CNestedMessage_Result, goSlice []NestedMessage_Result) {
	for i := 0; i < len(goSlice); i++ {
		NestedMessage_ResultTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
