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

#include <px4_msgs/msg/vehicle_status_flags.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/VehicleStatusFlags", VehicleStatusFlagsTypeSupport)
}

// Do not create instances of this type directly. Always use NewVehicleStatusFlags
// function instead.
type VehicleStatusFlags struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	ConditionCalibrationEnabled bool `yaml:"condition_calibration_enabled"`
	ConditionSystemSensorsInitialized bool `yaml:"condition_system_sensors_initialized"`
	ConditionSystemHotplugTimeout bool `yaml:"condition_system_hotplug_timeout"`// true if the hotplug sensor search is over
	ConditionSystemReturnedToHome bool `yaml:"condition_system_returned_to_home"`
	ConditionAutoMissionAvailable bool `yaml:"condition_auto_mission_available"`
	ConditionAngularVelocityValid bool `yaml:"condition_angular_velocity_valid"`
	ConditionAttitudeValid bool `yaml:"condition_attitude_valid"`
	ConditionLocalAltitudeValid bool `yaml:"condition_local_altitude_valid"`
	ConditionLocalPositionValid bool `yaml:"condition_local_position_valid"`// set to true by the commander app if the quality of the local position estimate is good enough to use for navigation
	ConditionLocalVelocityValid bool `yaml:"condition_local_velocity_valid"`// set to true by the commander app if the quality of the local horizontal velocity data is good enough to use for navigation
	ConditionGlobalPositionValid bool `yaml:"condition_global_position_valid"`// set to true by the commander app if the quality of the global position estimate is good enough to use for navigation
	ConditionHomePositionValid bool `yaml:"condition_home_position_valid"`// indicates a valid home position (a valid home position is not always a valid launch)
	ConditionPowerInputValid bool `yaml:"condition_power_input_valid"`// set if input power is valid
	ConditionBatteryHealthy bool `yaml:"condition_battery_healthy"`// set if battery is available and not low
	ConditionEscsError bool `yaml:"condition_escs_error"`// set to true if one or more ESCs reporting esc_status are offline
	CircuitBreakerEngagedPowerCheck bool `yaml:"circuit_breaker_engaged_power_check"`
	CircuitBreakerEngagedAirspdCheck bool `yaml:"circuit_breaker_engaged_airspd_check"`
	CircuitBreakerEngagedEnginefailureCheck bool `yaml:"circuit_breaker_engaged_enginefailure_check"`
	CircuitBreakerFlightTerminationDisabled bool `yaml:"circuit_breaker_flight_termination_disabled"`
	CircuitBreakerEngagedUsbCheck bool `yaml:"circuit_breaker_engaged_usb_check"`
	CircuitBreakerEngagedPosfailureCheck bool `yaml:"circuit_breaker_engaged_posfailure_check"`// set to true when the position valid checks have been disabled
	CircuitBreakerVtolFwArmingCheck bool `yaml:"circuit_breaker_vtol_fw_arming_check"`// set to true if for VTOLs arming in fixed-wing mode should be allowed
	OffboardControlSignalFoundOnce bool `yaml:"offboard_control_signal_found_once"`
	OffboardControlSignalLost bool `yaml:"offboard_control_signal_lost"`
	RcSignalFoundOnce bool `yaml:"rc_signal_found_once"`
	RcInputBlocked bool `yaml:"rc_input_blocked"`// set if RC input should be ignored temporarily
	RcCalibrationValid bool `yaml:"rc_calibration_valid"`// set if RC calibration is valid
	VtolTransitionFailure bool `yaml:"vtol_transition_failure"`// Set to true if vtol transition failed
	UsbConnected bool `yaml:"usb_connected"`// status of the USB power supply
	AvoidanceSystemRequired bool `yaml:"avoidance_system_required"`// Set to true if avoidance system is enabled via COM_OBS_AVOID parameter
	AvoidanceSystemValid bool `yaml:"avoidance_system_valid"`// Status of the obstacle avoidance system
}

// NewVehicleStatusFlags creates a new VehicleStatusFlags with default values.
func NewVehicleStatusFlags() *VehicleStatusFlags {
	self := VehicleStatusFlags{}
	self.SetDefaults()
	return &self
}

