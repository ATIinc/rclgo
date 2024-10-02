/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package action_msgs_msg
import (
	"unsafe"

	"github.com/ATIinc/rclgo/pkg/rclgo"
	"github.com/ATIinc/rclgo/pkg/rclgo/types"
	"github.com/ATIinc/rclgo/pkg/rclgo/typemap"
	builtin_interfaces_msg "github.com/ATIinc/rclgo/internal/msgs/builtin_interfaces/msg"
	unique_identifier_msgs_msg "github.com/ATIinc/rclgo/internal/msgs/unique_identifier_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <action_msgs/msg/goal_info.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("action_msgs/GoalInfo", GoalInfoTypeSupport)
	typemap.RegisterMessage("action_msgs/msg/GoalInfo", GoalInfoTypeSupport)
}

type GoalInfo struct {
	GoalId unique_identifier_msgs_msg.UUID `yaml:"goal_id"`// Goal ID
	Stamp builtin_interfaces_msg.Time `yaml:"stamp"`// Time when the goal was accepted
}

// NewGoalInfo creates a new GoalInfo with default values.
func NewGoalInfo() *GoalInfo {
	self := GoalInfo{}
	self.SetDefaults()
	return &self
}

func (t *GoalInfo) Clone() *GoalInfo {
	c := &GoalInfo{}
	c.GoalId = *t.GoalId.Clone()
	c.Stamp = *t.Stamp.Clone()
	return c
}

func (t *GoalInfo) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GoalInfo) SetDefaults() {
	t.GoalId.SetDefaults()
	t.Stamp.SetDefaults()
}

func (t *GoalInfo) GetTypeSupport() types.MessageTypeSupport {
	return GoalInfoTypeSupport
}

// GoalInfoPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type GoalInfoPublisher struct {
	*rclgo.Publisher
}

// NewGoalInfoPublisher creates and returns a new publisher for the
// GoalInfo
func NewGoalInfoPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*GoalInfoPublisher, error) {
	pub, err := node.NewPublisher(topic_name, GoalInfoTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &GoalInfoPublisher{pub}, nil
}

func (p *GoalInfoPublisher) Publish(msg *GoalInfo) error {
	return p.Publisher.Publish(msg)
}

// GoalInfoSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type GoalInfoSubscription struct {
	*rclgo.Subscription
}

// GoalInfoSubscriptionCallback type is used to provide a subscription
// handler function for a GoalInfoSubscription.
type GoalInfoSubscriptionCallback func(msg *GoalInfo, info *rclgo.MessageInfo, err error)

// NewGoalInfoSubscription creates and returns a new subscription for the
// GoalInfo
func NewGoalInfoSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback GoalInfoSubscriptionCallback) (*GoalInfoSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg GoalInfo
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, GoalInfoTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &GoalInfoSubscription{sub}, nil
}

func (s *GoalInfoSubscription) TakeMessage(out *GoalInfo) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneGoalInfoSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGoalInfoSlice(dst, src []GoalInfo) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GoalInfoTypeSupport types.MessageTypeSupport = _GoalInfoTypeSupport{}

type _GoalInfoTypeSupport struct{}

func (t _GoalInfoTypeSupport) New() types.Message {
	return NewGoalInfo()
}

func (t _GoalInfoTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.action_msgs__msg__GoalInfo
	return (unsafe.Pointer)(C.action_msgs__msg__GoalInfo__create())
}

func (t _GoalInfoTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.action_msgs__msg__GoalInfo__destroy((*C.action_msgs__msg__GoalInfo)(pointer_to_free))
}

func (t _GoalInfoTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GoalInfo)
	mem := (*C.action_msgs__msg__GoalInfo)(dst)
	unique_identifier_msgs_msg.UUIDTypeSupport.AsCStruct(unsafe.Pointer(&mem.goal_id), &m.GoalId)
	builtin_interfaces_msg.TimeTypeSupport.AsCStruct(unsafe.Pointer(&mem.stamp), &m.Stamp)
}

func (t _GoalInfoTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GoalInfo)
	mem := (*C.action_msgs__msg__GoalInfo)(ros2_message_buffer)
	unique_identifier_msgs_msg.UUIDTypeSupport.AsGoStruct(&m.GoalId, unsafe.Pointer(&mem.goal_id))
	builtin_interfaces_msg.TimeTypeSupport.AsGoStruct(&m.Stamp, unsafe.Pointer(&mem.stamp))
}

func (t _GoalInfoTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__action_msgs__msg__GoalInfo())
}

type CGoalInfo = C.action_msgs__msg__GoalInfo
type CGoalInfo__Sequence = C.action_msgs__msg__GoalInfo__Sequence

func GoalInfo__Sequence_to_Go(goSlice *[]GoalInfo, cSlice CGoalInfo__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GoalInfo, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		GoalInfoTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func GoalInfo__Sequence_to_C(cSlice *CGoalInfo__Sequence, goSlice []GoalInfo) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.action_msgs__msg__GoalInfo)(C.malloc(C.sizeof_struct_action_msgs__msg__GoalInfo * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		GoalInfoTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func GoalInfo__Array_to_Go(goSlice []GoalInfo, cSlice []CGoalInfo) {
	for i := 0; i < len(cSlice); i++ {
		GoalInfoTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GoalInfo__Array_to_C(cSlice []CGoalInfo, goSlice []GoalInfo) {
	for i := 0; i < len(goSlice); i++ {
		GoalInfoTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
