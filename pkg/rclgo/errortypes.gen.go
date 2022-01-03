/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package rclgo

/*
#include <rcl/types.h>
#include <rmw/ret_types.h>
*/
import "C"
import (
	"runtime"
)

func errorsCastC(rcl_ret_t C.rcl_ret_t, context string) error {
	stackTraceBuffer := make([]byte, 2048)
	runtime.Stack(stackTraceBuffer, false) // Get stack trace of the current running thread only

	// https://stackoverflow.com/questions/9928221/table-of-functions-vs-switch-in-golang
	// switch-case is faster thanks to compiler optimization than a dispatcher?
	switch rcl_ret_t {
	case C.RCL_RET_ALREADY_INIT:
		return &AlreadyInit{rclRetStruct: rclRetStruct{rclRetCode: 100, trace: string(stackTraceBuffer), context: errorsBuildContext(&AlreadyInit{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_NOT_INIT:
		return &NotInit{rclRetStruct: rclRetStruct{rclRetCode: 101, trace: string(stackTraceBuffer), context: errorsBuildContext(&NotInit{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_MISMATCHED_RMW_ID:
		return &MismatchedRmwId{rclRetStruct: rclRetStruct{rclRetCode: 102, trace: string(stackTraceBuffer), context: errorsBuildContext(&MismatchedRmwId{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_TOPIC_NAME_INVALID:
		return &TopicNameInvalid{rclRetStruct: rclRetStruct{rclRetCode: 103, trace: string(stackTraceBuffer), context: errorsBuildContext(&TopicNameInvalid{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_SERVICE_NAME_INVALID:
		return &ServiceNameInvalid{rclRetStruct: rclRetStruct{rclRetCode: 104, trace: string(stackTraceBuffer), context: errorsBuildContext(&ServiceNameInvalid{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_UNKNOWN_SUBSTITUTION:
		return &UnknownSubstitution{rclRetStruct: rclRetStruct{rclRetCode: 105, trace: string(stackTraceBuffer), context: errorsBuildContext(&UnknownSubstitution{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_ALREADY_SHUTDOWN:
		return &AlreadyShutdown{rclRetStruct: rclRetStruct{rclRetCode: 106, trace: string(stackTraceBuffer), context: errorsBuildContext(&AlreadyShutdown{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_NODE_INVALID:
		return &NodeInvalid{rclRetStruct: rclRetStruct{rclRetCode: 200, trace: string(stackTraceBuffer), context: errorsBuildContext(&NodeInvalid{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_NODE_INVALID_NAME:
		return &NodeInvalidName{rclRetStruct: rclRetStruct{rclRetCode: 201, trace: string(stackTraceBuffer), context: errorsBuildContext(&NodeInvalidName{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_NODE_INVALID_NAMESPACE:
		return &NodeInvalidNamespace{rclRetStruct: rclRetStruct{rclRetCode: 202, trace: string(stackTraceBuffer), context: errorsBuildContext(&NodeInvalidNamespace{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_NODE_NAME_NON_EXISTENT:
		return &NodeNameNonExistent{rclRetStruct: rclRetStruct{rclRetCode: 203, trace: string(stackTraceBuffer), context: errorsBuildContext(&NodeNameNonExistent{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_PUBLISHER_INVALID:
		return &PublisherInvalid{rclRetStruct: rclRetStruct{rclRetCode: 300, trace: string(stackTraceBuffer), context: errorsBuildContext(&PublisherInvalid{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_SUBSCRIPTION_INVALID:
		return &SubscriptionInvalid{rclRetStruct: rclRetStruct{rclRetCode: 400, trace: string(stackTraceBuffer), context: errorsBuildContext(&SubscriptionInvalid{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_SUBSCRIPTION_TAKE_FAILED:
		return &SubscriptionTakeFailed{rclRetStruct: rclRetStruct{rclRetCode: 401, trace: string(stackTraceBuffer), context: errorsBuildContext(&SubscriptionTakeFailed{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_CLIENT_INVALID:
		return &ClientInvalid{rclRetStruct: rclRetStruct{rclRetCode: 500, trace: string(stackTraceBuffer), context: errorsBuildContext(&ClientInvalid{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_CLIENT_TAKE_FAILED:
		return &ClientTakeFailed{rclRetStruct: rclRetStruct{rclRetCode: 501, trace: string(stackTraceBuffer), context: errorsBuildContext(&ClientTakeFailed{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_SERVICE_INVALID:
		return &ServiceInvalid{rclRetStruct: rclRetStruct{rclRetCode: 600, trace: string(stackTraceBuffer), context: errorsBuildContext(&ServiceInvalid{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_SERVICE_TAKE_FAILED:
		return &ServiceTakeFailed{rclRetStruct: rclRetStruct{rclRetCode: 601, trace: string(stackTraceBuffer), context: errorsBuildContext(&ServiceTakeFailed{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_TIMER_INVALID:
		return &TimerInvalid{rclRetStruct: rclRetStruct{rclRetCode: 800, trace: string(stackTraceBuffer), context: errorsBuildContext(&TimerInvalid{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_TIMER_CANCELED:
		return &TimerCanceled{rclRetStruct: rclRetStruct{rclRetCode: 801, trace: string(stackTraceBuffer), context: errorsBuildContext(&TimerCanceled{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_WAIT_SET_INVALID:
		return &WaitSetInvalid{rclRetStruct: rclRetStruct{rclRetCode: 900, trace: string(stackTraceBuffer), context: errorsBuildContext(&WaitSetInvalid{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_WAIT_SET_EMPTY:
		return &WaitSetEmpty{rclRetStruct: rclRetStruct{rclRetCode: 901, trace: string(stackTraceBuffer), context: errorsBuildContext(&WaitSetEmpty{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_WAIT_SET_FULL:
		return &WaitSetFull{rclRetStruct: rclRetStruct{rclRetCode: 902, trace: string(stackTraceBuffer), context: errorsBuildContext(&WaitSetFull{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_INVALID_REMAP_RULE:
		return &InvalidRemapRule{rclRetStruct: rclRetStruct{rclRetCode: 1001, trace: string(stackTraceBuffer), context: errorsBuildContext(&InvalidRemapRule{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_WRONG_LEXEME:
		return &WrongLexeme{rclRetStruct: rclRetStruct{rclRetCode: 1002, trace: string(stackTraceBuffer), context: errorsBuildContext(&WrongLexeme{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_INVALID_ROS_ARGS:
		return &InvalidRosArgs{rclRetStruct: rclRetStruct{rclRetCode: 1003, trace: string(stackTraceBuffer), context: errorsBuildContext(&InvalidRosArgs{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_INVALID_PARAM_RULE:
		return &InvalidParamRule{rclRetStruct: rclRetStruct{rclRetCode: 1010, trace: string(stackTraceBuffer), context: errorsBuildContext(&InvalidParamRule{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_INVALID_LOG_LEVEL_RULE:
		return &InvalidLogLevelRule{rclRetStruct: rclRetStruct{rclRetCode: 1020, trace: string(stackTraceBuffer), context: errorsBuildContext(&InvalidLogLevelRule{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_EVENT_INVALID:
		return &EventInvalid{rclRetStruct: rclRetStruct{rclRetCode: 2000, trace: string(stackTraceBuffer), context: errorsBuildContext(&EventInvalid{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_EVENT_TAKE_FAILED:
		return &EventTakeFailed{rclRetStruct: rclRetStruct{rclRetCode: 2001, trace: string(stackTraceBuffer), context: errorsBuildContext(&EventTakeFailed{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_LIFECYCLE_STATE_REGISTERED:
		return &LifecycleStateRegistered{rclRetStruct: rclRetStruct{rclRetCode: 3000, trace: string(stackTraceBuffer), context: errorsBuildContext(&LifecycleStateRegistered{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_LIFECYCLE_STATE_NOT_REGISTERED:
		return &LifecycleStateNotRegistered{rclRetStruct: rclRetStruct{rclRetCode: 3001, trace: string(stackTraceBuffer), context: errorsBuildContext(&LifecycleStateNotRegistered{}, context, string(stackTraceBuffer))}}
	case C.RMW_RET_OK:
		return &RmwOk{rclRetStruct: rclRetStruct{rclRetCode: 0, trace: string(stackTraceBuffer), context: errorsBuildContext(&RmwOk{}, context, string(stackTraceBuffer))}}
	case C.RMW_RET_ERROR:
		return &RmwError{rclRetStruct: rclRetStruct{rclRetCode: 1, trace: string(stackTraceBuffer), context: errorsBuildContext(&RmwError{}, context, string(stackTraceBuffer))}}
	case C.RMW_RET_TIMEOUT:
		return &RmwTimeout{rclRetStruct: rclRetStruct{rclRetCode: 2, trace: string(stackTraceBuffer), context: errorsBuildContext(&RmwTimeout{}, context, string(stackTraceBuffer))}}
	case C.RMW_RET_UNSUPPORTED:
		return &RmwUnsupported{rclRetStruct: rclRetStruct{rclRetCode: 3, trace: string(stackTraceBuffer), context: errorsBuildContext(&RmwUnsupported{}, context, string(stackTraceBuffer))}}
	case C.RMW_RET_BAD_ALLOC:
		return &RmwBadAlloc{rclRetStruct: rclRetStruct{rclRetCode: 10, trace: string(stackTraceBuffer), context: errorsBuildContext(&RmwBadAlloc{}, context, string(stackTraceBuffer))}}
	case C.RMW_RET_INVALID_ARGUMENT:
		return &RmwInvalidArgument{rclRetStruct: rclRetStruct{rclRetCode: 11, trace: string(stackTraceBuffer), context: errorsBuildContext(&RmwInvalidArgument{}, context, string(stackTraceBuffer))}}
	case C.RMW_RET_INCORRECT_RMW_IMPLEMENTATION:
		return &RmwIncorrectRmwImplementation{rclRetStruct: rclRetStruct{rclRetCode: 12, trace: string(stackTraceBuffer), context: errorsBuildContext(&RmwIncorrectRmwImplementation{}, context, string(stackTraceBuffer))}}
	
	default:
		return &UnknownReturnCode{rclRetStruct: rclRetStruct{rclRetCode: int(rcl_ret_t), context: context}}
	}
}

type UnknownReturnCode struct {
	rclRetStruct
}


// AlreadyInit rcl specific ret codes start at 100rcl_init() already called return code.
type AlreadyInit struct {
	rclRetStruct
}

// NotInit rcl_init() not yet called return code.
type NotInit struct {
	rclRetStruct
}

// MismatchedRmwId Mismatched rmw identifier return code.
type MismatchedRmwId struct {
	rclRetStruct
}

// TopicNameInvalid Topic name does not pass validation.
type TopicNameInvalid struct {
	rclRetStruct
}

// ServiceNameInvalid Service name (same as topic name) does not pass validation.
type ServiceNameInvalid struct {
	rclRetStruct
}

// UnknownSubstitution Topic name substitution is unknown.
type UnknownSubstitution struct {
	rclRetStruct
}

// AlreadyShutdown rcl_shutdown() already called return code.
type AlreadyShutdown struct {
	rclRetStruct
}

// NodeInvalid rcl node specific ret codes in 2XXInvalid rcl_node_t given return code.
type NodeInvalid struct {
	rclRetStruct
}

// NodeInvalidName Invalid node name return code.
type NodeInvalidName struct {
	rclRetStruct
}

// NodeInvalidNamespace Invalid node namespace return code.
type NodeInvalidNamespace struct {
	rclRetStruct
}

// NodeNameNonExistent Failed to find node name
type NodeNameNonExistent struct {
	rclRetStruct
}

// PublisherInvalid rcl publisher specific ret codes in 3XXInvalid rcl_publisher_t given return code.
type PublisherInvalid struct {
	rclRetStruct
}

// SubscriptionInvalid rcl subscription specific ret codes in 4XXInvalid rcl_subscription_t given return code.
type SubscriptionInvalid struct {
	rclRetStruct
}

// SubscriptionTakeFailed Failed to take a message from the subscription return code.
type SubscriptionTakeFailed struct {
	rclRetStruct
}

// ClientInvalid rcl service client specific ret codes in 5XXInvalid rcl_client_t given return code.
type ClientInvalid struct {
	rclRetStruct
}

// ClientTakeFailed Failed to take a response from the client return code.
type ClientTakeFailed struct {
	rclRetStruct
}

// ServiceInvalid rcl service server specific ret codes in 6XXInvalid rcl_service_t given return code.
type ServiceInvalid struct {
	rclRetStruct
}

// ServiceTakeFailed Failed to take a request from the service return code.
type ServiceTakeFailed struct {
	rclRetStruct
}

// TimerInvalid rcl timer specific ret codes in 8XXInvalid rcl_timer_t given return code.
type TimerInvalid struct {
	rclRetStruct
}

// TimerCanceled Given timer was canceled return code.
type TimerCanceled struct {
	rclRetStruct
}

// WaitSetInvalid rcl wait and wait set specific ret codes in 9XXInvalid rcl_wait_set_t given return code.
type WaitSetInvalid struct {
	rclRetStruct
}

// WaitSetEmpty Given rcl_wait_set_t is empty return code.
type WaitSetEmpty struct {
	rclRetStruct
}

// WaitSetFull Given rcl_wait_set_t is full return code.
type WaitSetFull struct {
	rclRetStruct
}

// InvalidRemapRule rcl argument parsing specific ret codes in 1XXXArgument is not a valid remap rule
type InvalidRemapRule struct {
	rclRetStruct
}

// WrongLexeme Expected one type of lexeme but got another
type WrongLexeme struct {
	rclRetStruct
}

// InvalidRosArgs Found invalid ros argument while parsing
type InvalidRosArgs struct {
	rclRetStruct
}

// InvalidParamRule Argument is not a valid parameter rule
type InvalidParamRule struct {
	rclRetStruct
}

// InvalidLogLevelRule Argument is not a valid log level rule
type InvalidLogLevelRule struct {
	rclRetStruct
}

// EventInvalid rcl event specific ret codes in 20XXInvalid rcl_event_t given return code.
type EventInvalid struct {
	rclRetStruct
}

// EventTakeFailed Failed to take an event from the event handle
type EventTakeFailed struct {
	rclRetStruct
}

// LifecycleStateRegistered rcl_lifecycle state register ret codes in 30XXrcl_lifecycle state registered
type LifecycleStateRegistered struct {
	rclRetStruct
}

// LifecycleStateNotRegistered rcl_lifecycle state not registered
type LifecycleStateNotRegistered struct {
	rclRetStruct
}

// RmwOk Return code for rmw functionsThe operation ran as expected
type RmwOk struct {
	rclRetStruct
}

// RmwError Generic error to indicate operation could not complete successfully
type RmwError struct {
	rclRetStruct
}

// RmwTimeout The operation was halted early because it exceeded its timeout critera
type RmwTimeout struct {
	rclRetStruct
}

// RmwUnsupported The operation or event handling is not supported.
type RmwUnsupported struct {
	rclRetStruct
}

// RmwBadAlloc Failed to allocate memory
type RmwBadAlloc struct {
	rclRetStruct
}

// RmwInvalidArgument Argument to function was invalid
type RmwInvalidArgument struct {
	rclRetStruct
}

// RmwIncorrectRmwImplementation Incorrect rmw implementation.
type RmwIncorrectRmwImplementation struct {
	rclRetStruct
}

// RmwNodeNameNonExistent rmw node specific ret codes in 2XXFailed to find node nameUsing same return code than in rcl
type RmwNodeNameNonExistent struct {
	rclRetStruct
}



// Ok Success return code.
type Ok = RmwOk

// Error Unspecified error return code.
type Error = RmwError

// Timeout Timeout occurred return code.
type Timeout = RmwTimeout

// BadAlloc Failed to allocate memory return code.
type BadAlloc = RmwBadAlloc

// InvalidArgument Invalid argument return code.
type InvalidArgument = RmwInvalidArgument

// Unsupported Unsupported return code.
type Unsupported = RmwUnsupported


