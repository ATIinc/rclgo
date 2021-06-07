/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package nav_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	geometry_msgs_msg "github.com/tiiuae/rclgo/pkg/ros2/msgs/geometry_msgs/msg"
	std_msgs_msg "github.com/tiiuae/rclgo/pkg/ros2/msgs/std_msgs/msg"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lnav_msgs__rosidl_typesupport_c -lnav_msgs__rosidl_generator_c
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <nav_msgs/msg/odometry.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("nav_msgs/Odometry", OdometryTypeSupport)
}

// Do not create instances of this type directly. Always use NewOdometry
// function instead.
type Odometry struct {
	Header std_msgs_msg.Header `yaml:"header"`// Includes the frame id of the pose parent.
	ChildFrameId string `yaml:"child_frame_id"`// Frame id the pose points to. The twist is in this coordinate frame.
	Pose geometry_msgs_msg.PoseWithCovariance `yaml:"pose"`// Estimated pose that is typically relative to a fixed world frame.
	Twist geometry_msgs_msg.TwistWithCovariance `yaml:"twist"`// Estimated linear and angular velocity relative to child_frame_id.
}

// NewOdometry creates a new Odometry with default values.
func NewOdometry() *Odometry {
	self := Odometry{}
	self.SetDefaults()
	return &self
}

func (t *Odometry) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *Odometry) SetDefaults() {
	t.Header.SetDefaults()
	t.Pose.SetDefaults()
	t.Twist.SetDefaults()
	
}

// Modifying this variable is undefined behavior.
var OdometryTypeSupport types.MessageTypeSupport = _OdometryTypeSupport{}

type _OdometryTypeSupport struct{}

func (t _OdometryTypeSupport) New() types.Message {
	return NewOdometry()
}

func (t _OdometryTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.nav_msgs__msg__Odometry
	return (unsafe.Pointer)(C.nav_msgs__msg__Odometry__create())
}

func (t _OdometryTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.nav_msgs__msg__Odometry__destroy((*C.nav_msgs__msg__Odometry)(pointer_to_free))
}

func (t _OdometryTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Odometry)
	mem := (*C.nav_msgs__msg__Odometry)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	rosidl_runtime_c.StringAsCStruct(unsafe.Pointer(&mem.child_frame_id), m.ChildFrameId)
	geometry_msgs_msg.PoseWithCovarianceTypeSupport.AsCStruct(unsafe.Pointer(&mem.pose), &m.Pose)
	geometry_msgs_msg.TwistWithCovarianceTypeSupport.AsCStruct(unsafe.Pointer(&mem.twist), &m.Twist)
}

func (t _OdometryTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Odometry)
	mem := (*C.nav_msgs__msg__Odometry)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	rosidl_runtime_c.StringAsGoStruct(&m.ChildFrameId, unsafe.Pointer(&mem.child_frame_id))
	geometry_msgs_msg.PoseWithCovarianceTypeSupport.AsGoStruct(&m.Pose, unsafe.Pointer(&mem.pose))
	geometry_msgs_msg.TwistWithCovarianceTypeSupport.AsGoStruct(&m.Twist, unsafe.Pointer(&mem.twist))
}

func (t _OdometryTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__nav_msgs__msg__Odometry())
}

type COdometry = C.nav_msgs__msg__Odometry
type COdometry__Sequence = C.nav_msgs__msg__Odometry__Sequence

func Odometry__Sequence_to_Go(goSlice *[]Odometry, cSlice COdometry__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Odometry, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.nav_msgs__msg__Odometry__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_nav_msgs__msg__Odometry * uintptr(i)),
		))
		OdometryTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Odometry__Sequence_to_C(cSlice *COdometry__Sequence, goSlice []Odometry) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.nav_msgs__msg__Odometry)(C.malloc((C.size_t)(C.sizeof_struct_nav_msgs__msg__Odometry * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.nav_msgs__msg__Odometry)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_nav_msgs__msg__Odometry * uintptr(i)),
		))
		OdometryTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Odometry__Array_to_Go(goSlice []Odometry, cSlice []COdometry) {
	for i := 0; i < len(cSlice); i++ {
		OdometryTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Odometry__Array_to_C(cSlice []COdometry, goSlice []Odometry) {
	for i := 0; i < len(goSlice); i++ {
		OdometryTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
