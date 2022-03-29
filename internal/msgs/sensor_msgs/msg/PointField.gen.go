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
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/point_field.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/PointField", PointFieldTypeSupport)
	typemap.RegisterMessage("sensor_msgs/msg/PointField", PointFieldTypeSupport)
}
const (
	PointField_INT8 uint8 = 1// This message holds the description of one point entry in thePointCloud2 message format.
	PointField_UINT8 uint8 = 2
	PointField_INT16 uint8 = 3
	PointField_UINT16 uint8 = 4
	PointField_INT32 uint8 = 5
	PointField_UINT32 uint8 = 6
	PointField_FLOAT32 uint8 = 7
	PointField_FLOAT64 uint8 = 8
)

// Do not create instances of this type directly. Always use NewPointField
// function instead.
type PointField struct {
	Name string `yaml:"name"`// Name of field. Common PointField names are x, y, z, intensity, rgb, rgba
	Offset uint32 `yaml:"offset"`// Offset from start of point struct
	Datatype uint8 `yaml:"datatype"`// Datatype enumeration, see above
	Count uint32 `yaml:"count"`// How many elements in the field
}

// NewPointField creates a new PointField with default values.
func NewPointField() *PointField {
	self := PointField{}
	self.SetDefaults()
	return &self
}

func (t *PointField) Clone() *PointField {
	c := &PointField{}
	c.Name = t.Name
	c.Offset = t.Offset
	c.Datatype = t.Datatype
	c.Count = t.Count
	return c
}

func (t *PointField) CloneMsg() types.Message {
	return t.Clone()
}

func (t *PointField) SetDefaults() {
	t.Name = ""
	t.Offset = 0
	t.Datatype = 0
	t.Count = 0
}

func (t *PointField) GetTypeSupport() types.MessageTypeSupport {
	return PointFieldTypeSupport
}

// PointFieldPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type PointFieldPublisher struct {
	*rclgo.Publisher
}

// NewPointFieldPublisher creates and returns a new publisher for the
// PointField
func NewPointFieldPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*PointFieldPublisher, error) {
	pub, err := node.NewPublisher(topic_name, PointFieldTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &PointFieldPublisher{pub}, nil
}

func (p *PointFieldPublisher) Publish(msg *PointField) error {
	return p.Publisher.Publish(msg)
}

// PointFieldSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type PointFieldSubscription struct {
	*rclgo.Subscription
}

// PointFieldSubscriptionCallback type is used to provide a subscription
// handler function for a PointFieldSubscription.
type PointFieldSubscriptionCallback func(msg *PointField, info *rclgo.RmwMessageInfo, err error)

// NewPointFieldSubscription creates and returns a new subscription for the
// PointField
func NewPointFieldSubscription(node *rclgo.Node, topic_name string, subscriptionCallback PointFieldSubscriptionCallback) (*PointFieldSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg PointField
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, PointFieldTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &PointFieldSubscription{sub}, nil
}

func (s *PointFieldSubscription) TakeMessage(out *PointField) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// ClonePointFieldSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePointFieldSlice(dst, src []PointField) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PointFieldTypeSupport types.MessageTypeSupport = _PointFieldTypeSupport{}

type _PointFieldTypeSupport struct{}

func (t _PointFieldTypeSupport) New() types.Message {
	return NewPointField()
}

func (t _PointFieldTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__PointField
	return (unsafe.Pointer)(C.sensor_msgs__msg__PointField__create())
}

func (t _PointFieldTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__PointField__destroy((*C.sensor_msgs__msg__PointField)(pointer_to_free))
}

func (t _PointFieldTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*PointField)
	mem := (*C.sensor_msgs__msg__PointField)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.name), m.Name)
	mem.offset = C.uint32_t(m.Offset)
	mem.datatype = C.uint8_t(m.Datatype)
	mem.count = C.uint32_t(m.Count)
}

func (t _PointFieldTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*PointField)
	mem := (*C.sensor_msgs__msg__PointField)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.Name, unsafe.Pointer(&mem.name))
	m.Offset = uint32(mem.offset)
	m.Datatype = uint8(mem.datatype)
	m.Count = uint32(mem.count)
}

func (t _PointFieldTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__PointField())
}

type CPointField = C.sensor_msgs__msg__PointField
type CPointField__Sequence = C.sensor_msgs__msg__PointField__Sequence

func PointField__Sequence_to_Go(goSlice *[]PointField, cSlice CPointField__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PointField, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		PointFieldTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func PointField__Sequence_to_C(cSlice *CPointField__Sequence, goSlice []PointField) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__PointField)(C.malloc(C.sizeof_struct_sensor_msgs__msg__PointField * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		PointFieldTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func PointField__Array_to_Go(goSlice []PointField, cSlice []CPointField) {
	for i := 0; i < len(cSlice); i++ {
		PointFieldTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func PointField__Array_to_C(cSlice []CPointField, goSlice []PointField) {
	for i := 0; i < len(goSlice); i++ {
		PointFieldTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
