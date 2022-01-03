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

#include <geometry_msgs/msg/twist_with_covariance_stamped.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/TwistWithCovarianceStamped", TwistWithCovarianceStampedTypeSupport)
}

// Do not create instances of this type directly. Always use NewTwistWithCovarianceStamped
// function instead.
type TwistWithCovarianceStamped struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Twist TwistWithCovariance `yaml:"twist"`
}

// NewTwistWithCovarianceStamped creates a new TwistWithCovarianceStamped with default values.
func NewTwistWithCovarianceStamped() *TwistWithCovarianceStamped {
	self := TwistWithCovarianceStamped{}
	self.SetDefaults()
	return &self
}

func (t *TwistWithCovarianceStamped) Clone() *TwistWithCovarianceStamped {
	c := &TwistWithCovarianceStamped{}
	c.Header = *t.Header.Clone()
	c.Twist = *t.Twist.Clone()
	return c
}

func (t *TwistWithCovarianceStamped) CloneMsg() types.Message {
	return t.Clone()
}

func (t *TwistWithCovarianceStamped) SetDefaults() {
	t.Header.SetDefaults()
	t.Twist.SetDefaults()
}

// TwistWithCovarianceStampedPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type TwistWithCovarianceStampedPublisher struct {
	*rclgo.Publisher
}

// NewTwistWithCovarianceStampedPublisher creates and returns a new publisher for the
// TwistWithCovarianceStamped
func NewTwistWithCovarianceStampedPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*TwistWithCovarianceStampedPublisher, error) {
	pub, err := node.NewPublisher(topic_name, TwistWithCovarianceStampedTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &TwistWithCovarianceStampedPublisher{pub}, nil
}

func (p *TwistWithCovarianceStampedPublisher) Publish(msg *TwistWithCovarianceStamped) error {
	return p.Publisher.Publish(msg)
}

// TwistWithCovarianceStampedSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type TwistWithCovarianceStampedSubscription struct {
	*rclgo.Subscription
}

// TwistWithCovarianceStampedSubscriptionCallback type is used to provide a subscription
// handler function for a TwistWithCovarianceStampedSubscription.
type TwistWithCovarianceStampedSubscriptionCallback func(msg *TwistWithCovarianceStamped, info *rclgo.RmwMessageInfo, err error)

// NewTwistWithCovarianceStampedSubscription creates and returns a new subscription for the
// TwistWithCovarianceStamped
func NewTwistWithCovarianceStampedSubscription(node *rclgo.Node, topic_name string, subscriptionCallback TwistWithCovarianceStampedSubscriptionCallback) (*TwistWithCovarianceStampedSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg TwistWithCovarianceStamped
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, TwistWithCovarianceStampedTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &TwistWithCovarianceStampedSubscription{sub}, nil
}

func (s *TwistWithCovarianceStampedSubscription) TakeMessage(out *TwistWithCovarianceStamped) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}


// CloneTwistWithCovarianceStampedSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneTwistWithCovarianceStampedSlice(dst, src []TwistWithCovarianceStamped) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var TwistWithCovarianceStampedTypeSupport types.MessageTypeSupport = _TwistWithCovarianceStampedTypeSupport{}

type _TwistWithCovarianceStampedTypeSupport struct{}

func (t _TwistWithCovarianceStampedTypeSupport) New() types.Message {
	return NewTwistWithCovarianceStamped()
}

func (t _TwistWithCovarianceStampedTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__TwistWithCovarianceStamped
	return (unsafe.Pointer)(C.geometry_msgs__msg__TwistWithCovarianceStamped__create())
}

func (t _TwistWithCovarianceStampedTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__TwistWithCovarianceStamped__destroy((*C.geometry_msgs__msg__TwistWithCovarianceStamped)(pointer_to_free))
}

func (t _TwistWithCovarianceStampedTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*TwistWithCovarianceStamped)
	mem := (*C.geometry_msgs__msg__TwistWithCovarianceStamped)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	TwistWithCovarianceTypeSupport.AsCStruct(unsafe.Pointer(&mem.twist), &m.Twist)
}

func (t _TwistWithCovarianceStampedTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*TwistWithCovarianceStamped)
	mem := (*C.geometry_msgs__msg__TwistWithCovarianceStamped)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	TwistWithCovarianceTypeSupport.AsGoStruct(&m.Twist, unsafe.Pointer(&mem.twist))
}

func (t _TwistWithCovarianceStampedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__TwistWithCovarianceStamped())
}

type CTwistWithCovarianceStamped = C.geometry_msgs__msg__TwistWithCovarianceStamped
type CTwistWithCovarianceStamped__Sequence = C.geometry_msgs__msg__TwistWithCovarianceStamped__Sequence

func TwistWithCovarianceStamped__Sequence_to_Go(goSlice *[]TwistWithCovarianceStamped, cSlice CTwistWithCovarianceStamped__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TwistWithCovarianceStamped, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__TwistWithCovarianceStamped__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__TwistWithCovarianceStamped * uintptr(i)),
		))
		TwistWithCovarianceStampedTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func TwistWithCovarianceStamped__Sequence_to_C(cSlice *CTwistWithCovarianceStamped__Sequence, goSlice []TwistWithCovarianceStamped) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__TwistWithCovarianceStamped)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__TwistWithCovarianceStamped * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__TwistWithCovarianceStamped)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__TwistWithCovarianceStamped * uintptr(i)),
		))
		TwistWithCovarianceStampedTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func TwistWithCovarianceStamped__Array_to_Go(goSlice []TwistWithCovarianceStamped, cSlice []CTwistWithCovarianceStamped) {
	for i := 0; i < len(cSlice); i++ {
		TwistWithCovarianceStampedTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func TwistWithCovarianceStamped__Array_to_C(cSlice []CTwistWithCovarianceStamped, goSlice []TwistWithCovarianceStamped) {
	for i := 0; i < len(goSlice); i++ {
		TwistWithCovarianceStampedTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
