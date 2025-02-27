// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/agent.proto

package admin

import (
	fmt "fmt"
	core "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The state of the execution is used to control its visibility in the UI/CLI.
type State int32

const (
	State_RETRYABLE_FAILURE State = 0
	State_PERMANENT_FAILURE State = 1
	State_PENDING           State = 2
	State_RUNNING           State = 3
	State_SUCCEEDED         State = 4
)

var State_name = map[int32]string{
	0: "RETRYABLE_FAILURE",
	1: "PERMANENT_FAILURE",
	2: "PENDING",
	3: "RUNNING",
	4: "SUCCEEDED",
}

var State_value = map[string]int32{
	"RETRYABLE_FAILURE": 0,
	"PERMANENT_FAILURE": 1,
	"PENDING":           2,
	"RUNNING":           3,
	"SUCCEEDED":         4,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{0}
}

// Represents a subset of runtime task execution metadata that are relevant to external plugins.
type TaskExecutionMetadata struct {
	// ID of the task execution
	TaskExecutionId *core.TaskExecutionIdentifier `protobuf:"bytes,1,opt,name=task_execution_id,json=taskExecutionId,proto3" json:"task_execution_id,omitempty"`
	// k8s namespace where the task is executed in
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Labels attached to the task execution
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Annotations attached to the task execution
	Annotations map[string]string `protobuf:"bytes,4,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// k8s service account associated with the task execution
	K8SServiceAccount string `protobuf:"bytes,5,opt,name=k8s_service_account,json=k8sServiceAccount,proto3" json:"k8s_service_account,omitempty"`
	// Environment variables attached to the task execution
	EnvironmentVariables map[string]string `protobuf:"bytes,6,rep,name=environment_variables,json=environmentVariables,proto3" json:"environment_variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TaskExecutionMetadata) Reset()         { *m = TaskExecutionMetadata{} }
func (m *TaskExecutionMetadata) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionMetadata) ProtoMessage()    {}
func (*TaskExecutionMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{0}
}

func (m *TaskExecutionMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionMetadata.Unmarshal(m, b)
}
func (m *TaskExecutionMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionMetadata.Marshal(b, m, deterministic)
}
func (m *TaskExecutionMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionMetadata.Merge(m, src)
}
func (m *TaskExecutionMetadata) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionMetadata.Size(m)
}
func (m *TaskExecutionMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionMetadata proto.InternalMessageInfo

func (m *TaskExecutionMetadata) GetTaskExecutionId() *core.TaskExecutionIdentifier {
	if m != nil {
		return m.TaskExecutionId
	}
	return nil
}

func (m *TaskExecutionMetadata) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *TaskExecutionMetadata) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *TaskExecutionMetadata) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *TaskExecutionMetadata) GetK8SServiceAccount() string {
	if m != nil {
		return m.K8SServiceAccount
	}
	return ""
}

func (m *TaskExecutionMetadata) GetEnvironmentVariables() map[string]string {
	if m != nil {
		return m.EnvironmentVariables
	}
	return nil
}

// Represents a request structure to create task.
type CreateTaskRequest struct {
	// The inputs required to start the execution. All required inputs must be
	// included in this map. If not required and not provided, defaults apply.
	// +optional
	Inputs *core.LiteralMap `protobuf:"bytes,1,opt,name=inputs,proto3" json:"inputs,omitempty"`
	// Template of the task that encapsulates all the metadata of the task.
	Template *core.TaskTemplate `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	// Prefix for where task output data will be written. (e.g. s3://my-bucket/randomstring)
	OutputPrefix string `protobuf:"bytes,3,opt,name=output_prefix,json=outputPrefix,proto3" json:"output_prefix,omitempty"`
	// subset of runtime task execution metadata.
	TaskExecutionMetadata *TaskExecutionMetadata `protobuf:"bytes,4,opt,name=task_execution_metadata,json=taskExecutionMetadata,proto3" json:"task_execution_metadata,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *CreateTaskRequest) Reset()         { *m = CreateTaskRequest{} }
func (m *CreateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTaskRequest) ProtoMessage()    {}
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{1}
}

