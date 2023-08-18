/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package test_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/msg/bounded_plain_sequences.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/BoundedPlainSequences", BoundedPlainSequencesTypeSupport)
	typemap.RegisterMessage("test_msgs/msg/BoundedPlainSequences", BoundedPlainSequencesTypeSupport)
}

type BoundedPlainSequences struct {
	BoolValues []bool `yaml:"bool_values"`// Bounded sequences of different POD types
	ByteValues []byte `yaml:"byte_values"`
	CharValues []byte `yaml:"char_values"`
	Float32Values []float32 `yaml:"float32_values"`
	Float64Values []float64 `yaml:"float64_values"`
	Int8Values []int8 `yaml:"int8_values"`
	Uint8Values []uint8 `yaml:"uint8_values"`
	Int16Values []int16 `yaml:"int16_values"`
	Uint16Values []uint16 `yaml:"uint16_values"`
	Int32Values []int32 `yaml:"int32_values"`
	Uint32Values []uint32 `yaml:"uint32_values"`
	Int64Values []int64 `yaml:"int64_values"`
	Uint64Values []uint64 `yaml:"uint64_values"`
	BasicTypesValues []BasicTypes `yaml:"basic_types_values"`
	ConstantsValues []Constants `yaml:"constants_values"`
	DefaultsValues []Defaults `yaml:"defaults_values"`
	BoolValuesDefault []bool `yaml:"bool_values_default"`
	ByteValuesDefault []byte `yaml:"byte_values_default"`
	CharValuesDefault []byte `yaml:"char_values_default"`
	Float32ValuesDefault []float32 `yaml:"float32_values_default"`
	Float64ValuesDefault []float64 `yaml:"float64_values_default"`
	Int8ValuesDefault []int8 `yaml:"int8_values_default"`
	Uint8ValuesDefault []uint8 `yaml:"uint8_values_default"`
	Int16ValuesDefault []int16 `yaml:"int16_values_default"`
	Uint16ValuesDefault []uint16 `yaml:"uint16_values_default"`
	Int32ValuesDefault []int32 `yaml:"int32_values_default"`
	Uint32ValuesDefault []uint32 `yaml:"uint32_values_default"`
	Int64ValuesDefault []int64 `yaml:"int64_values_default"`
	Uint64ValuesDefault []uint64 `yaml:"uint64_values_default"`
	AlignmentCheck int32 `yaml:"alignment_check"`// Regression test: check alignment of basic field after a sequence field is correct
}

// NewBoundedPlainSequences creates a new BoundedPlainSequences with default values.
func NewBoundedPlainSequences() *BoundedPlainSequences {
	self := BoundedPlainSequences{}
	self.SetDefaults()
	return &self
}

