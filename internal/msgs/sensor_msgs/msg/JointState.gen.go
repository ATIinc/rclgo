/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package sensor_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/std_msgs/msg"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/joint_state.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/JointState", JointStateTypeSupport)
	typemap.RegisterMessage("sensor_msgs/msg/JointState", JointStateTypeSupport)
}

type JointState struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Name []string `yaml:"name"`
	Position []float64 `yaml:"position"`
	Velocity []float64 `yaml:"velocity"`
	Effort []float64 `yaml:"effort"`
}

// NewJointState creates a new JointState with default values.
func NewJointState() *JointState {
	self := JointState{}
	self.SetDefaults()
	return &self
}

func (t *JointState) Clone() *JointState {
	c := &JointState{}
	c.Header = *t.Header.Clone()
	if t.Name != nil {
		c.Name = make([]string, len(t.Name))
		copy(c.Name, t.Name)
	}
	if t.Position != nil {
		c.Position = make([]float64, len(t.Position))
		copy(c.Position, t.Position)
	}
	if t.Velocity != nil {
		c.Velocity = make([]float64, len(t.Velocity))
		copy(c.Velocity, t.Velocity)
	}
	if t.Effort != nil {
		c.Effort = make([]float64, len(t.Effort))
		copy(c.Effort, t.Effort)
	}
	return c
}

func (t *JointState) CloneMsg() types.Message {
	return t.Clone()
}

func (t *JointState) SetDefaults() {
	t.Header.SetDefaults()
	t.Name = nil
	t.Position = nil
	t.Velocity = nil
	t.Effort = nil
}

func (t *JointState) GetTypeSupport() types.MessageTypeSupport {
	return JointStateTypeSupport
}

// JointStatePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type JointStatePublisher struct {
	*rclgo.Publisher
}

// NewJointStatePublisher creates and returns a new publisher for the
// JointState
func NewJointStatePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*JointStatePublisher, error) {
	pub, err := node.NewPublisher(topic_name, JointStateTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &JointStatePublisher{pub}, nil
}

func (p *JointStatePublisher) Publish(msg *JointState) error {
	return p.Publisher.Publish(msg)
}

// JointStateSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type JointStateSubscription struct {
	*rclgo.Subscription
}

// JointStateSubscriptionCallback type is used to provide a subscription
// handler function for a JointStateSubscription.
type JointStateSubscriptionCallback func(msg *JointState, info *rclgo.MessageInfo, err error)

// NewJointStateSubscription creates and returns a new subscription for the
// JointState
func NewJointStateSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback JointStateSubscriptionCallback) (*JointStateSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg JointState
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, JointStateTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &JointStateSubscription{sub}, nil
}

func (s *JointStateSubscription) TakeMessage(out *JointState) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneJointStateSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneJointStateSlice(dst, src []JointState) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var JointStateTypeSupport types.MessageTypeSupport = _JointStateTypeSupport{}

type _JointStateTypeSupport struct{}

func (t _JointStateTypeSupport) New() types.Message {
	return NewJointState()
}

func (t _JointStateTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__JointState
	return (unsafe.Pointer)(C.sensor_msgs__msg__JointState__create())
}

func (t _JointStateTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__JointState__destroy((*C.sensor_msgs__msg__JointState)(pointer_to_free))
}

func (t _JointStateTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*JointState)
	mem := (*C.sensor_msgs__msg__JointState)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	primitives.String__Sequence_to_C((*primitives.CString__Sequence)(unsafe.Pointer(&mem.name)), m.Name)
	primitives.Float64__Sequence_to_C((*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.position)), m.Position)
	primitives.Float64__Sequence_to_C((*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.velocity)), m.Velocity)
	primitives.Float64__Sequence_to_C((*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.effort)), m.Effort)
}

func (t _JointStateTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*JointState)
	mem := (*C.sensor_msgs__msg__JointState)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	primitives.String__Sequence_to_Go(&m.Name, *(*primitives.CString__Sequence)(unsafe.Pointer(&mem.name)))
	primitives.Float64__Sequence_to_Go(&m.Position, *(*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.position)))
	primitives.Float64__Sequence_to_Go(&m.Velocity, *(*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.velocity)))
	primitives.Float64__Sequence_to_Go(&m.Effort, *(*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.effort)))
}

func (t _JointStateTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__JointState())
}

type CJointState = C.sensor_msgs__msg__JointState
type CJointState__Sequence = C.sensor_msgs__msg__JointState__Sequence

func JointState__Sequence_to_Go(goSlice *[]JointState, cSlice CJointState__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]JointState, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		JointStateTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func JointState__Sequence_to_C(cSlice *CJointState__Sequence, goSlice []JointState) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__JointState)(C.malloc(C.sizeof_struct_sensor_msgs__msg__JointState * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		JointStateTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func JointState__Array_to_Go(goSlice []JointState, cSlice []CJointState) {
	for i := 0; i < len(cSlice); i++ {
		JointStateTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func JointState__Array_to_C(cSlice []CJointState, goSlice []JointState) {
	for i := 0; i < len(goSlice); i++ {
		JointStateTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
