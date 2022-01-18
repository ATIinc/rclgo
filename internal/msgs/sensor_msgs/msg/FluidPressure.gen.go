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
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/fluid_pressure.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/FluidPressure", FluidPressureTypeSupport)
}

// Do not create instances of this type directly. Always use NewFluidPressure
// function instead.
type FluidPressure struct {
	Header std_msgs_msg.Header `yaml:"header"`// timestamp of the measurement
	FluidPressure float64 `yaml:"fluid_pressure"`// Absolute pressure reading in Pascals.
	Variance float64 `yaml:"variance"`// 0 is interpreted as variance unknown
}

// NewFluidPressure creates a new FluidPressure with default values.
func NewFluidPressure() *FluidPressure {
	self := FluidPressure{}
	self.SetDefaults()
	return &self
}

func (t *FluidPressure) Clone() *FluidPressure {
	c := &FluidPressure{}
	c.Header = *t.Header.Clone()
	c.FluidPressure = t.FluidPressure
	c.Variance = t.Variance
	return c
}

func (t *FluidPressure) CloneMsg() types.Message {
	return t.Clone()
}

func (t *FluidPressure) SetDefaults() {
	t.Header.SetDefaults()
	t.FluidPressure = 0
	t.Variance = 0
}

// FluidPressurePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type FluidPressurePublisher struct {
	*rclgo.Publisher
}

// NewFluidPressurePublisher creates and returns a new publisher for the
// FluidPressure
func NewFluidPressurePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*FluidPressurePublisher, error) {
	pub, err := node.NewPublisher(topic_name, FluidPressureTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &FluidPressurePublisher{pub}, nil
}

func (p *FluidPressurePublisher) Publish(msg *FluidPressure) error {
	return p.Publisher.Publish(msg)
}

// FluidPressureSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type FluidPressureSubscription struct {
	*rclgo.Subscription
}

// FluidPressureSubscriptionCallback type is used to provide a subscription
// handler function for a FluidPressureSubscription.
type FluidPressureSubscriptionCallback func(msg *FluidPressure, info *rclgo.RmwMessageInfo, err error)

// NewFluidPressureSubscription creates and returns a new subscription for the
// FluidPressure
func NewFluidPressureSubscription(node *rclgo.Node, topic_name string, subscriptionCallback FluidPressureSubscriptionCallback) (*FluidPressureSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg FluidPressure
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, FluidPressureTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &FluidPressureSubscription{sub}, nil
}

func (s *FluidPressureSubscription) TakeMessage(out *FluidPressure) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneFluidPressureSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneFluidPressureSlice(dst, src []FluidPressure) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var FluidPressureTypeSupport types.MessageTypeSupport = _FluidPressureTypeSupport{}

type _FluidPressureTypeSupport struct{}

func (t _FluidPressureTypeSupport) New() types.Message {
	return NewFluidPressure()
}

func (t _FluidPressureTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__FluidPressure
	return (unsafe.Pointer)(C.sensor_msgs__msg__FluidPressure__create())
}

func (t _FluidPressureTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__FluidPressure__destroy((*C.sensor_msgs__msg__FluidPressure)(pointer_to_free))
}

func (t _FluidPressureTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*FluidPressure)
	mem := (*C.sensor_msgs__msg__FluidPressure)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.fluid_pressure = C.double(m.FluidPressure)
	mem.variance = C.double(m.Variance)
}

func (t _FluidPressureTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*FluidPressure)
	mem := (*C.sensor_msgs__msg__FluidPressure)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.FluidPressure = float64(mem.fluid_pressure)
	m.Variance = float64(mem.variance)
}

func (t _FluidPressureTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__FluidPressure())
}

type CFluidPressure = C.sensor_msgs__msg__FluidPressure
type CFluidPressure__Sequence = C.sensor_msgs__msg__FluidPressure__Sequence

func FluidPressure__Sequence_to_Go(goSlice *[]FluidPressure, cSlice CFluidPressure__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]FluidPressure, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.sensor_msgs__msg__FluidPressure__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__FluidPressure * uintptr(i)),
		))
		FluidPressureTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func FluidPressure__Sequence_to_C(cSlice *CFluidPressure__Sequence, goSlice []FluidPressure) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__FluidPressure)(C.malloc((C.size_t)(C.sizeof_struct_sensor_msgs__msg__FluidPressure * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.sensor_msgs__msg__FluidPressure)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__FluidPressure * uintptr(i)),
		))
		FluidPressureTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func FluidPressure__Array_to_Go(goSlice []FluidPressure, cSlice []CFluidPressure) {
	for i := 0; i < len(cSlice); i++ {
		FluidPressureTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func FluidPressure__Array_to_C(cSlice []CFluidPressure, goSlice []FluidPressure) {
	for i := 0; i < len(goSlice); i++ {
		FluidPressureTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
