/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
    http://www.apache.org/licenses/LICENSE-2.0
*/

package gogen

import (
	"strings"
	"text/template"
)

var templateFuncMap template.FuncMap = template.FuncMap{
	"lc":                          strings.ToLower,
	"camelToSnake":                camelToSnake,
	"defaultCode":                 defaultCode,
	"ucFirst":                     ucFirst,
	"srvNameFromSrvMsgName":       srvNameFromSrvMsgName,
	"actionNameFromActionMsgName": actionNameFromActionMsgName,
	"actionNameFromActionSrvName": actionNameFromActionSrvName,
	"cReturnCodeNameToGo":         cReturnCodeNameToGo,
	"cloneCode":                   cloneCode,
	"actionHasSuffix":             actionHasSuffix,
	"matchMsg":                    matchMsg,
	"sanitizeValue":               defaultValueSanitizer,
}

var ros2PackageCommonTemplate = template.Must(
	template.New("ros2PackageCommonTemplate").Funcs(templateFuncMap).Parse(
		`/* This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package {{ .GoPackage }}

/*
{{range $dir := .RootPaths -}}
#cgo LDFLAGS: "-L{{$dir}}/lib" "-Wl,-rpath={{$dir}}/lib"
{{end}}
#cgo LDFLAGS: -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -l{{.CPackage}}__rosidl_typesupport_c -l{{.CPackage}}__rosidl_generator_c
{{range $k, $v := .CImports -}}
#cgo LDFLAGS: -l{{$k}}__rosidl_typesupport_c -l{{$k}}__rosidl_generator_c
{{""}}
{{- end}}
{{range $dir := .RootPaths -}}
#cgo CFLAGS: "-I{{$dir}}/include"
{{end}}
*/
import "C"
`,
	),
)