func (t *BoundedPlainSequences) Clone() *BoundedPlainSequences {
	c := &BoundedPlainSequences{}
	if t.BoolValues != nil {
		c.BoolValues = make([]bool, len(t.BoolValues))
		copy(c.BoolValues, t.BoolValues)
	}
	if t.ByteValues != nil {
		c.ByteValues = make([]byte, len(t.ByteValues))
		copy(c.ByteValues, t.ByteValues)
	}
	if t.CharValues != nil {
		c.CharValues = make([]byte, len(t.CharValues))
		copy(c.CharValues, t.CharValues)
	}
	if t.Float32Values != nil {
		c.Float32Values = make([]float32, len(t.Float32Values))
		copy(c.Float32Values, t.Float32Values)
	}
	if t.Float64Values != nil {
		c.Float64Values = make([]float64, len(t.Float64Values))
		copy(c.Float64Values, t.Float64Values)
	}
	if t.Int8Values != nil {
		c.Int8Values = make([]int8, len(t.Int8Values))
		copy(c.Int8Values, t.Int8Values)
	}
	if t.Uint8Values != nil {
		c.Uint8Values = make([]uint8, len(t.Uint8Values))
		copy(c.Uint8Values, t.Uint8Values)
	}
	if t.Int16Values != nil {
		c.Int16Values = make([]int16, len(t.Int16Values))
		copy(c.Int16Values, t.Int16Values)
	}
	if t.Uint16Values != nil {
		c.Uint16Values = make([]uint16, len(t.Uint16Values))
		copy(c.Uint16Values, t.Uint16Values)
	}
	if t.Int32Values != nil {
		c.Int32Values = make([]int32, len(t.Int32Values))
		copy(c.Int32Values, t.Int32Values)
	}
	if t.Uint32Values != nil {
		c.Uint32Values = make([]uint32, len(t.Uint32Values))
		copy(c.Uint32Values, t.Uint32Values)
	}
	if t.Int64Values != nil {
		c.Int64Values = make([]int64, len(t.Int64Values))
		copy(c.Int64Values, t.Int64Values)
	}
	if t.Uint64Values != nil {
		c.Uint64Values = make([]uint64, len(t.Uint64Values))
		copy(c.Uint64Values, t.Uint64Values)
	}
	if t.BasicTypesValues != nil {
		c.BasicTypesValues = make([]BasicTypes, len(t.BasicTypesValues))
		CloneBasicTypesSlice(c.BasicTypesValues, t.BasicTypesValues)
	}
	if t.ConstantsValues != nil {
		c.ConstantsValues = make([]Constants, len(t.ConstantsValues))
		CloneConstantsSlice(c.ConstantsValues, t.ConstantsValues)
	}
	if t.DefaultsValues != nil {
		c.DefaultsValues = make([]Defaults, len(t.DefaultsValues))
		CloneDefaultsSlice(c.DefaultsValues, t.DefaultsValues)
	}
	if t.BoolValuesDefault != nil {
		c.BoolValuesDefault = make([]bool, len(t.BoolValuesDefault))
		copy(c.BoolValuesDefault, t.BoolValuesDefault)
	}
	if t.ByteValuesDefault != nil {
		c.ByteValuesDefault = make([]byte, len(t.ByteValuesDefault))
		copy(c.ByteValuesDefault, t.ByteValuesDefault)
	}
	if t.CharValuesDefault != nil {
		c.CharValuesDefault = make([]byte, len(t.CharValuesDefault))
		copy(c.CharValuesDefault, t.CharValuesDefault)
	}
	if t.Float32ValuesDefault != nil {
		c.Float32ValuesDefault = make([]float32, len(t.Float32ValuesDefault))
		copy(c.Float32ValuesDefault, t.Float32ValuesDefault)
	}
	if t.Float64ValuesDefault != nil {
		c.Float64ValuesDefault = make([]float64, len(t.Float64ValuesDefault))
		copy(c.Float64ValuesDefault, t.Float64ValuesDefault)
	}
	if t.Int8ValuesDefault != nil {
		c.Int8ValuesDefault = make([]int8, len(t.Int8ValuesDefault))
		copy(c.Int8ValuesDefault, t.Int8ValuesDefault)
	}
	if t.Uint8ValuesDefault != nil {
		c.Uint8ValuesDefault = make([]uint8, len(t.Uint8ValuesDefault))
		copy(c.Uint8ValuesDefault, t.Uint8ValuesDefault)
	}
	if t.Int16ValuesDefault != nil {
		c.Int16ValuesDefault = make([]int16, len(t.Int16ValuesDefault))
		copy(c.Int16ValuesDefault, t.Int16ValuesDefault)
	}
	if t.Uint16ValuesDefault != nil {
		c.Uint16ValuesDefault = make([]uint16, len(t.Uint16ValuesDefault))
		copy(c.Uint16ValuesDefault, t.Uint16ValuesDefault)
	}
	if t.Int32ValuesDefault != nil {
		c.Int32ValuesDefault = make([]int32, len(t.Int32ValuesDefault))
		copy(c.Int32ValuesDefault, t.Int32ValuesDefault)
	}
	if t.Uint32ValuesDefault != nil {
		c.Uint32ValuesDefault = make([]uint32, len(t.Uint32ValuesDefault))
		copy(c.Uint32ValuesDefault, t.Uint32ValuesDefault)
	}
	if t.Int64ValuesDefault != nil {
		c.Int64ValuesDefault = make([]int64, len(t.Int64ValuesDefault))
		copy(c.Int64ValuesDefault, t.Int64ValuesDefault)
	}
	if t.Uint64ValuesDefault != nil {
		c.Uint64ValuesDefault = make([]uint64, len(t.Uint64ValuesDefault))
		copy(c.Uint64ValuesDefault, t.Uint64ValuesDefault)
	}
	c.AlignmentCheck = t.AlignmentCheck
	return c
}

