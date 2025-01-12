// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: plugins/base/proto/v1/base.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PluginType int32

const (
	PluginType_PLUGIN_TYPE_UNSPECIFIED PluginType = 0
	PluginType_PLUGIN_TYPE_APM         PluginType = 1
	PluginType_PLUGIN_TYPE_STRATEGY    PluginType = 2
	PluginType_PLUGIN_TYPE_TARGET      PluginType = 3
)

// Enum value maps for PluginType.
var (
	PluginType_name = map[int32]string{
		0: "PLUGIN_TYPE_UNSPECIFIED",
		1: "PLUGIN_TYPE_APM",
		2: "PLUGIN_TYPE_STRATEGY",
		3: "PLUGIN_TYPE_TARGET",
	}
	PluginType_value = map[string]int32{
		"PLUGIN_TYPE_UNSPECIFIED": 0,
		"PLUGIN_TYPE_APM":         1,
		"PLUGIN_TYPE_STRATEGY":    2,
		"PLUGIN_TYPE_TARGET":      3,
	}
)

func (x PluginType) Enum() *PluginType {
	p := new(PluginType)
	*p = x
	return p
}

func (x PluginType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PluginType) Descriptor() protoreflect.EnumDescriptor {
	return file_plugins_base_proto_v1_base_proto_enumTypes[0].Descriptor()
}

func (PluginType) Type() protoreflect.EnumType {
	return &file_plugins_base_proto_v1_base_proto_enumTypes[0]
}

func (x PluginType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PluginType.Descriptor instead.
func (PluginType) EnumDescriptor() ([]byte, []int) {
	return file_plugins_base_proto_v1_base_proto_rawDescGZIP(), []int{0}
}

type PluginInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PluginInfoRequest) Reset() {
	*x = PluginInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_base_proto_v1_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginInfoRequest) ProtoMessage() {}

func (x *PluginInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_base_proto_v1_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginInfoRequest.ProtoReflect.Descriptor instead.
func (*PluginInfoRequest) Descriptor() ([]byte, []int) {
	return file_plugins_base_proto_v1_base_proto_rawDescGZIP(), []int{0}
}

type PluginInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type PluginType `protobuf:"varint,2,opt,name=type,proto3,enum=hashicorp.nomad_autoscaler.plugins.base.proto.v1.PluginType" json:"type,omitempty"`
}

func (x *PluginInfoResponse) Reset() {
	*x = PluginInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_base_proto_v1_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginInfoResponse) ProtoMessage() {}

