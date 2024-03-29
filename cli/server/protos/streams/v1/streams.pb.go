//
//  Copyright 2020 Docker Compose CLI authors

//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at

//      http://www.apache.org/licenses/LICENSE-2.0

//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.7
// source: cli/server/protos/streams/v1/streams.proto

package v1

import (
	context "context"
	reflect "reflect"
	sync "sync"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IOStream int32

const (
	IOStream_STDIN  IOStream = 0
	IOStream_STDOUT IOStream = 1
	IOStream_STDERR IOStream = 2
)

// Enum value maps for IOStream.
var (
	IOStream_name = map[int32]string{
		0: "STDIN",
		1: "STDOUT",
		2: "STDERR",
	}
	IOStream_value = map[string]int32{
		"STDIN":  0,
		"STDOUT": 1,
		"STDERR": 2,
	}
)

func (x IOStream) Enum() *IOStream {
	p := new(IOStream)
	*p = x
	return p
}

func (x IOStream) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IOStream) Descriptor() protoreflect.EnumDescriptor {
	return file_cli_server_protos_streams_v1_streams_proto_enumTypes[0].Descriptor()
}

func (IOStream) Type() protoreflect.EnumType {
	return &file_cli_server_protos_streams_v1_streams_proto_enumTypes[0]
}

func (x IOStream) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IOStream.Descriptor instead.
func (IOStream) EnumDescriptor() ([]byte, []int) {
	return file_cli_server_protos_streams_v1_streams_proto_rawDescGZIP(), []int{0}
}

type BytesMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  IOStream `protobuf:"varint,1,opt,name=type,proto3,enum=com.docker.api.protos.streams.v1.IOStream" json:"type,omitempty"`
	Value []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BytesMessage) Reset() {
	*x = BytesMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cli_server_protos_streams_v1_streams_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BytesMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesMessage) ProtoMessage() {}

