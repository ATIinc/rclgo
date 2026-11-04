/*
This file is part of rclgo

# Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");

	http://www.apache.org/licenses/LICENSE-2.0
*/
package rclgo

/*
#cgo LDFLAGS: -ldl

#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include <dlfcn.h>

#include <rosidl_runtime_c/message_type_support_struct.h>

char* formatString(const char* format, ...) {
	va_list args1, args2;
	va_start(args1, format);
	va_copy(args2, args1);
	size_t len = 1 + vsnprintf(NULL, 0, format, args1);
	va_end(args1);
	char* buf = malloc(len);
	if (buf != NULL) {
		vsprintf(buf, format, args2);
	}
	va_end(args2);
	return buf;
}

typedef rosidl_message_type_support_t * (*GetTypeSupportFunc)();

// dlerrorCopy copies the dlerror message before a later dl call can overwrite it.
static char* dlerrorCopy() {
	const char* msg = dlerror();
	return strdup(msg != NULL ? msg : "unknown dynamic loading error");
}

// loadTypeSupport returns NULL on success, otherwise an error message the caller frees.
char* loadTypeSupport(
	const char* pkgName,
	const char* subfolder,
	const char* ifaceName,
	void** lib,
	void** typeSupport
) {
	char* libName = formatString("lib%s__rosidl_typesupport_c.so", pkgName);
	if (libName == NULL) {
		return strdup("allocation failed");
	}
	*lib = dlopen(libName, RTLD_LAZY);
	free(libName);
	if (*lib == NULL) {
		return dlerrorCopy();
	}
	char* tsName = formatString(
		"rosidl_typesupport_c__get_message_type_support_handle__%s__%s__%s",
		pkgName, subfolder, ifaceName
	);
	if (tsName == NULL) {
		dlclose(*lib);
		return strdup("allocation failed");
	}
	void* tsSym = dlsym(*lib, tsName);
	free(tsName);
	if (tsSym == NULL) {
		char* err = dlerrorCopy();
		dlclose(*lib);
		return err;
	}
	*typeSupport = ((GetTypeSupportFunc)tsSym)();
	return NULL;
}
*/
import "C"

import (
	"fmt"
	"runtime"
	"slices"
	"strings"
	"unsafe"

	"github.com/ATIinc/rclgo/pkg/rclgo/types"
)

type dynamicMessageTypeSupport struct {
	lib         unsafe.Pointer // void*
	typeSupport unsafe.Pointer // rosidl_message_type_support_t*
}

// LoadDynamicMessageTypeSupport loads a message type support implementation
// dynamically.
//
// MessageTypeSupport instances returned by LoadDynamicMessageTypeSupport
// support use cases related to handling only serialized messages. Methods New,
// PrepareMemory, ReleaseMemory, AsCStruct and AsGoStruct will panic.
//
// Backwards compatibility is not guaranteed for this API. Use it only if
// necessary.
func LoadDynamicMessageTypeSupport(pkgName, msgName string) (types.MessageTypeSupport, error) {
	return loadDynamicTypeSupport(pkgName, "msg", msgName)
}

// LoadDynamicMessageTypeSupportByName is LoadDynamicMessageTypeSupport for a
// message type named as graph queries report it: <package>/<msg|srv|action>/<Name>.
// Besides plain messages this covers the messages services and actions publish,
// such as std_srvs/srv/Trigger_Event and
// example_interfaces/action/Fibonacci_FeedbackMessage.
func LoadDynamicMessageTypeSupportByName(typeName string) (types.MessageTypeSupport, error) {
	parts := strings.Split(typeName, "/")
	if len(parts) != 3 || parts[0] == "" || parts[2] == "" ||
		!slices.Contains([]string{"msg", "srv", "action"}, parts[1]) {
		return nil, fmt.Errorf("invalid message type name %q: want <package>/<msg|srv|action>/<Name>", typeName)
	}
	return loadDynamicTypeSupport(parts[0], parts[1], parts[2])
}

func loadDynamicTypeSupport(pkgName, subfolder, name string) (types.MessageTypeSupport, error) {
	cPkgName := C.CString(pkgName)
	defer C.free(unsafe.Pointer(cPkgName))
	cSubfolder := C.CString(subfolder)
	defer C.free(unsafe.Pointer(cSubfolder))
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	ts := new(dynamicMessageTypeSupport)
	if cErr := C.loadTypeSupport(cPkgName, cSubfolder, cName, &ts.lib, &ts.typeSupport); cErr != nil {
		defer C.free(unsafe.Pointer(cErr))
		return nil, fmt.Errorf("failed to load type support for %s/%s/%s: %s", pkgName, subfolder, name, C.GoString(cErr))
	}
	runtime.SetFinalizer(ts, func(g *dynamicMessageTypeSupport) {
		C.dlclose(g.lib)
	})
	return ts, nil
}

func (g *dynamicMessageTypeSupport) New() types.Message {
	panic("not supported")
}

func (g *dynamicMessageTypeSupport) PrepareMemory() unsafe.Pointer {
	panic("not supported")
}

func (g *dynamicMessageTypeSupport) ReleaseMemory(unsafe.Pointer) {
	panic("not supported")
}

func (g *dynamicMessageTypeSupport) AsCStruct(unsafe.Pointer, types.Message) {
	panic("not supported")
}

func (g *dynamicMessageTypeSupport) AsGoStruct(types.Message, unsafe.Pointer) {
	panic("not supported")
}

func (g *dynamicMessageTypeSupport) TypeSupport() unsafe.Pointer { // *C.rosidl_message_type_support_t
	return g.typeSupport
}