var ros2MsgToGolangTypeTemplate = template.Must(
	template.New("ros2MsgToGolangTypeTemplate").Funcs(templateFuncMap).Parse(
		`/*{{ $Md := .Message }}
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package {{ $Md.GoPackage }}
import (
	"unsafe"

	"{{.Config.RclgoImportPath}}/pkg/rclgo"
	"{{.Config.RclgoImportPath}}/pkg/rclgo/types"
	"{{.Config.RclgoImportPath}}/pkg/rclgo/typemap"
	{{range $path, $name := $Md.GoImports -}}
	{{$name}} "{{$path}}"
	{{""}}{{- end}}
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>
{{if eq $Md.Type "msg"}}
#include <{{$Md.Package}}/msg/{{$Md.Name | camelToSnake}}.h>
{{else if eq $Md.Type "srv"}}
#include <{{$Md.Package}}/srv/{{$Md.Name | srvNameFromSrvMsgName | camelToSnake}}.h>
{{else if eq $Md.Type "action"}}
#include <{{$Md.Package}}/action/{{$Md.Name | actionNameFromActionMsgName | camelToSnake}}.h>
{{end}}
*/
import "C"

func init() {
	typemap.RegisterMessage("{{$Md.Package}}/{{$Md.Name}}", {{$Md.Name}}TypeSupport)
	typemap.RegisterMessage("{{$Md.Package}}/{{$Md.Type}}/{{$Md.Name}}", {{$Md.Name}}TypeSupport)
}

{{- if $Md.Constants }}
const (
{{- range $Md.Constants }}
	{{$Md.Name}}_{{.RosName}} {{.GoPkgReference}}{{.GoType}} = {{sanitizeValue .RosType .Value}}{{if .Comment -}} // {{.Comment}}{{- end}}
{{- end }}
)
{{- end }}

// Do not create instances of this type directly. Always use New{{$Md.Name}}
// function instead.
type {{$Md.Name}} struct {
	{{- range $k, $v := $Md.Fields }}
	{{$v.GoName }} {{$v.TypeArray}}{{$v.GoPkgReference}}{{$v.GoType}}` +
			"{{\"\"}} `yaml:\"{{$v.RosName}}\"`" + `{{if .Comment -}} // {{.Comment}}{{- end}}
	{{- end }}
}

// New{{$Md.Name}} creates a new {{$Md.Name}} with default values.
func New{{$Md.Name}}() *{{$Md.Name}} {
	self := {{$Md.Name}}{}
	self.SetDefaults()
	return &self
}

func (t *{{$Md.Name}}) Clone() *{{$Md.Name}} {
	c := &{{$Md.Name}}{}
	{{- range $f := $Md.Fields }}
	{{cloneCode $f}}
	{{- end }}
	return c
}

func (t *{{$Md.Name}}) CloneMsg() types.Message {
	return t.Clone()
}

func (t *{{$Md.Name}}) SetDefaults() {
	{{- range $k, $v := $Md.Fields }}
	{{defaultCode $v}}
	{{- end }}
}

func (t *{{$Md.Name}}) GetTypeSupport() types.MessageTypeSupport {
	return {{$Md.Name}}TypeSupport
}

{{- /* Some special cased methods to avoid cyclic dependency in actions */ -}}

{{- if actionHasSuffix $Md 
	"_SendGoal_Request"
	"_GetResult_Request"
	"_CancelGoal_Request"
	"_FeedbackMessage"
}}
func (t *{{$Md.Name}}) GetGoalID() *types.GoalID {
	return (*types.GoalID)(&t.GoalID.Uuid)
}

func (t *{{$Md.Name}}) SetGoalID(id *types.GoalID) {
	t.GoalID.Uuid = *id
}
{{- end -}}

{{- if actionHasSuffix $Md "_SendGoal_Request" }}
func (t *{{$Md.Name}}) GetGoalDescription() types.Message {
	return &t.Goal
}

func (t *{{$Md.Name}}) SetGoalDescription(desc types.Message) {
	t.Goal = *desc.(*{{$Md.Name | actionNameFromActionMsgName}}_Goal)
}
{{- end -}}

{{- if actionHasSuffix $Md "_SendGoal_Response" }}
func (t *{{$Md.Name}}) GetGoalAccepted() bool {
	return t.Accepted
}
{{- end -}}

{{- if matchMsg $Md "action_msgs_srv" "CancelGoal_Response" }}
func (t *{{$Md.Name}}) CallForEach(f func(interface{})) {
	for i := range t.GoalsCanceling {
		f((*types.GoalID)(&t.GoalsCanceling[i].GoalId.Uuid))
	}
}
{{- else if matchMsg $Md "action_msgs_msg" "GoalStatus" }}
func (t *{{$Md.Name}}) GetGoalID() *types.GoalID {
	return (*types.GoalID)(&t.GoalInfo.GoalId.Uuid)
}

func (t *{{$Md.Name}}) SetGoalID(id *types.GoalID) {
	t.GoalInfo.GoalId.Uuid = *id
}
{{- else if matchMsg $Md "action_msgs_msg" "GoalStatusArray" }}
func (t *{{$Md.Name}}) CallForEach(f func(interface{})) {
	for i := range t.StatusList {
		f(&t.StatusList[i])
	}
}
{{- end }}

// {{$Md.Name}}Publisher wraps rclgo.Publisher to provide type safe helper
// functions
type {{$Md.Name}}Publisher struct {
	*rclgo.Publisher
}

// New{{$Md.Name}}Publisher creates and returns a new publisher for the
// {{$Md.Name}}
func New{{$Md.Name}}Publisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*{{$Md.Name}}Publisher, error) {
	pub, err := node.NewPublisher(topic_name, {{$Md.Name}}TypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &{{$Md.Name}}Publisher{pub}, nil
}

func (p *{{$Md.Name}}Publisher) Publish(msg *{{$Md.Name}}) error {
	return p.Publisher.Publish(msg)
}

// {{$Md.Name}}Subscription wraps rclgo.Subscription to provide type safe helper
// functions
type {{$Md.Name}}Subscription struct {
	*rclgo.Subscription
}

// {{$Md.Name}}SubscriptionCallback type is used to provide a subscription
// handler function for a {{$Md.Name}}Subscription.
type {{$Md.Name}}SubscriptionCallback func(msg *{{$Md.Name}}, info *rclgo.RmwMessageInfo, err error)

// New{{$Md.Name}}Subscription creates and returns a new subscription for the
// {{$Md.Name}}
func New{{$Md.Name}}Subscription(node *rclgo.Node, topic_name string, subscriptionCallback {{$Md.Name}}SubscriptionCallback) (*{{$Md.Name}}Subscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg {{$Md.Name}}
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, {{$Md.Name}}TypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &{{$Md.Name}}Subscription{sub}, nil
}

func (s *{{$Md.Name}}Subscription) TakeMessage(out *{{$Md.Name}}) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// Clone{{$Md.Name}}Slice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func Clone{{$Md.Name}}Slice(dst, src []{{$Md.Name}}) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var {{$Md.Name}}TypeSupport types.MessageTypeSupport = _{{$Md.Name}}TypeSupport{}

type _{{$Md.Name}}TypeSupport struct{}

func (t _{{$Md.Name}}TypeSupport) New() types.Message {
	return New{{$Md.Name}}()
}

func (t _{{$Md.Name}}TypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.{{$Md.Package}}__{{$Md.Type}}__{{$Md.Name}}
	return (unsafe.Pointer)(C.{{$Md.Package}}__{{$Md.Type}}__{{$Md.Name}}__create())
}

func (t _{{$Md.Name}}TypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.{{$Md.Package}}__{{$Md.Type}}__{{$Md.Name}}__destroy((*C.{{$Md.Package}}__{{$Md.Type}}__{{$Md.Name}})(pointer_to_free))
}

func (t _{{$Md.Name}}TypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	{{ if $Md.Fields -}}
	m := msg.(*{{$Md.Name}})
	mem := (*C.{{$Md.Package}}__{{$Md.Type}}__{{$Md.Name}})(dst)
	{{- range $Md.Fields }}
	{{call $.cSerializationCode . $Md}}
	{{- end }}
	{{- end }}
}

func (t _{{$Md.Name}}TypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	{{if $Md.Fields -}}
	m := msg.(*{{$Md.Name}})
	mem := (*C.{{$Md.Package}}__{{$Md.Type}}__{{$Md.Name}})(ros2_message_buffer)
	{{- range $Md.Fields }}
	{{call $.goSerializationCode . $Md}}
	{{- end }}
	{{- end }}
}

func (t _{{$Md.Name}}TypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__{{$Md.Package}}__{{$Md.Type}}__{{$Md.Name}}())
}

type C{{$Md.Name}} = C.{{$Md.Package}}__{{$Md.Type}}__{{$Md.Name}}
type C{{$Md.Name}}__Sequence = C.{{$Md.Package}}__{{$Md.Type}}__{{$Md.Name}}__Sequence

func {{$Md.Name}}__Sequence_to_Go(goSlice *[]{{$Md.Name}}, cSlice C{{$Md.Name}}__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]{{$Md.Name}}, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		{{$Md.Name}}TypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func {{$Md.Name}}__Sequence_to_C(cSlice *C{{$Md.Name}}__Sequence, goSlice []{{$Md.Name}}) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.{{$Md.Package}}__{{$Md.Type}}__{{$Md.Name}})(C.malloc(C.sizeof_struct_{{$Md.Package}}__{{$Md.Type}}__{{$Md.Name}} * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		{{$Md.Name}}TypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func {{$Md.Name}}__Array_to_Go(goSlice []{{$Md.Name}}, cSlice []C{{$Md.Name}}) {
	for i := 0; i < len(cSlice); i++ {
		{{$Md.Name}}TypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func {{$Md.Name}}__Array_to_C(cSlice []C{{$Md.Name}}, goSlice []{{$Md.Name}}) {
	for i := 0; i < len(goSlice); i++ {
		{{$Md.Name}}TypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
`),
)

