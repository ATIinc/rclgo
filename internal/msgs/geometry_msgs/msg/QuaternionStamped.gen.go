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
	std_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/quaternion_stamped.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/QuaternionStamped", QuaternionStampedTypeSupport)
}

// Do not create instances of this type directly. Always use NewQuaternionStamped
// function instead.
type QuaternionStamped struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Quaternion Quaternion `yaml:"quaternion"`
}

// NewQuaternionStamped creates a new QuaternionStamped with default values.
func NewQuaternionStamped() *QuaternionStamped {
	self := QuaternionStamped{}
	self.SetDefaults()
	return &self
}

func (t *QuaternionStamped) Clone() *QuaternionStamped {
	c := &QuaternionStamped{}
	c.Header = *t.Header.Clone()
	c.Quaternion = *t.Quaternion.Clone()
	return c
}

func (t *QuaternionStamped) CloneMsg() types.Message {
	return t.Clone()
}

func (t *QuaternionStamped) SetDefaults() {
	t.Header.SetDefaults()
	t.Quaternion.SetDefaults()
}

// QuaternionStampedPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type QuaternionStampedPublisher struct {
	*rclgo.Publisher
}

// NewQuaternionStampedPublisher creates and returns a new publisher for the
// QuaternionStamped
func NewQuaternionStampedPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*QuaternionStampedPublisher, error) {
	pub, err := node.NewPublisher(topic_name, QuaternionStampedTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &QuaternionStampedPublisher{pub}, nil
}

func (p *QuaternionStampedPublisher) Publish(msg *QuaternionStamped) error {
	return p.Publisher.Publish(msg)
}

// QuaternionStampedSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type QuaternionStampedSubscription struct {
	*rclgo.Subscription
}

// QuaternionStampedSubscriptionCallback type is used to provide a subscription
// handler function for a QuaternionStampedSubscription.
type QuaternionStampedSubscriptionCallback func(msg *QuaternionStamped, info *rclgo.RmwMessageInfo, err error)

// NewQuaternionStampedSubscription creates and returns a new subscription for the
// QuaternionStamped
func NewQuaternionStampedSubscription(node *rclgo.Node, topic_name string, subscriptionCallback QuaternionStampedSubscriptionCallback) (*QuaternionStampedSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg QuaternionStamped
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, QuaternionStampedTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &QuaternionStampedSubscription{sub}, nil
}

func (s *QuaternionStampedSubscription) TakeMessage(out *QuaternionStamped) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}


// CloneQuaternionStampedSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneQuaternionStampedSlice(dst, src []QuaternionStamped) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var QuaternionStampedTypeSupport types.MessageTypeSupport = _QuaternionStampedTypeSupport{}

type _QuaternionStampedTypeSupport struct{}

func (t _QuaternionStampedTypeSupport) New() types.Message {
	return NewQuaternionStamped()
}

func (t _QuaternionStampedTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__QuaternionStamped
	return (unsafe.Pointer)(C.geometry_msgs__msg__QuaternionStamped__create())
}

func (t _QuaternionStampedTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__QuaternionStamped__destroy((*C.geometry_msgs__msg__QuaternionStamped)(pointer_to_free))
}

func (t _QuaternionStampedTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*QuaternionStamped)
	mem := (*C.geometry_msgs__msg__QuaternionStamped)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	QuaternionTypeSupport.AsCStruct(unsafe.Pointer(&mem.quaternion), &m.Quaternion)
}

func (t _QuaternionStampedTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*QuaternionStamped)
	mem := (*C.geometry_msgs__msg__QuaternionStamped)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	QuaternionTypeSupport.AsGoStruct(&m.Quaternion, unsafe.Pointer(&mem.quaternion))
}

func (t _QuaternionStampedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__QuaternionStamped())
}

type CQuaternionStamped = C.geometry_msgs__msg__QuaternionStamped
type CQuaternionStamped__Sequence = C.geometry_msgs__msg__QuaternionStamped__Sequence

func QuaternionStamped__Sequence_to_Go(goSlice *[]QuaternionStamped, cSlice CQuaternionStamped__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]QuaternionStamped, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__QuaternionStamped__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__QuaternionStamped * uintptr(i)),
		))
		QuaternionStampedTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func QuaternionStamped__Sequence_to_C(cSlice *CQuaternionStamped__Sequence, goSlice []QuaternionStamped) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__QuaternionStamped)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__QuaternionStamped * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__QuaternionStamped)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__QuaternionStamped * uintptr(i)),
		))
		QuaternionStampedTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func QuaternionStamped__Array_to_Go(goSlice []QuaternionStamped, cSlice []CQuaternionStamped) {
	for i := 0; i < len(cSlice); i++ {
		QuaternionStampedTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func QuaternionStamped__Array_to_C(cSlice []CQuaternionStamped, goSlice []QuaternionStamped) {
	for i := 0; i < len(goSlice); i++ {
		QuaternionStampedTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