func (x *PluginInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_base_proto_v1_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginInfoResponse.ProtoReflect.Descriptor instead.
func (*PluginInfoResponse) Descriptor() ([]byte, []int) {
	return file_plugins_base_proto_v1_base_proto_rawDescGZIP(), []int{1}
}

func (x *PluginInfoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PluginInfoResponse) GetType() PluginType {
	if x != nil {
		return x.Type
	}
	return PluginType_PLUGIN_TYPE_UNSPECIFIED
}

type SetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config map[string]string `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SetConfigRequest) Reset() {
	*x = SetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_base_proto_v1_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetConfigRequest) ProtoMessage() {}

func (x *SetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_base_proto_v1_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetConfigRequest.ProtoReflect.Descriptor instead.
func (*SetConfigRequest) Descriptor() ([]byte, []int) {
	return file_plugins_base_proto_v1_base_proto_rawDescGZIP(), []int{2}
}

func (x *SetConfigRequest) GetConfig() map[string]string {
	if x != nil {
		return x.Config
	}
	return nil
}

type SetConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetConfigResponse) Reset() {
	*x = SetConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_base_proto_v1_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetConfigResponse) ProtoMessage() {}

func (x *SetConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_base_proto_v1_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetConfigResponse.ProtoReflect.Descriptor instead.
func (*SetConfigResponse) Descriptor() ([]byte, []int) {
	return file_plugins_base_proto_v1_base_proto_rawDescGZIP(), []int{3}
}

var File_plugins_base_proto_v1_base_proto protoreflect.FileDescriptor

var file_plugins_base_proto_v1_base_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x30, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x6e, 0x6f,
	0x6d, 0x61, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x22, 0x13, 0x0a, 0x11, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x7a, 0x0a, 0x12, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x50, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x3c, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x6e, 0x6f,
	0x6d, 0x61, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xb5, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x66, 0x0a, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x61, 0x64, 0x5f, 0x61, 0x75, 0x74,
	0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x1a, 0x39, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x13, 0x0a,
	0x11, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2a, 0x70, 0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x17, 0x50, 0x4c, 0x55, 0x47, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a,
	0x0f, 0x50, 0x4c, 0x55, 0x47, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x50, 0x4d,
	0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x4c, 0x55, 0x47, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12,
	0x50, 0x4c, 0x55, 0x47, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x41, 0x52, 0x47,
	0x45, 0x54, 0x10, 0x03, 0x32, 0xc8, 0x02, 0x0a, 0x11, 0x42, 0x61, 0x73, 0x65, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x99, 0x01, 0x0a, 0x0a, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x43, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x61, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x6f,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44,
	0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x61, 0x64,
	0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x96, 0x01, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x42, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x6e, 0x6f, 0x6d, 0x61, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x61, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x07, 0x5a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugins_base_proto_v1_base_proto_rawDescOnce sync.Once
	file_plugins_base_proto_v1_base_proto_rawDescData = file_plugins_base_proto_v1_base_proto_rawDesc
)

func file_plugins_base_proto_v1_base_proto_rawDescGZIP() []byte {
	file_plugins_base_proto_v1_base_proto_rawDescOnce.Do(func() {
		file_plugins_base_proto_v1_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugins_base_proto_v1_base_proto_rawDescData)
	})
	return file_plugins_base_proto_v1_base_proto_rawDescData
}

var file_plugins_base_proto_v1_base_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_plugins_base_proto_v1_base_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_plugins_base_proto_v1_base_proto_goTypes = []interface{}{
	(PluginType)(0),            // 0: hashicorp.nomad_autoscaler.plugins.base.proto.v1.PluginType
	(*PluginInfoRequest)(nil),  // 1: hashicorp.nomad_autoscaler.plugins.base.proto.v1.PluginInfoRequest
	(*PluginInfoResponse)(nil), // 2: hashicorp.nomad_autoscaler.plugins.base.proto.v1.PluginInfoResponse
	(*SetConfigRequest)(nil),   // 3: hashicorp.nomad_autoscaler.plugins.base.proto.v1.SetConfigRequest
	(*SetConfigResponse)(nil),  // 4: hashicorp.nomad_autoscaler.plugins.base.proto.v1.SetConfigResponse
	nil,                        // 5: hashicorp.nomad_autoscaler.plugins.base.proto.v1.SetConfigRequest.ConfigEntry
}
var file_plugins_base_proto_v1_base_proto_depIdxs = []int32{
	0, // 0: hashicorp.nomad_autoscaler.plugins.base.proto.v1.PluginInfoResponse.type:type_name -> hashicorp.nomad_autoscaler.plugins.base.proto.v1.PluginType
	5, // 1: hashicorp.nomad_autoscaler.plugins.base.proto.v1.SetConfigRequest.config:type_name -> hashicorp.nomad_autoscaler.plugins.base.proto.v1.SetConfigRequest.ConfigEntry
	1, // 2: hashicorp.nomad_autoscaler.plugins.base.proto.v1.BasePluginService.PluginInfo:input_type -> hashicorp.nomad_autoscaler.plugins.base.proto.v1.PluginInfoRequest
	3, // 3: hashicorp.nomad_autoscaler.plugins.base.proto.v1.BasePluginService.SetConfig:input_type -> hashicorp.nomad_autoscaler.plugins.base.proto.v1.SetConfigRequest
	2, // 4: hashicorp.nomad_autoscaler.plugins.base.proto.v1.BasePluginService.PluginInfo:output_type -> hashicorp.nomad_autoscaler.plugins.base.proto.v1.PluginInfoResponse
	4, // 5: hashicorp.nomad_autoscaler.plugins.base.proto.v1.BasePluginService.SetConfig:output_type -> hashicorp.nomad_autoscaler.plugins.base.proto.v1.SetConfigResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_plugins_base_proto_v1_base_proto_init() }
func file_plugins_base_proto_v1_base_proto_init() {
	if File_plugins_base_proto_v1_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugins_base_proto_v1_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginInfoRequest); i {
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
		file_plugins_base_proto_v1_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginInfoResponse); i {
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
		file_plugins_base_proto_v1_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetConfigRequest); i {
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
		file_plugins_base_proto_v1_base_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetConfigResponse); i {
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
			RawDescriptor: file_plugins_base_proto_v1_base_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plugins_base_proto_v1_base_proto_goTypes,
		DependencyIndexes: file_plugins_base_proto_v1_base_proto_depIdxs,
		EnumInfos:         file_plugins_base_proto_v1_base_proto_enumTypes,
		MessageInfos:      file_plugins_base_proto_v1_base_proto_msgTypes,
	}.Build()
	File_plugins_base_proto_v1_base_proto = out.File
	file_plugins_base_proto_v1_base_proto_rawDesc = nil
	file_plugins_base_proto_v1_base_proto_goTypes = nil
	file_plugins_base_proto_v1_base_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BasePluginServiceClient is the client API for BasePluginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BasePluginServiceClient interface {
	PluginInfo(ctx context.Context, in *PluginInfoRequest, opts ...grpc.CallOption) (*PluginInfoResponse, error)
	SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error)
}

type basePluginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBasePluginServiceClient(cc grpc.ClientConnInterface) BasePluginServiceClient {
	return &basePluginServiceClient{cc}
}

func (c *basePluginServiceClient) PluginInfo(ctx context.Context, in *PluginInfoRequest, opts ...grpc.CallOption) (*PluginInfoResponse, error) {
	out := new(PluginInfoResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad_autoscaler.plugins.base.proto.v1.BasePluginService/PluginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basePluginServiceClient) SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error) {
	out := new(SetConfigResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad_autoscaler.plugins.base.proto.v1.BasePluginService/SetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasePluginServiceServer is the server API for BasePluginService service.
type BasePluginServiceServer interface {
	PluginInfo(context.Context, *PluginInfoRequest) (*PluginInfoResponse, error)
	SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error)
}

// UnimplementedBasePluginServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBasePluginServiceServer struct {
}

func (*UnimplementedBasePluginServiceServer) PluginInfo(context.Context, *PluginInfoRequest) (*PluginInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginInfo not implemented")
}
func (*UnimplementedBasePluginServiceServer) SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConfig not implemented")
}

func RegisterBasePluginServiceServer(s *grpc.Server, srv BasePluginServiceServer) {
	s.RegisterService(&_BasePluginService_serviceDesc, srv)
}

func _BasePluginService_PluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasePluginServiceServer).PluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad_autoscaler.plugins.base.proto.v1.BasePluginService/PluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasePluginServiceServer).PluginInfo(ctx, req.(*PluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasePluginService_SetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasePluginServiceServer).SetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad_autoscaler.plugins.base.proto.v1.BasePluginService/SetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasePluginServiceServer).SetConfig(ctx, req.(*SetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BasePluginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.nomad_autoscaler.plugins.base.proto.v1.BasePluginService",
	HandlerType: (*BasePluginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PluginInfo",
			Handler:    _BasePluginService_PluginInfo_Handler,
		},
		{
			MethodName: "SetConfig",
			Handler:    _BasePluginService_SetConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugins/base/proto/v1/base.proto",
}
