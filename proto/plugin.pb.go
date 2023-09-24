// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: plugin.proto

package proto

import (
	context "context"
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

type EventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PluginInterface string `protobuf:"bytes,1,opt,name=pluginInterface,proto3" json:"pluginInterface,omitempty"`
}

func (x *EventsRequest) Reset() {
	*x = EventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventsRequest) ProtoMessage() {}

func (x *EventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventsRequest.ProtoReflect.Descriptor instead.
func (*EventsRequest) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *EventsRequest) GetPluginInterface() string {
	if x != nil {
		return x.PluginInterface
	}
	return ""
}

type EventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EventsResponse) Reset() {
	*x = EventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventsResponse) ProtoMessage() {}

func (x *EventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventsResponse.ProtoReflect.Descriptor instead.
func (*EventsResponse) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *EventsResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type InteractionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PluginInterface string `protobuf:"bytes,1,opt,name=pluginInterface,proto3" json:"pluginInterface,omitempty"`
	Function        string `protobuf:"bytes,2,opt,name=function,proto3" json:"function,omitempty"`
}

func (x *InteractionRequest) Reset() {
	*x = InteractionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InteractionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InteractionRequest) ProtoMessage() {}

func (x *InteractionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InteractionRequest.ProtoReflect.Descriptor instead.
func (*InteractionRequest) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *InteractionRequest) GetPluginInterface() string {
	if x != nil {
		return x.PluginInterface
	}
	return ""
}

func (x *InteractionRequest) GetFunction() string {
	if x != nil {
		return x.Function
	}
	return ""
}

type InteractionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *InteractionResponse) Reset() {
	*x = InteractionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InteractionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InteractionResponse) ProtoMessage() {}

func (x *InteractionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InteractionResponse.ProtoReflect.Descriptor instead.
func (*InteractionResponse) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{3}
}

func (x *InteractionResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_plugin_proto protoreflect.FileDescriptor

var file_plugin_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x22, 0x26, 0x0a, 0x0e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5a, 0x0a, 0x12, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28,
	0x0a, 0x0f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2b, 0x0a, 0x13, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x32, 0xdf, 0x03, 0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x35, 0x0a, 0x06,
	0x4f, 0x6e, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x4f, 0x6e, 0x55, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09,
	0x4f, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x4f, 0x6e, 0x55, 0x6e, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x4f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x4f, 0x6e,
	0x53, 0x74, 0x6f, 0x70, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x44, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x55, 0x49, 0x46, 0x72, 0x61,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugin_proto_rawDescOnce sync.Once
	file_plugin_proto_rawDescData = file_plugin_proto_rawDesc
)

func file_plugin_proto_rawDescGZIP() []byte {
	file_plugin_proto_rawDescOnce.Do(func() {
		file_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugin_proto_rawDescData)
	})
	return file_plugin_proto_rawDescData
}

