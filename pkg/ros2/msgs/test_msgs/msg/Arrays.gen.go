/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package test_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ltest_msgs__rosidl_typesupport_c -ltest_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/msg/arrays.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("test_msgs/Arrays", ArraysTypeSupport)
}

// Do not create instances of this type directly. Always use NewArrays
// function instead.
type Arrays struct {
	BoolValues [3]bool `yaml:"bool_values"`// Arrays of different types
	ByteValues [3]byte `yaml:"byte_values"`// Arrays of different types
	CharValues [3]byte `yaml:"char_values"`// Arrays of different types
	Float32Values [3]float32 `yaml:"float32_values"`// Arrays of different types
	Float64Values [3]float64 `yaml:"float64_values"`// Arrays of different types
	Int8Values [3]int8 `yaml:"int8_values"`// Arrays of different types
	Uint8Values [3]uint8 `yaml:"uint8_values"`// Arrays of different types
	Int16Values [3]int16 `yaml:"int16_values"`// Arrays of different types
	Uint16Values [3]uint16 `yaml:"uint16_values"`// Arrays of different types
	Int32Values [3]int32 `yaml:"int32_values"`// Arrays of different types
	Uint32Values [3]uint32 `yaml:"uint32_values"`// Arrays of different types
	Int64Values [3]int64 `yaml:"int64_values"`// Arrays of different types
	Uint64Values [3]uint64 `yaml:"uint64_values"`// Arrays of different types
	StringValues [3]string `yaml:"string_values"`// Arrays of different types
	BasicTypesValues [3]BasicTypes `yaml:"basic_types_values"`// Arrays of different types
	ConstantsValues [3]Constants `yaml:"constants_values"`// Arrays of different types
	DefaultsValues [3]Defaults `yaml:"defaults_values"`// Arrays of different types
	BoolValuesDefault [3]bool `yaml:"bool_values_default"`// Arrays of different types
	ByteValuesDefault [3]byte `yaml:"byte_values_default"`// Arrays of different types
	CharValuesDefault [3]byte `yaml:"char_values_default"`// Arrays of different types
	Float32ValuesDefault [3]float32 `yaml:"float32_values_default"`// Arrays of different types
	Float64ValuesDefault [3]float64 `yaml:"float64_values_default"`// Arrays of different types
	Int8ValuesDefault [3]int8 `yaml:"int8_values_default"`// Arrays of different types
	Uint8ValuesDefault [3]uint8 `yaml:"uint8_values_default"`// Arrays of different types
	Int16ValuesDefault [3]int16 `yaml:"int16_values_default"`// Arrays of different types
	Uint16ValuesDefault [3]uint16 `yaml:"uint16_values_default"`// Arrays of different types
	Int32ValuesDefault [3]int32 `yaml:"int32_values_default"`// Arrays of different types
	Uint32ValuesDefault [3]uint32 `yaml:"uint32_values_default"`// Arrays of different types
	Int64ValuesDefault [3]int64 `yaml:"int64_values_default"`// Arrays of different types
	Uint64ValuesDefault [3]uint64 `yaml:"uint64_values_default"`// Arrays of different types
	StringValuesDefault [3]string `yaml:"string_values_default"`// Arrays of different types
	AlignmentCheck int32 `yaml:"alignment_check"`// Arrays of different typesRegression test: check alignment of basic field after an array field is correct
}

// NewArrays creates a new Arrays with default values.
func NewArrays() *Arrays {
	self := Arrays{}
	self.SetDefaults()
	return &self
}

func (t *Arrays) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *Arrays) SetDefaults() {
	t.BasicTypesValues[0].SetDefaults()
	t.BasicTypesValues[1].SetDefaults()
	t.BasicTypesValues[2].SetDefaults()
	t.ConstantsValues[0].SetDefaults()
	t.ConstantsValues[1].SetDefaults()
	t.ConstantsValues[2].SetDefaults()
	t.DefaultsValues[0].SetDefaults()
	t.DefaultsValues[1].SetDefaults()
	t.DefaultsValues[2].SetDefaults()
	t.BoolValuesDefault = [3]bool{false,true,false}
	t.ByteValuesDefault = [3]byte{0,1,255}
	t.CharValuesDefault = [3]byte{0,1,127}
	t.Float32ValuesDefault = [3]float32{1.125,0.0,-1.125}
	t.Float64ValuesDefault = [3]float64{3.1415,0.0,-3.1415}
	t.Int8ValuesDefault = [3]int8{0,127,-128}
	t.Uint8ValuesDefault = [3]uint8{0,1,255}
	t.Int16ValuesDefault = [3]int16{0,32767,-32768}
	t.Uint16ValuesDefault = [3]uint16{0,1,65535}
	t.Int32ValuesDefault = [3]int32{0,2147483647,-2147483648}
	t.Uint32ValuesDefault = [3]uint32{0,1,4294967295}
	t.Int64ValuesDefault = [3]int64{0,9223372036854775807,-9223372036854775808}
	t.Uint64ValuesDefault = [3]uint64{0,1,18446744073709551615}
	t.StringValuesDefault = [3]string{"","max value","min value"}
	
}

