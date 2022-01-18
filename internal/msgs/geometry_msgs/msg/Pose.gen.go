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

#include <geometry_msgs/msg/pose.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/Pose", PoseTypeSupport)
}

// Do not create instances of this type directly. Always use NewPose
// function instead.
type Pose struct {
	Position Point `yaml:"position"`
	Orientation Quaternion `yaml:"orientation"`
}

// NewPose creates a new Pose with default values.
func NewPose() *Pose {
	self := Pose{}
	self.SetDefaults()
	return &self
}

func (t *Pose) Clone() *Pose {
	c := &Pose{}
	c.Position = *t.Position.Clone()
	c.Orientation = *t.Orientation.Clone()
	return c
}

func (t *Pose) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Pose) SetDefaults() {
	t.Position.SetDefaults()
	t.Orientation.SetDefaults()
}

// PosePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type PosePublisher struct {
	*rclgo.Publisher
}

// NewPosePublisher creates and returns a new publisher for the
// Pose
func NewPosePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*PosePublisher, error) {
	pub, err := node.NewPublisher(topic_name, PoseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &PosePublisher{pub}, nil
}

func (p *PosePublisher) Publish(msg *Pose) error {
	return p.Publisher.Publish(msg)
}

// PoseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type PoseSubscription struct {
	*rclgo.Subscription
}

// PoseSubscriptionCallback type is used to provide a subscription
// handler function for a PoseSubscription.
type PoseSubscriptionCallback func(msg *Pose, info *rclgo.RmwMessageInfo, err error)

// NewPoseSubscription creates and returns a new subscription for the
// Pose
func NewPoseSubscription(node *rclgo.Node, topic_name string, subscriptionCallback PoseSubscriptionCallback) (*PoseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Pose
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, PoseTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &PoseSubscription{sub}, nil
}

func (s *PoseSubscription) TakeMessage(out *Pose) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// ClonePoseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePoseSlice(dst, src []Pose) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PoseTypeSupport types.MessageTypeSupport = _PoseTypeSupport{}

type _PoseTypeSupport struct{}

func (t _PoseTypeSupport) New() types.Message {
	return NewPose()
}

func (t _PoseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__Pose
	return (unsafe.Pointer)(C.geometry_msgs__msg__Pose__create())
}

func (t _PoseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__Pose__destroy((*C.geometry_msgs__msg__Pose)(pointer_to_free))
}

func (t _PoseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Pose)
	mem := (*C.geometry_msgs__msg__Pose)(dst)
	PointTypeSupport.AsCStruct(unsafe.Pointer(&mem.position), &m.Position)
	QuaternionTypeSupport.AsCStruct(unsafe.Pointer(&mem.orientation), &m.Orientation)
}

func (t _PoseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Pose)
	mem := (*C.geometry_msgs__msg__Pose)(ros2_message_buffer)
	PointTypeSupport.AsGoStruct(&m.Position, unsafe.Pointer(&mem.position))
	QuaternionTypeSupport.AsGoStruct(&m.Orientation, unsafe.Pointer(&mem.orientation))
}

func (t _PoseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__Pose())
}

type CPose = C.geometry_msgs__msg__Pose
type CPose__Sequence = C.geometry_msgs__msg__Pose__Sequence

func Pose__Sequence_to_Go(goSlice *[]Pose, cSlice CPose__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Pose, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__Pose__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Pose * uintptr(i)),
		))
		PoseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Pose__Sequence_to_C(cSlice *CPose__Sequence, goSlice []Pose) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__Pose)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__Pose * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__Pose)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Pose * uintptr(i)),
		))
		PoseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Pose__Array_to_Go(goSlice []Pose, cSlice []CPose) {
	for i := 0; i < len(cSlice); i++ {
		PoseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Pose__Array_to_C(cSlice []CPose, goSlice []Pose) {
	for i := 0; i < len(goSlice); i++ {
		PoseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