func (t *VehicleStatusFlags) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *VehicleStatusFlags) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var VehicleStatusFlagsTypeSupport types.MessageTypeSupport = _VehicleStatusFlagsTypeSupport{}

type _VehicleStatusFlagsTypeSupport struct{}

func (t _VehicleStatusFlagsTypeSupport) New() types.Message {
	return NewVehicleStatusFlags()
}

func (t _VehicleStatusFlagsTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleStatusFlags
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleStatusFlags__create())
}

func (t _VehicleStatusFlagsTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleStatusFlags__destroy((*C.px4_msgs__msg__VehicleStatusFlags)(pointer_to_free))
}

func (t _VehicleStatusFlagsTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*VehicleStatusFlags)
	mem := (*C.px4_msgs__msg__VehicleStatusFlags)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.condition_calibration_enabled = C.bool(m.ConditionCalibrationEnabled)
	mem.condition_system_sensors_initialized = C.bool(m.ConditionSystemSensorsInitialized)
	mem.condition_system_hotplug_timeout = C.bool(m.ConditionSystemHotplugTimeout)
	mem.condition_system_returned_to_home = C.bool(m.ConditionSystemReturnedToHome)
	mem.condition_auto_mission_available = C.bool(m.ConditionAutoMissionAvailable)
	mem.condition_angular_velocity_valid = C.bool(m.ConditionAngularVelocityValid)
	mem.condition_attitude_valid = C.bool(m.ConditionAttitudeValid)
	mem.condition_local_altitude_valid = C.bool(m.ConditionLocalAltitudeValid)
	mem.condition_local_position_valid = C.bool(m.ConditionLocalPositionValid)
	mem.condition_local_velocity_valid = C.bool(m.ConditionLocalVelocityValid)
	mem.condition_global_position_valid = C.bool(m.ConditionGlobalPositionValid)
	mem.condition_home_position_valid = C.bool(m.ConditionHomePositionValid)
	mem.condition_power_input_valid = C.bool(m.ConditionPowerInputValid)
	mem.condition_battery_healthy = C.bool(m.ConditionBatteryHealthy)
	mem.condition_escs_error = C.bool(m.ConditionEscsError)
	mem.circuit_breaker_engaged_power_check = C.bool(m.CircuitBreakerEngagedPowerCheck)
	mem.circuit_breaker_engaged_airspd_check = C.bool(m.CircuitBreakerEngagedAirspdCheck)
	mem.circuit_breaker_engaged_enginefailure_check = C.bool(m.CircuitBreakerEngagedEnginefailureCheck)
	mem.circuit_breaker_flight_termination_disabled = C.bool(m.CircuitBreakerFlightTerminationDisabled)
	mem.circuit_breaker_engaged_usb_check = C.bool(m.CircuitBreakerEngagedUsbCheck)
	mem.circuit_breaker_engaged_posfailure_check = C.bool(m.CircuitBreakerEngagedPosfailureCheck)
	mem.circuit_breaker_vtol_fw_arming_check = C.bool(m.CircuitBreakerVtolFwArmingCheck)
	mem.offboard_control_signal_found_once = C.bool(m.OffboardControlSignalFoundOnce)
	mem.offboard_control_signal_lost = C.bool(m.OffboardControlSignalLost)
	mem.rc_signal_found_once = C.bool(m.RcSignalFoundOnce)
	mem.rc_input_blocked = C.bool(m.RcInputBlocked)
	mem.rc_calibration_valid = C.bool(m.RcCalibrationValid)
	mem.vtol_transition_failure = C.bool(m.VtolTransitionFailure)
	mem.usb_connected = C.bool(m.UsbConnected)
	mem.avoidance_system_required = C.bool(m.AvoidanceSystemRequired)
	mem.avoidance_system_valid = C.bool(m.AvoidanceSystemValid)
}

