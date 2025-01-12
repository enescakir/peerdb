// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: route.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ValidatePeerStatus int32

const (
	ValidatePeerStatus_CREATION_UNKNOWN ValidatePeerStatus = 0
	ValidatePeerStatus_VALID            ValidatePeerStatus = 1
	ValidatePeerStatus_INVALID          ValidatePeerStatus = 2
)

// Enum value maps for ValidatePeerStatus.
var (
	ValidatePeerStatus_name = map[int32]string{
		0: "CREATION_UNKNOWN",
		1: "VALID",
		2: "INVALID",
	}
	ValidatePeerStatus_value = map[string]int32{
		"CREATION_UNKNOWN": 0,
		"VALID":            1,
		"INVALID":          2,
	}
)

func (x ValidatePeerStatus) Enum() *ValidatePeerStatus {
	p := new(ValidatePeerStatus)
	*p = x
	return p
}

func (x ValidatePeerStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ValidatePeerStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_route_proto_enumTypes[0].Descriptor()
}

func (ValidatePeerStatus) Type() protoreflect.EnumType {
	return &file_route_proto_enumTypes[0]
}

func (x ValidatePeerStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ValidatePeerStatus.Descriptor instead.
func (ValidatePeerStatus) EnumDescriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{0}
}

type CreatePeerStatus int32

const (
	CreatePeerStatus_VALIDATION_UNKNOWN CreatePeerStatus = 0
	CreatePeerStatus_CREATED            CreatePeerStatus = 1
	CreatePeerStatus_FAILED             CreatePeerStatus = 2
)

// Enum value maps for CreatePeerStatus.
var (
	CreatePeerStatus_name = map[int32]string{
		0: "VALIDATION_UNKNOWN",
		1: "CREATED",
		2: "FAILED",
	}
	CreatePeerStatus_value = map[string]int32{
		"VALIDATION_UNKNOWN": 0,
		"CREATED":            1,
		"FAILED":             2,
	}
)

func (x CreatePeerStatus) Enum() *CreatePeerStatus {
	p := new(CreatePeerStatus)
	*p = x
	return p
}

func (x CreatePeerStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreatePeerStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_route_proto_enumTypes[1].Descriptor()
}

func (CreatePeerStatus) Type() protoreflect.EnumType {
	return &file_route_proto_enumTypes[1]
}

func (x CreatePeerStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreatePeerStatus.Descriptor instead.
func (CreatePeerStatus) EnumDescriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{1}
}

type CreateCDCFlowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionConfigs *FlowConnectionConfigs `protobuf:"bytes,1,opt,name=connection_configs,json=connectionConfigs,proto3" json:"connection_configs,omitempty"`
}

func (x *CreateCDCFlowRequest) Reset() {
	*x = CreateCDCFlowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCDCFlowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCDCFlowRequest) ProtoMessage() {}

func (x *CreateCDCFlowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCDCFlowRequest.ProtoReflect.Descriptor instead.
func (*CreateCDCFlowRequest) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCDCFlowRequest) GetConnectionConfigs() *FlowConnectionConfigs {
	if x != nil {
		return x.ConnectionConfigs
	}
	return nil
}

type CreateCDCFlowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorflowId string `protobuf:"bytes,1,opt,name=worflow_id,json=worflowId,proto3" json:"worflow_id,omitempty"`
}

func (x *CreateCDCFlowResponse) Reset() {
	*x = CreateCDCFlowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCDCFlowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCDCFlowResponse) ProtoMessage() {}

func (x *CreateCDCFlowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCDCFlowResponse.ProtoReflect.Descriptor instead.
func (*CreateCDCFlowResponse) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCDCFlowResponse) GetWorflowId() string {
	if x != nil {
		return x.WorflowId
	}
	return ""
}

type CreateQRepFlowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QrepConfig *QRepConfig `protobuf:"bytes,1,opt,name=qrep_config,json=qrepConfig,proto3" json:"qrep_config,omitempty"`
}

func (x *CreateQRepFlowRequest) Reset() {
	*x = CreateQRepFlowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQRepFlowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQRepFlowRequest) ProtoMessage() {}

func (x *CreateQRepFlowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQRepFlowRequest.ProtoReflect.Descriptor instead.
func (*CreateQRepFlowRequest) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{2}
}

func (x *CreateQRepFlowRequest) GetQrepConfig() *QRepConfig {
	if x != nil {
		return x.QrepConfig
	}
	return nil
}

type CreateQRepFlowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorflowId string `protobuf:"bytes,1,opt,name=worflow_id,json=worflowId,proto3" json:"worflow_id,omitempty"`
}

func (x *CreateQRepFlowResponse) Reset() {
	*x = CreateQRepFlowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQRepFlowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQRepFlowResponse) ProtoMessage() {}

func (x *CreateQRepFlowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQRepFlowResponse.ProtoReflect.Descriptor instead.
func (*CreateQRepFlowResponse) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{3}
}

func (x *CreateQRepFlowResponse) GetWorflowId() string {
	if x != nil {
		return x.WorflowId
	}
	return ""
}

type ShutdownRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowId      string `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	FlowJobName     string `protobuf:"bytes,2,opt,name=flow_job_name,json=flowJobName,proto3" json:"flow_job_name,omitempty"`
	SourcePeer      *Peer  `protobuf:"bytes,3,opt,name=source_peer,json=sourcePeer,proto3" json:"source_peer,omitempty"`
	DestinationPeer *Peer  `protobuf:"bytes,4,opt,name=destination_peer,json=destinationPeer,proto3" json:"destination_peer,omitempty"`
}

func (x *ShutdownRequest) Reset() {
	*x = ShutdownRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShutdownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShutdownRequest) ProtoMessage() {}

func (x *ShutdownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShutdownRequest.ProtoReflect.Descriptor instead.
func (*ShutdownRequest) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{4}
}

func (x *ShutdownRequest) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *ShutdownRequest) GetFlowJobName() string {
	if x != nil {
		return x.FlowJobName
	}
	return ""
}

func (x *ShutdownRequest) GetSourcePeer() *Peer {
	if x != nil {
		return x.SourcePeer
	}
	return nil
}

func (x *ShutdownRequest) GetDestinationPeer() *Peer {
	if x != nil {
		return x.DestinationPeer
	}
	return nil
}

type ShutdownResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok           bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *ShutdownResponse) Reset() {
	*x = ShutdownResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShutdownResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShutdownResponse) ProtoMessage() {}

func (x *ShutdownResponse) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShutdownResponse.ProtoReflect.Descriptor instead.
func (*ShutdownResponse) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{5}
}

func (x *ShutdownResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *ShutdownResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type ListPeersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPeersRequest) Reset() {
	*x = ListPeersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPeersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPeersRequest) ProtoMessage() {}

func (x *ListPeersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPeersRequest.ProtoReflect.Descriptor instead.
func (*ListPeersRequest) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{6}
}

type ListPeersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peers []*Peer `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *ListPeersResponse) Reset() {
	*x = ListPeersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPeersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPeersResponse) ProtoMessage() {}

func (x *ListPeersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPeersResponse.ProtoReflect.Descriptor instead.
func (*ListPeersResponse) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{7}
}

func (x *ListPeersResponse) GetPeers() []*Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

type ValidatePeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peer *Peer `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`
}

func (x *ValidatePeerRequest) Reset() {
	*x = ValidatePeerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatePeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatePeerRequest) ProtoMessage() {}

func (x *ValidatePeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatePeerRequest.ProtoReflect.Descriptor instead.
func (*ValidatePeerRequest) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{8}
}

func (x *ValidatePeerRequest) GetPeer() *Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

type CreatePeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peer *Peer `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`
}

func (x *CreatePeerRequest) Reset() {
	*x = CreatePeerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePeerRequest) ProtoMessage() {}

func (x *CreatePeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePeerRequest.ProtoReflect.Descriptor instead.
func (*CreatePeerRequest) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{9}
}

func (x *CreatePeerRequest) GetPeer() *Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

type ValidatePeerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  ValidatePeerStatus `protobuf:"varint,1,opt,name=status,proto3,enum=peerdb_route.ValidatePeerStatus" json:"status,omitempty"`
	Message string             `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ValidatePeerResponse) Reset() {
	*x = ValidatePeerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatePeerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatePeerResponse) ProtoMessage() {}

func (x *ValidatePeerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatePeerResponse.ProtoReflect.Descriptor instead.
func (*ValidatePeerResponse) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{10}
}

func (x *ValidatePeerResponse) GetStatus() ValidatePeerStatus {
	if x != nil {
		return x.Status
	}
	return ValidatePeerStatus_CREATION_UNKNOWN
}

func (x *ValidatePeerResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreatePeerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  CreatePeerStatus `protobuf:"varint,1,opt,name=status,proto3,enum=peerdb_route.CreatePeerStatus" json:"status,omitempty"`
	Message string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreatePeerResponse) Reset() {
	*x = CreatePeerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePeerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePeerResponse) ProtoMessage() {}

func (x *CreatePeerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePeerResponse.ProtoReflect.Descriptor instead.
func (*CreatePeerResponse) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{11}
}

func (x *CreatePeerResponse) GetStatus() CreatePeerStatus {
	if x != nil {
		return x.Status
	}
	return CreatePeerStatus_VALIDATION_UNKNOWN
}

func (x *CreatePeerResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_route_proto protoreflect.FileDescriptor

var file_route_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70,
	0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x70, 0x65,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x44, 0x43, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x51, 0x0a,
	0x12, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x65, 0x65, 0x72,
	0x64, 0x62, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x11, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x22, 0x36, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x44, 0x43, 0x46, 0x6c, 0x6f,
	0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x51, 0x52, 0x65, 0x70, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x38, 0x0a, 0x0b, 0x71, 0x72, 0x65, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x51, 0x52, 0x65, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x0a, 0x71, 0x72, 0x65, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x37, 0x0a, 0x16, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x52, 0x65, 0x70, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x49, 0x64, 0x22, 0xca, 0x01, 0x0a, 0x0f, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x66, 0x6c, 0x6f,
	0x77, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x66, 0x6c, 0x6f, 0x77, 0x4a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a,
	0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x70, 0x65, 0x65, 0x72,
	0x73, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x65,
	0x65, 0x72, 0x12, 0x3d, 0x0a, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x65, 0x65, 0x72,
	0x52, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x65,
	0x72, 0x22, 0x47, 0x0a, 0x10, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3d,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x70, 0x65, 0x65, 0x72,
	0x73, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x22, 0x3d, 0x0a,
	0x13, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x70, 0x65, 0x65, 0x72,
	0x73, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x22, 0x3b, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x26, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x50,
	0x65, 0x65, 0x72, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x22, 0x6a, 0x0a, 0x14, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x20, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x66, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x70, 0x65,
	0x65, 0x72, 0x64, 0x62, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x42, 0x0a,
	0x12, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0x02, 0x2a, 0x43, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x32, 0x95, 0x04, 0x0a, 0x0b, 0x46, 0x6c, 0x6f, 0x77, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65,
	0x65, 0x72, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x65, 0x65, 0x72,
	0x64, 0x62, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x51, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x12, 0x1f, 0x2e,
	0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x44, 0x43, 0x46,
	0x6c, 0x6f, 0x77, 0x12, 0x22, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x44, 0x43, 0x46, 0x6c, 0x6f, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62,
	0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x44, 0x43,
	0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x52, 0x65, 0x70, 0x46, 0x6c, 0x6f, 0x77,
	0x12, 0x23, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x52, 0x65, 0x70, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x52, 0x65, 0x70, 0x46,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a,
	0x0c, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x46, 0x6c, 0x6f, 0x77, 0x12, 0x1d, 0x2e,
	0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x53, 0x68, 0x75,
	0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70,
	0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x53, 0x68, 0x75, 0x74,
	0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x7c,
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x42, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x10, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x50, 0x65, 0x65, 0x72, 0x64,
	0x62, 0x52, 0x6f, 0x75, 0x74, 0x65, 0xca, 0x02, 0x0b, 0x50, 0x65, 0x65, 0x72, 0x64, 0x62, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0xe2, 0x02, 0x17, 0x50, 0x65, 0x65, 0x72, 0x64, 0x62, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0b, 0x50, 0x65, 0x65, 0x72, 0x64, 0x62, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_route_proto_rawDescOnce sync.Once
	file_route_proto_rawDescData = file_route_proto_rawDesc
)

func file_route_proto_rawDescGZIP() []byte {
	file_route_proto_rawDescOnce.Do(func() {
		file_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_route_proto_rawDescData)
	})
	return file_route_proto_rawDescData
}

var file_route_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_route_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_route_proto_goTypes = []interface{}{
	(ValidatePeerStatus)(0),        // 0: peerdb_route.ValidatePeerStatus
	(CreatePeerStatus)(0),          // 1: peerdb_route.CreatePeerStatus
	(*CreateCDCFlowRequest)(nil),   // 2: peerdb_route.CreateCDCFlowRequest
	(*CreateCDCFlowResponse)(nil),  // 3: peerdb_route.CreateCDCFlowResponse
	(*CreateQRepFlowRequest)(nil),  // 4: peerdb_route.CreateQRepFlowRequest
	(*CreateQRepFlowResponse)(nil), // 5: peerdb_route.CreateQRepFlowResponse
	(*ShutdownRequest)(nil),        // 6: peerdb_route.ShutdownRequest
	(*ShutdownResponse)(nil),       // 7: peerdb_route.ShutdownResponse
	(*ListPeersRequest)(nil),       // 8: peerdb_route.ListPeersRequest
	(*ListPeersResponse)(nil),      // 9: peerdb_route.ListPeersResponse
	(*ValidatePeerRequest)(nil),    // 10: peerdb_route.ValidatePeerRequest
	(*CreatePeerRequest)(nil),      // 11: peerdb_route.CreatePeerRequest
	(*ValidatePeerResponse)(nil),   // 12: peerdb_route.ValidatePeerResponse
	(*CreatePeerResponse)(nil),     // 13: peerdb_route.CreatePeerResponse
	(*FlowConnectionConfigs)(nil),  // 14: peerdb_flow.FlowConnectionConfigs
	(*QRepConfig)(nil),             // 15: peerdb_flow.QRepConfig
	(*Peer)(nil),                   // 16: peerdb_peers.Peer
}
var file_route_proto_depIdxs = []int32{
	14, // 0: peerdb_route.CreateCDCFlowRequest.connection_configs:type_name -> peerdb_flow.FlowConnectionConfigs
	15, // 1: peerdb_route.CreateQRepFlowRequest.qrep_config:type_name -> peerdb_flow.QRepConfig
	16, // 2: peerdb_route.ShutdownRequest.source_peer:type_name -> peerdb_peers.Peer
	16, // 3: peerdb_route.ShutdownRequest.destination_peer:type_name -> peerdb_peers.Peer
	16, // 4: peerdb_route.ListPeersResponse.peers:type_name -> peerdb_peers.Peer
	16, // 5: peerdb_route.ValidatePeerRequest.peer:type_name -> peerdb_peers.Peer
	16, // 6: peerdb_route.CreatePeerRequest.peer:type_name -> peerdb_peers.Peer
	0,  // 7: peerdb_route.ValidatePeerResponse.status:type_name -> peerdb_route.ValidatePeerStatus
	1,  // 8: peerdb_route.CreatePeerResponse.status:type_name -> peerdb_route.CreatePeerStatus
	8,  // 9: peerdb_route.FlowService.ListPeers:input_type -> peerdb_route.ListPeersRequest
	10, // 10: peerdb_route.FlowService.ValidatePeer:input_type -> peerdb_route.ValidatePeerRequest
	11, // 11: peerdb_route.FlowService.CreatePeer:input_type -> peerdb_route.CreatePeerRequest
	2,  // 12: peerdb_route.FlowService.CreateCDCFlow:input_type -> peerdb_route.CreateCDCFlowRequest
	4,  // 13: peerdb_route.FlowService.CreateQRepFlow:input_type -> peerdb_route.CreateQRepFlowRequest
	6,  // 14: peerdb_route.FlowService.ShutdownFlow:input_type -> peerdb_route.ShutdownRequest
	9,  // 15: peerdb_route.FlowService.ListPeers:output_type -> peerdb_route.ListPeersResponse
	12, // 16: peerdb_route.FlowService.ValidatePeer:output_type -> peerdb_route.ValidatePeerResponse
	13, // 17: peerdb_route.FlowService.CreatePeer:output_type -> peerdb_route.CreatePeerResponse
	3,  // 18: peerdb_route.FlowService.CreateCDCFlow:output_type -> peerdb_route.CreateCDCFlowResponse
	5,  // 19: peerdb_route.FlowService.CreateQRepFlow:output_type -> peerdb_route.CreateQRepFlowResponse
	7,  // 20: peerdb_route.FlowService.ShutdownFlow:output_type -> peerdb_route.ShutdownResponse
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_route_proto_init() }
func file_route_proto_init() {
	if File_route_proto != nil {
		return
	}
	file_peers_proto_init()
	file_flow_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCDCFlowRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCDCFlowResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQRepFlowRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQRepFlowResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShutdownRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShutdownResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPeersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPeersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatePeerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePeerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatePeerResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePeerResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_route_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_route_proto_goTypes,
		DependencyIndexes: file_route_proto_depIdxs,
		EnumInfos:         file_route_proto_enumTypes,
		MessageInfos:      file_route_proto_msgTypes,
	}.Build()
	File_route_proto = out.File
	file_route_proto_rawDesc = nil
	file_route_proto_goTypes = nil
	file_route_proto_depIdxs = nil
}