var ros2ServiceToGolangTypeTemplate = template.Must(
	template.New("ros2ServiceToGolangTypeTemplate").
		Funcs(templateFuncMap).
		Parse(
			`/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package {{ .Service.GoPackage }}

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
{{- if eq .Service.Type "action"}}
#include <{{.Service.Package}}/action/{{.Service.Name | actionNameFromActionSrvName | camelToSnake}}.h>
{{- else}}
#include <{{.Service.Package}}/srv/{{.Service.Name | camelToSnake}}.h>
{{- end}}
*/
import "C"

import (
	"context"
	"errors"
	"unsafe"

	"{{.Config.RclgoImportPath}}/pkg/rclgo"
	"{{.Config.RclgoImportPath}}/pkg/rclgo/typemap"
	"{{.Config.RclgoImportPath}}/pkg/rclgo/types"
)

func init() {
	typemap.RegisterService("{{.Service.Package}}/{{.Service.Name}}", {{ .Service.Name }}TypeSupport)
	typemap.RegisterService("{{.Service.Package}}/{{.Service.Type}}/{{.Service.Name}}", {{ .Service.Name }}TypeSupport)
}

type _{{.Service.Name}}TypeSupport struct {}

func (s _{{.Service.Name}}TypeSupport) Request() types.MessageTypeSupport {
	return {{.Service.Request.Name}}TypeSupport
}

func (s _{{.Service.Name}}TypeSupport) Response() types.MessageTypeSupport {
	return {{.Service.Response.Name}}TypeSupport
}

func (s _{{.Service.Name}}TypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__{{.Service.Package}}__{{.Service.Type}}__{{.Service.Name}}())
}

// Modifying this variable is undefined behavior.
var {{ .Service.Name }}TypeSupport types.ServiceTypeSupport = _{{.Service.Name}}TypeSupport{}

// {{.Service.Name}}Client wraps rclgo.Client to provide type safe helper
// functions
type {{.Service.Name}}Client struct {
	*rclgo.Client
}

// New{{.Service.Name}}Client creates and returns a new client for the
// {{.Service.Name}}
func New{{.Service.Name}}Client(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*{{.Service.Name}}Client, error) {
	client, err := node.NewClient(serviceName, {{.Service.Name}}TypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &{{.Service.Name}}Client{client}, nil
}

func (s *{{.Service.Name}}Client) Send(ctx context.Context, req *{{.Service.Request.Name}}) (*{{.Service.Response.Name}}, *rclgo.RmwServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*{{.Service.Response.Name}})
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type {{.Service.Name}}ServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s {{.Service.Name}}ServiceResponseSender) SendResponse(resp *{{.Service.Response.Name}}) error {
	return s.sender.SendResponse(resp)
}

type {{.Service.Name}}ServiceRequestHandler func(*rclgo.RmwServiceInfo, *{{.Service.Request.Name}}, {{.Service.Name}}ServiceResponseSender)

// {{.Service.Name}}Service wraps rclgo.Service to provide type safe helper
// functions
type {{.Service.Name}}Service struct {
	*rclgo.Service
}

// New{{.Service.Name}}Service creates and returns a new service for the
// {{.Service.Name}}
func New{{.Service.Name}}Service(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler {{.Service.Name}}ServiceRequestHandler) (*{{.Service.Name}}Service, error) {
	h := func(rmw *rclgo.RmwServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*{{.Service.Request.Name}})
		responseSender := {{.Service.Name}}ServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, {{.Service.Name}}TypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &{{.Service.Name}}Service{service}, nil
}`),
)

