// Code generated by protoc-gen-go. DO NOT EDIT.
// source: robotserver/robotserver.proto

package robotserver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ServoRequest struct {
	Angle                int64    `protobuf:"varint,1,opt,name=angle,proto3" json:"angle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServoRequest) Reset()         { *m = ServoRequest{} }
func (m *ServoRequest) String() string { return proto.CompactTextString(m) }
func (*ServoRequest) ProtoMessage()    {}
func (*ServoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_robotserver_addc601ae3ce8290, []int{0}
}
func (m *ServoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServoRequest.Unmarshal(m, b)
}
func (m *ServoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServoRequest.Marshal(b, m, deterministic)
}
func (dst *ServoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServoRequest.Merge(dst, src)
}
func (m *ServoRequest) XXX_Size() int {
	return xxx_messageInfo_ServoRequest.Size(m)
}
func (m *ServoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServoRequest proto.InternalMessageInfo

func (m *ServoRequest) GetAngle() int64 {
	if m != nil {
		return m.Angle
	}
	return 0
}

type LEDRequest struct {
	On                   bool     `protobuf:"varint,1,opt,name=on,proto3" json:"on,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LEDRequest) Reset()         { *m = LEDRequest{} }
func (m *LEDRequest) String() string { return proto.CompactTextString(m) }
func (*LEDRequest) ProtoMessage()    {}
func (*LEDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_robotserver_addc601ae3ce8290, []int{1}
}
func (m *LEDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LEDRequest.Unmarshal(m, b)
}
func (m *LEDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LEDRequest.Marshal(b, m, deterministic)
}
func (dst *LEDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LEDRequest.Merge(dst, src)
}
func (m *LEDRequest) XXX_Size() int {
	return xxx_messageInfo_LEDRequest.Size(m)
}
func (m *LEDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LEDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LEDRequest proto.InternalMessageInfo

func (m *LEDRequest) GetOn() bool {
	if m != nil {
		return m.On
	}
	return false
}

func init() {
	proto.RegisterType((*ServoRequest)(nil), "ServoRequest")
	proto.RegisterType((*LEDRequest)(nil), "LEDRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RobotServerClient is the client API for RobotServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RobotServerClient interface {
	SetServo(ctx context.Context, in *ServoRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SetLED(ctx context.Context, in *LEDRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type robotServerClient struct {
	cc *grpc.ClientConn
}

func NewRobotServerClient(cc *grpc.ClientConn) RobotServerClient {
	return &robotServerClient{cc}
}

func (c *robotServerClient) SetServo(ctx context.Context, in *ServoRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/RobotServer/SetServo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServerClient) SetLED(ctx context.Context, in *LEDRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/RobotServer/SetLED", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RobotServerServer is the server API for RobotServer service.
type RobotServerServer interface {
	SetServo(context.Context, *ServoRequest) (*empty.Empty, error)
	SetLED(context.Context, *LEDRequest) (*empty.Empty, error)
}

func RegisterRobotServerServer(s *grpc.Server, srv RobotServerServer) {
	s.RegisterService(&_RobotServer_serviceDesc, srv)
}

func _RobotServer_SetServo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServerServer).SetServo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RobotServer/SetServo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServerServer).SetServo(ctx, req.(*ServoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotServer_SetLED_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LEDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServerServer).SetLED(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RobotServer/SetLED",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServerServer).SetLED(ctx, req.(*LEDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RobotServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RobotServer",
	HandlerType: (*RobotServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetServo",
			Handler:    _RobotServer_SetServo_Handler,
		},
		{
			MethodName: "SetLED",
			Handler:    _RobotServer_SetLED_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "robotserver/robotserver.proto",
}

func init() {
	proto.RegisterFile("robotserver/robotserver.proto", fileDescriptor_robotserver_addc601ae3ce8290)
}

var fileDescriptor_robotserver_addc601ae3ce8290 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0xca, 0x4f, 0xca,
	0x2f, 0x29, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x47, 0x62, 0xeb, 0x15, 0x14, 0xe5, 0x97, 0xe4,
	0x4b, 0x49, 0xa7, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x83, 0x79, 0x49, 0xa5, 0x69, 0xfa, 0xa9,
	0xb9, 0x05, 0x25, 0x95, 0x10, 0x49, 0x25, 0x15, 0x2e, 0x9e, 0xe0, 0xd4, 0xa2, 0xb2, 0xfc, 0xa0,
	0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x11, 0x2e, 0xd6, 0xc4, 0xbc, 0xf4, 0x9c, 0x54, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x08, 0x47, 0x49, 0x86, 0x8b, 0xcb, 0xc7, 0xd5, 0x05, 0xa6,
	0x86, 0x8f, 0x8b, 0x29, 0x3f, 0x0f, 0xac, 0x80, 0x23, 0x88, 0x29, 0x3f, 0xcf, 0xa8, 0x98, 0x8b,
	0x3b, 0x08, 0x64, 0x6b, 0x30, 0xd8, 0x56, 0x21, 0x63, 0x2e, 0x8e, 0xe0, 0x54, 0x30, 0x27, 0x5f,
	0x88, 0x57, 0x0f, 0xd9, 0x74, 0x29, 0x31, 0x3d, 0x88, 0x5b, 0xf4, 0x60, 0x6e, 0xd1, 0x73, 0x05,
	0xb9, 0x45, 0x89, 0x41, 0x48, 0x9f, 0x8b, 0x2d, 0x38, 0xb5, 0xc4, 0xc7, 0xd5, 0x45, 0x88, 0x5b,
	0x0f, 0x61, 0x15, 0x6e, 0x0d, 0x49, 0x6c, 0x60, 0x11, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x10, 0xd5, 0xc4, 0x5f, 0xfd, 0x00, 0x00, 0x00,
}
