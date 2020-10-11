// Code generated by protoc-gen-go. DO NOT EDIT.
// source: translate.proto

package model

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// GoTransReq is request message for GoTrans grpc
type GoTransReq struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Text                 string   `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoTransReq) Reset()         { *m = GoTransReq{} }
func (m *GoTransReq) String() string { return proto.CompactTextString(m) }
func (*GoTransReq) ProtoMessage()    {}
func (*GoTransReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_badecccca6d3291a, []int{0}
}

func (m *GoTransReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoTransReq.Unmarshal(m, b)
}
func (m *GoTransReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoTransReq.Marshal(b, m, deterministic)
}
func (m *GoTransReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoTransReq.Merge(m, src)
}
func (m *GoTransReq) XXX_Size() int {
	return xxx_messageInfo_GoTransReq.Size(m)
}
func (m *GoTransReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GoTransReq.DiscardUnknown(m)
}

var xxx_messageInfo_GoTransReq proto.InternalMessageInfo

func (m *GoTransReq) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *GoTransReq) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *GoTransReq) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type TranslateData struct {
	ErrorMessage         string   `protobuf:"bytes,1,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	TranslateText        string   `protobuf:"bytes,2,opt,name=translateText,proto3" json:"translateText,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranslateData) Reset()         { *m = TranslateData{} }
func (m *TranslateData) String() string { return proto.CompactTextString(m) }
func (*TranslateData) ProtoMessage()    {}
func (*TranslateData) Descriptor() ([]byte, []int) {
	return fileDescriptor_badecccca6d3291a, []int{1}
}

func (m *TranslateData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslateData.Unmarshal(m, b)
}
func (m *TranslateData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslateData.Marshal(b, m, deterministic)
}
func (m *TranslateData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslateData.Merge(m, src)
}
func (m *TranslateData) XXX_Size() int {
	return xxx_messageInfo_TranslateData.Size(m)
}
func (m *TranslateData) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslateData.DiscardUnknown(m)
}

var xxx_messageInfo_TranslateData proto.InternalMessageInfo

func (m *TranslateData) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *TranslateData) GetTranslateText() string {
	if m != nil {
		return m.TranslateText
	}
	return ""
}