var ros2ActionToGolangTypeTemplate = template.Must(
	template.New("ros2ServiceToGolangTypeTemplate").
		Funcs(templateFuncMap).
		Parse(
			`/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package {{ .Action.GoPackage }}

/*
#cgo LDFLAGS: -L/opt/ros/${ROS_DISTRO}/lib -Wl,-rpath=/opt/ros/${ROS_DISTRO}/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -l{{.Action.Package}}__rosidl_typesupport_c -l{{.Action.Package}}__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/${ROS_DISTRO}/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <{{.Action.Package}}/action/{{.Action.Name | actionNameFromActionSrvName | camelToSnake}}.h>
*/
import "C"

import (
	"context"
	"time"
	"unsafe"

	"{{.Config.RclgoImportPath}}/pkg/rclgo"
	"{{.Config.RclgoImportPath}}/pkg/rclgo/typemap"
	"{{.Config.RclgoImportPath}}/pkg/rclgo/types"

	action_msgs_msg "{{.Config.MessageModulePrefix}}/action_msgs/msg"
	action_msgs_srv "{{.Config.MessageModulePrefix}}/action_msgs/srv"
)

func init() {
	typemap.RegisterAction("{{.Action.Package}}/{{.Action.Name}}", {{ .Action.Name }}TypeSupport)
	typemap.RegisterAction("{{.Action.Package}}/{{.Action.Type}}/{{.Action.Name}}", {{ .Action.Name }}TypeSupport)
}

type _{{.Action.Name}}TypeSupport struct {}

func (s _{{.Action.Name}}TypeSupport) Goal() types.MessageTypeSupport {
	return {{.Action.Goal.Name}}TypeSupport
}

func (s _{{.Action.Name}}TypeSupport) SendGoal() types.ServiceTypeSupport {
	return {{.Action.SendGoal.Name}}TypeSupport
}

func (s _{{.Action.Name}}TypeSupport) NewSendGoalResponse(accepted bool, stamp time.Duration) types.Message {
	msg := New{{.Action.Name}}_SendGoal_Response()
	msg.Accepted = accepted
	secs := stamp.Truncate(time.Second)
	msg.Stamp.Sec = int32(secs)
	msg.Stamp.Nanosec = uint32(stamp - secs)
	return msg
}

func (s _{{.Action.Name}}TypeSupport) Result() types.MessageTypeSupport {
	return {{.Action.Result.Name}}TypeSupport
}

func (s _{{.Action.Name}}TypeSupport) GetResult() types.ServiceTypeSupport {
	return {{.Action.GetResult.Name}}TypeSupport
}

func (s _{{.Action.Name}}TypeSupport) NewGetResultResponse(status int8, result types.Message) types.Message {
	msg := New{{.Action.Name}}_GetResult_Response()
	msg.Status = status
	if result == nil {
		msg.Result = *New{{.Action.Name}}_Result()
	} else {
		msg.Result = *result.(*{{.Action.Name}}_Result)
	}
	return msg
}

func (s _{{.Action.Name}}TypeSupport) CancelGoal() types.ServiceTypeSupport {
	return action_msgs_srv.CancelGoalTypeSupport
}

func (s _{{.Action.Name}}TypeSupport) Feedback() types.MessageTypeSupport {
	return {{.Action.Feedback.Name}}TypeSupport
}

func (s _{{.Action.Name}}TypeSupport) FeedbackMessage() types.MessageTypeSupport {
	return {{.Action.FeedbackMessage.Name}}TypeSupport
}

func (s _{{.Action.Name}}TypeSupport) NewFeedbackMessage(goalID *types.GoalID, feedback types.Message) types.Message {
	msg := New{{.Action.Name}}_FeedbackMessage()
	msg.GoalID.Uuid = *goalID
	msg.Feedback = *feedback.(*{{.Action.Name}}_Feedback)
	return msg
}

func (s _{{.Action.Name}}TypeSupport) GoalStatusArray() types.MessageTypeSupport {
	return action_msgs_msg.GoalStatusArrayTypeSupport
}

func (s _{{.Action.Name}}TypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_action_type_support_handle__{{.Action.Package}}__{{.Action.Type}}__{{.Action.Name}}())
}

// Modifying this variable is undefined behavior.
var {{.Action.Name}}TypeSupport types.ActionTypeSupport = _{{.Action.Name}}TypeSupport{}

type {{.Action.Name}}FeedbackSender struct {
	sender rclgo.FeedbackSender
}

func (s *{{.Action.Name}}FeedbackSender) Send(msg *{{.Action.Name}}_Feedback) error {
	return s.sender.Send(msg)
}

type {{.Action.Name}}GoalHandle struct{
	*rclgo.GoalHandle

	Description *{{.Action.Name}}_Goal
}

func (g *{{.Action.Name}}GoalHandle) Accept() (*{{.Action.Name}}FeedbackSender, error) {
	s, err := g.GoalHandle.Accept()
	if err != nil {
		return nil, err
	}
	return &{{.Action.Name}}FeedbackSender{*s}, nil
}

type {{.Action.Name}}Action interface {
	ExecuteGoal(context.Context, *{{.Action.Name}}GoalHandle) (*{{.Action.Name}}_Result, error)
}

func New{{.Action.Name}}Action(
	executeGoal func(context.Context, *{{.Action.Name}}GoalHandle) (*{{.Action.Name}}_Result, error),
) {{.Action.Name}}Action {
	return _{{.Action.Name}}FuncAction(executeGoal)
}

type _{{.Action.Name}}FuncAction func(context.Context, *{{.Action.Name}}GoalHandle) (*{{.Action.Name}}_Result, error)

func (a _{{.Action.Name}}FuncAction) ExecuteGoal(
	ctx context.Context, goal *{{.Action.Name}}GoalHandle,
) (*{{.Action.Name}}_Result, error) {
	return a(ctx, goal)
}

type _{{.Action.Name}}Action struct {
	action {{.Action.Name}}Action
}

func (a _{{.Action.Name}}Action) ExecuteGoal(ctx context.Context, handle *rclgo.GoalHandle) (types.Message, error) {
	return a.action.ExecuteGoal(ctx, &{{.Action.Name}}GoalHandle{
		GoalHandle:  handle,
		Description: handle.Description.(*{{.Action.Name}}_Goal),
	})
}

func (a _{{.Action.Name}}Action) TypeSupport() types.ActionTypeSupport {
	return {{.Action.Name}}TypeSupport
}

type {{.Action.Name}}Server struct{
	*rclgo.ActionServer
}

func New{{.Action.Name}}Server(node *rclgo.Node, name string, action {{.Action.Name}}Action, opts *rclgo.ActionServerOptions) (*{{.Action.Name}}Server, error) {
	server, err := node.NewActionServer(name, _{{.Action.Name}}Action{action}, opts)
	if err != nil {
		return nil, err
	}
	return &{{.Action.Name}}Server{server}, nil
}

type {{.Action.Name}}FeedbackHandler func(context.Context, *{{.Action.Name}}_FeedbackMessage)

type {{.Action.Name}}StatusHandler func(context.Context, *action_msgs_msg.GoalStatus)

type {{.Action.Name}}Client struct{
	*rclgo.ActionClient
}

func New{{.Action.Name}}Client(node *rclgo.Node, name string, opts *rclgo.ActionClientOptions) (*{{.Action.Name}}Client, error) {
	client, err := node.NewActionClient(name, {{.Action.Name}}TypeSupport, opts)
	if err != nil {
		return nil, err
	}
	return &{{.Action.Name}}Client{client}, nil
}

func (c *{{.Action.Name}}Client) WatchGoal(ctx context.Context, goal *{{.Action.Name}}_Goal, onFeedback {{.Action.Name}}FeedbackHandler) (*{{.Action.Name}}_GetResult_Response, error) {
	var resp types.Message
	var err error
	if onFeedback == nil {
		resp, err = c.ActionClient.WatchGoal(ctx, goal, nil)
	} else {
		resp, err = c.ActionClient.WatchGoal(ctx, goal, func(ctx context.Context, msg types.Message) {
			onFeedback(ctx, msg.(*{{.Action.Name}}_FeedbackMessage))
		})
	}
	if r, ok := resp.(*{{.Action.Name}}_GetResult_Response); ok {
		return r, err
	}
	return nil, err
}

func (c *{{.Action.Name}}Client) SendGoal(ctx context.Context, goal *{{.Action.Name}}_Goal) (*{{.Action.Name}}_SendGoal_Response, *types.GoalID, error) {
	resp, id, err := c.ActionClient.SendGoal(ctx, goal)
	if r, ok := resp.(*{{.Action.Name}}_SendGoal_Response); ok {
		return r, id, err
	}
	return nil, id, err
}

func (c *{{.Action.Name}}Client) SendGoalRequest(ctx context.Context, request *{{.Action.Name}}_SendGoal_Request) (*{{.Action.Name}}_SendGoal_Response, error) {
	resp, err := c.ActionClient.SendGoalRequest(ctx, request)
	if r, ok := resp.(*{{.Action.Name}}_SendGoal_Response); ok {
		return r, err
	}
	return nil, err
}

func (c *{{.Action.Name}}Client) GetResult(ctx context.Context, goalID *types.GoalID) (*{{.Action.Name}}_GetResult_Response, error) {
	resp, err := c.ActionClient.GetResult(ctx, goalID)
	if r, ok := resp.(*{{.Action.Name}}_GetResult_Response); ok {
		return r, err
	}
	return nil, err
}

func (c *{{.Action.Name}}Client) CancelGoal(ctx context.Context, request *action_msgs_srv.CancelGoal_Request) (*action_msgs_srv.CancelGoal_Response, error) {
	resp, err := c.ActionClient.CancelGoal(ctx, request)
	if r, ok := resp.(*action_msgs_srv.CancelGoal_Response); ok {
		return r, err
	}
	return nil, err
}

func (c *{{.Action.Name}}Client) WatchFeedback(ctx context.Context, goalID *types.GoalID, handler {{.Action.Name}}FeedbackHandler) <-chan error {
	return c.ActionClient.WatchFeedback(ctx, goalID, func(ctx context.Context, msg types.Message) {
		handler(ctx, msg.(*{{.Action.Name}}_FeedbackMessage))
	})
}

func (c *{{.Action.Name}}Client) WatchStatus(ctx context.Context, goalID *types.GoalID, handler {{.Action.Name}}StatusHandler) <-chan error {
	return c.ActionClient.WatchStatus(ctx, goalID, func(ctx context.Context, msg types.Message) {
		handler(ctx, msg.(*action_msgs_msg.GoalStatus))
	})
}
`))