// Modifying this variable is undefined behavior.
var ArraysTypeSupport types.MessageTypeSupport = _ArraysTypeSupport{}

type _ArraysTypeSupport struct{}

func (t _ArraysTypeSupport) New() types.Message {
	return NewArrays()
}

func (t _ArraysTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__msg__Arrays
	return (unsafe.Pointer)(C.test_msgs__msg__Arrays__create())
}

func (t _ArraysTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__msg__Arrays__destroy((*C.test_msgs__msg__Arrays)(pointer_to_free))
}

func (t _ArraysTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Arrays)
	mem := (*C.test_msgs__msg__Arrays)(dst)
	cSlice_bool_values := mem.bool_values[:]
	rosidl_runtime_c.Bool__Array_to_C(*(*[]rosidl_runtime_c.CBool)(unsafe.Pointer(&cSlice_bool_values)), m.BoolValues[:])
	cSlice_byte_values := mem.byte_values[:]
	rosidl_runtime_c.Byte__Array_to_C(*(*[]rosidl_runtime_c.CByte)(unsafe.Pointer(&cSlice_byte_values)), m.ByteValues[:])
	cSlice_char_values := mem.char_values[:]
	rosidl_runtime_c.Char__Array_to_C(*(*[]rosidl_runtime_c.CChar)(unsafe.Pointer(&cSlice_char_values)), m.CharValues[:])
	cSlice_float32_values := mem.float32_values[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_float32_values)), m.Float32Values[:])
	cSlice_float64_values := mem.float64_values[:]
	rosidl_runtime_c.Float64__Array_to_C(*(*[]rosidl_runtime_c.CFloat64)(unsafe.Pointer(&cSlice_float64_values)), m.Float64Values[:])
	cSlice_int8_values := mem.int8_values[:]
	rosidl_runtime_c.Int8__Array_to_C(*(*[]rosidl_runtime_c.CInt8)(unsafe.Pointer(&cSlice_int8_values)), m.Int8Values[:])
	cSlice_uint8_values := mem.uint8_values[:]
	rosidl_runtime_c.Uint8__Array_to_C(*(*[]rosidl_runtime_c.CUint8)(unsafe.Pointer(&cSlice_uint8_values)), m.Uint8Values[:])
	cSlice_int16_values := mem.int16_values[:]
	rosidl_runtime_c.Int16__Array_to_C(*(*[]rosidl_runtime_c.CInt16)(unsafe.Pointer(&cSlice_int16_values)), m.Int16Values[:])
	cSlice_uint16_values := mem.uint16_values[:]
	rosidl_runtime_c.Uint16__Array_to_C(*(*[]rosidl_runtime_c.CUint16)(unsafe.Pointer(&cSlice_uint16_values)), m.Uint16Values[:])
	cSlice_int32_values := mem.int32_values[:]
	rosidl_runtime_c.Int32__Array_to_C(*(*[]rosidl_runtime_c.CInt32)(unsafe.Pointer(&cSlice_int32_values)), m.Int32Values[:])
	cSlice_uint32_values := mem.uint32_values[:]
	rosidl_runtime_c.Uint32__Array_to_C(*(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_uint32_values)), m.Uint32Values[:])
	cSlice_int64_values := mem.int64_values[:]
	rosidl_runtime_c.Int64__Array_to_C(*(*[]rosidl_runtime_c.CInt64)(unsafe.Pointer(&cSlice_int64_values)), m.Int64Values[:])
	cSlice_uint64_values := mem.uint64_values[:]
	rosidl_runtime_c.Uint64__Array_to_C(*(*[]rosidl_runtime_c.CUint64)(unsafe.Pointer(&cSlice_uint64_values)), m.Uint64Values[:])
	cSlice_string_values := mem.string_values[:]
	rosidl_runtime_c.String__Array_to_C(*(*[]rosidl_runtime_c.CString)(unsafe.Pointer(&cSlice_string_values)), m.StringValues[:])
	BasicTypes__Array_to_C(mem.basic_types_values[:], m.BasicTypesValues[:])
	Constants__Array_to_C(mem.constants_values[:], m.ConstantsValues[:])
	Defaults__Array_to_C(mem.defaults_values[:], m.DefaultsValues[:])
	cSlice_bool_values_default := mem.bool_values_default[:]
	rosidl_runtime_c.Bool__Array_to_C(*(*[]rosidl_runtime_c.CBool)(unsafe.Pointer(&cSlice_bool_values_default)), m.BoolValuesDefault[:])
	cSlice_byte_values_default := mem.byte_values_default[:]
	rosidl_runtime_c.Byte__Array_to_C(*(*[]rosidl_runtime_c.CByte)(unsafe.Pointer(&cSlice_byte_values_default)), m.ByteValuesDefault[:])
	cSlice_char_values_default := mem.char_values_default[:]
	rosidl_runtime_c.Char__Array_to_C(*(*[]rosidl_runtime_c.CChar)(unsafe.Pointer(&cSlice_char_values_default)), m.CharValuesDefault[:])
	cSlice_float32_values_default := mem.float32_values_default[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_float32_values_default)), m.Float32ValuesDefault[:])
	cSlice_float64_values_default := mem.float64_values_default[:]
	rosidl_runtime_c.Float64__Array_to_C(*(*[]rosidl_runtime_c.CFloat64)(unsafe.Pointer(&cSlice_float64_values_default)), m.Float64ValuesDefault[:])
	cSlice_int8_values_default := mem.int8_values_default[:]
	rosidl_runtime_c.Int8__Array_to_C(*(*[]rosidl_runtime_c.CInt8)(unsafe.Pointer(&cSlice_int8_values_default)), m.Int8ValuesDefault[:])
	cSlice_uint8_values_default := mem.uint8_values_default[:]
	rosidl_runtime_c.Uint8__Array_to_C(*(*[]rosidl_runtime_c.CUint8)(unsafe.Pointer(&cSlice_uint8_values_default)), m.Uint8ValuesDefault[:])
	cSlice_int16_values_default := mem.int16_values_default[:]
	rosidl_runtime_c.Int16__Array_to_C(*(*[]rosidl_runtime_c.CInt16)(unsafe.Pointer(&cSlice_int16_values_default)), m.Int16ValuesDefault[:])
	cSlice_uint16_values_default := mem.uint16_values_default[:]
	rosidl_runtime_c.Uint16__Array_to_C(*(*[]rosidl_runtime_c.CUint16)(unsafe.Pointer(&cSlice_uint16_values_default)), m.Uint16ValuesDefault[:])
	cSlice_int32_values_default := mem.int32_values_default[:]
	rosidl_runtime_c.Int32__Array_to_C(*(*[]rosidl_runtime_c.CInt32)(unsafe.Pointer(&cSlice_int32_values_default)), m.Int32ValuesDefault[:])
	cSlice_uint32_values_default := mem.uint32_values_default[:]
	rosidl_runtime_c.Uint32__Array_to_C(*(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_uint32_values_default)), m.Uint32ValuesDefault[:])
	cSlice_int64_values_default := mem.int64_values_default[:]
	rosidl_runtime_c.Int64__Array_to_C(*(*[]rosidl_runtime_c.CInt64)(unsafe.Pointer(&cSlice_int64_values_default)), m.Int64ValuesDefault[:])
	cSlice_uint64_values_default := mem.uint64_values_default[:]
	rosidl_runtime_c.Uint64__Array_to_C(*(*[]rosidl_runtime_c.CUint64)(unsafe.Pointer(&cSlice_uint64_values_default)), m.Uint64ValuesDefault[:])
	cSlice_string_values_default := mem.string_values_default[:]
	rosidl_runtime_c.String__Array_to_C(*(*[]rosidl_runtime_c.CString)(unsafe.Pointer(&cSlice_string_values_default)), m.StringValuesDefault[:])
	mem.alignment_check = C.int32_t(m.AlignmentCheck)
}

