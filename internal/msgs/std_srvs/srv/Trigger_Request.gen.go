/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package std_srvs_srv
import (
	"unsafe"

	"github.com/ATIinc/rclgo/pkg/rclgo"
	"github.com/ATIinc/rclgo/pkg/rclgo/types"
	"github.com/ATIinc/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <std_srvs/srv/trigger.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("std_srvs/Trigger_Request", Trigger_RequestTypeSupport)
	typemap.RegisterMessage("std_srvs/srv/Trigger_Request", Trigger_RequestTypeSupport)
}

type Trigger_Request struct {
}

// NewTrigger_Request creates a new Trigger_Request with default values.
func NewTrigger_Request() *Trigger_Request {
	self := Trigger_Request{}
	self.SetDefaults()
	return &self
}

func (t *Trigger_Request) Clone() *Trigger_Request {
	c := &Trigger_Request{}
	return c
}

func (t *Trigger_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Trigger_Request) SetDefaults() {
}

func (t *Trigger_Request) GetTypeSupport() types.MessageTypeSupport {
	return Trigger_RequestTypeSupport
}

// Trigger_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type Trigger_RequestPublisher struct {
	*rclgo.Publisher
}

// NewTrigger_RequestPublisher creates and returns a new publisher for the
// Trigger_Request
func NewTrigger_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*Trigger_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, Trigger_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Trigger_RequestPublisher{pub}, nil
}

func (p *Trigger_RequestPublisher) Publish(msg *Trigger_Request) error {
	return p.Publisher.Publish(msg)
}

// Trigger_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type Trigger_RequestSubscription struct {
	*rclgo.Subscription
}

// Trigger_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a Trigger_RequestSubscription.
type Trigger_RequestSubscriptionCallback func(msg *Trigger_Request, info *rclgo.MessageInfo, err error)

// NewTrigger_RequestSubscription creates and returns a new subscription for the
// Trigger_Request
func NewTrigger_RequestSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback Trigger_RequestSubscriptionCallback) (*Trigger_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Trigger_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, Trigger_RequestTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &Trigger_RequestSubscription{sub}, nil
}

func (s *Trigger_RequestSubscription) TakeMessage(out *Trigger_Request) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneTrigger_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneTrigger_RequestSlice(dst, src []Trigger_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Trigger_RequestTypeSupport types.MessageTypeSupport = _Trigger_RequestTypeSupport{}

type _Trigger_RequestTypeSupport struct{}

func (t _Trigger_RequestTypeSupport) New() types.Message {
	return NewTrigger_Request()
}

func (t _Trigger_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_srvs__srv__Trigger_Request
	return (unsafe.Pointer)(C.std_srvs__srv__Trigger_Request__create())
}

func (t _Trigger_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_srvs__srv__Trigger_Request__destroy((*C.std_srvs__srv__Trigger_Request)(pointer_to_free))
}

func (t _Trigger_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _Trigger_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _Trigger_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_srvs__srv__Trigger_Request())
}

type CTrigger_Request = C.std_srvs__srv__Trigger_Request
type CTrigger_Request__Sequence = C.std_srvs__srv__Trigger_Request__Sequence

func Trigger_Request__Sequence_to_Go(goSlice *[]Trigger_Request, cSlice CTrigger_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Trigger_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		Trigger_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Trigger_Request__Sequence_to_C(cSlice *CTrigger_Request__Sequence, goSlice []Trigger_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.std_srvs__srv__Trigger_Request)(C.malloc(C.sizeof_struct_std_srvs__srv__Trigger_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		Trigger_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Trigger_Request__Array_to_Go(goSlice []Trigger_Request, cSlice []CTrigger_Request) {
	for i := 0; i < len(cSlice); i++ {
		Trigger_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Trigger_Request__Array_to_C(cSlice []CTrigger_Request, goSlice []Trigger_Request) {
	for i := 0; i < len(goSlice); i++ {
		Trigger_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