var primitiveTypes = template.Must(
	template.New("primitiveTypes").Funcs(templateFuncMap).Parse(
		`/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package primitives

/*
#cgo LDFLAGS: -L/opt/ros/${ROS_DISTRO}/lib -Wl,-rpath=/opt/ros/${ROS_DISTRO}/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo CFLAGS: -I/opt/ros/${ROS_DISTRO}/include

#include "rosidl_runtime_c/string.h"
#include "rosidl_runtime_c/primitives_sequence.h"

*/
import "C"
import (
	"unsafe"
){{range $k, $v := .PMap -}}{{if .SkipAutogen}}{{- else -}}
{{""}}
{{""}}
// {{.RosType | ucFirst}}
type C{{.RosType | ucFirst}} = C.{{.CType}}
type C{{.RosType | ucFirst}}__Sequence = C.rosidl_runtime_c__{{.CStructName}}__Sequence

func {{.RosType | ucFirst}}__Sequence_to_Go(goSlice *[]{{.GoType}}, cSlice C{{.RosType | ucFirst}}__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]{{.GoType}}, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		(*goSlice)[i] = {{.GoType}}(src[i])
	}
}
func {{.RosType | ucFirst}}__Sequence_to_C(cSlice *C{{.RosType | ucFirst}}__Sequence, goSlice []{{.GoType}}) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.{{.CType}})(C.malloc(C.sizeof_{{.CType}} * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		dst[i] = C.{{.CType}}(goSlice[i])
	}
}
func {{.RosType | ucFirst}}__Array_to_Go(goSlice []{{.GoType}}, cSlice []C{{.RosType | ucFirst}}) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i] = {{.GoType}}(cSlice[i])
	}
}
func {{.RosType | ucFirst}}__Array_to_C(cSlice []C{{.RosType | ucFirst}}, goSlice []{{.GoType}}) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = C.{{.CType}}(goSlice[i])
	}
}
{{- end}}{{- end}}
`),
)

