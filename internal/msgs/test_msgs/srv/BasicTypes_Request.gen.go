/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package test_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/srv/basic_types.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/BasicTypes_Request", BasicTypes_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewBasicTypes_Request
// function instead.
type BasicTypes_Request struct {
	BoolValue bool `yaml:"bool_value"`
	ByteValue byte `yaml:"byte_value"`
	CharValue byte `yaml:"char_value"`
	Float32Value float32 `yaml:"float32_value"`
	Float64Value float64 `yaml:"float64_value"`
	Int8Value int8 `yaml:"int8_value"`
	Uint8Value uint8 `yaml:"uint8_value"`
	Int16Value int16 `yaml:"int16_value"`
	Uint16Value uint16 `yaml:"uint16_value"`
	Int32Value int32 `yaml:"int32_value"`
	Uint32Value uint32 `yaml:"uint32_value"`
	Int64Value int64 `yaml:"int64_value"`
	Uint64Value uint64 `yaml:"uint64_value"`
	StringValue string `yaml:"string_value"`
}

// NewBasicTypes_Request creates a new BasicTypes_Request with default values.
func NewBasicTypes_Request() *BasicTypes_Request {
	self := BasicTypes_Request{}
	self.SetDefaults()
	return &self
}

func (t *BasicTypes_Request) Clone() *BasicTypes_Request {
	c := &BasicTypes_Request{}
	c.BoolValue = t.BoolValue
	c.ByteValue = t.ByteValue
	c.CharValue = t.CharValue
	c.Float32Value = t.Float32Value
	c.Float64Value = t.Float64Value
	c.Int8Value = t.Int8Value
	c.Uint8Value = t.Uint8Value
	c.Int16Value = t.Int16Value
	c.Uint16Value = t.Uint16Value
	c.Int32Value = t.Int32Value
	c.Uint32Value = t.Uint32Value
	c.Int64Value = t.Int64Value
	c.Uint64Value = t.Uint64Value
	c.StringValue = t.StringValue
	return c
}

func (t *BasicTypes_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *BasicTypes_Request) SetDefaults() {
	t.BoolValue = false
	t.ByteValue = 0
	t.CharValue = 0
	t.Float32Value = 0
	t.Float64Value = 0
	t.Int8Value = 0
	t.Uint8Value = 0
	t.Int16Value = 0
	t.Uint16Value = 0
	t.Int32Value = 0
	t.Uint32Value = 0
	t.Int64Value = 0
	t.Uint64Value = 0
	t.StringValue = ""
}

// BasicTypes_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type BasicTypes_RequestPublisher struct {
	*rclgo.Publisher
}

// NewBasicTypes_RequestPublisher creates and returns a new publisher for the
// BasicTypes_Request
func NewBasicTypes_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*BasicTypes_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, BasicTypes_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &BasicTypes_RequestPublisher{pub}, nil
}

func (p *BasicTypes_RequestPublisher) Publish(msg *BasicTypes_Request) error {
	return p.Publisher.Publish(msg)
}

// BasicTypes_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type BasicTypes_RequestSubscription struct {
	*rclgo.Subscription
}

// BasicTypes_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a BasicTypes_RequestSubscription.
type BasicTypes_RequestSubscriptionCallback func(msg *BasicTypes_Request, info *rclgo.RmwMessageInfo, err error)

// NewBasicTypes_RequestSubscription creates and returns a new subscription for the
// BasicTypes_Request
func NewBasicTypes_RequestSubscription(node *rclgo.Node, topic_name string, subscriptionCallback BasicTypes_RequestSubscriptionCallback) (*BasicTypes_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg BasicTypes_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, BasicTypes_RequestTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &BasicTypes_RequestSubscription{sub}, nil
}

func (s *BasicTypes_RequestSubscription) TakeMessage(out *BasicTypes_Request) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneBasicTypes_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneBasicTypes_RequestSlice(dst, src []BasicTypes_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var BasicTypes_RequestTypeSupport types.MessageTypeSupport = _BasicTypes_RequestTypeSupport{}

type _BasicTypes_RequestTypeSupport struct{}

func (t _BasicTypes_RequestTypeSupport) New() types.Message {
	return NewBasicTypes_Request()
}

func (t _BasicTypes_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__srv__BasicTypes_Request
	return (unsafe.Pointer)(C.test_msgs__srv__BasicTypes_Request__create())
}

func (t _BasicTypes_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__srv__BasicTypes_Request__destroy((*C.test_msgs__srv__BasicTypes_Request)(pointer_to_free))
}

func (t _BasicTypes_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*BasicTypes_Request)
	mem := (*C.test_msgs__srv__BasicTypes_Request)(dst)
	mem.bool_value = C.bool(m.BoolValue)
	mem.byte_value = C.uint8_t(m.ByteValue)
	mem.char_value = C.uchar(m.CharValue)
	mem.float32_value = C.float(m.Float32Value)
	mem.float64_value = C.double(m.Float64Value)
	mem.int8_value = C.int8_t(m.Int8Value)
	mem.uint8_value = C.uint8_t(m.Uint8Value)
	mem.int16_value = C.int16_t(m.Int16Value)
	mem.uint16_value = C.uint16_t(m.Uint16Value)
	mem.int32_value = C.int32_t(m.Int32Value)
	mem.uint32_value = C.uint32_t(m.Uint32Value)
	mem.int64_value = C.int64_t(m.Int64Value)
	mem.uint64_value = C.uint64_t(m.Uint64Value)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.string_value), m.StringValue)
}

func (t _BasicTypes_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*BasicTypes_Request)
	mem := (*C.test_msgs__srv__BasicTypes_Request)(ros2_message_buffer)
	m.BoolValue = bool(mem.bool_value)
	m.ByteValue = byte(mem.byte_value)
	m.CharValue = byte(mem.char_value)
	m.Float32Value = float32(mem.float32_value)
	m.Float64Value = float64(mem.float64_value)
	m.Int8Value = int8(mem.int8_value)
	m.Uint8Value = uint8(mem.uint8_value)
	m.Int16Value = int16(mem.int16_value)
	m.Uint16Value = uint16(mem.uint16_value)
	m.Int32Value = int32(mem.int32_value)
	m.Uint32Value = uint32(mem.uint32_value)
	m.Int64Value = int64(mem.int64_value)
	m.Uint64Value = uint64(mem.uint64_value)
	primitives.StringAsGoStruct(&m.StringValue, unsafe.Pointer(&mem.string_value))
}

func (t _BasicTypes_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__srv__BasicTypes_Request())
}

type CBasicTypes_Request = C.test_msgs__srv__BasicTypes_Request
type CBasicTypes_Request__Sequence = C.test_msgs__srv__BasicTypes_Request__Sequence

func BasicTypes_Request__Sequence_to_Go(goSlice *[]BasicTypes_Request, cSlice CBasicTypes_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]BasicTypes_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		BasicTypes_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func BasicTypes_Request__Sequence_to_C(cSlice *CBasicTypes_Request__Sequence, goSlice []BasicTypes_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.test_msgs__srv__BasicTypes_Request)(C.malloc(C.sizeof_struct_test_msgs__srv__BasicTypes_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		BasicTypes_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func BasicTypes_Request__Array_to_Go(goSlice []BasicTypes_Request, cSlice []CBasicTypes_Request) {
	for i := 0; i < len(cSlice); i++ {
		BasicTypes_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func BasicTypes_Request__Array_to_C(cSlice []CBasicTypes_Request, goSlice []BasicTypes_Request) {
	for i := 0; i < len(goSlice); i++ {
		BasicTypes_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
