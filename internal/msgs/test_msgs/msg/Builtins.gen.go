/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package test_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	builtin_interfaces_msg "github.com/tiiuae/rclgo/internal/msgs/builtin_interfaces/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/msg/builtins.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/Builtins", BuiltinsTypeSupport)
}

// Do not create instances of this type directly. Always use NewBuiltins
// function instead.
type Builtins struct {
	DurationValue builtin_interfaces_msg.Duration `yaml:"duration_value"`
	TimeValue builtin_interfaces_msg.Time `yaml:"time_value"`
}

// NewBuiltins creates a new Builtins with default values.
func NewBuiltins() *Builtins {
	self := Builtins{}
	self.SetDefaults()
	return &self
}

func (t *Builtins) Clone() *Builtins {
	c := &Builtins{}
	c.DurationValue = *t.DurationValue.Clone()
	c.TimeValue = *t.TimeValue.Clone()
	return c
}

func (t *Builtins) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Builtins) SetDefaults() {
	t.DurationValue.SetDefaults()
	t.TimeValue.SetDefaults()
}

// BuiltinsPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type BuiltinsPublisher struct {
	*rclgo.Publisher
}

// NewBuiltinsPublisher creates and returns a new publisher for the
// Builtins
func NewBuiltinsPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*BuiltinsPublisher, error) {
	pub, err := node.NewPublisher(topic_name, BuiltinsTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &BuiltinsPublisher{pub}, nil
}

func (p *BuiltinsPublisher) Publish(msg *Builtins) error {
	return p.Publisher.Publish(msg)
}

// BuiltinsSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type BuiltinsSubscription struct {
	*rclgo.Subscription
}

// BuiltinsSubscriptionCallback type is used to provide a subscription
// handler function for a BuiltinsSubscription.
type BuiltinsSubscriptionCallback func(msg *Builtins, info *rclgo.RmwMessageInfo, err error)

// NewBuiltinsSubscription creates and returns a new subscription for the
// Builtins
func NewBuiltinsSubscription(node *rclgo.Node, topic_name string, subscriptionCallback BuiltinsSubscriptionCallback) (*BuiltinsSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Builtins
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, BuiltinsTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &BuiltinsSubscription{sub}, nil
}

func (s *BuiltinsSubscription) TakeMessage(out *Builtins) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneBuiltinsSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneBuiltinsSlice(dst, src []Builtins) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var BuiltinsTypeSupport types.MessageTypeSupport = _BuiltinsTypeSupport{}

type _BuiltinsTypeSupport struct{}

func (t _BuiltinsTypeSupport) New() types.Message {
	return NewBuiltins()
}

func (t _BuiltinsTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__msg__Builtins
	return (unsafe.Pointer)(C.test_msgs__msg__Builtins__create())
}

func (t _BuiltinsTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__msg__Builtins__destroy((*C.test_msgs__msg__Builtins)(pointer_to_free))
}

func (t _BuiltinsTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Builtins)
	mem := (*C.test_msgs__msg__Builtins)(dst)
	builtin_interfaces_msg.DurationTypeSupport.AsCStruct(unsafe.Pointer(&mem.duration_value), &m.DurationValue)
	builtin_interfaces_msg.TimeTypeSupport.AsCStruct(unsafe.Pointer(&mem.time_value), &m.TimeValue)
}

func (t _BuiltinsTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Builtins)
	mem := (*C.test_msgs__msg__Builtins)(ros2_message_buffer)
	builtin_interfaces_msg.DurationTypeSupport.AsGoStruct(&m.DurationValue, unsafe.Pointer(&mem.duration_value))
	builtin_interfaces_msg.TimeTypeSupport.AsGoStruct(&m.TimeValue, unsafe.Pointer(&mem.time_value))
}

func (t _BuiltinsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__msg__Builtins())
}

type CBuiltins = C.test_msgs__msg__Builtins
type CBuiltins__Sequence = C.test_msgs__msg__Builtins__Sequence

func Builtins__Sequence_to_Go(goSlice *[]Builtins, cSlice CBuiltins__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Builtins, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.test_msgs__msg__Builtins__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__msg__Builtins * uintptr(i)),
		))
		BuiltinsTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Builtins__Sequence_to_C(cSlice *CBuiltins__Sequence, goSlice []Builtins) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.test_msgs__msg__Builtins)(C.malloc((C.size_t)(C.sizeof_struct_test_msgs__msg__Builtins * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.test_msgs__msg__Builtins)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__msg__Builtins * uintptr(i)),
		))
		BuiltinsTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Builtins__Array_to_Go(goSlice []Builtins, cSlice []CBuiltins) {
	for i := 0; i < len(cSlice); i++ {
		BuiltinsTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Builtins__Array_to_C(cSlice []CBuiltins, goSlice []Builtins) {
	for i := 0; i < len(goSlice); i++ {
		BuiltinsTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
