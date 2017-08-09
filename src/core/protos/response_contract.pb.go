// Code generated by protoc-gen-go.
// source: response_contract.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ResponseContract struct {
	// response code, 9bit
	Code int32 `protobuf:"varint,1,opt,name=code" json:"code"`
	// response message
	Msg string `protobuf:"bytes,2,opt,name=msg" json:"msg"`
	// response data
	Result *Contract `protobuf:"bytes,3,opt,name=result" json:"result"`
}

func (m *ResponseContract) Reset()                    { *m = ResponseContract{} }
func (m *ResponseContract) String() string            { return proto.CompactTextString(m) }
func (*ResponseContract) ProtoMessage()               {}
func (*ResponseContract) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *ResponseContract) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResponseContract) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ResponseContract) GetResult() *Contract {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*ResponseContract)(nil), "protos.ResponseContract")
}

func init() { proto.RegisterFile("response_contract.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x4a, 0x2d, 0x2e,
	0xc8, 0xcf, 0x2b, 0x4e, 0x8d, 0x4f, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0xd1, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x52, 0x7c, 0xa8, 0xe2, 0x4a, 0x49, 0x5c, 0x02,
	0x41, 0x50, 0x2d, 0xce, 0x50, 0x19, 0x21, 0x21, 0x2e, 0x96, 0xe4, 0xfc, 0x94, 0x54, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xd6, 0x20, 0x30, 0x5b, 0x48, 0x80, 0x8b, 0x39, 0xb7, 0x38, 0x5d, 0x82, 0x49,
	0x81, 0x51, 0x83, 0x33, 0x08, 0xc4, 0x14, 0xd2, 0xe0, 0x62, 0x2b, 0x4a, 0x2d, 0x2e, 0xcd, 0x29,
	0x91, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x36, 0x12, 0x80, 0x98, 0x58, 0xac, 0x07, 0x33, 0x27, 0x08,
	0x2a, 0xef, 0xa4, 0xcb, 0x25, 0x92, 0x9c, 0x9f, 0xab, 0x57, 0x9a, 0x97, 0x99, 0x93, 0x9a, 0x92,
	0x9e, 0x5a, 0x04, 0x55, 0xe8, 0x24, 0x1a, 0x00, 0xa2, 0xd1, 0xad, 0x4f, 0x82, 0x38, 0xd5, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x60, 0xe3, 0x95, 0xa8, 0xcc, 0x00, 0x00, 0x00,
}
