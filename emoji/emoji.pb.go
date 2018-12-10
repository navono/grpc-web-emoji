// Code generated by protoc-gen-go. DO NOT EDIT.
// source: emoji.proto

package emoji

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type EmojizeRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmojizeRequest) Reset()         { *m = EmojizeRequest{} }
func (m *EmojizeRequest) String() string { return proto.CompactTextString(m) }
func (*EmojizeRequest) ProtoMessage()    {}
func (*EmojizeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_890939a4de8c1374, []int{0}
}

func (m *EmojizeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmojizeRequest.Unmarshal(m, b)
}
func (m *EmojizeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmojizeRequest.Marshal(b, m, deterministic)
}
func (m *EmojizeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmojizeRequest.Merge(m, src)
}
func (m *EmojizeRequest) XXX_Size() int {
	return xxx_messageInfo_EmojizeRequest.Size(m)
}
func (m *EmojizeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmojizeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmojizeRequest proto.InternalMessageInfo

func (m *EmojizeRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type EmojizeReply struct {
	EmojizedText         string   `protobuf:"bytes,1,opt,name=emojized_text,json=emojizedText,proto3" json:"emojized_text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmojizeReply) Reset()         { *m = EmojizeReply{} }
func (m *EmojizeReply) String() string { return proto.CompactTextString(m) }
func (*EmojizeReply) ProtoMessage()    {}
func (*EmojizeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_890939a4de8c1374, []int{1}
}

func (m *EmojizeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmojizeReply.Unmarshal(m, b)
}
func (m *EmojizeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmojizeReply.Marshal(b, m, deterministic)
}
func (m *EmojizeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmojizeReply.Merge(m, src)
}
func (m *EmojizeReply) XXX_Size() int {
	return xxx_messageInfo_EmojizeReply.Size(m)
}
func (m *EmojizeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EmojizeReply.DiscardUnknown(m)
}

var xxx_messageInfo_EmojizeReply proto.InternalMessageInfo

func (m *EmojizeReply) GetEmojizedText() string {
	if m != nil {
		return m.EmojizedText
	}
	return ""
}

func init() {
	proto.RegisterType((*EmojizeRequest)(nil), "emoji.EmojizeRequest")
	proto.RegisterType((*EmojizeReply)(nil), "emoji.EmojizeReply")
}

func init() { proto.RegisterFile("emoji.proto", fileDescriptor_890939a4de8c1374) }

var fileDescriptor_890939a4de8c1374 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcd, 0xcd, 0xcf,
	0xca, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x54, 0xb8, 0xf8, 0x5c,
	0x41, 0x8c, 0xaa, 0xd4, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0x92,
	0xd4, 0x8a, 0x12, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0xc9, 0x98, 0x8b, 0x07,
	0xae, 0xaa, 0x20, 0xa7, 0x52, 0x48, 0x99, 0x8b, 0x37, 0x15, 0xc2, 0x4f, 0x89, 0x47, 0x52, 0xcc,
	0x03, 0x13, 0x0c, 0x49, 0xad, 0x28, 0x31, 0x72, 0x85, 0x6a, 0x0a, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c,
	0x4e, 0x15, 0x32, 0xe5, 0x62, 0x87, 0x1a, 0x22, 0x24, 0xaa, 0x07, 0x71, 0x0a, 0xaa, 0xd5, 0x52,
	0xc2, 0xe8, 0xc2, 0x05, 0x39, 0x95, 0x49, 0x6c, 0x60, 0xf7, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x67, 0x48, 0xb0, 0x7c, 0xbe, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EmojiServiceClient is the client API for EmojiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmojiServiceClient interface {
	Emojize(ctx context.Context, in *EmojizeRequest, opts ...grpc.CallOption) (*EmojizeReply, error)
}

type emojiServiceClient struct {
	cc *grpc.ClientConn
}

func NewEmojiServiceClient(cc *grpc.ClientConn) EmojiServiceClient {
	return &emojiServiceClient{cc}
}

func (c *emojiServiceClient) Emojize(ctx context.Context, in *EmojizeRequest, opts ...grpc.CallOption) (*EmojizeReply, error) {
	out := new(EmojizeReply)
	err := c.cc.Invoke(ctx, "/emoji.EmojiService/Emojize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmojiServiceServer is the server API for EmojiService service.
type EmojiServiceServer interface {
	Emojize(context.Context, *EmojizeRequest) (*EmojizeReply, error)
}

func RegisterEmojiServiceServer(s *grpc.Server, srv EmojiServiceServer) {
	s.RegisterService(&_EmojiService_serviceDesc, srv)
}

func _EmojiService_Emojize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmojizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmojiServiceServer).Emojize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emoji.EmojiService/Emojize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmojiServiceServer).Emojize(ctx, req.(*EmojizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmojiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "emoji.EmojiService",
	HandlerType: (*EmojiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Emojize",
			Handler:    _EmojiService_Emojize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "emoji.proto",
}