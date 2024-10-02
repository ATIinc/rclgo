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

	"github.com/ATIinc/rclgo/pkg/rclgo"
	"github.com/ATIinc/rclgo/pkg/rclgo/types"
	"github.com/ATIinc/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/srv/empty.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/Empty_Response", Empty_ResponseTypeSupport)
	typemap.RegisterMessage("test_msgs/srv/Empty_Response", Empty_ResponseTypeSupport)
}

type Empty_Response struct {
}

// NewEmpty_Response creates a new Empty_Response with default values.
func NewEmpty_Response() *Empty_Response {
	self := Empty_Response{}
	self.SetDefaults()
	return &self
}

func (t *Empty_Response) Clone() *Empty_Response {
	c := &Empty_Response{}
	return c
}

func (t *Empty_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Empty_Response) SetDefaults() {
}

func (t *Empty_Response) GetTypeSupport() types.MessageTypeSupport {
	return Empty_ResponseTypeSupport
}

// Empty_ResponsePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type Empty_ResponsePublisher struct {
	*rclgo.Publisher
}

// NewEmpty_ResponsePublisher creates and returns a new publisher for the
// Empty_Response
func NewEmpty_ResponsePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*Empty_ResponsePublisher, error) {
	pub, err := node.NewPublisher(topic_name, Empty_ResponseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Empty_ResponsePublisher{pub}, nil
}

func (p *Empty_ResponsePublisher) Publish(msg *Empty_Response) error {
	return p.Publisher.Publish(msg)
}

// Empty_ResponseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type Empty_ResponseSubscription struct {
	*rclgo.Subscription
}

// Empty_ResponseSubscriptionCallback type is used to provide a subscription
// handler function for a Empty_ResponseSubscription.
type Empty_ResponseSubscriptionCallback func(msg *Empty_Response, info *rclgo.MessageInfo, err error)

// NewEmpty_ResponseSubscription creates and returns a new subscription for the
// Empty_Response
func NewEmpty_ResponseSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback Empty_ResponseSubscriptionCallback) (*Empty_ResponseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Empty_Response
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, Empty_ResponseTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &Empty_ResponseSubscription{sub}, nil
}

func (s *Empty_ResponseSubscription) TakeMessage(out *Empty_Response) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneEmpty_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneEmpty_ResponseSlice(dst, src []Empty_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Empty_ResponseTypeSupport types.MessageTypeSupport = _Empty_ResponseTypeSupport{}

type _Empty_ResponseTypeSupport struct{}

func (t _Empty_ResponseTypeSupport) New() types.Message {
	return NewEmpty_Response()
}

func (t _Empty_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__srv__Empty_Response
	return (unsafe.Pointer)(C.test_msgs__srv__Empty_Response__create())
}

func (t _Empty_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__srv__Empty_Response__destroy((*C.test_msgs__srv__Empty_Response)(pointer_to_free))
}

func (t _Empty_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _Empty_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _Empty_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__srv__Empty_Response())
}

type CEmpty_Response = C.test_msgs__srv__Empty_Response
type CEmpty_Response__Sequence = C.test_msgs__srv__Empty_Response__Sequence

func Empty_Response__Sequence_to_Go(goSlice *[]Empty_Response, cSlice CEmpty_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Empty_Response, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		Empty_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Empty_Response__Sequence_to_C(cSlice *CEmpty_Response__Sequence, goSlice []Empty_Response) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.test_msgs__srv__Empty_Response)(C.malloc(C.sizeof_struct_test_msgs__srv__Empty_Response * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		Empty_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Empty_Response__Array_to_Go(goSlice []Empty_Response, cSlice []CEmpty_Response) {
	for i := 0; i < len(cSlice); i++ {
		Empty_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Empty_Response__Array_to_C(cSlice []CEmpty_Response, goSlice []Empty_Response) {
	for i := 0; i < len(goSlice); i++ {
		Empty_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