func (t *BoundedPlainSequences) CloneMsg() types.Message {
	return t.Clone()
}

func (t *BoundedPlainSequences) SetDefaults() {
	t.BoolValues = nil
	t.ByteValues = nil
	t.CharValues = nil
	t.Float32Values = nil
	t.Float64Values = nil
	t.Int8Values = nil
	t.Uint8Values = nil
	t.Int16Values = nil
	t.Uint16Values = nil
	t.Int32Values = nil
	t.Uint32Values = nil
	t.Int64Values = nil
	t.Uint64Values = nil
	t.BasicTypesValues = nil
	t.ConstantsValues = nil
	t.DefaultsValues = nil
	t.BoolValuesDefault = []bool{false,true,false}
	t.ByteValuesDefault = []byte{0,1,255}
	t.CharValuesDefault = []byte{0,1,127}
	t.Float32ValuesDefault = []float32{1.125,0.0,-1.125}
	t.Float64ValuesDefault = []float64{3.1415,0.0,-3.1415}
	t.Int8ValuesDefault = []int8{0,127,-128}
	t.Uint8ValuesDefault = []uint8{0,1,255}
	t.Int16ValuesDefault = []int16{0,32767,-32768}
	t.Uint16ValuesDefault = []uint16{0,1,65535}
	t.Int32ValuesDefault = []int32{0,2147483647,-2147483648}
	t.Uint32ValuesDefault = []uint32{0,1,4294967295}
	t.Int64ValuesDefault = []int64{0,9223372036854775807,-9223372036854775808}
	t.Uint64ValuesDefault = []uint64{0,1,18446744073709551615}
	t.AlignmentCheck = 0
}

func (t *BoundedPlainSequences) GetTypeSupport() types.MessageTypeSupport {
	return BoundedPlainSequencesTypeSupport
}

// BoundedPlainSequencesPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type BoundedPlainSequencesPublisher struct {
	*rclgo.Publisher
}

// NewBoundedPlainSequencesPublisher creates and returns a new publisher for the
// BoundedPlainSequences
func NewBoundedPlainSequencesPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*BoundedPlainSequencesPublisher, error) {
	pub, err := node.NewPublisher(topic_name, BoundedPlainSequencesTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &BoundedPlainSequencesPublisher{pub}, nil
}

func (p *BoundedPlainSequencesPublisher) Publish(msg *BoundedPlainSequences) error {
	return p.Publisher.Publish(msg)
}

// BoundedPlainSequencesSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type BoundedPlainSequencesSubscription struct {
	*rclgo.Subscription
}

// BoundedPlainSequencesSubscriptionCallback type is used to provide a subscription
// handler function for a BoundedPlainSequencesSubscription.
type BoundedPlainSequencesSubscriptionCallback func(msg *BoundedPlainSequences, info *rclgo.MessageInfo, err error)

// NewBoundedPlainSequencesSubscription creates and returns a new subscription for the
// BoundedPlainSequences
func NewBoundedPlainSequencesSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback BoundedPlainSequencesSubscriptionCallback) (*BoundedPlainSequencesSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg BoundedPlainSequences
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, BoundedPlainSequencesTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &BoundedPlainSequencesSubscription{sub}, nil
}

