/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package px4_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/test_motor.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/TestMotor", TestMotorTypeSupport)
}
const (
	TestMotor_NUM_MOTOR_OUTPUTS uint8 = 8
	TestMotor_ACTION_STOP uint8 = 0// stop all motors (disable motor test mode)
	TestMotor_ACTION_RUN uint8 = 1// run motor(s) (enable motor test mode)
	TestMotor_ORB_QUEUE_LENGTH uint8 = 4
)

// Do not create instances of this type directly. Always use NewTestMotor
// function instead.
type TestMotor struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Action uint8 `yaml:"action"`// one of ACTION_* (applies to all motors)
	MotorNumber uint32 `yaml:"motor_number"`// number of motor to spin [0..N-1]
	Value float32 `yaml:"value"`// output power, range [0..1], -1 to stop individual motor
	TimeoutMs uint32 `yaml:"timeout_ms"`// timeout in ms after which to exit test mode (if 0, do not time out)
	DriverInstance uint8 `yaml:"driver_instance"`// select output driver (for boards with multiple outputs, like IO+FMU)
}

// NewTestMotor creates a new TestMotor with default values.
func NewTestMotor() *TestMotor {
	self := TestMotor{}
	self.SetDefaults()
	return &self
}

func (t *TestMotor) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *TestMotor) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var TestMotorTypeSupport types.MessageTypeSupport = _TestMotorTypeSupport{}

type _TestMotorTypeSupport struct{}

func (t _TestMotorTypeSupport) New() types.Message {
	return NewTestMotor()
}

func (t _TestMotorTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__TestMotor
	return (unsafe.Pointer)(C.px4_msgs__msg__TestMotor__create())
}

func (t _TestMotorTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__TestMotor__destroy((*C.px4_msgs__msg__TestMotor)(pointer_to_free))
}

func (t _TestMotorTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*TestMotor)
	mem := (*C.px4_msgs__msg__TestMotor)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.action = C.uint8_t(m.Action)
	mem.motor_number = C.uint32_t(m.MotorNumber)
	mem.value = C.float(m.Value)
	mem.timeout_ms = C.uint32_t(m.TimeoutMs)
	mem.driver_instance = C.uint8_t(m.DriverInstance)
}

func (t _TestMotorTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*TestMotor)
	mem := (*C.px4_msgs__msg__TestMotor)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.Action = uint8(mem.action)
	m.MotorNumber = uint32(mem.motor_number)
	m.Value = float32(mem.value)
	m.TimeoutMs = uint32(mem.timeout_ms)
	m.DriverInstance = uint8(mem.driver_instance)
}

func (t _TestMotorTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__TestMotor())
}

type CTestMotor = C.px4_msgs__msg__TestMotor
type CTestMotor__Sequence = C.px4_msgs__msg__TestMotor__Sequence

func TestMotor__Sequence_to_Go(goSlice *[]TestMotor, cSlice CTestMotor__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TestMotor, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__TestMotor__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__TestMotor * uintptr(i)),
		))
		TestMotorTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func TestMotor__Sequence_to_C(cSlice *CTestMotor__Sequence, goSlice []TestMotor) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__TestMotor)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__TestMotor * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__TestMotor)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__TestMotor * uintptr(i)),
		))
		TestMotorTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func TestMotor__Array_to_Go(goSlice []TestMotor, cSlice []CTestMotor) {
	for i := 0; i < len(cSlice); i++ {
		TestMotorTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func TestMotor__Array_to_C(cSlice []CTestMotor, goSlice []TestMotor) {
	for i := 0; i < len(goSlice); i++ {
		TestMotorTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