// GoTransResp is response message for GoTrans grpc
type GoTransResp struct {
	IsSuccess            bool           `protobuf:"varint,1,opt,name=isSuccess,proto3" json:"isSuccess,omitempty"`
	TranslateData        *TranslateData `protobuf:"bytes,2,opt,name=translateData,proto3" json:"translateData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GoTransResp) Reset()         { *m = GoTransResp{} }
func (m *GoTransResp) String() string { return proto.CompactTextString(m) }
func (*GoTransResp) ProtoMessage()    {}
func (*GoTransResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_badecccca6d3291a, []int{2}
}

func (m *GoTransResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoTransResp.Unmarshal(m, b)
}
func (m *GoTransResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoTransResp.Marshal(b, m, deterministic)
}
func (m *GoTransResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoTransResp.Merge(m, src)
}
func (m *GoTransResp) XXX_Size() int {
	return xxx_messageInfo_GoTransResp.Size(m)
}
func (m *GoTransResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GoTransResp.DiscardUnknown(m)
}

var xxx_messageInfo_GoTransResp proto.InternalMessageInfo

func (m *GoTransResp) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func (m *GoTransResp) GetTranslateData() *TranslateData {
	if m != nil {
		return m.TranslateData
	}
	return nil
}

func init() {
	proto.RegisterType((*GoTransReq)(nil), "model.GoTransReq")
	proto.RegisterType((*TranslateData)(nil), "model.TranslateData")
	proto.RegisterType((*GoTransResp)(nil), "model.GoTransResp")
}

func init() { proto.RegisterFile("translate.proto", fileDescriptor_badecccca6d3291a) }

var fileDescriptor_badecccca6d3291a = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xbf, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x69, 0xf8, 0x21, 0x72, 0xa1, 0x20, 0x4e, 0x0c, 0x11, 0x62, 0x40, 0x16, 0x03, 0x53,
	0x86, 0x22, 0x31, 0xb0, 0x52, 0xa9, 0x13, 0x8b, 0xc9, 0xc2, 0x68, 0xc2, 0x51, 0x21, 0xb5, 0x9c,
	0xeb, 0xbb, 0x4a, 0xfc, 0xf9, 0xc8, 0x6e, 0x48, 0x70, 0xb7, 0xd3, 0x77, 0x4f, 0x9f, 0x9f, 0x0f,
	0x2e, 0x34, 0xb8, 0x6f, 0x59, 0x39, 0xa5, 0xc6, 0x07, 0x56, 0xc6, 0xe3, 0x35, 0x7f, 0xd0, 0xca,
	0xcc, 0x01, 0x16, 0xdc, 0xc6, 0x9d, 0xa5, 0x0d, 0x22, 0x1c, 0x7d, 0x06, 0x5e, 0xd7, 0x93, 0xdb,
	0xc9, 0x7d, 0x69, 0xd3, 0x8c, 0xe7, 0x50, 0x28, 0xd7, 0x45, 0x22, 0x85, 0x72, 0xcc, 0x28, 0xfd,
	0x68, 0x7d, 0xb8, 0xcb, 0xc4, 0xd9, 0xbc, 0xc1, 0xb4, 0xfd, 0xf3, 0xcf, 0x9d, 0x3a, 0x34, 0x70,
	0x46, 0x21, 0x70, 0x78, 0x21, 0x11, 0xb7, 0xa4, 0x5e, 0x98, 0x31, 0xbc, 0x83, 0xe9, 0x50, 0xaa,
	0x8d, 0xc6, 0xdd, 0x1b, 0x39, 0x34, 0x4b, 0xa8, 0x86, 0x82, 0xe2, 0xf1, 0x06, 0xca, 0x2f, 0x79,
	0xdd, 0x76, 0x1d, 0x89, 0x24, 0xeb, 0xa9, 0x1d, 0x01, 0x3e, 0xfd, 0x53, 0xc6, 0x1e, 0x49, 0x59,
	0xcd, 0xae, 0x9a, 0xf4, 0xd9, 0x26, 0xeb, 0x68, 0xf3, 0xe8, 0xec, 0x19, 0xca, 0x61, 0x8f, 0x8f,
	0x50, 0x59, 0xda, 0x6c, 0x49, 0x74, 0x11, 0x7c, 0x87, 0x97, 0xbd, 0x60, 0x3c, 0xd5, 0x35, 0xee,
	0x23, 0xf1, 0xe6, 0xe0, 0xfd, 0x24, 0x1d, 0xf7, 0xe1, 0x37, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x21,
	0x1b, 0x17, 0x6f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TranslateClient is the client API for Translate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TranslateClient interface {
	// RequestGrpc is grpc request for GoTrans
	RequestGrpc(ctx context.Context, in *GoTransReq, opts ...grpc.CallOption) (*GoTransResp, error)
}

type translateClient struct {
	cc *grpc.ClientConn
}

func NewTranslateClient(cc *grpc.ClientConn) TranslateClient {
	return &translateClient{cc}
}

func (c *translateClient) RequestGrpc(ctx context.Context, in *GoTransReq, opts ...grpc.CallOption) (*GoTransResp, error) {
	out := new(GoTransResp)
	err := c.cc.Invoke(ctx, "/model.Translate/RequestGrpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranslateServer is the server API for Translate service.
type TranslateServer interface {
	// RequestGrpc is grpc request for GoTrans
	RequestGrpc(context.Context, *GoTransReq) (*GoTransResp, error)
}

// UnimplementedTranslateServer can be embedded to have forward compatible implementations.
type UnimplementedTranslateServer struct {
}

func (*UnimplementedTranslateServer) RequestGrpc(ctx context.Context, req *GoTransReq) (*GoTransResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestGrpc not implemented")
}

func RegisterTranslateServer(s *grpc.Server, srv TranslateServer) {
	s.RegisterService(&_Translate_serviceDesc, srv)
}

func _Translate_RequestGrpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoTransReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslateServer).RequestGrpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Translate/RequestGrpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslateServer).RequestGrpc(ctx, req.(*GoTransReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Translate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.Translate",
	HandlerType: (*TranslateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestGrpc",
			Handler:    _Translate_RequestGrpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "translate.proto",
}
