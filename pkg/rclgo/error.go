/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

package rclgo

/*
#include <rcl/types.h>
#include <rcutils/error_handling.h>

// Reads and clears in one call so a later failure that sets no message cannot
// report this one's.
static rcutils_error_string_t rclgo_take_error_string(void) {
	rcutils_error_string_t s = rcutils_get_error_string();
	rcutils_reset_error();
	return s;
}
*/
import "C"

import (
	"fmt"
	"runtime"
	"slices"
)

type rclError struct {
	rclRetCode int
	context    string
	trace      string
}

func (e *rclError) Error() string {
	return e.context
}

/// Return the error message followed by `, at <file>:<line>` if set, else "error not set".
/**
 * This function is "safe" because it returns a copy of the current error
 * string or one containing the string "error not set" if no error was set.
 * This ensures that the copy is owned by the calling thread and is therefore
 * never invalidated by other error handling calls, and that the C string
 * inside is always valid and null terminated.
 *
 * \return The current error string, with file and line number, or "error not set" if not set.
 */
func errorString() string {
	rcutils_error_string_str := C.rclgo_take_error_string().str

	// Because the C string is null-terminated, we need to find the NULL-character to know where the string ends.
	// Otherwise we create a Go string of length 1024 of NULLs and gibberish
	bytes := make([]byte, len(rcutils_error_string_str))
	for i := range rcutils_error_string_str {
		if byte(rcutils_error_string_str[i]) == 0x00 {
			return string(bytes[:i]) // ending slice offset is exclusive
		}
		bytes[i] = byte(rcutils_error_string_str[i])
	}
	return string(bytes)

	// This would be much faster I guess.
	// upt := (*[1024]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&rcutils_error_string_str))))
	// return string((*upt)[:])
}

func errorsBuildContext(e error, ctx string, stackTrace string) string {
	return fmt.Sprintf("[%T]", e) + " " + ctx + " " + errorString() + "\n" + stackTrace + "\n"
}

// rcutils error state is thread-local, so a failing call and the read of its
// error string must run on one OS thread. rclCall and rclCallRC lock the
// goroutine to its thread for exactly that span; hot paths inline the same.

// rclCall runs call and, if it does not return RCL_RET_OK, builds the error
// from the rcutils error state with context.
func rclCall(context string, call func() C.rcl_ret_t) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	if rc := call(); rc != C.RCL_RET_OK {
		return errorsCastC(rc, context)
	}
	return nil
}

// rclCallRC is rclCall for callers that branch on the return code. Codes in
// expected are returned with a nil error and leave the error state unread.
func rclCallRC(context string, call func() C.rcl_ret_t, expected ...C.rcl_ret_t) (C.rcl_ret_t, error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	rc := call()
	if rc == C.RCL_RET_OK || slices.Contains(expected, rc) {
		return rc, nil
	}
	return rc, errorsCastC(rc, context)
}

// rclIgnore runs call, whose failure the caller cannot act on, and consumes
// the error state on failure so a later failure cannot report it.
func rclIgnore(logger *Logger, context string, call func() C.rcl_ret_t) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	if rc := call(); rc != C.RCL_RET_OK {
		if logger != nil {
			logger.Debug(errorsCastC(rc, context))
		} else {
			C.rcutils_reset_error()
		}
	}
}

// rclCallPtr is rclCall for calls that return nil on failure. It returns the
// consumed error string when they do.
func rclCallPtr[T any](call func() *T) (*T, string) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	if p := call(); p != nil {
		return p, ""
	}
	return nil, errorString()
}

func errorsCast(rcl_ret_t C.rcl_ret_t) error {
	return errorsCastC(rcl_ret_t, "")
}

func onErr(err *error, f func() error) {
	if *err != nil {
		f() //nolint:errcheck
	}
}

type closeError string

func (e closeError) Error() string {
	return string(e)
}

func closeErr(s string) error {
	return closeError("tried to close a closed " + s)
}