var file_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_plugin_proto_goTypes = []interface{}{
	(*EventsRequest)(nil),       // 0: proto.EventsRequest
	(*EventsResponse)(nil),      // 1: proto.EventsResponse
	(*InteractionRequest)(nil),  // 2: proto.InteractionRequest
	(*InteractionResponse)(nil), // 3: proto.InteractionResponse
}
var file_plugin_proto_depIdxs = []int32{
	0, // 0: proto.Plugin.OnLoad:input_type -> proto.EventsRequest
	0, // 1: proto.Plugin.OnUnload:input_type -> proto.EventsRequest
	0, // 2: proto.Plugin.OnInstall:input_type -> proto.EventsRequest
	0, // 3: proto.Plugin.OnUninstall:input_type -> proto.EventsRequest
	0, // 4: proto.Plugin.OnStart:input_type -> proto.EventsRequest
	0, // 5: proto.Plugin.OnStop:input_type -> proto.EventsRequest
	2, // 6: proto.Plugin.Interaction:input_type -> proto.InteractionRequest
	0, // 7: proto.Plugin.UIFramework:input_type -> proto.EventsRequest
	1, // 8: proto.Plugin.OnLoad:output_type -> proto.EventsResponse
	1, // 9: proto.Plugin.OnUnload:output_type -> proto.EventsResponse
	1, // 10: proto.Plugin.OnInstall:output_type -> proto.EventsResponse
	1, // 11: proto.Plugin.OnUninstall:output_type -> proto.EventsResponse
	1, // 12: proto.Plugin.OnStart:output_type -> proto.EventsResponse
	1, // 13: proto.Plugin.OnStop:output_type -> proto.EventsResponse
	3, // 14: proto.Plugin.Interaction:output_type -> proto.InteractionResponse
	1, // 15: proto.Plugin.UIFramework:output_type -> proto.EventsResponse
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_plugin_proto_init() }
func file_plugin_proto_init() {
	if File_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventsRequest); i {
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
		file_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventsResponse); i {
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
		file_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InteractionRequest); i {
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
		file_plugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InteractionResponse); i {
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
			RawDescriptor: file_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plugin_proto_goTypes,
		DependencyIndexes: file_plugin_proto_depIdxs,
		MessageInfos:      file_plugin_proto_msgTypes,
	}.Build()
	File_plugin_proto = out.File
	file_plugin_proto_rawDesc = nil
	file_plugin_proto_goTypes = nil
	file_plugin_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PluginClient is the client API for Plugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PluginClient interface {
	OnLoad(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (*EventsResponse, error)
	OnUnload(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (*EventsResponse, error)
	OnInstall(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (*EventsResponse, error)
	OnUninstall(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (*EventsResponse, error)
	OnStart(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (*EventsResponse, error)
	OnStop(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (*EventsResponse, error)
	Interaction(ctx context.Context, in *InteractionRequest, opts ...grpc.CallOption) (*InteractionResponse, error)
	UIFramework(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (*EventsResponse, error)
}

type pluginClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginClient(cc grpc.ClientConnInterface) PluginClient {
	return &pluginClient{cc}
}

func (c *pluginClient) OnLoad(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/proto.Plugin/OnLoad", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) OnUnload(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/proto.Plugin/OnUnload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) OnInstall(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/proto.Plugin/OnInstall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) OnUninstall(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/proto.Plugin/OnUninstall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) OnStart(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/proto.Plugin/OnStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) OnStop(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/proto.Plugin/OnStop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Interaction(ctx context.Context, in *InteractionRequest, opts ...grpc.CallOption) (*InteractionResponse, error) {
	out := new(InteractionResponse)
	err := c.cc.Invoke(ctx, "/proto.Plugin/Interaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) UIFramework(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/proto.Plugin/UIFramework", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServer is the server API for Plugin service.
type PluginServer interface {
	OnLoad(context.Context, *EventsRequest) (*EventsResponse, error)
	OnUnload(context.Context, *EventsRequest) (*EventsResponse, error)
	OnInstall(context.Context, *EventsRequest) (*EventsResponse, error)
	OnUninstall(context.Context, *EventsRequest) (*EventsResponse, error)
	OnStart(context.Context, *EventsRequest) (*EventsResponse, error)
	OnStop(context.Context, *EventsRequest) (*EventsResponse, error)
	Interaction(context.Context, *InteractionRequest) (*InteractionResponse, error)
	UIFramework(context.Context, *EventsRequest) (*EventsResponse, error)
}

// UnimplementedPluginServer can be embedded to have forward compatible implementations.
type UnimplementedPluginServer struct {
}

func (*UnimplementedPluginServer) OnLoad(context.Context, *EventsRequest) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnLoad not implemented")
}
func (*UnimplementedPluginServer) OnUnload(context.Context, *EventsRequest) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnUnload not implemented")
}
func (*UnimplementedPluginServer) OnInstall(context.Context, *EventsRequest) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnInstall not implemented")
}
func (*UnimplementedPluginServer) OnUninstall(context.Context, *EventsRequest) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnUninstall not implemented")
}
func (*UnimplementedPluginServer) OnStart(context.Context, *EventsRequest) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnStart not implemented")
}
func (*UnimplementedPluginServer) OnStop(context.Context, *EventsRequest) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnStop not implemented")
}
func (*UnimplementedPluginServer) Interaction(context.Context, *InteractionRequest) (*InteractionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Interaction not implemented")
}
func (*UnimplementedPluginServer) UIFramework(context.Context, *EventsRequest) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UIFramework not implemented")
}

func RegisterPluginServer(s *grpc.Server, srv PluginServer) {
	s.RegisterService(&_Plugin_serviceDesc, srv)
}

func _Plugin_OnLoad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).OnLoad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/OnLoad",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).OnLoad(ctx, req.(*EventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_OnUnload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).OnUnload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/OnUnload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).OnUnload(ctx, req.(*EventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_OnInstall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).OnInstall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/OnInstall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).OnInstall(ctx, req.(*EventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_OnUninstall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).OnUninstall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/OnUninstall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).OnUninstall(ctx, req.(*EventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_OnStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).OnStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/OnStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).OnStart(ctx, req.(*EventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_OnStop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).OnStop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/OnStop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).OnStop(ctx, req.(*EventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Interaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InteractionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Interaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/Interaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Interaction(ctx, req.(*InteractionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_UIFramework_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).UIFramework(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/UIFramework",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).UIFramework(ctx, req.(*EventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Plugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Plugin",
	HandlerType: (*PluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnLoad",
			Handler:    _Plugin_OnLoad_Handler,
		},
		{
			MethodName: "OnUnload",
			Handler:    _Plugin_OnUnload_Handler,
		},
		{
			MethodName: "OnInstall",
			Handler:    _Plugin_OnInstall_Handler,
		},
		{
			MethodName: "OnUninstall",
			Handler:    _Plugin_OnUninstall_Handler,
		},
		{
			MethodName: "OnStart",
			Handler:    _Plugin_OnStart_Handler,
		},
		{
			MethodName: "OnStop",
			Handler:    _Plugin_OnStop_Handler,
		},
		{
			MethodName: "Interaction",
			Handler:    _Plugin_Interaction_Handler,
		},
		{
			MethodName: "UIFramework",
			Handler:    _Plugin_UIFramework_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin.proto",
}