var ros2MsgImportAllPackage = template.Must(
	template.New("ros2MsgToGolangTypeTemplate").Funcs(templateFuncMap).Parse(
		`/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package msgs

import (
	{{- range $import, $unused := .Packages }}
	_ "{{$.Config.MessageModulePrefix}}/{{$import}}" //
	{{- end }}
)
`),
)

var ros2ErrorCodes = template.Must(
	template.New("ros2ErrorCodes").Funcs(templateFuncMap).Parse(
		`/*{{ $P := . }}
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package rclgo

/*
{{- range $file := .includes }}
#include <{{$file}}>
{{- end }}
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
	{{range $e := .errorTypes -}}{{if $e.Rcl_ret_t -}}{{if not (index $P.dedupFilter $e.Name) -}}
	case C.{{$e.Name}}:
		return &{{$e.Name|cReturnCodeNameToGo}}{rclError: rclError{rclRetCode: {{$e.Rcl_ret_t}}, trace: string(stackTraceBuffer), context: errorsBuildContext(&{{$e.Name|cReturnCodeNameToGo}}{}, context, string(stackTraceBuffer))}}
	{{""}}
	{{- end}}{{- end}}{{- end}}
	default:
		return &UnknownReturnCode{rclError: rclError{rclRetCode: int(rcl_ret_t), context: context}}
	}
}

type UnknownReturnCode struct {
	rclError
}

{{range $e := .errorTypes -}}{{if $e.Rcl_ret_t}}
// {{$e.Name|cReturnCodeNameToGo}} {{$e.Comment}}
type {{$e.Name|cReturnCodeNameToGo}} struct {
	rclError
}
{{""}}
{{- end}}{{- end}}

{{range $e := .errorTypes -}}{{if $e.Reference}}
// {{$e.Name|cReturnCodeNameToGo}} {{$e.Comment}}
type {{$e.Name|cReturnCodeNameToGo}} = {{$e.Reference|cReturnCodeNameToGo}}
{{""}}
{{- end}}{{- end}}

`),
)
