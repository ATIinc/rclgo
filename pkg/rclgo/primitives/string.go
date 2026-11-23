/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
    http://www.apache.org/licenses/LICENSE-2.0
*/

package primitives

// #include <stdlib.h>
// #include "rosidl_runtime_c/string.h"
// #include "rosidl_runtime_c/string_functions.h"
import "C"

import (
	"strings"
	"unsafe"
)

// StringAsCStruct writes m into dst, which must be an initialized or zeroed
// rosidl_runtime_c__String. Assigning reallocates the buffer __init allocated
// instead of leaking it.
func StringAsCStruct(dst unsafe.Pointer, m string) {
	mem := (*C.rosidl_runtime_c__String)(dst)
	// assignn rejects a NULL value, and the data pointer of "" may be nil.
	var empty C.char
	value := &empty
	if len(m) > 0 {
		value = (*C.char)(unsafe.Pointer(unsafe.StringData(m)))
	}
	if !C.rosidl_runtime_c__String__assignn(mem, value, C.size_t(len(m))) {
		panic("rclgo: out of memory converting a string to C")
	}
}

func StringAsGoStruct(m *string, ros2_message_buffer unsafe.Pointer) {
	mem := (*C.rosidl_runtime_c__String)(ros2_message_buffer)
	sb := strings.Builder{}
	sb.Grow(int(mem.size))
	sb.Write(unsafe.Slice((*byte)(unsafe.Pointer(mem.data)), mem.size))
	*m = sb.String()
}

type (
	CString           = C.rosidl_runtime_c__String
	CString__Sequence = C.rosidl_runtime_c__String__Sequence
)

func String__Sequence_to_Go(goSlice *[]string, cSlice CString__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]string, int64(cSlice.size))
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range int(cSlice.size) {
		StringAsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}

// String__Sequence_to_C frees what cSlice holds, so cSlice must be initialized or zeroed.
func String__Sequence_to_C(cSlice *CString__Sequence, goSlice []string) {
	// Frees what __init allocated, such as an IDL default, including each
	// element's buffer.
	C.rosidl_runtime_c__String__Sequence__fini(cSlice)
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
	cSlice.data = (*C.rosidl_runtime_c__String)(C.calloc(C.size_t(len(goSlice)), C.sizeof_struct_rosidl_runtime_c__String))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		StringAsCStruct(unsafe.Pointer(&dst[i]), goSlice[i])
	}
}

func String__Array_to_Go(goSlice []string, cSlice []CString) {
	for i := range cSlice {
		StringAsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}

func String__Array_to_C(cSlice []CString, goSlice []string) {
	for i := range goSlice {
		StringAsCStruct(unsafe.Pointer(&cSlice[i]), goSlice[i])
	}
}