func (s *BoundedPlainSequencesSubscription) TakeMessage(out *BoundedPlainSequences) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneBoundedPlainSequencesSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneBoundedPlainSequencesSlice(dst, src []BoundedPlainSequences) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var BoundedPlainSequencesTypeSupport types.MessageTypeSupport = _BoundedPlainSequencesTypeSupport{}

type _BoundedPlainSequencesTypeSupport struct{}

func (t _BoundedPlainSequencesTypeSupport) New() types.Message {
	return NewBoundedPlainSequences()
}

func (t _BoundedPlainSequencesTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__msg__BoundedPlainSequences
	return (unsafe.Pointer)(C.test_msgs__msg__BoundedPlainSequences__create())
}

func (t _BoundedPlainSequencesTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__msg__BoundedPlainSequences__destroy((*C.test_msgs__msg__BoundedPlainSequences)(pointer_to_free))
}

func (t _BoundedPlainSequencesTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*BoundedPlainSequences)
	mem := (*C.test_msgs__msg__BoundedPlainSequences)(dst)
	primitives.Bool__Sequence_to_C((*primitives.CBool__Sequence)(unsafe.Pointer(&mem.bool_values)), m.BoolValues)
	primitives.Byte__Sequence_to_C((*primitives.CByte__Sequence)(unsafe.Pointer(&mem.byte_values)), m.ByteValues)
	primitives.Char__Sequence_to_C((*primitives.CChar__Sequence)(unsafe.Pointer(&mem.char_values)), m.CharValues)
	primitives.Float32__Sequence_to_C((*primitives.CFloat32__Sequence)(unsafe.Pointer(&mem.float32_values)), m.Float32Values)
	primitives.Float64__Sequence_to_C((*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.float64_values)), m.Float64Values)
	primitives.Int8__Sequence_to_C((*primitives.CInt8__Sequence)(unsafe.Pointer(&mem.int8_values)), m.Int8Values)
	primitives.Uint8__Sequence_to_C((*primitives.CUint8__Sequence)(unsafe.Pointer(&mem.uint8_values)), m.Uint8Values)
	primitives.Int16__Sequence_to_C((*primitives.CInt16__Sequence)(unsafe.Pointer(&mem.int16_values)), m.Int16Values)
	primitives.Uint16__Sequence_to_C((*primitives.CUint16__Sequence)(unsafe.Pointer(&mem.uint16_values)), m.Uint16Values)
	primitives.Int32__Sequence_to_C((*primitives.CInt32__Sequence)(unsafe.Pointer(&mem.int32_values)), m.Int32Values)
	primitives.Uint32__Sequence_to_C((*primitives.CUint32__Sequence)(unsafe.Pointer(&mem.uint32_values)), m.Uint32Values)
	primitives.Int64__Sequence_to_C((*primitives.CInt64__Sequence)(unsafe.Pointer(&mem.int64_values)), m.Int64Values)
	primitives.Uint64__Sequence_to_C((*primitives.CUint64__Sequence)(unsafe.Pointer(&mem.uint64_values)), m.Uint64Values)
	BasicTypes__Sequence_to_C(&mem.basic_types_values, m.BasicTypesValues)
	Constants__Sequence_to_C(&mem.constants_values, m.ConstantsValues)
	Defaults__Sequence_to_C(&mem.defaults_values, m.DefaultsValues)
	primitives.Bool__Sequence_to_C((*primitives.CBool__Sequence)(unsafe.Pointer(&mem.bool_values_default)), m.BoolValuesDefault)
	primitives.Byte__Sequence_to_C((*primitives.CByte__Sequence)(unsafe.Pointer(&mem.byte_values_default)), m.ByteValuesDefault)
	primitives.Char__Sequence_to_C((*primitives.CChar__Sequence)(unsafe.Pointer(&mem.char_values_default)), m.CharValuesDefault)
	primitives.Float32__Sequence_to_C((*primitives.CFloat32__Sequence)(unsafe.Pointer(&mem.float32_values_default)), m.Float32ValuesDefault)
	primitives.Float64__Sequence_to_C((*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.float64_values_default)), m.Float64ValuesDefault)
	primitives.Int8__Sequence_to_C((*primitives.CInt8__Sequence)(unsafe.Pointer(&mem.int8_values_default)), m.Int8ValuesDefault)
	primitives.Uint8__Sequence_to_C((*primitives.CUint8__Sequence)(unsafe.Pointer(&mem.uint8_values_default)), m.Uint8ValuesDefault)
	primitives.Int16__Sequence_to_C((*primitives.CInt16__Sequence)(unsafe.Pointer(&mem.int16_values_default)), m.Int16ValuesDefault)
	primitives.Uint16__Sequence_to_C((*primitives.CUint16__Sequence)(unsafe.Pointer(&mem.uint16_values_default)), m.Uint16ValuesDefault)
	primitives.Int32__Sequence_to_C((*primitives.CInt32__Sequence)(unsafe.Pointer(&mem.int32_values_default)), m.Int32ValuesDefault)
	primitives.Uint32__Sequence_to_C((*primitives.CUint32__Sequence)(unsafe.Pointer(&mem.uint32_values_default)), m.Uint32ValuesDefault)
	primitives.Int64__Sequence_to_C((*primitives.CInt64__Sequence)(unsafe.Pointer(&mem.int64_values_default)), m.Int64ValuesDefault)
	primitives.Uint64__Sequence_to_C((*primitives.CUint64__Sequence)(unsafe.Pointer(&mem.uint64_values_default)), m.Uint64ValuesDefault)
	mem.alignment_check = C.int32_t(m.AlignmentCheck)
}

