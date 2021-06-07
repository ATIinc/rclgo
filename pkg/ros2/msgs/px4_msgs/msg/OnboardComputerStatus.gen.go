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
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/onboard_computer_status.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/OnboardComputerStatus", OnboardComputerStatusTypeSupport)
}

// Do not create instances of this type directly. Always use NewOnboardComputerStatus
// function instead.
type OnboardComputerStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// [us] time since system start (microseconds). ONBOARD_COMPUTER_STATUS message data
	Uptime uint32 `yaml:"uptime"`// [ms] time since system boot of the companion (milliseconds). ONBOARD_COMPUTER_STATUS message data
	Type uint8 `yaml:"type"`// type of onboard computer 0: Mission computer primary, 1: Mission computer backup 1, 2: Mission computer backup 2, 3: Compute node, 4-5: Compute spares, 6-9: Payload computers.
	CpuCores [8]uint8 `yaml:"cpu_cores"`// CPU usage on the component in percent
	CpuCombined [10]uint8 `yaml:"cpu_combined"`// Combined CPU usage as the last 10 slices of 100 MS
	GpuCores [4]uint8 `yaml:"gpu_cores"`// GPU usage on the component in percent
	GpuCombined [10]uint8 `yaml:"gpu_combined"`// Combined GPU usage as the last 10 slices of 100 MS
	TemperatureBoard int8 `yaml:"temperature_board"`// [degC] Temperature of the board
	TemperatureCore [8]int8 `yaml:"temperature_core"`// [degC] Temperature of the CPU core
	FanSpeed [4]int16 `yaml:"fan_speed"`// [rpm] Fan speeds
	RamUsage uint32 `yaml:"ram_usage"`// [MB] Amount of used RAM on the component system
	RamTotal uint32 `yaml:"ram_total"`// [MB] Total amount of RAM on the component system
	StorageType [4]uint32 `yaml:"storage_type"`// Storage type: 0: HDD, 1: SSD, 2: EMMC, 3: SD card (non-removable), 4: SD card (removable)
	StorageUsage [4]uint32 `yaml:"storage_usage"`// [MB] Amount of used storage space on the component system
	StorageTotal [4]uint32 `yaml:"storage_total"`// [MB] Total amount of storage space on the component system
	LinkType [6]uint32 `yaml:"link_type"`// [Kb/s] Link type: 0-9: UART, 10-19: Wired network, 20-29: Wifi, 30-39: Point-to-point proprietary, 40-49: Mesh proprietary
	LinkTxRate [6]uint32 `yaml:"link_tx_rate"`// [Kb/s] Network traffic from the component system
	LinkRxRate [6]uint32 `yaml:"link_rx_rate"`// [Kb/s] Network traffic to the component system
	LinkTxMax [6]uint32 `yaml:"link_tx_max"`// [Kb/s] Network capacity from the component system
	LinkRxMax [6]uint32 `yaml:"link_rx_max"`// [Kb/s] Network capacity to the component system
}

// NewOnboardComputerStatus creates a new OnboardComputerStatus with default values.
func NewOnboardComputerStatus() *OnboardComputerStatus {
	self := OnboardComputerStatus{}
	self.SetDefaults()
	return &self
}

func (t *OnboardComputerStatus) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *OnboardComputerStatus) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var OnboardComputerStatusTypeSupport types.MessageTypeSupport = _OnboardComputerStatusTypeSupport{}

type _OnboardComputerStatusTypeSupport struct{}

func (t _OnboardComputerStatusTypeSupport) New() types.Message {
	return NewOnboardComputerStatus()
}

func (t _OnboardComputerStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__OnboardComputerStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__OnboardComputerStatus__create())
}

func (t _OnboardComputerStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__OnboardComputerStatus__destroy((*C.px4_msgs__msg__OnboardComputerStatus)(pointer_to_free))
}

