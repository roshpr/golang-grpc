// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: bootstrap.proto

package bootstrap

import (
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

type BootStrapConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Arch  string `protobuf:"bytes,2,opt,name=arch,proto3" json:"arch,omitempty"`
	Cores int64  `protobuf:"varint,3,opt,name=cores,proto3" json:"cores,omitempty"`
}

func (x *BootStrapConfig) Reset() {
	*x = BootStrapConfig{}
	mi := &file_bootstrap_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BootStrapConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootStrapConfig) ProtoMessage() {}

func (x *BootStrapConfig) ProtoReflect() protoreflect.Message {
	mi := &file_bootstrap_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootStrapConfig.ProtoReflect.Descriptor instead.
func (*BootStrapConfig) Descriptor() ([]byte, []int) {
	return file_bootstrap_proto_rawDescGZIP(), []int{0}
}

func (x *BootStrapConfig) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *BootStrapConfig) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

func (x *BootStrapConfig) GetCores() int64 {
	if x != nil {
		return x.Cores
	}
	return 0
}

type BootStrapResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Arch  string `protobuf:"bytes,2,opt,name=arch,proto3" json:"arch,omitempty"`
	Cores int64  `protobuf:"varint,3,opt,name=cores,proto3" json:"cores,omitempty"`
}

func (x *BootStrapResponse) Reset() {
	*x = BootStrapResponse{}
	mi := &file_bootstrap_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BootStrapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootStrapResponse) ProtoMessage() {}

func (x *BootStrapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bootstrap_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootStrapResponse.ProtoReflect.Descriptor instead.
func (*BootStrapResponse) Descriptor() ([]byte, []int) {
	return file_bootstrap_proto_rawDescGZIP(), []int{1}
}

func (x *BootStrapResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *BootStrapResponse) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

func (x *BootStrapResponse) GetCores() int64 {
	if x != nil {
		return x.Cores
	}
	return 0
}

var File_bootstrap_proto protoreflect.FileDescriptor

var file_bootstrap_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4f, 0x0a, 0x0f, 0x42, 0x6f, 0x6f, 0x74, 0x53, 0x74, 0x72, 0x61, 0x70, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x63, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x22, 0x51, 0x0a, 0x11, 0x42, 0x6f, 0x6f, 0x74, 0x53, 0x74, 0x72, 0x61, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x72, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x6f, 0x72, 0x65, 0x73, 0x32, 0x3b, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72,
	0x61, 0x70, 0x12, 0x2e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x42,
	0x6f, 0x6f, 0x74, 0x53, 0x74, 0x72, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x12,
	0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x53, 0x74, 0x72, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x72, 0x6f, 0x73, 0x68, 0x70, 0x72, 0x2e, 0x6e, 0x65, 0x74,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x63, 0x63, 0x2d, 0x62, 0x6f, 0x6f,
	0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74,
	0x72, 0x61, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bootstrap_proto_rawDescOnce sync.Once
	file_bootstrap_proto_rawDescData = file_bootstrap_proto_rawDesc
)

func file_bootstrap_proto_rawDescGZIP() []byte {
	file_bootstrap_proto_rawDescOnce.Do(func() {
		file_bootstrap_proto_rawDescData = protoimpl.X.CompressGZIP(file_bootstrap_proto_rawDescData)
	})
	return file_bootstrap_proto_rawDescData
}

var file_bootstrap_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bootstrap_proto_goTypes = []any{
	(*BootStrapConfig)(nil),   // 0: BootStrapConfig
	(*BootStrapResponse)(nil), // 1: BootStrapResponse
}
var file_bootstrap_proto_depIdxs = []int32{
	0, // 0: Bootstrap.Create:input_type -> BootStrapConfig
	1, // 1: Bootstrap.Create:output_type -> BootStrapResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bootstrap_proto_init() }
func file_bootstrap_proto_init() {
	if File_bootstrap_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bootstrap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bootstrap_proto_goTypes,
		DependencyIndexes: file_bootstrap_proto_depIdxs,
		MessageInfos:      file_bootstrap_proto_msgTypes,
	}.Build()
	File_bootstrap_proto = out.File
	file_bootstrap_proto_rawDesc = nil
	file_bootstrap_proto_goTypes = nil
	file_bootstrap_proto_depIdxs = nil
}
