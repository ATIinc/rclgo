/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
    http://www.apache.org/licenses/LICENSE-2.0
*/

package primitives

//#include <stdlib.h>
//#include "rosidl_runtime_c/u16string.h"
//#include "rosidl_runtime_c/u16string_functions.h"
import "C"

import (
	"unicode/utf16"
	"unsafe"
)

// U16StringAsCStruct writes m into dst, which must be an initialized or zeroed
// rosidl_runtime_c__U16String. Assigning reallocates the buffer __init
// allocated instead of leaking it.
func U16StringAsCStruct(dst unsafe.Pointer, m string) {
	mem := (*C.rosidl_runtime_c__U16String)(dst)
	encoded := utf16.Encode([]rune(m))
	// assignn rejects a NULL value, which an empty slice has no element to give.
	var empty C.uint16_t
	value := &empty
	if len(encoded) > 0 {
		value = (*C.uint16_t)(unsafe.Pointer(&encoded[0]))
	}
	if !C.rosidl_runtime_c__U16String__assignn(mem, value, C.size_t(len(encoded))) {
		panic("rclgo: out of memory converting a wstring to C")
	}
}

func U16StringAsGoStruct(msg *string, ros2_message_buffer unsafe.Pointer) {
	mem := (*C.rosidl_runtime_c__U16String)(ros2_message_buffer)

	*msg = string(utf16.Decode(unsafe.Slice((*uint16)(mem.data), mem.size)))
}

type (
	CU16String           = C.rosidl_runtime_c__U16String
	CU16String__Sequence = C.rosidl_runtime_c__U16String__Sequence
)

func U16String__Sequence_to_Go(goSlice *[]string, cSlice CU16String__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]string, int64(cSlice.size))
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range int(cSlice.size) {
		U16StringAsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}

// U16String__Sequence_to_C frees what cSlice holds, so cSlice must be initialized or zeroed.
func U16String__Sequence_to_C(cSlice *CU16String__Sequence, goSlice []string) {
	// Frees what __init allocated, such as an IDL default, including each
	// element's buffer.
	C.rosidl_runtime_c__U16String__Sequence__fini(cSlice)
	// Set every field rather than rely on what fini leaves behind. data below
	// is a plain allocation, never a rosidl::Buffer.
	cSlice.data = nil
	cSlice.size = 0
	cSlice.capacity = 0
	cSlice.is_rosidl_buffer = false
	cSlice.owns_rosidl_buffer = false
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.rosidl_runtime_c__U16String)(C.calloc(C.size_t(len(goSlice)), C.sizeof_struct_rosidl_runtime_c__U16String))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		U16StringAsCStruct(unsafe.Pointer(&dst[i]), goSlice[i])
	}
}

func U16String__Array_to_Go(goSlice []string, cSlice []CU16String) {
	for i := range cSlice {
		U16StringAsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}

func U16String__Array_to_C(cSlice []CU16String, goSlice []string) {
	for i := range goSlice {
		U16StringAsCStruct(unsafe.Pointer(&cSlice[i]), goSlice[i])
	}
}
