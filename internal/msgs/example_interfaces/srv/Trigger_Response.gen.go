/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package example_interfaces_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/srv/trigger.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/Trigger_Response", Trigger_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewTrigger_Response
// function instead.
type Trigger_Response struct {
	Success bool `yaml:"success"`// indicate successful run of triggered service
	Message string `yaml:"message"`// informational, e.g. for error messages.
}

// NewTrigger_Response creates a new Trigger_Response with default values.
func NewTrigger_Response() *Trigger_Response {
	self := Trigger_Response{}
	self.SetDefaults()
	return &self
}

func (t *Trigger_Response) Clone() *Trigger_Response {
	c := &Trigger_Response{}
	c.Success = t.Success
	c.Message = t.Message
	return c
}

func (t *Trigger_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Trigger_Response) SetDefaults() {
	t.Success = false
	t.Message = ""
}

// Trigger_ResponsePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type Trigger_ResponsePublisher struct {
	*rclgo.Publisher
}

// NewTrigger_ResponsePublisher creates and returns a new publisher for the
// Trigger_Response
func NewTrigger_ResponsePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*Trigger_ResponsePublisher, error) {
	pub, err := node.NewPublisher(topic_name, Trigger_ResponseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Trigger_ResponsePublisher{pub}, nil
}

func (p *Trigger_ResponsePublisher) Publish(msg *Trigger_Response) error {
	return p.Publisher.Publish(msg)
}

// Trigger_ResponseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type Trigger_ResponseSubscription struct {
	*rclgo.Subscription
}

// Trigger_ResponseSubscriptionCallback type is used to provide a subscription
// handler function for a Trigger_ResponseSubscription.
type Trigger_ResponseSubscriptionCallback func(msg *Trigger_Response, info *rclgo.RmwMessageInfo, err error)

// NewTrigger_ResponseSubscription creates and returns a new subscription for the
// Trigger_Response
func NewTrigger_ResponseSubscription(node *rclgo.Node, topic_name string, subscriptionCallback Trigger_ResponseSubscriptionCallback) (*Trigger_ResponseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Trigger_Response
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, Trigger_ResponseTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &Trigger_ResponseSubscription{sub}, nil
}

func (s *Trigger_ResponseSubscription) TakeMessage(out *Trigger_Response) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneTrigger_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneTrigger_ResponseSlice(dst, src []Trigger_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Trigger_ResponseTypeSupport types.MessageTypeSupport = _Trigger_ResponseTypeSupport{}

type _Trigger_ResponseTypeSupport struct{}

func (t _Trigger_ResponseTypeSupport) New() types.Message {
	return NewTrigger_Response()
}

func (t _Trigger_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__srv__Trigger_Response
	return (unsafe.Pointer)(C.example_interfaces__srv__Trigger_Response__create())
}

func (t _Trigger_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__srv__Trigger_Response__destroy((*C.example_interfaces__srv__Trigger_Response)(pointer_to_free))
}

func (t _Trigger_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Trigger_Response)
	mem := (*C.example_interfaces__srv__Trigger_Response)(dst)
	mem.success = C.bool(m.Success)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.message), m.Message)
}

func (t _Trigger_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Trigger_Response)
	mem := (*C.example_interfaces__srv__Trigger_Response)(ros2_message_buffer)
	m.Success = bool(mem.success)
	primitives.StringAsGoStruct(&m.Message, unsafe.Pointer(&mem.message))
}

func (t _Trigger_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__srv__Trigger_Response())
}

type CTrigger_Response = C.example_interfaces__srv__Trigger_Response
type CTrigger_Response__Sequence = C.example_interfaces__srv__Trigger_Response__Sequence

func Trigger_Response__Sequence_to_Go(goSlice *[]Trigger_Response, cSlice CTrigger_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Trigger_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__srv__Trigger_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__srv__Trigger_Response * uintptr(i)),
		))
		Trigger_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Trigger_Response__Sequence_to_C(cSlice *CTrigger_Response__Sequence, goSlice []Trigger_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__srv__Trigger_Response)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__srv__Trigger_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__srv__Trigger_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__srv__Trigger_Response * uintptr(i)),
		))
		Trigger_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Trigger_Response__Array_to_Go(goSlice []Trigger_Response, cSlice []CTrigger_Response) {
	for i := 0; i < len(cSlice); i++ {
		Trigger_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Trigger_Response__Array_to_C(cSlice []CTrigger_Response, goSlice []Trigger_Response) {
	for i := 0; i < len(goSlice); i++ {
		Trigger_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