func (x *BytesMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cli_server_protos_streams_v1_streams_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BytesMessage.ProtoReflect.Descriptor instead.
func (*BytesMessage) Descriptor() ([]byte, []int) {
	return file_cli_server_protos_streams_v1_streams_proto_rawDescGZIP(), []int{0}
}

func (x *BytesMessage) GetType() IOStream {
	if x != nil {
		return x.Type
	}
	return IOStream_STDIN
}

func (x *BytesMessage) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type ResizeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width  uint32 `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height uint32 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *ResizeMessage) Reset() {
	*x = ResizeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cli_server_protos_streams_v1_streams_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResizeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResizeMessage) ProtoMessage() {}

func (x *ResizeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cli_server_protos_streams_v1_streams_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResizeMessage.ProtoReflect.Descriptor instead.
func (*ResizeMessage) Descriptor() ([]byte, []int) {
	return file_cli_server_protos_streams_v1_streams_proto_rawDescGZIP(), []int{1}
}

func (x *ResizeMessage) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ResizeMessage) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type ExitMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ExitMessage) Reset() {
	*x = ExitMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cli_server_protos_streams_v1_streams_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExitMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExitMessage) ProtoMessage() {}

func (x *ExitMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cli_server_protos_streams_v1_streams_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExitMessage.ProtoReflect.Descriptor instead.
func (*ExitMessage) Descriptor() ([]byte, []int) {
	return file_cli_server_protos_streams_v1_streams_proto_rawDescGZIP(), []int{2}
}

func (x *ExitMessage) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_cli_server_protos_streams_v1_streams_proto protoreflect.FileDescriptor

var file_cli_server_protos_streams_v1_streams_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6c, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x0c, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3e, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f,
	0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x4f, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x3d, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x25,
	0x0a, 0x0b, 0x45, 0x78, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x2d, 0x0a, 0x08, 0x49, 0x4f, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x44, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x53, 0x54, 0x44, 0x4f, 0x55, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x44, 0x45,
	0x52, 0x52, 0x10, 0x02, 0x32, 0x48, 0x0a, 0x09, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x12, 0x3b, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x28, 0x01, 0x30, 0x01, 0x42, 0x3f,
	0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2d, 0x63, 0x6c, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cli_server_protos_streams_v1_streams_proto_rawDescOnce sync.Once
	file_cli_server_protos_streams_v1_streams_proto_rawDescData = file_cli_server_protos_streams_v1_streams_proto_rawDesc
)

func file_cli_server_protos_streams_v1_streams_proto_rawDescGZIP() []byte {
	file_cli_server_protos_streams_v1_streams_proto_rawDescOnce.Do(func() {
		file_cli_server_protos_streams_v1_streams_proto_rawDescData = protoimpl.X.CompressGZIP(file_cli_server_protos_streams_v1_streams_proto_rawDescData)
	})
	return file_cli_server_protos_streams_v1_streams_proto_rawDescData
}

var file_cli_server_protos_streams_v1_streams_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cli_server_protos_streams_v1_streams_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cli_server_protos_streams_v1_streams_proto_goTypes = []interface{}{
	(IOStream)(0),         // 0: com.docker.api.protos.streams.v1.IOStream
	(*BytesMessage)(nil),  // 1: com.docker.api.protos.streams.v1.BytesMessage
	(*ResizeMessage)(nil), // 2: com.docker.api.protos.streams.v1.ResizeMessage
	(*ExitMessage)(nil),   // 3: com.docker.api.protos.streams.v1.ExitMessage
	(*anypb.Any)(nil),     // 4: google.protobuf.Any
}
var file_cli_server_protos_streams_v1_streams_proto_depIdxs = []int32{
	0, // 0: com.docker.api.protos.streams.v1.BytesMessage.type:type_name -> com.docker.api.protos.streams.v1.IOStream
	4, // 1: com.docker.api.protos.streams.v1.Streaming.NewStream:input_type -> google.protobuf.Any
	4, // 2: com.docker.api.protos.streams.v1.Streaming.NewStream:output_type -> google.protobuf.Any
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cli_server_protos_streams_v1_streams_proto_init() }
func file_cli_server_protos_streams_v1_streams_proto_init() {
	if File_cli_server_protos_streams_v1_streams_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cli_server_protos_streams_v1_streams_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BytesMessage); i {
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
		file_cli_server_protos_streams_v1_streams_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResizeMessage); i {
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
		file_cli_server_protos_streams_v1_streams_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExitMessage); i {
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
			RawDescriptor: file_cli_server_protos_streams_v1_streams_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cli_server_protos_streams_v1_streams_proto_goTypes,
		DependencyIndexes: file_cli_server_protos_streams_v1_streams_proto_depIdxs,
		EnumInfos:         file_cli_server_protos_streams_v1_streams_proto_enumTypes,
		MessageInfos:      file_cli_server_protos_streams_v1_streams_proto_msgTypes,
	}.Build()
	File_cli_server_protos_streams_v1_streams_proto = out.File
	file_cli_server_protos_streams_v1_streams_proto_rawDesc = nil
	file_cli_server_protos_streams_v1_streams_proto_goTypes = nil
	file_cli_server_protos_streams_v1_streams_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StreamingClient is the client API for Streaming service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamingClient interface {
	NewStream(ctx context.Context, opts ...grpc.CallOption) (Streaming_NewStreamClient, error)
}

type streamingClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamingClient(cc grpc.ClientConnInterface) StreamingClient {
	return &streamingClient{cc}
}

func (c *streamingClient) NewStream(ctx context.Context, opts ...grpc.CallOption) (Streaming_NewStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Streaming_serviceDesc.Streams[0], "/com.docker.api.protos.streams.v1.Streaming/NewStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingNewStreamClient{stream}
	return x, nil
}

type Streaming_NewStreamClient interface {
	Send(*anypb.Any) error
	Recv() (*anypb.Any, error)
	grpc.ClientStream
}

type streamingNewStreamClient struct {
	grpc.ClientStream
}

func (x *streamingNewStreamClient) Send(m *anypb.Any) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamingNewStreamClient) Recv() (*anypb.Any, error) {
	m := new(anypb.Any)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingServer is the server API for Streaming service.
type StreamingServer interface {
	NewStream(Streaming_NewStreamServer) error
}

// UnimplementedStreamingServer can be embedded to have forward compatible implementations.
type UnimplementedStreamingServer struct {
}

func (*UnimplementedStreamingServer) NewStream(Streaming_NewStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method NewStream not implemented")
}

func RegisterStreamingServer(s *grpc.Server, srv StreamingServer) {
	s.RegisterService(&_Streaming_serviceDesc, srv)
}

func _Streaming_NewStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamingServer).NewStream(&streamingNewStreamServer{stream})
}

type Streaming_NewStreamServer interface {
	Send(*anypb.Any) error
	Recv() (*anypb.Any, error)
	grpc.ServerStream
}

type streamingNewStreamServer struct {
	grpc.ServerStream
}

func (x *streamingNewStreamServer) Send(m *anypb.Any) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamingNewStreamServer) Recv() (*anypb.Any, error) {
	m := new(anypb.Any)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Streaming_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.docker.api.protos.streams.v1.Streaming",
	HandlerType: (*StreamingServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NewStream",
			Handler:       _Streaming_NewStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cli/server/protos/streams/v1/streams.proto",
}
