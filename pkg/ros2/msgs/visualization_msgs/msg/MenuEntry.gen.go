/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package visualization_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lvisualization_msgs__rosidl_typesupport_c -lvisualization_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <visualization_msgs/msg/menu_entry.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("visualization_msgs/MenuEntry", MenuEntryTypeSupport)
}
const (
	MenuEntry_FEEDBACK uint8 = 0// Command_type stores the type of response desired when this menuentry is clicked.FEEDBACK: send an InteractiveMarkerFeedback message with menu_entry_id set to this entry's id.ROSRUN: execute "rosrun" with arguments given in the command field (above).ROSLAUNCH: execute "roslaunch" with arguments given in the command field (above).
	MenuEntry_ROSRUN uint8 = 1// Command_type stores the type of response desired when this menuentry is clicked.FEEDBACK: send an InteractiveMarkerFeedback message with menu_entry_id set to this entry's id.ROSRUN: execute "rosrun" with arguments given in the command field (above).ROSLAUNCH: execute "roslaunch" with arguments given in the command field (above).
	MenuEntry_ROSLAUNCH uint8 = 2// Command_type stores the type of response desired when this menuentry is clicked.FEEDBACK: send an InteractiveMarkerFeedback message with menu_entry_id set to this entry's id.ROSRUN: execute "rosrun" with arguments given in the command field (above).ROSLAUNCH: execute "roslaunch" with arguments given in the command field (above).
)

// Do not create instances of this type directly. Always use NewMenuEntry
// function instead.
type MenuEntry struct {
	Id uint32 `yaml:"id"`// ID is a number for each menu entry.  Must be unique within thecontrol, and should never be 0.
	ParentId uint32 `yaml:"parent_id"`// ID of the parent of this menu entry, if it is a submenu.  If thismenu entry is a top-level entry, set parent_id to 0.
	Title string `yaml:"title"`// menu / entry title
	Command string `yaml:"command"`// Arguments to command indicated by command_type (below)
	CommandType uint8 `yaml:"command_type"`// Command_type stores the type of response desired when this menuentry is clicked.FEEDBACK: send an InteractiveMarkerFeedback message with menu_entry_id set to this entry's id.ROSRUN: execute "rosrun" with arguments given in the command field (above).ROSLAUNCH: execute "roslaunch" with arguments given in the command field (above).
}

// NewMenuEntry creates a new MenuEntry with default values.
func NewMenuEntry() *MenuEntry {
	self := MenuEntry{}
	self.SetDefaults()
	return &self
}

func (t *MenuEntry) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *MenuEntry) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var MenuEntryTypeSupport types.MessageTypeSupport = _MenuEntryTypeSupport{}

type _MenuEntryTypeSupport struct{}

func (t _MenuEntryTypeSupport) New() types.Message {
	return NewMenuEntry()
}

func (t _MenuEntryTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.visualization_msgs__msg__MenuEntry
	return (unsafe.Pointer)(C.visualization_msgs__msg__MenuEntry__create())
}

func (t _MenuEntryTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.visualization_msgs__msg__MenuEntry__destroy((*C.visualization_msgs__msg__MenuEntry)(pointer_to_free))
}

func (t _MenuEntryTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*MenuEntry)
	mem := (*C.visualization_msgs__msg__MenuEntry)(dst)
	mem.id = C.uint32_t(m.Id)
	mem.parent_id = C.uint32_t(m.ParentId)
	rosidl_runtime_c.StringAsCStruct(unsafe.Pointer(&mem.title), m.Title)
	rosidl_runtime_c.StringAsCStruct(unsafe.Pointer(&mem.command), m.Command)
	mem.command_type = C.uint8_t(m.CommandType)
}

func (t _MenuEntryTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*MenuEntry)
	mem := (*C.visualization_msgs__msg__MenuEntry)(ros2_message_buffer)
	m.Id = uint32(mem.id)
	m.ParentId = uint32(mem.parent_id)
	rosidl_runtime_c.StringAsGoStruct(&m.Title, unsafe.Pointer(&mem.title))
	rosidl_runtime_c.StringAsGoStruct(&m.Command, unsafe.Pointer(&mem.command))
	m.CommandType = uint8(mem.command_type)
}

func (t _MenuEntryTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__visualization_msgs__msg__MenuEntry())
}

type CMenuEntry = C.visualization_msgs__msg__MenuEntry
type CMenuEntry__Sequence = C.visualization_msgs__msg__MenuEntry__Sequence

func MenuEntry__Sequence_to_Go(goSlice *[]MenuEntry, cSlice CMenuEntry__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MenuEntry, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.visualization_msgs__msg__MenuEntry__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__MenuEntry * uintptr(i)),
		))
		MenuEntryTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func MenuEntry__Sequence_to_C(cSlice *CMenuEntry__Sequence, goSlice []MenuEntry) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.visualization_msgs__msg__MenuEntry)(C.malloc((C.size_t)(C.sizeof_struct_visualization_msgs__msg__MenuEntry * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.visualization_msgs__msg__MenuEntry)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__MenuEntry * uintptr(i)),
		))
		MenuEntryTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func MenuEntry__Array_to_Go(goSlice []MenuEntry, cSlice []CMenuEntry) {
	for i := 0; i < len(cSlice); i++ {
		MenuEntryTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func MenuEntry__Array_to_C(cSlice []CMenuEntry, goSlice []MenuEntry) {
	for i := 0; i < len(goSlice); i++ {
		MenuEntryTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
