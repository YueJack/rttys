// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rtty.proto

/*
Package rtty is a generated protocol buffer package.

It is generated from these files:
	rtty.proto

It has these top-level messages:
	RttyMessage
*/
package rtty

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RttyMessage_Type int32

const (
	RttyMessage_UNKNOWN  RttyMessage_Type = 0
	RttyMessage_LOGIN    RttyMessage_Type = 1
	RttyMessage_LOGINACK RttyMessage_Type = 2
	RttyMessage_LOGOUT   RttyMessage_Type = 3
	RttyMessage_TTY      RttyMessage_Type = 4
	RttyMessage_UPFILE   RttyMessage_Type = 5
	RttyMessage_DOWNFILE RttyMessage_Type = 6
	RttyMessage_COMMAND  RttyMessage_Type = 7
)

var RttyMessage_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "LOGIN",
	2: "LOGINACK",
	3: "LOGOUT",
	4: "TTY",
	5: "UPFILE",
	6: "DOWNFILE",
	7: "COMMAND",
}
var RttyMessage_Type_value = map[string]int32{
	"UNKNOWN":  0,
	"LOGIN":    1,
	"LOGINACK": 2,
	"LOGOUT":   3,
	"TTY":      4,
	"UPFILE":   5,
	"DOWNFILE": 6,
	"COMMAND":  7,
}