func (t _BoundedPlainSequencesTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*BoundedPlainSequences)
	mem := (*C.test_msgs__msg__BoundedPlainSequences)(ros2_message_buffer)
	primitives.Bool__Sequence_to_Go(&m.BoolValues, *(*primitives.CBool__Sequence)(unsafe.Pointer(&mem.bool_values)))
	primitives.Byte__Sequence_to_Go(&m.ByteValues, *(*primitives.CByte__Sequence)(unsafe.Pointer(&mem.byte_values)))
	primitives.Char__Sequence_to_Go(&m.CharValues, *(*primitives.CChar__Sequence)(unsafe.Pointer(&mem.char_values)))
	primitives.Float32__Sequence_to_Go(&m.Float32Values, *(*primitives.CFloat32__Sequence)(unsafe.Pointer(&mem.float32_values)))
	primitives.Float64__Sequence_to_Go(&m.Float64Values, *(*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.float64_values)))
	primitives.Int8__Sequence_to_Go(&m.Int8Values, *(*primitives.CInt8__Sequence)(unsafe.Pointer(&mem.int8_values)))
	primitives.Uint8__Sequence_to_Go(&m.Uint8Values, *(*primitives.CUint8__Sequence)(unsafe.Pointer(&mem.uint8_values)))
	primitives.Int16__Sequence_to_Go(&m.Int16Values, *(*primitives.CInt16__Sequence)(unsafe.Pointer(&mem.int16_values)))
	primitives.Uint16__Sequence_to_Go(&m.Uint16Values, *(*primitives.CUint16__Sequence)(unsafe.Pointer(&mem.uint16_values)))
	primitives.Int32__Sequence_to_Go(&m.Int32Values, *(*primitives.CInt32__Sequence)(unsafe.Pointer(&mem.int32_values)))
	primitives.Uint32__Sequence_to_Go(&m.Uint32Values, *(*primitives.CUint32__Sequence)(unsafe.Pointer(&mem.uint32_values)))
	primitives.Int64__Sequence_to_Go(&m.Int64Values, *(*primitives.CInt64__Sequence)(unsafe.Pointer(&mem.int64_values)))
	primitives.Uint64__Sequence_to_Go(&m.Uint64Values, *(*primitives.CUint64__Sequence)(unsafe.Pointer(&mem.uint64_values)))
	BasicTypes__Sequence_to_Go(&m.BasicTypesValues, mem.basic_types_values)
	Constants__Sequence_to_Go(&m.ConstantsValues, mem.constants_values)
	Defaults__Sequence_to_Go(&m.DefaultsValues, mem.defaults_values)
	primitives.Bool__Sequence_to_Go(&m.BoolValuesDefault, *(*primitives.CBool__Sequence)(unsafe.Pointer(&mem.bool_values_default)))
	primitives.Byte__Sequence_to_Go(&m.ByteValuesDefault, *(*primitives.CByte__Sequence)(unsafe.Pointer(&mem.byte_values_default)))
	primitives.Char__Sequence_to_Go(&m.CharValuesDefault, *(*primitives.CChar__Sequence)(unsafe.Pointer(&mem.char_values_default)))
	primitives.Float32__Sequence_to_Go(&m.Float32ValuesDefault, *(*primitives.CFloat32__Sequence)(unsafe.Pointer(&mem.float32_values_default)))
	primitives.Float64__Sequence_to_Go(&m.Float64ValuesDefault, *(*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.float64_values_default)))
	primitives.Int8__Sequence_to_Go(&m.Int8ValuesDefault, *(*primitives.CInt8__Sequence)(unsafe.Pointer(&mem.int8_values_default)))
	primitives.Uint8__Sequence_to_Go(&m.Uint8ValuesDefault, *(*primitives.CUint8__Sequence)(unsafe.Pointer(&mem.uint8_values_default)))
	primitives.Int16__Sequence_to_Go(&m.Int16ValuesDefault, *(*primitives.CInt16__Sequence)(unsafe.Pointer(&mem.int16_values_default)))
	primitives.Uint16__Sequence_to_Go(&m.Uint16ValuesDefault, *(*primitives.CUint16__Sequence)(unsafe.Pointer(&mem.uint16_values_default)))
	primitives.Int32__Sequence_to_Go(&m.Int32ValuesDefault, *(*primitives.CInt32__Sequence)(unsafe.Pointer(&mem.int32_values_default)))
	primitives.Uint32__Sequence_to_Go(&m.Uint32ValuesDefault, *(*primitives.CUint32__Sequence)(unsafe.Pointer(&mem.uint32_values_default)))
	primitives.Int64__Sequence_to_Go(&m.Int64ValuesDefault, *(*primitives.CInt64__Sequence)(unsafe.Pointer(&mem.int64_values_default)))
	primitives.Uint64__Sequence_to_Go(&m.Uint64ValuesDefault, *(*primitives.CUint64__Sequence)(unsafe.Pointer(&mem.uint64_values_default)))
	m.AlignmentCheck = int32(mem.alignment_check)
}