func (m *CreateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskRequest.Unmarshal(m, b)
}
func (m *CreateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskRequest.Marshal(b, m, deterministic)
}
func (m *CreateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskRequest.Merge(m, src)
}
func (m *CreateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTaskRequest.Size(m)
}
func (m *CreateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskRequest proto.InternalMessageInfo

func (m *CreateTaskRequest) GetInputs() *core.LiteralMap {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *CreateTaskRequest) GetTemplate() *core.TaskTemplate {
	if m != nil {
		return m.Template
	}
	return nil
}

func (m *CreateTaskRequest) GetOutputPrefix() string {
	if m != nil {
		return m.OutputPrefix
	}
	return ""
}

func (m *CreateTaskRequest) GetTaskExecutionMetadata() *TaskExecutionMetadata {
	if m != nil {
		return m.TaskExecutionMetadata
	}
	return nil
}

// Represents a create response structure.
type CreateTaskResponse struct {
	// Metadata is created by the agent. It could be a string (jobId) or a dict (more complex metadata).
	ResourceMeta         []byte   `protobuf:"bytes,1,opt,name=resource_meta,json=resourceMeta,proto3" json:"resource_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTaskResponse) Reset()         { *m = CreateTaskResponse{} }
func (m *CreateTaskResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTaskResponse) ProtoMessage()    {}
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{2}
}

func (m *CreateTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskResponse.Unmarshal(m, b)
}
func (m *CreateTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskResponse.Marshal(b, m, deterministic)
}
func (m *CreateTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskResponse.Merge(m, src)
}
func (m *CreateTaskResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTaskResponse.Size(m)
}
func (m *CreateTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskResponse proto.InternalMessageInfo

func (m *CreateTaskResponse) GetResourceMeta() []byte {
	if m != nil {
		return m.ResourceMeta
	}
	return nil
}

// A message used to fetch a job resource from flyte agent server.
type GetTaskRequest struct {
	// A predefined yet extensible Task type identifier.
	TaskType string `protobuf:"bytes,1,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	// Metadata about the resource to be pass to the agent.
	ResourceMeta         []byte   `protobuf:"bytes,2,opt,name=resource_meta,json=resourceMeta,proto3" json:"resource_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTaskRequest) Reset()         { *m = GetTaskRequest{} }
func (m *GetTaskRequest) String() string { return proto.CompactTextString(m) }
func (*GetTaskRequest) ProtoMessage()    {}
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{3}
}

func (m *GetTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTaskRequest.Unmarshal(m, b)
}
func (m *GetTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTaskRequest.Marshal(b, m, deterministic)
}
func (m *GetTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTaskRequest.Merge(m, src)
}
func (m *GetTaskRequest) XXX_Size() int {
	return xxx_messageInfo_GetTaskRequest.Size(m)
}
func (m *GetTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTaskRequest proto.InternalMessageInfo

func (m *GetTaskRequest) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *GetTaskRequest) GetResourceMeta() []byte {
	if m != nil {
		return m.ResourceMeta
	}
	return nil
}

// Response to get an individual task resource.
type GetTaskResponse struct {
	Resource *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// log information for the task execution.
	LogLinks             []*core.TaskLog `protobuf:"bytes,2,rep,name=log_links,json=logLinks,proto3" json:"log_links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetTaskResponse) Reset()         { *m = GetTaskResponse{} }
func (m *GetTaskResponse) String() string { return proto.CompactTextString(m) }
func (*GetTaskResponse) ProtoMessage()    {}
func (*GetTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{4}
}

func (m *GetTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTaskResponse.Unmarshal(m, b)
}
func (m *GetTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTaskResponse.Marshal(b, m, deterministic)
}
func (m *GetTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTaskResponse.Merge(m, src)
}
func (m *GetTaskResponse) XXX_Size() int {
	return xxx_messageInfo_GetTaskResponse.Size(m)
}
func (m *GetTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTaskResponse proto.InternalMessageInfo

func (m *GetTaskResponse) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *GetTaskResponse) GetLogLinks() []*core.TaskLog {
	if m != nil {
		return m.LogLinks
	}
	return nil
}

type Resource struct {
	// The state of the execution is used to control its visibility in the UI/CLI.
	State State `protobuf:"varint,1,opt,name=state,proto3,enum=flyteidl.admin.State" json:"state,omitempty"`
	// The outputs of the execution. It's typically used by sql task. Agent service will create a
	// Structured dataset pointing to the query result table.
	// +optional
	Outputs *core.LiteralMap `protobuf:"bytes,2,opt,name=outputs,proto3" json:"outputs,omitempty"`
	// A descriptive message for the current state. e.g. waiting for cluster.
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{5}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetState() State {
	if m != nil {
		return m.State
	}
	return State_RETRYABLE_FAILURE
}

func (m *Resource) GetOutputs() *core.LiteralMap {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *Resource) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// A message used to delete a task.
type DeleteTaskRequest struct {
	// A predefined yet extensible Task type identifier.
	TaskType string `protobuf:"bytes,1,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	// Metadata about the resource to be pass to the agent.
	ResourceMeta         []byte   `protobuf:"bytes,2,opt,name=resource_meta,json=resourceMeta,proto3" json:"resource_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTaskRequest) Reset()         { *m = DeleteTaskRequest{} }
func (m *DeleteTaskRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTaskRequest) ProtoMessage()    {}
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{6}
}

func (m *DeleteTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTaskRequest.Unmarshal(m, b)
}
func (m *DeleteTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTaskRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTaskRequest.Merge(m, src)
}
func (m *DeleteTaskRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTaskRequest.Size(m)
}
func (m *DeleteTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTaskRequest proto.InternalMessageInfo

func (m *DeleteTaskRequest) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *DeleteTaskRequest) GetResourceMeta() []byte {
	if m != nil {
		return m.ResourceMeta
	}
	return nil
}

// Response to delete a task.
type DeleteTaskResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTaskResponse) Reset()         { *m = DeleteTaskResponse{} }
func (m *DeleteTaskResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTaskResponse) ProtoMessage()    {}
func (*DeleteTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{7}
}

func (m *DeleteTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTaskResponse.Unmarshal(m, b)
}
func (m *DeleteTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTaskResponse.Marshal(b, m, deterministic)
}
func (m *DeleteTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTaskResponse.Merge(m, src)
}
func (m *DeleteTaskResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTaskResponse.Size(m)
}
func (m *DeleteTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTaskResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("flyteidl.admin.State", State_name, State_value)
	proto.RegisterType((*TaskExecutionMetadata)(nil), "flyteidl.admin.TaskExecutionMetadata")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.admin.TaskExecutionMetadata.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.admin.TaskExecutionMetadata.EnvironmentVariablesEntry")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.admin.TaskExecutionMetadata.LabelsEntry")
	proto.RegisterType((*CreateTaskRequest)(nil), "flyteidl.admin.CreateTaskRequest")
	proto.RegisterType((*CreateTaskResponse)(nil), "flyteidl.admin.CreateTaskResponse")
	proto.RegisterType((*GetTaskRequest)(nil), "flyteidl.admin.GetTaskRequest")
	proto.RegisterType((*GetTaskResponse)(nil), "flyteidl.admin.GetTaskResponse")
	proto.RegisterType((*Resource)(nil), "flyteidl.admin.Resource")
	proto.RegisterType((*DeleteTaskRequest)(nil), "flyteidl.admin.DeleteTaskRequest")
	proto.RegisterType((*DeleteTaskResponse)(nil), "flyteidl.admin.DeleteTaskResponse")
}

func init() { proto.RegisterFile("flyteidl/admin/agent.proto", fileDescriptor_c434e52bb0028071) }

var fileDescriptor_c434e52bb0028071 = []byte{
	// 773 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x7f, 0x6f, 0xe3, 0x44,
	0x10, 0x25, 0x49, 0xd3, 0x26, 0x93, 0x5e, 0x2f, 0x59, 0x1a, 0x70, 0x73, 0x07, 0xaa, 0x82, 0x40,
	0x15, 0x08, 0x47, 0xd7, 0x22, 0xe8, 0x81, 0x00, 0xe5, 0x5a, 0x53, 0x45, 0x4a, 0xa3, 0x6a, 0x9b,
	0x20, 0x40, 0x02, 0x6b, 0xe3, 0x4c, 0x8c, 0x15, 0x67, 0xd7, 0x78, 0xd7, 0xd1, 0x45, 0xe2, 0x3f,
	0x24, 0x3e, 0x03, 0x1f, 0x17, 0x79, 0xfd, 0xa3, 0x49, 0x64, 0x50, 0x4f, 0xf7, 0x9f, 0x77, 0xde,
	0x9b, 0x37, 0x6f, 0x67, 0xc6, 0x5a, 0xe8, 0xcc, 0xfd, 0xb5, 0x42, 0x6f, 0xe6, 0xf7, 0xd8, 0x6c,
	0xe9, 0xf1, 0x1e, 0x73, 0x91, 0x2b, 0x33, 0x08, 0x85, 0x12, 0xe4, 0x28, 0xc3, 0x4c, 0x8d, 0x75,
	0x9e, 0xe7, 0x5c, 0x47, 0x84, 0xd8, 0xf3, 0x3d, 0x85, 0x21, 0xf3, 0x65, 0xc2, 0xee, 0x9c, 0x6c,
	0xa3, 0x8a, 0xc9, 0x45, 0x06, 0x7d, 0xb0, 0x0d, 0x79, 0x5c, 0x61, 0x38, 0x67, 0x0e, 0xa6, 0xf0,
	0x87, 0x3b, 0xf0, 0x0c, 0xb9, 0xf2, 0xe6, 0x1e, 0x86, 0xc5, 0xe9, 0xf8, 0x1a, 0x9d, 0x48, 0x79,
	0x82, 0x27, 0x70, 0xf7, 0x9f, 0x2a, 0xb4, 0xc7, 0x4c, 0x2e, 0xac, 0x2c, 0x7e, 0x8b, 0x8a, 0xcd,
	0x98, 0x62, 0x84, 0x42, 0x2b, 0xb6, 0x61, 0xe7, 0x19, 0xb6, 0x37, 0x33, 0x4a, 0xa7, 0xa5, 0xb3,
	0xc6, 0xf9, 0x27, 0x66, 0x7e, 0xb9, 0x58, 0xd4, 0xdc, 0x12, 0x18, 0xe4, 0x0e, 0xe8, 0x53, 0xb5,
	0x0d, 0x90, 0xe7, 0x50, 0xe7, 0x6c, 0x89, 0x32, 0x60, 0x0e, 0x1a, 0xe5, 0xd3, 0xd2, 0x59, 0x9d,
	0x3e, 0x04, 0xc8, 0x00, 0xf6, 0x7d, 0x36, 0x45, 0x5f, 0x1a, 0x95, 0xd3, 0xca, 0x59, 0xe3, 0xfc,
	0x85, 0xb9, 0xdd, 0x43, 0xb3, 0xd0, 0xa8, 0x39, 0xd4, 0x39, 0x16, 0x57, 0xe1, 0x9a, 0xa6, 0x02,
	0xe4, 0x27, 0x68, 0x30, 0xce, 0x85, 0x62, 0x31, 0x53, 0x1a, 0x7b, 0x5a, 0xef, 0xcb, 0xc7, 0xe9,
	0xf5, 0x1f, 0x12, 0x13, 0xd1, 0x4d, 0x29, 0x62, 0xc2, 0xbb, 0x8b, 0x4b, 0x69, 0x4b, 0x0c, 0x57,
	0x9e, 0x83, 0x36, 0x73, 0x1c, 0x11, 0x71, 0x65, 0x54, 0xf5, 0x65, 0x5a, 0x8b, 0x4b, 0x79, 0x9f,
	0x20, 0xfd, 0x04, 0x20, 0x0a, 0xda, 0xc8, 0x57, 0x5e, 0x28, 0xf8, 0x12, 0xb9, 0xb2, 0x57, 0x2c,
	0xf4, 0xd8, 0xd4, 0x47, 0x69, 0xec, 0x6b, 0x4f, 0xdf, 0x3f, 0xce, 0x93, 0xf5, 0x20, 0xf1, 0x63,
	0xa6, 0x90, 0x98, 0x3b, 0xc6, 0x02, 0xa8, 0xf3, 0x12, 0x1a, 0x1b, 0x6d, 0x21, 0x4d, 0xa8, 0x2c,
	0x70, 0xad, 0xa7, 0x57, 0xa7, 0xf1, 0x27, 0x39, 0x86, 0xea, 0x8a, 0xf9, 0x51, 0x36, 0x85, 0xe4,
	0xf0, 0x75, 0xf9, 0xb2, 0xd4, 0xf9, 0x0e, 0x9a, 0xbb, 0x1d, 0x78, 0xa3, 0xfc, 0x1b, 0x38, 0xf9,
	0x4f, 0xb7, 0x6f, 0x22, 0xd4, 0xfd, 0xab, 0x0c, 0xad, 0xab, 0x10, 0x99, 0xc2, 0xb8, 0x27, 0x14,
	0xff, 0x88, 0x50, 0x2a, 0xf2, 0x02, 0xf6, 0x3d, 0x1e, 0x44, 0x4a, 0xa6, 0xbb, 0x78, 0xb2, 0xb3,
	0x8b, 0xc3, 0xe4, 0xc7, 0xba, 0x65, 0x01, 0x4d, 0x89, 0xe4, 0x2b, 0xa8, 0x29, 0x5c, 0x06, 0x3e,
	0x53, 0x49, 0x95, 0xc6, 0xf9, 0xb3, 0x82, 0x05, 0x1e, 0xa7, 0x14, 0x9a, 0x93, 0xc9, 0x47, 0xf0,
	0x44, 0x44, 0x2a, 0x88, 0x94, 0x1d, 0x84, 0x38, 0xf7, 0x5e, 0x1b, 0x15, 0xed, 0xf1, 0x30, 0x09,
	0xde, 0xe9, 0x18, 0xf9, 0x15, 0xde, 0xdf, 0xf9, 0x4f, 0x96, 0xe9, 0xd4, 0x8c, 0x3d, 0x5d, 0xec,
	0xe3, 0x47, 0x8d, 0x98, 0xb6, 0x55, 0x51, 0xb8, 0xfb, 0x12, 0xc8, 0x66, 0x13, 0x64, 0x20, 0xb8,
	0xd4, 0xce, 0x42, 0x94, 0x22, 0x0a, 0x1d, 0xd4, 0xe5, 0x74, 0x33, 0x0e, 0xe9, 0x61, 0x16, 0x8c,
	0xd3, 0xbb, 0x14, 0x8e, 0x6e, 0x50, 0x6d, 0x36, 0xef, 0x19, 0xd4, 0xb5, 0x57, 0xb5, 0x0e, 0x30,
	0x1d, 0x42, 0x2d, 0x0e, 0x8c, 0xd7, 0x41, 0x81, 0x66, 0xb9, 0x40, 0xf3, 0x4f, 0x78, 0x9a, 0x6b,
	0xa6, 0x5e, 0xbe, 0x80, 0x5a, 0x46, 0x49, 0x67, 0x62, 0xec, 0xde, 0x98, 0xa6, 0x38, 0xcd, 0x99,
	0xe4, 0x02, 0xea, 0xbe, 0x70, 0x6d, 0xdf, 0xe3, 0x0b, 0x69, 0x94, 0xf5, 0xbf, 0xf0, 0x5e, 0xc1,
	0x54, 0x86, 0xc2, 0xa5, 0x35, 0x5f, 0xb8, 0xc3, 0x98, 0xd7, 0xfd, 0xbb, 0x04, 0xb5, 0x4c, 0x8b,
	0x7c, 0x06, 0x55, 0xa9, 0xe2, 0x99, 0xc6, 0x45, 0x8f, 0xce, 0xdb, 0xbb, 0x45, 0xef, 0x63, 0x90,
	0x26, 0x1c, 0x72, 0x01, 0x07, 0xc9, 0xd4, 0x64, 0xba, 0x02, 0xff, 0xb3, 0x37, 0x19, 0x93, 0x18,
	0x70, 0xb0, 0x44, 0x29, 0x99, 0x8b, 0xe9, 0xe4, 0xb3, 0x63, 0x77, 0x02, 0xad, 0x6b, 0xf4, 0x71,
	0x7b, 0x35, 0xdf, 0xbe, 0xbb, 0xc7, 0x40, 0x36, 0x65, 0x93, 0x06, 0x7f, 0xfa, 0x1b, 0x54, 0xf5,
	0x5d, 0x48, 0x1b, 0x5a, 0xd4, 0x1a, 0xd3, 0x9f, 0xfb, 0xaf, 0x86, 0x96, 0xfd, 0x43, 0x7f, 0x30,
	0x9c, 0x50, 0xab, 0xf9, 0x4e, 0x1c, 0xbe, 0xb3, 0xe8, 0x6d, 0x7f, 0x64, 0x8d, 0xc6, 0x79, 0xb8,
	0x44, 0x1a, 0x70, 0x70, 0x67, 0x8d, 0xae, 0x07, 0xa3, 0x9b, 0x66, 0x39, 0x3e, 0xd0, 0xc9, 0x68,
	0x14, 0x1f, 0x2a, 0xe4, 0x09, 0xd4, 0xef, 0x27, 0x57, 0x57, 0x96, 0x75, 0x6d, 0x5d, 0x37, 0xf7,
	0x5e, 0x7d, 0xfb, 0xcb, 0x37, 0xae, 0xa7, 0x7e, 0x8f, 0xa6, 0xa6, 0x23, 0x96, 0x3d, 0xdd, 0x16,
	0x11, 0xba, 0xc9, 0x47, 0x2f, 0x7f, 0x3e, 0x5c, 0xe4, 0xbd, 0x60, 0xfa, 0xb9, 0x2b, 0x7a, 0xdb,
	0xaf, 0xde, 0x74, 0x5f, 0xbf, 0x24, 0x17, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x8d, 0xa5,
	0xb7, 0x0e, 0x07, 0x00, 0x00,
}
