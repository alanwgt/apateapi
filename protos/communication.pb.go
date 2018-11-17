// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/communication.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ServerResponse_Status int32

const (
	ServerResponse_Ok    ServerResponse_Status = 0
	ServerResponse_ERROR ServerResponse_Status = 1
)

var ServerResponse_Status_name = map[int32]string{
	0: "Ok",
	1: "ERROR",
}

var ServerResponse_Status_value = map[string]int32{
	"Ok":    0,
	"ERROR": 1,
}

func (x ServerResponse_Status) String() string {
	return proto.EnumName(ServerResponse_Status_name, int32(x))
}

func (ServerResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a9db6ee45f90bde9, []int{0, 0}
}

// the request type
type DeviceRequest_Type int32

const (
	DeviceRequest_Handshake DeviceRequest_Type = 0
)

var DeviceRequest_Type_name = map[int32]string{
	0: "Handshake",
}

var DeviceRequest_Type_value = map[string]int32{
	"Handshake": 0,
}

func (x DeviceRequest_Type) String() string {
	return proto.EnumName(DeviceRequest_Type_name, int32(x))
}

func (DeviceRequest_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a9db6ee45f90bde9, []int{1, 0}
}

type ServerResponse struct {
	Status               ServerResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=protos.ServerResponse_Status" json:"status,omitempty"`
	StatusCode           int32                 `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message              string                `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ServerResponse) Reset()         { *m = ServerResponse{} }
func (m *ServerResponse) String() string { return proto.CompactTextString(m) }
func (*ServerResponse) ProtoMessage()    {}
func (*ServerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9db6ee45f90bde9, []int{0}
}

func (m *ServerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerResponse.Unmarshal(m, b)
}
func (m *ServerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerResponse.Marshal(b, m, deterministic)
}
func (m *ServerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerResponse.Merge(m, src)
}
func (m *ServerResponse) XXX_Size() int {
	return xxx_messageInfo_ServerResponse.Size(m)
}
func (m *ServerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerResponse proto.InternalMessageInfo

func (m *ServerResponse) GetStatus() ServerResponse_Status {
	if m != nil {
		return m.Status
	}
	return ServerResponse_Ok
}

func (m *ServerResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *ServerResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DeviceRequest struct {
	Type     DeviceRequest_Type `protobuf:"varint,1,opt,name=type,proto3,enum=protos.DeviceRequest_Type" json:"type,omitempty"`
	Username string             `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// the payload will be always be an encrypted message encoded in base64
	Paylod string `protobuf:"bytes,3,opt,name=paylod,proto3" json:"paylod,omitempty"`
	// the nonce is also encoded in base64
	Nonce                string   `protobuf:"bytes,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceRequest) Reset()         { *m = DeviceRequest{} }
func (m *DeviceRequest) String() string { return proto.CompactTextString(m) }
func (*DeviceRequest) ProtoMessage()    {}
func (*DeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9db6ee45f90bde9, []int{1}
}

func (m *DeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceRequest.Unmarshal(m, b)
}
func (m *DeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceRequest.Marshal(b, m, deterministic)
}
func (m *DeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceRequest.Merge(m, src)
}
func (m *DeviceRequest) XXX_Size() int {
	return xxx_messageInfo_DeviceRequest.Size(m)
}
func (m *DeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceRequest proto.InternalMessageInfo

func (m *DeviceRequest) GetType() DeviceRequest_Type {
	if m != nil {
		return m.Type
	}
	return DeviceRequest_Handshake
}

func (m *DeviceRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *DeviceRequest) GetPaylod() string {
	if m != nil {
		return m.Paylod
	}
	return ""
}

func (m *DeviceRequest) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func init() {
	proto.RegisterEnum("protos.ServerResponse_Status", ServerResponse_Status_name, ServerResponse_Status_value)
	proto.RegisterEnum("protos.DeviceRequest_Type", DeviceRequest_Type_name, DeviceRequest_Type_value)
	proto.RegisterType((*ServerResponse)(nil), "protos.ServerResponse")
	proto.RegisterType((*DeviceRequest)(nil), "protos.DeviceRequest")
}

func init() { proto.RegisterFile("protos/communication.proto", fileDescriptor_a9db6ee45f90bde9) }

var fileDescriptor_a9db6ee45f90bde9 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x51, 0x4a, 0xf4, 0x30,
	0x14, 0x85, 0x27, 0xf3, 0xb7, 0xfd, 0xed, 0x85, 0x19, 0xca, 0x45, 0xa5, 0x54, 0x94, 0xd2, 0xa7,
	0x3e, 0x55, 0x50, 0x5c, 0x81, 0x0a, 0xbe, 0x0d, 0x64, 0xdc, 0x40, 0x6c, 0x2f, 0x5a, 0xc6, 0x26,
	0xb1, 0x37, 0x1d, 0xe8, 0x66, 0xc4, 0xa5, 0x8a, 0x4d, 0x47, 0x9c, 0xb7, 0x9c, 0x93, 0x2f, 0x87,
	0x8f, 0x40, 0x66, 0x7b, 0xe3, 0x0c, 0x5f, 0xd7, 0xa6, 0xeb, 0x06, 0xdd, 0xd6, 0xca, 0xb5, 0x46,
	0x57, 0x53, 0x89, 0x91, 0xbf, 0x2b, 0x3e, 0x05, 0xac, 0xb7, 0xd4, 0xef, 0xa9, 0x97, 0xc4, 0xd6,
	0x68, 0x26, 0xbc, 0x83, 0x88, 0x9d, 0x72, 0x03, 0xa7, 0x22, 0x17, 0xe5, 0xfa, 0xe6, 0xd2, 0x3f,
	0xe1, 0xea, 0x98, 0xab, 0xb6, 0x13, 0x24, 0x67, 0x18, 0xaf, 0x00, 0xfc, 0xe9, 0xde, 0x34, 0x94,
	0x2e, 0x73, 0x51, 0x86, 0xf2, 0x4f, 0x83, 0x29, 0xfc, 0xef, 0x88, 0x59, 0xbd, 0x52, 0xfa, 0x2f,
	0x17, 0x65, 0x2c, 0x0f, 0xb1, 0xb8, 0x80, 0xc8, 0x6f, 0x61, 0x04, 0xcb, 0xcd, 0x2e, 0x59, 0x60,
	0x0c, 0xe1, 0xa3, 0x94, 0x1b, 0x99, 0x88, 0xe2, 0x4b, 0xc0, 0xea, 0x81, 0xf6, 0x6d, 0x4d, 0x92,
	0x3e, 0x06, 0x62, 0x87, 0x15, 0x04, 0x6e, 0xb4, 0x34, 0xdb, 0x65, 0x07, 0xbb, 0x23, 0xa8, 0x7a,
	0x1e, 0x2d, 0xc9, 0x89, 0xc3, 0x0c, 0x4e, 0x06, 0xa6, 0x5e, 0xab, 0xce, 0x6b, 0xc5, 0xf2, 0x37,
	0xe3, 0x39, 0x44, 0x56, 0x8d, 0xef, 0xa6, 0x99, 0x9d, 0xe6, 0x84, 0xa7, 0x10, 0x6a, 0xa3, 0x6b,
	0x4a, 0x83, 0xa9, 0xf6, 0xa1, 0x38, 0x83, 0xe0, 0x67, 0x17, 0x57, 0x10, 0x3f, 0x29, 0xdd, 0xf0,
	0x9b, 0xda, 0x51, 0xb2, 0x78, 0xf1, 0x7f, 0x79, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x95, 0x3e,
	0x3a, 0x7e, 0x70, 0x01, 0x00, 0x00,
}