func (t _OnboardComputerStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*OnboardComputerStatus)
	mem := (*C.px4_msgs__msg__OnboardComputerStatus)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.uptime = C.uint32_t(m.Uptime)
	mem._type = C.uint8_t(m.Type)
	cSlice_cpu_cores := mem.cpu_cores[:]
	rosidl_runtime_c.Uint8__Array_to_C(*(*[]rosidl_runtime_c.CUint8)(unsafe.Pointer(&cSlice_cpu_cores)), m.CpuCores[:])
	cSlice_cpu_combined := mem.cpu_combined[:]
	rosidl_runtime_c.Uint8__Array_to_C(*(*[]rosidl_runtime_c.CUint8)(unsafe.Pointer(&cSlice_cpu_combined)), m.CpuCombined[:])
	cSlice_gpu_cores := mem.gpu_cores[:]
	rosidl_runtime_c.Uint8__Array_to_C(*(*[]rosidl_runtime_c.CUint8)(unsafe.Pointer(&cSlice_gpu_cores)), m.GpuCores[:])
	cSlice_gpu_combined := mem.gpu_combined[:]
	rosidl_runtime_c.Uint8__Array_to_C(*(*[]rosidl_runtime_c.CUint8)(unsafe.Pointer(&cSlice_gpu_combined)), m.GpuCombined[:])
	mem.temperature_board = C.int8_t(m.TemperatureBoard)
	cSlice_temperature_core := mem.temperature_core[:]
	rosidl_runtime_c.Int8__Array_to_C(*(*[]rosidl_runtime_c.CInt8)(unsafe.Pointer(&cSlice_temperature_core)), m.TemperatureCore[:])
	cSlice_fan_speed := mem.fan_speed[:]
	rosidl_runtime_c.Int16__Array_to_C(*(*[]rosidl_runtime_c.CInt16)(unsafe.Pointer(&cSlice_fan_speed)), m.FanSpeed[:])
	mem.ram_usage = C.uint32_t(m.RamUsage)
	mem.ram_total = C.uint32_t(m.RamTotal)
	cSlice_storage_type := mem.storage_type[:]
	rosidl_runtime_c.Uint32__Array_to_C(*(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_storage_type)), m.StorageType[:])
	cSlice_storage_usage := mem.storage_usage[:]
	rosidl_runtime_c.Uint32__Array_to_C(*(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_storage_usage)), m.StorageUsage[:])
	cSlice_storage_total := mem.storage_total[:]
	rosidl_runtime_c.Uint32__Array_to_C(*(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_storage_total)), m.StorageTotal[:])
	cSlice_link_type := mem.link_type[:]
	rosidl_runtime_c.Uint32__Array_to_C(*(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_link_type)), m.LinkType[:])
	cSlice_link_tx_rate := mem.link_tx_rate[:]
	rosidl_runtime_c.Uint32__Array_to_C(*(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_link_tx_rate)), m.LinkTxRate[:])
	cSlice_link_rx_rate := mem.link_rx_rate[:]
	rosidl_runtime_c.Uint32__Array_to_C(*(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_link_rx_rate)), m.LinkRxRate[:])
	cSlice_link_tx_max := mem.link_tx_max[:]
	rosidl_runtime_c.Uint32__Array_to_C(*(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_link_tx_max)), m.LinkTxMax[:])
	cSlice_link_rx_max := mem.link_rx_max[:]
	rosidl_runtime_c.Uint32__Array_to_C(*(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_link_rx_max)), m.LinkRxMax[:])
}