func (t _BoundedPlainSequencesTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__msg__BoundedPlainSequences())
}

type CBoundedPlainSequences = C.test_msgs__msg__BoundedPlainSequences
type CBoundedPlainSequences__Sequence = C.test_msgs__msg__BoundedPlainSequences__Sequence

func BoundedPlainSequences__Sequence_to_Go(goSlice *[]BoundedPlainSequences, cSlice CBoundedPlainSequences__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]BoundedPlainSequences, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		BoundedPlainSequencesTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func BoundedPlainSequences__Sequence_to_C(cSlice *CBoundedPlainSequences__Sequence, goSlice []BoundedPlainSequences) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.test_msgs__msg__BoundedPlainSequences)(C.malloc(C.sizeof_struct_test_msgs__msg__BoundedPlainSequences * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		BoundedPlainSequencesTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func BoundedPlainSequences__Array_to_Go(goSlice []BoundedPlainSequences, cSlice []CBoundedPlainSequences) {
	for i := 0; i < len(cSlice); i++ {
		BoundedPlainSequencesTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func BoundedPlainSequences__Array_to_C(cSlice []CBoundedPlainSequences, goSlice []BoundedPlainSequences) {
	for i := 0; i < len(goSlice); i++ {
		BoundedPlainSequencesTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
