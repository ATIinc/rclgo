/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package geometry_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/transform.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/Transform", TransformTypeSupport)
}

// Do not create instances of this type directly. Always use NewTransform
// function instead.
type Transform struct {
	Translation Vector3 `yaml:"translation"`
	Rotation Quaternion `yaml:"rotation"`
}

// NewTransform creates a new Transform with default values.
func NewTransform() *Transform {
	self := Transform{}
	self.SetDefaults()
	return &self
}

func (t *Transform) Clone() *Transform {
	c := &Transform{}
	c.Translation = *t.Translation.Clone()
	c.Rotation = *t.Rotation.Clone()
	return c
}

func (t *Transform) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Transform) SetDefaults() {
	t.Translation.SetDefaults()
	t.Rotation.SetDefaults()
}

// TransformPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type TransformPublisher struct {
	*rclgo.Publisher
}

// NewTransformPublisher creates and returns a new publisher for the
// Transform
func NewTransformPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*TransformPublisher, error) {
	pub, err := node.NewPublisher(topic_name, TransformTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &TransformPublisher{pub}, nil
}

func (p *TransformPublisher) Publish(msg *Transform) error {
	return p.Publisher.Publish(msg)
}

// TransformSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type TransformSubscription struct {
	*rclgo.Subscription
}

// TransformSubscriptionCallback type is used to provide a subscription
// handler function for a TransformSubscription.
type TransformSubscriptionCallback func(msg *Transform, info *rclgo.RmwMessageInfo, err error)

// NewTransformSubscription creates and returns a new subscription for the
// Transform
func NewTransformSubscription(node *rclgo.Node, topic_name string, subscriptionCallback TransformSubscriptionCallback) (*TransformSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Transform
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, TransformTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &TransformSubscription{sub}, nil
}

func (s *TransformSubscription) TakeMessage(out *Transform) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneTransformSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneTransformSlice(dst, src []Transform) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var TransformTypeSupport types.MessageTypeSupport = _TransformTypeSupport{}

type _TransformTypeSupport struct{}

func (t _TransformTypeSupport) New() types.Message {
	return NewTransform()
}

func (t _TransformTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__Transform
	return (unsafe.Pointer)(C.geometry_msgs__msg__Transform__create())
}

func (t _TransformTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__Transform__destroy((*C.geometry_msgs__msg__Transform)(pointer_to_free))
}

func (t _TransformTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Transform)
	mem := (*C.geometry_msgs__msg__Transform)(dst)
	Vector3TypeSupport.AsCStruct(unsafe.Pointer(&mem.translation), &m.Translation)
	QuaternionTypeSupport.AsCStruct(unsafe.Pointer(&mem.rotation), &m.Rotation)
}

func (t _TransformTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Transform)
	mem := (*C.geometry_msgs__msg__Transform)(ros2_message_buffer)
	Vector3TypeSupport.AsGoStruct(&m.Translation, unsafe.Pointer(&mem.translation))
	QuaternionTypeSupport.AsGoStruct(&m.Rotation, unsafe.Pointer(&mem.rotation))
}

func (t _TransformTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__Transform())
}

type CTransform = C.geometry_msgs__msg__Transform
type CTransform__Sequence = C.geometry_msgs__msg__Transform__Sequence

func Transform__Sequence_to_Go(goSlice *[]Transform, cSlice CTransform__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Transform, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__Transform__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Transform * uintptr(i)),
		))
		TransformTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Transform__Sequence_to_C(cSlice *CTransform__Sequence, goSlice []Transform) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__Transform)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__Transform * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__Transform)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Transform * uintptr(i)),
		))
		TransformTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Transform__Array_to_Go(goSlice []Transform, cSlice []CTransform) {
	for i := 0; i < len(cSlice); i++ {
		TransformTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Transform__Array_to_C(cSlice []CTransform, goSlice []Transform) {
	for i := 0; i < len(goSlice); i++ {
		TransformTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