func (x RttyMessage_Type) String() string {
	return proto.EnumName(RttyMessage_Type_name, int32(x))
}
func (RttyMessage_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type RttyMessage_LoginCode int32

const (
	RttyMessage_OK      RttyMessage_LoginCode = 0
	RttyMessage_OFFLINE RttyMessage_LoginCode = 1
)

var RttyMessage_LoginCode_name = map[int32]string{
	0: "OK",
	1: "OFFLINE",
}
var RttyMessage_LoginCode_value = map[string]int32{
	"OK":      0,
	"OFFLINE": 1,
}

func (x RttyMessage_LoginCode) String() string {
	return proto.EnumName(RttyMessage_LoginCode_name, int32(x))
}
func (RttyMessage_LoginCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type RttyMessage_FileCode int32

const (
	RttyMessage_START     RttyMessage_FileCode = 0
	RttyMessage_RATELIMIT RttyMessage_FileCode = 1
	RttyMessage_FILEDATA  RttyMessage_FileCode = 2
	RttyMessage_CANCELED  RttyMessage_FileCode = 3
	RttyMessage_END       RttyMessage_FileCode = 4
)

var RttyMessage_FileCode_name = map[int32]string{
	0: "START",
	1: "RATELIMIT",
	2: "FILEDATA",
	3: "CANCELED",
	4: "END",
}
var RttyMessage_FileCode_value = map[string]int32{
	"START":     0,
	"RATELIMIT": 1,
	"FILEDATA":  2,
	"CANCELED":  3,
	"END":       4,
}

func (x RttyMessage_FileCode) String() string {
	return proto.EnumName(RttyMessage_FileCode_name, int32(x))
}
func (RttyMessage_FileCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

type RttyMessage_CommandErr int32

const (
	RttyMessage_NONE           RttyMessage_CommandErr = 0
	RttyMessage_TIMEOUT        RttyMessage_CommandErr = 1
	RttyMessage_NOTFOUND       RttyMessage_CommandErr = 2
	RttyMessage_READ           RttyMessage_CommandErr = 3
	RttyMessage_PERMISSION     RttyMessage_CommandErr = 4
	RttyMessage_SYSCALL        RttyMessage_CommandErr = 5
	RttyMessage_DEV_OFFLINE    RttyMessage_CommandErr = 6
	RttyMessage_CMD_REQUIRED   RttyMessage_CommandErr = 7
	RttyMessage_DEVID_REQUIRED RttyMessage_CommandErr = 8
)

var RttyMessage_CommandErr_name = map[int32]string{
	0: "NONE",
	1: "TIMEOUT",
	2: "NOTFOUND",
	3: "READ",
	4: "PERMISSION",
	5: "SYSCALL",
	6: "DEV_OFFLINE",
	7: "CMD_REQUIRED",
	8: "DEVID_REQUIRED",
}
var RttyMessage_CommandErr_value = map[string]int32{
	"NONE":           0,
	"TIMEOUT":        1,
	"NOTFOUND":       2,
	"READ":           3,
	"PERMISSION":     4,
	"SYSCALL":        5,
	"DEV_OFFLINE":    6,
	"CMD_REQUIRED":   7,
	"DEVID_REQUIRED": 8,
}

func (x RttyMessage_CommandErr) String() string {
	return proto.EnumName(RttyMessage_CommandErr_name, int32(x))
}
func (RttyMessage_CommandErr) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 3} }

type RttyMessage struct {
	Version  uint32            `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Type     RttyMessage_Type  `protobuf:"varint,2,opt,name=type,enum=RttyMessage_Type" json:"type,omitempty"`
	Sid      string            `protobuf:"bytes,3,opt,name=sid" json:"sid,omitempty"`
	Code     int32             `protobuf:"varint,4,opt,name=code" json:"code,omitempty"`
	Data     []byte            `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Name     string            `protobuf:"bytes,6,opt,name=name" json:"name,omitempty"`
	Size     uint32            `protobuf:"varint,7,opt,name=size" json:"size,omitempty"`
	Id       uint32            `protobuf:"varint,8,opt,name=id" json:"id,omitempty"`
	Err      int32             `protobuf:"varint,9,opt,name=err" json:"err,omitempty"`
	Username string            `protobuf:"bytes,10,opt,name=username" json:"username,omitempty"`
	Password string            `protobuf:"bytes,11,opt,name=password" json:"password,omitempty"`
	StdOut   string            `protobuf:"bytes,12,opt,name=std_out,json=stdOut" json:"std_out,omitempty"`
	StdErr   string            `protobuf:"bytes,13,opt,name=std_err,json=stdErr" json:"std_err,omitempty"`
	Params   []string          `protobuf:"bytes,14,rep,name=params" json:"params,omitempty"`
	Env      map[string]string `protobuf:"bytes,15,rep,name=env" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *RttyMessage) Reset()                    { *m = RttyMessage{} }
func (m *RttyMessage) String() string            { return proto.CompactTextString(m) }
func (*RttyMessage) ProtoMessage()               {}
func (*RttyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RttyMessage) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *RttyMessage) GetType() RttyMessage_Type {
	if m != nil {
		return m.Type
	}
	return RttyMessage_UNKNOWN
}

func (m *RttyMessage) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *RttyMessage) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RttyMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *RttyMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RttyMessage) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *RttyMessage) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RttyMessage) GetErr() int32 {
	if m != nil {
		return m.Err
	}
	return 0
}

func (m *RttyMessage) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RttyMessage) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RttyMessage) GetStdOut() string {
	if m != nil {
		return m.StdOut
	}
	return ""
}

func (m *RttyMessage) GetStdErr() string {
	if m != nil {
		return m.StdErr
	}
	return ""
}

func (m *RttyMessage) GetParams() []string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *RttyMessage) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

func init() {
	proto.RegisterType((*RttyMessage)(nil), "rtty_message")
	proto.RegisterEnum("RttyMessage_Type", RttyMessage_Type_name, RttyMessage_Type_value)
	proto.RegisterEnum("RttyMessage_LoginCode", RttyMessage_LoginCode_name, RttyMessage_LoginCode_value)
	proto.RegisterEnum("RttyMessage_FileCode", RttyMessage_FileCode_name, RttyMessage_FileCode_value)
	proto.RegisterEnum("RttyMessage_CommandErr", RttyMessage_CommandErr_name, RttyMessage_CommandErr_value)
}

func init() { proto.RegisterFile("rtty.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x53, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0x8e, 0xfe, 0xa5, 0xb1, 0xe3, 0x2c, 0x4b, 0x49, 0x97, 0x9c, 0x84, 0x0f, 0x45, 0xa7, 0x1c,
	0x52, 0x28, 0xa5, 0x37, 0x21, 0xad, 0x83, 0x88, 0xbc, 0x9b, 0xae, 0xe5, 0x84, 0x9c, 0x8c, 0x5a,
	0x2d, 0xa9, 0x68, 0x2c, 0x19, 0x49, 0x76, 0x71, 0x5f, 0xa2, 0xcf, 0xd0, 0x37, 0x2d, 0xbb, 0xc2,
	0x21, 0xbd, 0x7d, 0x3f, 0x33, 0xdf, 0xcc, 0x8e, 0x10, 0x40, 0x37, 0x0c, 0xc7, 0xeb, 0x5d, 0xd7,
	0x0e, 0xed, 0xfc, 0xaf, 0x0b, 0x53, 0x45, 0x37, 0x5b, 0xd9, 0xf7, 0xe5, 0xb3, 0xc4, 0x04, 0xbc,
	0x83, 0xec, 0xfa, 0xba, 0x6d, 0x88, 0x11, 0x1a, 0xd1, 0xb9, 0x38, 0x51, 0xfc, 0x01, 0xec, 0xe1,
	0xb8, 0x93, 0xc4, 0x0c, 0x8d, 0x68, 0x76, 0x83, 0xaf, 0xdf, 0xb6, 0x5d, 0x17, 0xc7, 0x9d, 0x14,
	0xda, 0xc7, 0x08, 0xac, 0xbe, 0xae, 0x88, 0x15, 0x1a, 0x51, 0x20, 0x14, 0xc4, 0x18, 0xec, 0xef,
	0x6d, 0x25, 0x89, 0x1d, 0x1a, 0x91, 0x23, 0x34, 0x56, 0x5a, 0x55, 0x0e, 0x25, 0x71, 0x42, 0x23,
	0x9a, 0x0a, 0x8d, 0x95, 0xd6, 0x94, 0x5b, 0x49, 0x5c, 0xdd, 0xaa, 0xb1, 0xd2, 0xfa, 0xfa, 0xb7,
	0x24, 0x9e, 0x5e, 0x46, 0x63, 0x3c, 0x03, 0xb3, 0xae, 0x88, 0xaf, 0x15, 0xb3, 0xae, 0xd4, 0x44,
	0xd9, 0x75, 0x24, 0xd0, 0xf1, 0x0a, 0xe2, 0x2b, 0xf0, 0xf7, 0xbd, 0xec, 0x74, 0x1a, 0xe8, 0xb4,
	0x57, 0xae, 0xbc, 0x5d, 0xd9, 0xf7, 0xbf, 0xda, 0xae, 0x22, 0x93, 0xd1, 0x3b, 0x71, 0xfc, 0x1e,
	0xbc, 0x7e, 0xa8, 0x36, 0xed, 0x7e, 0x20, 0x53, 0x6d, 0xb9, 0xfd, 0x50, 0xf1, 0xfd, 0x70, 0x32,
	0xd4, 0x98, 0xf3, 0x57, 0x83, 0x76, 0x1d, 0xbe, 0x04, 0x77, 0x57, 0x76, 0xe5, 0xb6, 0x27, 0xb3,
	0xd0, 0x52, 0xfa, 0xc8, 0x70, 0x04, 0x96, 0x6c, 0x0e, 0xe4, 0x22, 0xb4, 0xa2, 0xc9, 0xcd, 0xe5,
	0xff, 0xc7, 0xa2, 0xcd, 0x81, 0x36, 0x43, 0x77, 0x14, 0xaa, 0xe4, 0xea, 0x13, 0xf8, 0x27, 0x41,
	0xbd, 0xe4, 0xa7, 0x3c, 0xea, 0xcb, 0x07, 0x42, 0x41, 0xfc, 0x0e, 0x9c, 0x43, 0xf9, 0xb2, 0x1f,
	0xcf, 0x1e, 0x88, 0x91, 0x7c, 0x31, 0x3f, 0x1b, 0xf3, 0x1f, 0x60, 0xab, 0xab, 0xe3, 0x09, 0x78,
	0x6b, 0x76, 0xc7, 0xf8, 0x23, 0x43, 0x67, 0x38, 0x00, 0x27, 0xe7, 0xb7, 0x19, 0x43, 0x06, 0x9e,
	0x82, 0xaf, 0x61, 0x9c, 0xdc, 0x21, 0x13, 0x03, 0xb8, 0x39, 0xbf, 0xe5, 0xeb, 0x02, 0x59, 0xd8,
	0x03, 0xab, 0x28, 0x9e, 0x90, 0xad, 0xc4, 0xf5, 0xfd, 0x22, 0xcb, 0x29, 0x72, 0x54, 0x79, 0xca,
	0x1f, 0x99, 0x66, 0xae, 0x0a, 0x4d, 0xf8, 0x72, 0x19, 0xb3, 0x14, 0x79, 0xf3, 0x10, 0x82, 0xbc,
	0x7d, 0xae, 0x9b, 0x44, 0x7d, 0x38, 0x17, 0x4c, 0x7e, 0x87, 0xce, 0x54, 0x05, 0x5f, 0x2c, 0xf2,
	0x8c, 0x51, 0x64, 0xcc, 0x33, 0xf0, 0x17, 0xf5, 0x8b, 0xd4, 0x05, 0x01, 0x38, 0xab, 0x22, 0x16,
	0x05, 0x3a, 0xc3, 0xe7, 0x10, 0x88, 0xb8, 0xa0, 0x79, 0xb6, 0xcc, 0x8a, 0x71, 0x23, 0x15, 0x9f,
	0xc6, 0x45, 0x8c, 0x4c, 0xc5, 0x92, 0x98, 0x25, 0x34, 0xa7, 0xe9, 0xb8, 0x13, 0x65, 0x29, 0xb2,
	0xe7, 0x7f, 0x0c, 0x80, 0xa4, 0xdd, 0x6e, 0xcb, 0x46, 0xdf, 0xd7, 0x07, 0x9b, 0x71, 0x46, 0xc7,
	0x81, 0x45, 0xb6, 0xa4, 0xea, 0x09, 0x3a, 0x8a, 0xf1, 0x62, 0xc1, 0xd7, 0x2c, 0x45, 0xa6, 0x2a,
	0x12, 0x34, 0x56, 0x31, 0x33, 0x80, 0x7b, 0x2a, 0x96, 0xd9, 0x6a, 0x95, 0x71, 0x86, 0x6c, 0xd5,
	0xb4, 0x7a, 0x5a, 0x25, 0x71, 0x9e, 0x23, 0x07, 0x5f, 0xc0, 0x24, 0xa5, 0x0f, 0x9b, 0xd3, 0xda,
	0x2e, 0x46, 0x30, 0x4d, 0x96, 0xe9, 0x46, 0xd0, 0xaf, 0xeb, 0x4c, 0xd0, 0x14, 0x79, 0x18, 0xc3,
	0x2c, 0xa5, 0x0f, 0xd9, 0x1b, 0xcd, 0xff, 0xe6, 0xea, 0x5f, 0xe5, 0xe3, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x3f, 0x23, 0xf5, 0x17, 0x38, 0x03, 0x00, 0x00,
}