func (t _ArraysTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Arrays)
	mem := (*C.test_msgs__msg__Arrays)(ros2_message_buffer)
	cSlice_bool_values := mem.bool_values[:]
	rosidl_runtime_c.Bool__Array_to_Go(m.BoolValues[:], *(*[]rosidl_runtime_c.CBool)(unsafe.Pointer(&cSlice_bool_values)))
	cSlice_byte_values := mem.byte_values[:]
	rosidl_runtime_c.Byte__Array_to_Go(m.ByteValues[:], *(*[]rosidl_runtime_c.CByte)(unsafe.Pointer(&cSlice_byte_values)))
	cSlice_char_values := mem.char_values[:]
	rosidl_runtime_c.Char__Array_to_Go(m.CharValues[:], *(*[]rosidl_runtime_c.CChar)(unsafe.Pointer(&cSlice_char_values)))
	cSlice_float32_values := mem.float32_values[:]
	rosidl_runtime_c.Float32__Array_to_Go(m.Float32Values[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_float32_values)))
	cSlice_float64_values := mem.float64_values[:]
	rosidl_runtime_c.Float64__Array_to_Go(m.Float64Values[:], *(*[]rosidl_runtime_c.CFloat64)(unsafe.Pointer(&cSlice_float64_values)))
	cSlice_int8_values := mem.int8_values[:]
	rosidl_runtime_c.Int8__Array_to_Go(m.Int8Values[:], *(*[]rosidl_runtime_c.CInt8)(unsafe.Pointer(&cSlice_int8_values)))
	cSlice_uint8_values := mem.uint8_values[:]
	rosidl_runtime_c.Uint8__Array_to_Go(m.Uint8Values[:], *(*[]rosidl_runtime_c.CUint8)(unsafe.Pointer(&cSlice_uint8_values)))
	cSlice_int16_values := mem.int16_values[:]
	rosidl_runtime_c.Int16__Array_to_Go(m.Int16Values[:], *(*[]rosidl_runtime_c.CInt16)(unsafe.Pointer(&cSlice_int16_values)))
	cSlice_uint16_values := mem.uint16_values[:]
	rosidl_runtime_c.Uint16__Array_to_Go(m.Uint16Values[:], *(*[]rosidl_runtime_c.CUint16)(unsafe.Pointer(&cSlice_uint16_values)))
	cSlice_int32_values := mem.int32_values[:]
	rosidl_runtime_c.Int32__Array_to_Go(m.Int32Values[:], *(*[]rosidl_runtime_c.CInt32)(unsafe.Pointer(&cSlice_int32_values)))
	cSlice_uint32_values := mem.uint32_values[:]
	rosidl_runtime_c.Uint32__Array_to_Go(m.Uint32Values[:], *(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_uint32_values)))
	cSlice_int64_values := mem.int64_values[:]
	rosidl_runtime_c.Int64__Array_to_Go(m.Int64Values[:], *(*[]rosidl_runtime_c.CInt64)(unsafe.Pointer(&cSlice_int64_values)))
	cSlice_uint64_values := mem.uint64_values[:]
	rosidl_runtime_c.Uint64__Array_to_Go(m.Uint64Values[:], *(*[]rosidl_runtime_c.CUint64)(unsafe.Pointer(&cSlice_uint64_values)))
	cSlice_string_values := mem.string_values[:]
	rosidl_runtime_c.String__Array_to_Go(m.StringValues[:], *(*[]rosidl_runtime_c.CString)(unsafe.Pointer(&cSlice_string_values)))
	BasicTypes__Array_to_Go(m.BasicTypesValues[:], mem.basic_types_values[:])
	Constants__Array_to_Go(m.ConstantsValues[:], mem.constants_values[:])
	Defaults__Array_to_Go(m.DefaultsValues[:], mem.defaults_values[:])
	cSlice_bool_values_default := mem.bool_values_default[:]
	rosidl_runtime_c.Bool__Array_to_Go(m.BoolValuesDefault[:], *(*[]rosidl_runtime_c.CBool)(unsafe.Pointer(&cSlice_bool_values_default)))
	cSlice_byte_values_default := mem.byte_values_default[:]
	rosidl_runtime_c.Byte__Array_to_Go(m.ByteValuesDefault[:], *(*[]rosidl_runtime_c.CByte)(unsafe.Pointer(&cSlice_byte_values_default)))
	cSlice_char_values_default := mem.char_values_default[:]
	rosidl_runtime_c.Char__Array_to_Go(m.CharValuesDefault[:], *(*[]rosidl_runtime_c.CChar)(unsafe.Pointer(&cSlice_char_values_default)))
	cSlice_float32_values_default := mem.float32_values_default[:]
	rosidl_runtime_c.Float32__Array_to_Go(m.Float32ValuesDefault[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_float32_values_default)))
	cSlice_float64_values_default := mem.float64_values_default[:]
	rosidl_runtime_c.Float64__Array_to_Go(m.Float64ValuesDefault[:], *(*[]rosidl_runtime_c.CFloat64)(unsafe.Pointer(&cSlice_float64_values_default)))
	cSlice_int8_values_default := mem.int8_values_default[:]
	rosidl_runtime_c.Int8__Array_to_Go(m.Int8ValuesDefault[:], *(*[]rosidl_runtime_c.CInt8)(unsafe.Pointer(&cSlice_int8_values_default)))
	cSlice_uint8_values_default := mem.uint8_values_default[:]
	rosidl_runtime_c.Uint8__Array_to_Go(m.Uint8ValuesDefault[:], *(*[]rosidl_runtime_c.CUint8)(unsafe.Pointer(&cSlice_uint8_values_default)))
	cSlice_int16_values_default := mem.int16_values_default[:]
	rosidl_runtime_c.Int16__Array_to_Go(m.Int16ValuesDefault[:], *(*[]rosidl_runtime_c.CInt16)(unsafe.Pointer(&cSlice_int16_values_default)))
	cSlice_uint16_values_default := mem.uint16_values_default[:]
	rosidl_runtime_c.Uint16__Array_to_Go(m.Uint16ValuesDefault[:], *(*[]rosidl_runtime_c.CUint16)(unsafe.Pointer(&cSlice_uint16_values_default)))
	cSlice_int32_values_default := mem.int32_values_default[:]
	rosidl_runtime_c.Int32__Array_to_Go(m.Int32ValuesDefault[:], *(*[]rosidl_runtime_c.CInt32)(unsafe.Pointer(&cSlice_int32_values_default)))
	cSlice_uint32_values_default := mem.uint32_values_default[:]
	rosidl_runtime_c.Uint32__Array_to_Go(m.Uint32ValuesDefault[:], *(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_uint32_values_default)))
	cSlice_int64_values_default := mem.int64_values_default[:]
	rosidl_runtime_c.Int64__Array_to_Go(m.Int64ValuesDefault[:], *(*[]rosidl_runtime_c.CInt64)(unsafe.Pointer(&cSlice_int64_values_default)))
	cSlice_uint64_values_default := mem.uint64_values_default[:]
	rosidl_runtime_c.Uint64__Array_to_Go(m.Uint64ValuesDefault[:], *(*[]rosidl_runtime_c.CUint64)(unsafe.Pointer(&cSlice_uint64_values_default)))
	cSlice_string_values_default := mem.string_values_default[:]
	rosidl_runtime_c.String__Array_to_Go(m.StringValuesDefault[:], *(*[]rosidl_runtime_c.CString)(unsafe.Pointer(&cSlice_string_values_default)))
	m.AlignmentCheck = int32(mem.alignment_check)
}

func (t _ArraysTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__msg__Arrays())
}

type CArrays = C.test_msgs__msg__Arrays
type CArrays__Sequence = C.test_msgs__msg__Arrays__Sequence

func Arrays__Sequence_to_Go(goSlice *[]Arrays, cSlice CArrays__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Arrays, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.test_msgs__msg__Arrays__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__msg__Arrays * uintptr(i)),
		))
		ArraysTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Arrays__Sequence_to_C(cSlice *CArrays__Sequence, goSlice []Arrays) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.test_msgs__msg__Arrays)(C.malloc((C.size_t)(C.sizeof_struct_test_msgs__msg__Arrays * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.test_msgs__msg__Arrays)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__msg__Arrays * uintptr(i)),
		))
		ArraysTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Arrays__Array_to_Go(goSlice []Arrays, cSlice []CArrays) {
	for i := 0; i < len(cSlice); i++ {
		ArraysTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Arrays__Array_to_C(cSlice []CArrays, goSlice []Arrays) {
	for i := 0; i < len(goSlice); i++ {
		ArraysTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