func (t _VehicleStatusFlagsTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*VehicleStatusFlags)
	mem := (*C.px4_msgs__msg__VehicleStatusFlags)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.ConditionCalibrationEnabled = bool(mem.condition_calibration_enabled)
	m.ConditionSystemSensorsInitialized = bool(mem.condition_system_sensors_initialized)
	m.ConditionSystemHotplugTimeout = bool(mem.condition_system_hotplug_timeout)
	m.ConditionSystemReturnedToHome = bool(mem.condition_system_returned_to_home)
	m.ConditionAutoMissionAvailable = bool(mem.condition_auto_mission_available)
	m.ConditionAngularVelocityValid = bool(mem.condition_angular_velocity_valid)
	m.ConditionAttitudeValid = bool(mem.condition_attitude_valid)
	m.ConditionLocalAltitudeValid = bool(mem.condition_local_altitude_valid)
	m.ConditionLocalPositionValid = bool(mem.condition_local_position_valid)
	m.ConditionLocalVelocityValid = bool(mem.condition_local_velocity_valid)
	m.ConditionGlobalPositionValid = bool(mem.condition_global_position_valid)
	m.ConditionHomePositionValid = bool(mem.condition_home_position_valid)
	m.ConditionPowerInputValid = bool(mem.condition_power_input_valid)
	m.ConditionBatteryHealthy = bool(mem.condition_battery_healthy)
	m.ConditionEscsError = bool(mem.condition_escs_error)
	m.CircuitBreakerEngagedPowerCheck = bool(mem.circuit_breaker_engaged_power_check)
	m.CircuitBreakerEngagedAirspdCheck = bool(mem.circuit_breaker_engaged_airspd_check)
	m.CircuitBreakerEngagedEnginefailureCheck = bool(mem.circuit_breaker_engaged_enginefailure_check)
	m.CircuitBreakerFlightTerminationDisabled = bool(mem.circuit_breaker_flight_termination_disabled)
	m.CircuitBreakerEngagedUsbCheck = bool(mem.circuit_breaker_engaged_usb_check)
	m.CircuitBreakerEngagedPosfailureCheck = bool(mem.circuit_breaker_engaged_posfailure_check)
	m.CircuitBreakerVtolFwArmingCheck = bool(mem.circuit_breaker_vtol_fw_arming_check)
	m.OffboardControlSignalFoundOnce = bool(mem.offboard_control_signal_found_once)
	m.OffboardControlSignalLost = bool(mem.offboard_control_signal_lost)
	m.RcSignalFoundOnce = bool(mem.rc_signal_found_once)
	m.RcInputBlocked = bool(mem.rc_input_blocked)
	m.RcCalibrationValid = bool(mem.rc_calibration_valid)
	m.VtolTransitionFailure = bool(mem.vtol_transition_failure)
	m.UsbConnected = bool(mem.usb_connected)
	m.AvoidanceSystemRequired = bool(mem.avoidance_system_required)
	m.AvoidanceSystemValid = bool(mem.avoidance_system_valid)
}

func (t _VehicleStatusFlagsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleStatusFlags())
}

type CVehicleStatusFlags = C.px4_msgs__msg__VehicleStatusFlags
type CVehicleStatusFlags__Sequence = C.px4_msgs__msg__VehicleStatusFlags__Sequence

func VehicleStatusFlags__Sequence_to_Go(goSlice *[]VehicleStatusFlags, cSlice CVehicleStatusFlags__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleStatusFlags, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleStatusFlags__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleStatusFlags * uintptr(i)),
		))
		VehicleStatusFlagsTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func VehicleStatusFlags__Sequence_to_C(cSlice *CVehicleStatusFlags__Sequence, goSlice []VehicleStatusFlags) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleStatusFlags)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleStatusFlags * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleStatusFlags)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleStatusFlags * uintptr(i)),
		))
		VehicleStatusFlagsTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func VehicleStatusFlags__Array_to_Go(goSlice []VehicleStatusFlags, cSlice []CVehicleStatusFlags) {
	for i := 0; i < len(cSlice); i++ {
		VehicleStatusFlagsTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleStatusFlags__Array_to_C(cSlice []CVehicleStatusFlags, goSlice []VehicleStatusFlags) {
	for i := 0; i < len(goSlice); i++ {
		VehicleStatusFlagsTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
