# TODO

Follow-up work for service introspection (see PR #13).

## 1. Generate service event (`_Event`) message bindings for Go

The introspection publisher emits `<package>/srv/<Service>_Event` messages
(e.g. `example_interfaces/srv/AddTwoInts_Event`) on the hidden
`<service>/_service_event` topic. `rclgo-gen` currently generates only the
`_Request`/`_Response` bindings, and `service_msgs` (`ServiceEventInfo`) is not
generated at all.

- Teach `rclgo-gen` to emit the `_Event` message type for each service.
- Generate the `service_msgs` package (`ServiceEventInfo`, etc.) that the event
  message depends on.
- Regenerate `internal/msgs` and confirm the new bindings build.

## 2. Add a payload-level introspection unit test

Depends on task 1. Once the `_Event` bindings exist, add a test that proves
events actually flow end-to-end:

- Enable `ServiceIntrospectionContents` on a service and client.
- Subscribe to `<service>/_service_event` using the generated `_Event` type.
- Make a request/response round-trip and assert the received event messages
  contain the expected metadata and request/response contents.

This is the deterministic behavioral test that the current graph-discovery
approach can't provide reliably. See PR #13 discussion for context.

## 3. Add action introspection support

`rcl` exposes introspection configuration for actions as well
(`rcl_action_*` introspection hooks, built on the underlying service
introspection). Extend the `rclgo` action client/server API the same way the
service `Service`/`Client` support was added:

- Map the relevant `rcl_action` introspection functions through CGO.
- Add `ConfigureIntrospection`/`IntrospectionState` (or equivalent options) to
  the action client and server types.
- Mirror the concurrency, lifecycle, and allocator handling used by the
  service/client implementation.
