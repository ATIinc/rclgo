package rclgo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	action_msgs "github.com/ATIinc/rclgo/internal/msgs/action_msgs/msg"
	example_interfaces_action "github.com/ATIinc/rclgo/internal/msgs/example_interfaces/action"
	std_msgs "github.com/ATIinc/rclgo/internal/msgs/std_msgs/msg"
	std_srvs "github.com/ATIinc/rclgo/internal/msgs/std_srvs/srv"
	"github.com/ATIinc/rclgo/pkg/rclgo"
	"github.com/ATIinc/rclgo/pkg/rclgo/types"
)

// A dynamically loaded type support must be the same handle the generated bindings link against.
func TestLoadDynamicMessageTypeSupportByName(t *testing.T) {
	tests := map[string]types.MessageTypeSupport{
		"std_msgs/msg/String":                                 std_msgs.StringTypeSupport,
		"action_msgs/msg/GoalStatusArray":                     action_msgs.GoalStatusArrayTypeSupport,
		"std_srvs/srv/Trigger_Event":                          std_srvs.Trigger_EventTypeSupport,
		"example_interfaces/action/Fibonacci_FeedbackMessage": example_interfaces_action.Fibonacci_FeedbackMessageTypeSupport,
		"example_interfaces/action/Fibonacci_SendGoal_Event":  example_interfaces_action.Fibonacci_SendGoal_EventTypeSupport,
	}
	for name, generated := range tests {
		t.Run(name, func(t *testing.T) {
			ts, err := rclgo.LoadDynamicMessageTypeSupportByName(name)
			require.NoError(t, err)
			assert.Equal(t, generated.TypeSupport(), ts.TypeSupport())
		})
	}
}

func TestLoadDynamicMessageTypeSupport(t *testing.T) {
	ts, err := rclgo.LoadDynamicMessageTypeSupport("std_msgs", "String")
	require.NoError(t, err)
	assert.Equal(t, std_msgs.StringTypeSupport.TypeSupport(), ts.TypeSupport())
}

func TestLoadDynamicMessageTypeSupportByNameErrors(t *testing.T) {
	tests := map[string]string{
		"":                           "invalid message type name",
		"std_msgs/String":            "invalid message type name",
		"std_msgs/idl/String":        "invalid message type name",
		"/msg/String":                "invalid message type name",
		"std_msgs/msg/":              "invalid message type name",
		"std_msgs/msg/String/extra":  "invalid message type name",
		"no_such_pkg/msg/Thing":      "libno_such_pkg__rosidl_typesupport_c.so",
		"std_msgs/msg/NoSuchType":    "get_message_type_support_handle__std_msgs__msg__NoSuchType",
		"std_srvs/msg/Trigger_Event": "get_message_type_support_handle__std_srvs__msg__Trigger_Event",
	}
	for name, want := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := rclgo.LoadDynamicMessageTypeSupportByName(name)
			require.Error(t, err)
			assert.Contains(t, err.Error(), want)
		})
	}
}