func (t _OnboardComputerStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*OnboardComputerStatus)
	mem := (*C.px4_msgs__msg__OnboardComputerStatus)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.Uptime = uint32(mem.uptime)
	m.Type = uint8(mem._type)
	cSlice_cpu_cores := mem.cpu_cores[:]
	rosidl_runtime_c.Uint8__Array_to_Go(m.CpuCores[:], *(*[]rosidl_runtime_c.CUint8)(unsafe.Pointer(&cSlice_cpu_cores)))
	cSlice_cpu_combined := mem.cpu_combined[:]
	rosidl_runtime_c.Uint8__Array_to_Go(m.CpuCombined[:], *(*[]rosidl_runtime_c.CUint8)(unsafe.Pointer(&cSlice_cpu_combined)))
	cSlice_gpu_cores := mem.gpu_cores[:]
	rosidl_runtime_c.Uint8__Array_to_Go(m.GpuCores[:], *(*[]rosidl_runtime_c.CUint8)(unsafe.Pointer(&cSlice_gpu_cores)))
	cSlice_gpu_combined := mem.gpu_combined[:]
	rosidl_runtime_c.Uint8__Array_to_Go(m.GpuCombined[:], *(*[]rosidl_runtime_c.CUint8)(unsafe.Pointer(&cSlice_gpu_combined)))
	m.TemperatureBoard = int8(mem.temperature_board)
	cSlice_temperature_core := mem.temperature_core[:]
	rosidl_runtime_c.Int8__Array_to_Go(m.TemperatureCore[:], *(*[]rosidl_runtime_c.CInt8)(unsafe.Pointer(&cSlice_temperature_core)))
	cSlice_fan_speed := mem.fan_speed[:]
	rosidl_runtime_c.Int16__Array_to_Go(m.FanSpeed[:], *(*[]rosidl_runtime_c.CInt16)(unsafe.Pointer(&cSlice_fan_speed)))
	m.RamUsage = uint32(mem.ram_usage)
	m.RamTotal = uint32(mem.ram_total)
	cSlice_storage_type := mem.storage_type[:]
	rosidl_runtime_c.Uint32__Array_to_Go(m.StorageType[:], *(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_storage_type)))
	cSlice_storage_usage := mem.storage_usage[:]
	rosidl_runtime_c.Uint32__Array_to_Go(m.StorageUsage[:], *(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_storage_usage)))
	cSlice_storage_total := mem.storage_total[:]
	rosidl_runtime_c.Uint32__Array_to_Go(m.StorageTotal[:], *(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_storage_total)))
	cSlice_link_type := mem.link_type[:]
	rosidl_runtime_c.Uint32__Array_to_Go(m.LinkType[:], *(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_link_type)))
	cSlice_link_tx_rate := mem.link_tx_rate[:]
	rosidl_runtime_c.Uint32__Array_to_Go(m.LinkTxRate[:], *(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_link_tx_rate)))
	cSlice_link_rx_rate := mem.link_rx_rate[:]
	rosidl_runtime_c.Uint32__Array_to_Go(m.LinkRxRate[:], *(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_link_rx_rate)))
	cSlice_link_tx_max := mem.link_tx_max[:]
	rosidl_runtime_c.Uint32__Array_to_Go(m.LinkTxMax[:], *(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_link_tx_max)))
	cSlice_link_rx_max := mem.link_rx_max[:]
	rosidl_runtime_c.Uint32__Array_to_Go(m.LinkRxMax[:], *(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_link_rx_max)))
}

func (t _OnboardComputerStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__OnboardComputerStatus())
}

type COnboardComputerStatus = C.px4_msgs__msg__OnboardComputerStatus
type COnboardComputerStatus__Sequence = C.px4_msgs__msg__OnboardComputerStatus__Sequence

func OnboardComputerStatus__Sequence_to_Go(goSlice *[]OnboardComputerStatus, cSlice COnboardComputerStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]OnboardComputerStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__OnboardComputerStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__OnboardComputerStatus * uintptr(i)),
		))
		OnboardComputerStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func OnboardComputerStatus__Sequence_to_C(cSlice *COnboardComputerStatus__Sequence, goSlice []OnboardComputerStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__OnboardComputerStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__OnboardComputerStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__OnboardComputerStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__OnboardComputerStatus * uintptr(i)),
		))
		OnboardComputerStatusTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func OnboardComputerStatus__Array_to_Go(goSlice []OnboardComputerStatus, cSlice []COnboardComputerStatus) {
	for i := 0; i < len(cSlice); i++ {
		OnboardComputerStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func OnboardComputerStatus__Array_to_C(cSlice []COnboardComputerStatus, goSlice []OnboardComputerStatus) {
	for i := 0; i < len(goSlice); i++ {
		OnboardComputerStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
