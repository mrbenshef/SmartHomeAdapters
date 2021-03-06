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
	RobotId              string   `protobuf:"bytes,2,opt,name=robotId,proto3" json:"robotId,omitempty"`
	Angle                int64    `protobuf:"varint,1,opt,name=angle,proto3" json:"angle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServoRequest) Reset()         { *m = ServoRequest{} }
func (m *ServoRequest) String() string { return proto.CompactTextString(m) }
func (*ServoRequest) ProtoMessage()    {}
func (*ServoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_robotserver_c4d3f1f89a56371f, []int{0}
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

func (m *ServoRequest) GetRobotId() string {
	if m != nil {
		return m.RobotId
	}
	return ""
}

func (m *ServoRequest) GetAngle() int64 {
	if m != nil {
		return m.Angle
	}
	return 0
}

type LEDRequest struct {
	RobotId              string   `protobuf:"bytes,2,opt,name=robotId,proto3" json:"robotId,omitempty"`
	On                   bool     `protobuf:"varint,1,opt,name=on,proto3" json:"on,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LEDRequest) Reset()         { *m = LEDRequest{} }
func (m *LEDRequest) String() string { return proto.CompactTextString(m) }
func (*LEDRequest) ProtoMessage()    {}
func (*LEDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_robotserver_c4d3f1f89a56371f, []int{1}
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

func (m *LEDRequest) GetRobotId() string {
	if m != nil {
		return m.RobotId
	}
	return ""
}

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
	proto.RegisterFile("robotserver/robotserver.proto", fileDescriptor_robotserver_c4d3f1f89a56371f)
}

var fileDescriptor_robotserver_c4d3f1f89a56371f = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0xca, 0x4f, 0xca,
	0x2f, 0x29, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x47, 0x62, 0xeb, 0x15, 0x14, 0xe5, 0x97, 0xe4,
	0x4b, 0x49, 0xa7, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x83, 0x79, 0x49, 0xa5, 0x69, 0xfa, 0xa9,
	0xb9, 0x05, 0x25, 0x95, 0x10, 0x49, 0x25, 0x3b, 0x2e, 0x9e, 0xe0, 0xd4, 0xa2, 0xb2, 0xfc, 0xa0,
	0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x09, 0x2e, 0x76, 0xb0, 0x09, 0x9e, 0x29, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x90, 0x08, 0x17, 0x6b, 0x62, 0x5e, 0x7a, 0x4e, 0xaa,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x84, 0xa3, 0x64, 0xc6, 0xc5, 0xe5, 0xe3, 0xea, 0x42,
	0x58, 0x37, 0x1f, 0x17, 0x53, 0x7e, 0x1e, 0x58, 0x2b, 0x47, 0x10, 0x53, 0x7e, 0x9e, 0x51, 0x31,
	0x17, 0x77, 0x10, 0x48, 0x2a, 0x18, 0xec, 0x52, 0x21, 0x63, 0x2e, 0x8e, 0xe0, 0x54, 0x30, 0x27,
	0x5f, 0x88, 0x57, 0x0f, 0xd9, 0x45, 0x52, 0x62, 0x7a, 0x10, 0xf7, 0xeb, 0xc1, 0xdc, 0xaf, 0xe7,
	0x0a, 0x72, 0xbf, 0x12, 0x83, 0x90, 0x3e, 0x17, 0x5b, 0x70, 0x6a, 0x89, 0x8f, 0xab, 0x8b, 0x10,
	0xb7, 0x1e, 0xc2, 0x11, 0xb8, 0x35, 0x24, 0xb1, 0x81, 0x45, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xfe, 0x68, 0x51, 0x93, 0x31, 0x01, 0x00, 0x00,
}
