// Code generated by protoc-gen-go. DO NOT EDIT.
// source: SMSService.proto

package com_salty_protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SMSOperationType int32

const (
	SMSOperationType_UNDEFINED      SMSOperationType = 0
	SMSOperationType_REGISTER       SMSOperationType = 1
	SMSOperationType_LOGIN          SMSOperationType = 2
	SMSOperationType_RESET_PASSWORD SMSOperationType = 3
)

var SMSOperationType_name = map[int32]string{
	0: "UNDEFINED",
	1: "REGISTER",
	2: "LOGIN",
	3: "RESET_PASSWORD",
}

var SMSOperationType_value = map[string]int32{
	"UNDEFINED":      0,
	"REGISTER":       1,
	"LOGIN":          2,
	"RESET_PASSWORD": 3,
}

func (x SMSOperationType) String() string {
	return proto.EnumName(SMSOperationType_name, int32(x))
}

func (SMSOperationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d18f85bca1ab810, []int{0}
}

type ObtainTelephoneSMSCodeReq struct {
	OperationType        SMSOperationType `protobuf:"varint,1,opt,name=operationType,proto3,enum=com.salty.protos.SMSOperationType" json:"operationType,omitempty"`
	Telephone            string           `protobuf:"bytes,2,opt,name=telephone,proto3" json:"telephone,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ObtainTelephoneSMSCodeReq) Reset()         { *m = ObtainTelephoneSMSCodeReq{} }
func (m *ObtainTelephoneSMSCodeReq) String() string { return proto.CompactTextString(m) }
func (*ObtainTelephoneSMSCodeReq) ProtoMessage()    {}
func (*ObtainTelephoneSMSCodeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d18f85bca1ab810, []int{0}
}

func (m *ObtainTelephoneSMSCodeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObtainTelephoneSMSCodeReq.Unmarshal(m, b)
}
func (m *ObtainTelephoneSMSCodeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObtainTelephoneSMSCodeReq.Marshal(b, m, deterministic)
}
func (m *ObtainTelephoneSMSCodeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObtainTelephoneSMSCodeReq.Merge(m, src)
}
func (m *ObtainTelephoneSMSCodeReq) XXX_Size() int {
	return xxx_messageInfo_ObtainTelephoneSMSCodeReq.Size(m)
}
func (m *ObtainTelephoneSMSCodeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ObtainTelephoneSMSCodeReq.DiscardUnknown(m)
}

var xxx_messageInfo_ObtainTelephoneSMSCodeReq proto.InternalMessageInfo

func (m *ObtainTelephoneSMSCodeReq) GetOperationType() SMSOperationType {
	if m != nil {
		return m.OperationType
	}
	return SMSOperationType_UNDEFINED
}

func (m *ObtainTelephoneSMSCodeReq) GetTelephone() string {
	if m != nil {
		return m.Telephone
	}
	return ""
}

type ObtainTelephoneSMSCodeResp struct {
	SmsCodeLength        int32    `protobuf:"varint,1,opt,name=smsCodeLength,proto3" json:"smsCodeLength,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObtainTelephoneSMSCodeResp) Reset()         { *m = ObtainTelephoneSMSCodeResp{} }
func (m *ObtainTelephoneSMSCodeResp) String() string { return proto.CompactTextString(m) }
func (*ObtainTelephoneSMSCodeResp) ProtoMessage()    {}
func (*ObtainTelephoneSMSCodeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d18f85bca1ab810, []int{1}
}

func (m *ObtainTelephoneSMSCodeResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObtainTelephoneSMSCodeResp.Unmarshal(m, b)
}
func (m *ObtainTelephoneSMSCodeResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObtainTelephoneSMSCodeResp.Marshal(b, m, deterministic)
}
func (m *ObtainTelephoneSMSCodeResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObtainTelephoneSMSCodeResp.Merge(m, src)
}
func (m *ObtainTelephoneSMSCodeResp) XXX_Size() int {
	return xxx_messageInfo_ObtainTelephoneSMSCodeResp.Size(m)
}
func (m *ObtainTelephoneSMSCodeResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ObtainTelephoneSMSCodeResp.DiscardUnknown(m)
}

var xxx_messageInfo_ObtainTelephoneSMSCodeResp proto.InternalMessageInfo

func (m *ObtainTelephoneSMSCodeResp) GetSmsCodeLength() int32 {
	if m != nil {
		return m.SmsCodeLength
	}
	return 0
}

type VerifyTelephoneSMSCodeReq struct {
	OperationType        SMSOperationType `protobuf:"varint,1,opt,name=operationType,proto3,enum=com.salty.protos.SMSOperationType" json:"operationType,omitempty"`
	Telephone            string           `protobuf:"bytes,2,opt,name=telephone,proto3" json:"telephone,omitempty"`
	SmsCode              string           `protobuf:"bytes,3,opt,name=smsCode,proto3" json:"smsCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *VerifyTelephoneSMSCodeReq) Reset()         { *m = VerifyTelephoneSMSCodeReq{} }
func (m *VerifyTelephoneSMSCodeReq) String() string { return proto.CompactTextString(m) }
func (*VerifyTelephoneSMSCodeReq) ProtoMessage()    {}
func (*VerifyTelephoneSMSCodeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d18f85bca1ab810, []int{2}
}

func (m *VerifyTelephoneSMSCodeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyTelephoneSMSCodeReq.Unmarshal(m, b)
}
func (m *VerifyTelephoneSMSCodeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyTelephoneSMSCodeReq.Marshal(b, m, deterministic)
}
func (m *VerifyTelephoneSMSCodeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyTelephoneSMSCodeReq.Merge(m, src)
}
func (m *VerifyTelephoneSMSCodeReq) XXX_Size() int {
	return xxx_messageInfo_VerifyTelephoneSMSCodeReq.Size(m)
}
func (m *VerifyTelephoneSMSCodeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyTelephoneSMSCodeReq.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyTelephoneSMSCodeReq proto.InternalMessageInfo

func (m *VerifyTelephoneSMSCodeReq) GetOperationType() SMSOperationType {
	if m != nil {
		return m.OperationType
	}
	return SMSOperationType_UNDEFINED
}

func (m *VerifyTelephoneSMSCodeReq) GetTelephone() string {
	if m != nil {
		return m.Telephone
	}
	return ""
}

func (m *VerifyTelephoneSMSCodeReq) GetSmsCode() string {
	if m != nil {
		return m.SmsCode
	}
	return ""
}

type VerifyTelephoneSMSCodeResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyTelephoneSMSCodeResp) Reset()         { *m = VerifyTelephoneSMSCodeResp{} }
func (m *VerifyTelephoneSMSCodeResp) String() string { return proto.CompactTextString(m) }
func (*VerifyTelephoneSMSCodeResp) ProtoMessage()    {}
func (*VerifyTelephoneSMSCodeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d18f85bca1ab810, []int{3}
}

func (m *VerifyTelephoneSMSCodeResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyTelephoneSMSCodeResp.Unmarshal(m, b)
}
func (m *VerifyTelephoneSMSCodeResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyTelephoneSMSCodeResp.Marshal(b, m, deterministic)
}
func (m *VerifyTelephoneSMSCodeResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyTelephoneSMSCodeResp.Merge(m, src)
}
func (m *VerifyTelephoneSMSCodeResp) XXX_Size() int {
	return xxx_messageInfo_VerifyTelephoneSMSCodeResp.Size(m)
}
func (m *VerifyTelephoneSMSCodeResp) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyTelephoneSMSCodeResp.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyTelephoneSMSCodeResp proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("com.salty.protos.SMSOperationType", SMSOperationType_name, SMSOperationType_value)
	proto.RegisterType((*ObtainTelephoneSMSCodeReq)(nil), "com.salty.protos.ObtainTelephoneSMSCodeReq")
	proto.RegisterType((*ObtainTelephoneSMSCodeResp)(nil), "com.salty.protos.ObtainTelephoneSMSCodeResp")
	proto.RegisterType((*VerifyTelephoneSMSCodeReq)(nil), "com.salty.protos.VerifyTelephoneSMSCodeReq")
	proto.RegisterType((*VerifyTelephoneSMSCodeResp)(nil), "com.salty.protos.VerifyTelephoneSMSCodeResp")
}

func init() { proto.RegisterFile("SMSService.proto", fileDescriptor_0d18f85bca1ab810) }

var fileDescriptor_0d18f85bca1ab810 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0xc1, 0x4e, 0xfa, 0x40,
	0x10, 0xc6, 0x59, 0x08, 0xff, 0xbf, 0x9d, 0x00, 0x69, 0xe6, 0x60, 0x4a, 0xc3, 0x81, 0x34, 0x1e,
	0x88, 0x87, 0x1e, 0xf0, 0x09, 0x44, 0x2a, 0x92, 0x40, 0x8b, 0xbb, 0x55, 0x8f, 0x06, 0x70, 0x94,
	0x26, 0xd0, 0x5d, 0xba, 0x1b, 0x13, 0xce, 0x3e, 0x87, 0x2f, 0xe1, 0x13, 0x1a, 0xaa, 0x04, 0x51,
	0xb8, 0x78, 0xf1, 0xf8, 0x7d, 0xb3, 0xfb, 0xcd, 0x6f, 0x66, 0x17, 0x6c, 0x31, 0x14, 0x82, 0xb2,
	0xe7, 0x64, 0x4a, 0xbe, 0xca, 0xa4, 0x91, 0x68, 0x4f, 0xe5, 0xc2, 0xd7, 0xe3, 0xb9, 0x59, 0x7d,
	0x18, 0xda, 0xad, 0x50, 0x6a, 0x92, 0x8d, 0xf4, 0x5e, 0x18, 0xd4, 0xa3, 0x89, 0x19, 0x27, 0x69,
	0x4c, 0x73, 0x52, 0x33, 0x99, 0x92, 0x18, 0x8a, 0x0b, 0xf9, 0x40, 0x9c, 0x96, 0x78, 0x05, 0x55,
	0xa9, 0x28, 0x1b, 0x9b, 0x44, 0xa6, 0xf1, 0x4a, 0x91, 0xc3, 0x9a, 0xac, 0x55, 0x6b, 0x7b, 0xfe,
	0xf7, 0x54, 0x5f, 0x0c, 0x45, 0xf4, 0xf5, 0x24, 0xdf, 0xbd, 0x88, 0x0d, 0xb0, 0xcc, 0xa6, 0x81,
	0x53, 0x6c, 0xb2, 0x96, 0xc5, 0xb7, 0x86, 0xd7, 0x01, 0xf7, 0x10, 0x84, 0x56, 0x78, 0x02, 0x55,
	0xbd, 0xd0, 0x6b, 0x39, 0xa0, 0xf4, 0xc9, 0xcc, 0x72, 0x8a, 0x32, 0xdf, 0x35, 0xbd, 0x57, 0x06,
	0xf5, 0x5b, 0xca, 0x92, 0xc7, 0xd5, 0x1f, 0x4e, 0x82, 0x0e, 0xfc, 0xff, 0xc4, 0x72, 0x4a, 0x79,
	0x6d, 0x23, 0xbd, 0x06, 0xb8, 0x87, 0xf0, 0xb4, 0x3a, 0x0d, 0xf3, 0xb7, 0xdb, 0x69, 0x8c, 0x55,
	0xb0, 0x6e, 0xc2, 0x6e, 0x70, 0xd9, 0x0f, 0x83, 0xae, 0x5d, 0xc0, 0x0a, 0x1c, 0xf1, 0xa0, 0xd7,
	0x17, 0x71, 0xc0, 0x6d, 0x86, 0x16, 0x94, 0x07, 0x51, 0xaf, 0x1f, 0xda, 0x45, 0x44, 0xa8, 0xf1,
	0x40, 0x04, 0xf1, 0xfd, 0xe8, 0x5c, 0x88, 0xbb, 0x88, 0x77, 0xed, 0x52, 0xfb, 0x8d, 0x01, 0x6c,
	0x3f, 0x03, 0x5e, 0xc3, 0xf1, 0xfe, 0x05, 0x63, 0xfd, 0xe7, 0x06, 0x7a, 0x99, 0x9a, 0x72, 0x5a,
	0xba, 0xee, 0xa1, 0x92, 0x56, 0x5e, 0x61, 0x1d, 0xb9, 0x7f, 0x9e, 0x5f, 0x47, 0x76, 0x8a, 0x23,
	0x36, 0xf9, 0x97, 0xdb, 0x67, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x61, 0x4e, 0x3e, 0xe1, 0xcb,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SMSServiceClient is the client API for SMSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SMSServiceClient interface {
	//获取验证码
	ObtainTelephoneSMSCode(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error)
	//校验验证码
	VerifyTelephoneSMSCode(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error)
}

type sMSServiceClient struct {
	cc *grpc.ClientConn
}

func NewSMSServiceClient(cc *grpc.ClientConn) SMSServiceClient {
	return &sMSServiceClient{cc}
}

func (c *sMSServiceClient) ObtainTelephoneSMSCode(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error) {
	out := new(GrpcResp)
	err := c.cc.Invoke(ctx, "/com.salty.protos.SMSService/ObtainTelephoneSMSCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sMSServiceClient) VerifyTelephoneSMSCode(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error) {
	out := new(GrpcResp)
	err := c.cc.Invoke(ctx, "/com.salty.protos.SMSService/VerifyTelephoneSMSCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SMSServiceServer is the server API for SMSService service.
type SMSServiceServer interface {
	//获取验证码
	ObtainTelephoneSMSCode(context.Context, *GrpcReq) (*GrpcResp, error)
	//校验验证码
	VerifyTelephoneSMSCode(context.Context, *GrpcReq) (*GrpcResp, error)
}

// UnimplementedSMSServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSMSServiceServer struct {
}

func (*UnimplementedSMSServiceServer) ObtainTelephoneSMSCode(ctx context.Context, req *GrpcReq) (*GrpcResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObtainTelephoneSMSCode not implemented")
}
func (*UnimplementedSMSServiceServer) VerifyTelephoneSMSCode(ctx context.Context, req *GrpcReq) (*GrpcResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyTelephoneSMSCode not implemented")
}

func RegisterSMSServiceServer(s *grpc.Server, srv SMSServiceServer) {
	s.RegisterService(&_SMSService_serviceDesc, srv)
}

func _SMSService_ObtainTelephoneSMSCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrpcReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSServiceServer).ObtainTelephoneSMSCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.salty.protos.SMSService/ObtainTelephoneSMSCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSServiceServer).ObtainTelephoneSMSCode(ctx, req.(*GrpcReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SMSService_VerifyTelephoneSMSCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrpcReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSServiceServer).VerifyTelephoneSMSCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.salty.protos.SMSService/VerifyTelephoneSMSCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSServiceServer).VerifyTelephoneSMSCode(ctx, req.(*GrpcReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _SMSService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.salty.protos.SMSService",
	HandlerType: (*SMSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ObtainTelephoneSMSCode",
			Handler:    _SMSService_ObtainTelephoneSMSCode_Handler,
		},
		{
			MethodName: "VerifyTelephoneSMSCode",
			Handler:    _SMSService_VerifyTelephoneSMSCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "SMSService.proto",
}
