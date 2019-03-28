// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/grpc.proto

package grpc

import (
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoregistry "github.com/golang/protobuf/v2/reflect/protoregistry"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	sync "sync"
)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) ProtoReflect() protoreflect.Message {
	return xxx_File_grpc_grpc_proto_messageTypes[0].MessageOf(m)
}
func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return protoimpl.X.MessageStringOf(m) }
func (*Request) ProtoMessage()    {}

// Deprecated: Use Request.ProtoReflect.Type instead.
func (*Request) Descriptor() ([]byte, []int) {
	return xxx_File_grpc_grpc_proto_rawDescGZIP(), []int{0}
}

type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) ProtoReflect() protoreflect.Message {
	return xxx_File_grpc_grpc_proto_messageTypes[1].MessageOf(m)
}
func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return protoimpl.X.MessageStringOf(m) }
func (*Response) ProtoMessage()    {}

// Deprecated: Use Response.ProtoReflect.Type instead.
func (*Response) Descriptor() ([]byte, []int) {
	return xxx_File_grpc_grpc_proto_rawDescGZIP(), []int{1}
}

var File_grpc_grpc_proto protoreflect.FileDescriptor

var xxx_File_grpc_grpc_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x22, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x0a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc9, 0x02,
	0x0a, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49,
	0x0a, 0x0a, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x64, 0x6f, 0x77,
	0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x0d, 0x75,
	0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x4c, 0x0a, 0x09, 0x62,
	0x69, 0x64, 0x69, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	xxx_File_grpc_grpc_proto_rawDesc_once sync.Once
	xxx_File_grpc_grpc_proto_rawDesc_data = xxx_File_grpc_grpc_proto_rawDesc
)

func xxx_File_grpc_grpc_proto_rawDescGZIP() []byte {
	xxx_File_grpc_grpc_proto_rawDesc_once.Do(func() {
		xxx_File_grpc_grpc_proto_rawDesc_data = protoimpl.X.CompressGZIP(xxx_File_grpc_grpc_proto_rawDesc_data)
	})
	return xxx_File_grpc_grpc_proto_rawDesc_data
}

var xxx_File_grpc_grpc_proto_messageTypes = make([]protoimpl.MessageType, 2)
var xxx_File_grpc_grpc_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: goproto.protoc.grpc.Request
	(*Response)(nil), // 1: goproto.protoc.grpc.Response
}
var xxx_File_grpc_grpc_proto_depIdxs = []int32{
	0, // goproto.protoc.grpc.test_service.unary_call:input_type -> goproto.protoc.grpc.Request
	1, // goproto.protoc.grpc.test_service.unary_call:output_type -> goproto.protoc.grpc.Response
	0, // goproto.protoc.grpc.test_service.downstream_call:input_type -> goproto.protoc.grpc.Request
	1, // goproto.protoc.grpc.test_service.downstream_call:output_type -> goproto.protoc.grpc.Response
	0, // goproto.protoc.grpc.test_service.upstream_call:input_type -> goproto.protoc.grpc.Request
	1, // goproto.protoc.grpc.test_service.upstream_call:output_type -> goproto.protoc.grpc.Response
	0, // goproto.protoc.grpc.test_service.bidi_call:input_type -> goproto.protoc.grpc.Request
	1, // goproto.protoc.grpc.test_service.bidi_call:output_type -> goproto.protoc.grpc.Response
}

func init() { xxx_File_grpc_grpc_proto_init() }
func xxx_File_grpc_grpc_proto_init() {
	if File_grpc_grpc_proto != nil {
		return
	}
	File_grpc_grpc_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_grpc_grpc_proto_rawDesc,
		GoTypes:            xxx_File_grpc_grpc_proto_goTypes,
		DependencyIndexes:  xxx_File_grpc_grpc_proto_depIdxs,
		MessageOutputTypes: xxx_File_grpc_grpc_proto_messageTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	xxx_File_grpc_grpc_proto_rawDesc = nil
	xxx_File_grpc_grpc_proto_goTypes = nil
	xxx_File_grpc_grpc_proto_depIdxs = nil
}
