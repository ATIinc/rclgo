/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package action_msgs_srv
import (
	"unsafe"

	"github.com/ATIinc/rclgo/pkg/rclgo"
	"github.com/ATIinc/rclgo/pkg/rclgo/types"
	"github.com/ATIinc/rclgo/pkg/rclgo/typemap"
	action_msgs_msg "github.com/ATIinc/rclgo/internal/msgs/action_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <action_msgs/srv/cancel_goal.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("action_msgs/CancelGoal_Response", CancelGoal_ResponseTypeSupport)
	typemap.RegisterMessage("action_msgs/srv/CancelGoal_Response", CancelGoal_ResponseTypeSupport)
}
const (
	CancelGoal_Response_ERROR_NONE int8 = 0// Indicates the request was accepted without any errors.One or more goals have transitioned to the CANCELING state. Thegoals_canceling list is not empty.
	CancelGoal_Response_ERROR_REJECTED int8 = 1// Indicates the request was rejected.No goals have transitioned to the CANCELING state. The goals_canceling list isempty.
	CancelGoal_Response_ERROR_UNKNOWN_GOAL_ID int8 = 2// Indicates the requested goal ID does not exist.No goals have transitioned to the CANCELING state. The goals_canceling list isempty.
	CancelGoal_Response_ERROR_GOAL_TERMINATED int8 = 3// Indicates the goal is not cancelable because it is already in a terminal state.No goals have transitioned to the CANCELING state. The goals_canceling list isempty.
)

type CancelGoal_Response struct {
	ReturnCode int8 `yaml:"return_code"`// Return code, see above definitions.
	GoalsCanceling []action_msgs_msg.GoalInfo `yaml:"goals_canceling"`// Goals that accepted the cancel request.
}

// NewCancelGoal_Response creates a new CancelGoal_Response with default values.
func NewCancelGoal_Response() *CancelGoal_Response {
	self := CancelGoal_Response{}
	self.SetDefaults()
	return &self
}

func (t *CancelGoal_Response) Clone() *CancelGoal_Response {
	c := &CancelGoal_Response{}
	c.ReturnCode = t.ReturnCode
	if t.GoalsCanceling != nil {
		c.GoalsCanceling = make([]action_msgs_msg.GoalInfo, len(t.GoalsCanceling))
		action_msgs_msg.CloneGoalInfoSlice(c.GoalsCanceling, t.GoalsCanceling)
	}
	return c
}

func (t *CancelGoal_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *CancelGoal_Response) SetDefaults() {
	t.ReturnCode = 0
	t.GoalsCanceling = nil
}

func (t *CancelGoal_Response) GetTypeSupport() types.MessageTypeSupport {
	return CancelGoal_ResponseTypeSupport
}
func (t *CancelGoal_Response) CallForEach(f func(interface{})) {
	for i := range t.GoalsCanceling {
		f((*types.GoalID)(&t.GoalsCanceling[i].GoalId.Uuid))
	}
}

// CancelGoal_ResponsePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type CancelGoal_ResponsePublisher struct {
	*rclgo.Publisher
}

// NewCancelGoal_ResponsePublisher creates and returns a new publisher for the
// CancelGoal_Response
func NewCancelGoal_ResponsePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*CancelGoal_ResponsePublisher, error) {
	pub, err := node.NewPublisher(topic_name, CancelGoal_ResponseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &CancelGoal_ResponsePublisher{pub}, nil
}

func (p *CancelGoal_ResponsePublisher) Publish(msg *CancelGoal_Response) error {
	return p.Publisher.Publish(msg)
}

// CancelGoal_ResponseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type CancelGoal_ResponseSubscription struct {
	*rclgo.Subscription
}

// CancelGoal_ResponseSubscriptionCallback type is used to provide a subscription
// handler function for a CancelGoal_ResponseSubscription.
type CancelGoal_ResponseSubscriptionCallback func(msg *CancelGoal_Response, info *rclgo.MessageInfo, err error)

// NewCancelGoal_ResponseSubscription creates and returns a new subscription for the
// CancelGoal_Response
func NewCancelGoal_ResponseSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback CancelGoal_ResponseSubscriptionCallback) (*CancelGoal_ResponseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg CancelGoal_Response
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, CancelGoal_ResponseTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &CancelGoal_ResponseSubscription{sub}, nil
}

func (s *CancelGoal_ResponseSubscription) TakeMessage(out *CancelGoal_Response) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneCancelGoal_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneCancelGoal_ResponseSlice(dst, src []CancelGoal_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var CancelGoal_ResponseTypeSupport types.MessageTypeSupport = _CancelGoal_ResponseTypeSupport{}

type _CancelGoal_ResponseTypeSupport struct{}

func (t _CancelGoal_ResponseTypeSupport) New() types.Message {
	return NewCancelGoal_Response()
}

func (t _CancelGoal_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.action_msgs__srv__CancelGoal_Response
	return (unsafe.Pointer)(C.action_msgs__srv__CancelGoal_Response__create())
}

func (t _CancelGoal_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.action_msgs__srv__CancelGoal_Response__destroy((*C.action_msgs__srv__CancelGoal_Response)(pointer_to_free))
}

func (t _CancelGoal_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*CancelGoal_Response)
	mem := (*C.action_msgs__srv__CancelGoal_Response)(dst)
	mem.return_code = C.int8_t(m.ReturnCode)
	action_msgs_msg.GoalInfo__Sequence_to_C((*action_msgs_msg.CGoalInfo__Sequence)(unsafe.Pointer(&mem.goals_canceling)), m.GoalsCanceling)
}

func (t _CancelGoal_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*CancelGoal_Response)
	mem := (*C.action_msgs__srv__CancelGoal_Response)(ros2_message_buffer)
	m.ReturnCode = int8(mem.return_code)
	action_msgs_msg.GoalInfo__Sequence_to_Go(&m.GoalsCanceling, *(*action_msgs_msg.CGoalInfo__Sequence)(unsafe.Pointer(&mem.goals_canceling)))
}

func (t _CancelGoal_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__action_msgs__srv__CancelGoal_Response())
}

type CCancelGoal_Response = C.action_msgs__srv__CancelGoal_Response
type CCancelGoal_Response__Sequence = C.action_msgs__srv__CancelGoal_Response__Sequence

func CancelGoal_Response__Sequence_to_Go(goSlice *[]CancelGoal_Response, cSlice CCancelGoal_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]CancelGoal_Response, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		CancelGoal_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func CancelGoal_Response__Sequence_to_C(cSlice *CCancelGoal_Response__Sequence, goSlice []CancelGoal_Response) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.action_msgs__srv__CancelGoal_Response)(C.malloc(C.sizeof_struct_action_msgs__srv__CancelGoal_Response * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		CancelGoal_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func CancelGoal_Response__Array_to_Go(goSlice []CancelGoal_Response, cSlice []CCancelGoal_Response) {
	for i := 0; i < len(cSlice); i++ {
		CancelGoal_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func CancelGoal_Response__Array_to_C(cSlice []CCancelGoal_Response, goSlice []CancelGoal_Response) {
	for i := 0; i < len(goSlice); i++ {
		CancelGoal_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
