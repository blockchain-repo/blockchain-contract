// Code generated by protoc-gen-go.
// source: contract.proto
// DO NOT EDIT!

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	contract.proto

It has these top-level messages:
	ContractAttributes
	ContractSignature
	ContractAssert
	PlanTaskCondition
	Plan
	Task
	ContractComponents
	Contract
	ContractProto
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ContractAttributes struct {
	Name           string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	StartTimestamp int64  `protobuf:"varint,2,opt,name=start_timestamp,json=startTimestamp" json:"start_timestamp,omitempty"`
	EndTimestamp   int64  `protobuf:"varint,3,opt,name=end_timestamp,json=endTimestamp" json:"end_timestamp,omitempty"`
}

func (m *ContractAttributes) Reset()                    { *m = ContractAttributes{} }
func (m *ContractAttributes) String() string            { return proto.CompactTextString(m) }
func (*ContractAttributes) ProtoMessage()               {}
func (*ContractAttributes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ContractAttributes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContractAttributes) GetStartTimestamp() int64 {
	if m != nil {
		return m.StartTimestamp
	}
	return 0
}

func (m *ContractAttributes) GetEndTimestamp() int64 {
	if m != nil {
		return m.EndTimestamp
	}
	return 0
}

type ContractSignature struct {
	OwnerPubkey string `protobuf:"bytes,1,opt,name=owner_pubkey,json=ownerPubkey" json:"owner_pubkey,omitempty"`
	Signature   string `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	Timestamp   int64  `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *ContractSignature) Reset()                    { *m = ContractSignature{} }
func (m *ContractSignature) String() string            { return proto.CompactTextString(m) }
func (*ContractSignature) ProtoMessage()               {}
func (*ContractSignature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ContractSignature) GetOwnerPubkey() string {
	if m != nil {
		return m.OwnerPubkey
	}
	return ""
}

func (m *ContractSignature) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *ContractSignature) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type ContractAssert struct {
	Id       string               `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string               `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Amount   float32              `protobuf:"fixed32,3,opt,name=amount" json:"amount,omitempty"`
	Metadata *google_protobuf.Any `protobuf:"bytes,4,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *ContractAssert) Reset()                    { *m = ContractAssert{} }
func (m *ContractAssert) String() string            { return proto.CompactTextString(m) }
func (*ContractAssert) ProtoMessage()               {}
func (*ContractAssert) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ContractAssert) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ContractAssert) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContractAssert) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ContractAssert) GetMetadata() *google_protobuf.Any {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type PlanTaskCondition struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type        string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Value       string `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
}

func (m *PlanTaskCondition) Reset()                    { *m = PlanTaskCondition{} }
func (m *PlanTaskCondition) String() string            { return proto.CompactTextString(m) }
func (*PlanTaskCondition) ProtoMessage()               {}
func (*PlanTaskCondition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PlanTaskCondition) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PlanTaskCondition) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PlanTaskCondition) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PlanTaskCondition) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PlanTaskCondition) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Plan struct {
	Id           string               `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type         string               `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	State        string               `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
	Name         string               `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Description  string               `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	Condition    []*PlanTaskCondition `protobuf:"bytes,6,rep,name=condition" json:"condition,omitempty"`
	Level        int32                `protobuf:"varint,7,opt,name=level" json:"level,omitempty"`
	ContractType string               `protobuf:"bytes,8,opt,name=contract_type,json=contractType" json:"contract_type,omitempty"`
	NextTask     []string             `protobuf:"bytes,9,rep,name=next_task,json=nextTask" json:"next_task,omitempty"`
}

func (m *Plan) Reset()                    { *m = Plan{} }
func (m *Plan) String() string            { return proto.CompactTextString(m) }
func (*Plan) ProtoMessage()               {}
func (*Plan) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Plan) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Plan) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Plan) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Plan) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Plan) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Plan) GetCondition() []*PlanTaskCondition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *Plan) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Plan) GetContractType() string {
	if m != nil {
		return m.ContractType
	}
	return ""
}

func (m *Plan) GetNextTask() []string {
	if m != nil {
		return m.NextTask
	}
	return nil
}

type Task struct {
	Id           string               `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type         string               `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	State        string               `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
	Name         string               `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Description  string               `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	Condition    []*PlanTaskCondition `protobuf:"bytes,6,rep,name=condition" json:"condition,omitempty"`
	Level        int32                `protobuf:"varint,7,opt,name=level" json:"level,omitempty"`
	ContractType string               `protobuf:"bytes,8,opt,name=contract_type,json=contractType" json:"contract_type,omitempty"`
	NextTask     []string             `protobuf:"bytes,9,rep,name=next_task,json=nextTask" json:"next_task,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Task) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Task) GetCondition() []*PlanTaskCondition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *Task) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Task) GetContractType() string {
	if m != nil {
		return m.ContractType
	}
	return ""
}

func (m *Task) GetNextTask() []string {
	if m != nil {
		return m.NextTask
	}
	return nil
}

type ContractComponents struct {
	Plans []*Plan `protobuf:"bytes,1,rep,name=plans" json:"plans,omitempty"`
	Tasks []*Task `protobuf:"bytes,2,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *ContractComponents) Reset()                    { *m = ContractComponents{} }
func (m *ContractComponents) String() string            { return proto.CompactTextString(m) }
func (*ContractComponents) ProtoMessage()               {}
func (*ContractComponents) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ContractComponents) GetPlans() []*Plan {
	if m != nil {
		return m.Plans
	}
	return nil
}

func (m *ContractComponents) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type Contract struct {
	CreatorPubkey      string               `protobuf:"bytes,1,opt,name=creator_pubkey,json=creatorPubkey" json:"creator_pubkey,omitempty"`
	CreateTimestamp    int64                `protobuf:"varint,2,opt,name=create_timestamp,json=createTimestamp" json:"create_timestamp,omitempty"`
	Operation          string               `protobuf:"bytes,3,opt,name=operation" json:"operation,omitempty"`
	ContractAttributes *ContractAttributes  `protobuf:"bytes,4,opt,name=contract_attributes,json=contractAttributes" json:"contract_attributes,omitempty"`
	ContractOwners     []string             `protobuf:"bytes,5,rep,name=contract_owners,json=contractOwners" json:"contract_owners,omitempty"`
	ContractSignatures []*ContractSignature `protobuf:"bytes,6,rep,name=contract_signatures,json=contractSignatures" json:"contract_signatures,omitempty"`
	ContractAsserts    []*ContractAssert    `protobuf:"bytes,7,rep,name=contract_asserts,json=contractAsserts" json:"contract_asserts,omitempty"`
	ContractComponents *ContractComponents  `protobuf:"bytes,8,opt,name=contract_components,json=contractComponents" json:"contract_components,omitempty"`
}

func (m *Contract) Reset()                    { *m = Contract{} }
func (m *Contract) String() string            { return proto.CompactTextString(m) }
func (*Contract) ProtoMessage()               {}
func (*Contract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Contract) GetCreatorPubkey() string {
	if m != nil {
		return m.CreatorPubkey
	}
	return ""
}

func (m *Contract) GetCreateTimestamp() int64 {
	if m != nil {
		return m.CreateTimestamp
	}
	return 0
}

func (m *Contract) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *Contract) GetContractAttributes() *ContractAttributes {
	if m != nil {
		return m.ContractAttributes
	}
	return nil
}

func (m *Contract) GetContractOwners() []string {
	if m != nil {
		return m.ContractOwners
	}
	return nil
}

func (m *Contract) GetContractSignatures() []*ContractSignature {
	if m != nil {
		return m.ContractSignatures
	}
	return nil
}

func (m *Contract) GetContractAsserts() []*ContractAssert {
	if m != nil {
		return m.ContractAsserts
	}
	return nil
}

func (m *Contract) GetContractComponents() *ContractComponents {
	if m != nil {
		return m.ContractComponents
	}
	return nil
}

type ContractProto struct {
	Id         string    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	NodePubkey string    `protobuf:"bytes,2,opt,name=node_pubkey,json=nodePubkey" json:"node_pubkey,omitempty"`
	MainPubkey string    `protobuf:"bytes,3,opt,name=main_pubkey,json=mainPubkey" json:"main_pubkey,omitempty"`
	Signature  string    `protobuf:"bytes,4,opt,name=signature" json:"signature,omitempty"`
	Voters     []string  `protobuf:"bytes,5,rep,name=voters" json:"voters,omitempty"`
	Timestamp  int64     `protobuf:"varint,6,opt,name=timestamp" json:"timestamp,omitempty"`
	Version    string    `protobuf:"bytes,7,opt,name=version" json:"version,omitempty"`
	Contract   *Contract `protobuf:"bytes,8,opt,name=contract" json:"contract,omitempty"`
}

func (m *ContractProto) Reset()                    { *m = ContractProto{} }
func (m *ContractProto) String() string            { return proto.CompactTextString(m) }
func (*ContractProto) ProtoMessage()               {}
func (*ContractProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ContractProto) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ContractProto) GetNodePubkey() string {
	if m != nil {
		return m.NodePubkey
	}
	return ""
}

func (m *ContractProto) GetMainPubkey() string {
	if m != nil {
		return m.MainPubkey
	}
	return ""
}

func (m *ContractProto) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *ContractProto) GetVoters() []string {
	if m != nil {
		return m.Voters
	}
	return nil
}

func (m *ContractProto) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ContractProto) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ContractProto) GetContract() *Contract {
	if m != nil {
		return m.Contract
	}
	return nil
}

func init() {
	proto.RegisterType((*ContractAttributes)(nil), "protos.ContractAttributes")
	proto.RegisterType((*ContractSignature)(nil), "protos.ContractSignature")
	proto.RegisterType((*ContractAssert)(nil), "protos.ContractAssert")
	proto.RegisterType((*PlanTaskCondition)(nil), "protos.PlanTaskCondition")
	proto.RegisterType((*Plan)(nil), "protos.Plan")
	proto.RegisterType((*Task)(nil), "protos.Task")
	proto.RegisterType((*ContractComponents)(nil), "protos.ContractComponents")
	proto.RegisterType((*Contract)(nil), "protos.Contract")
	proto.RegisterType((*ContractProto)(nil), "protos.ContractProto")
}

func init() { proto.RegisterFile("contract.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 696 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x55, 0xcd, 0x6e, 0xd4, 0x3c,
	0x14, 0x55, 0x32, 0xbf, 0xb9, 0x33, 0x9d, 0xb6, 0xfe, 0xaa, 0x2a, 0xed, 0x57, 0xa9, 0x21, 0x08,
	0x75, 0x90, 0xd0, 0x14, 0x95, 0x05, 0xeb, 0xaa, 0x3b, 0x58, 0x50, 0x99, 0x2e, 0x91, 0x46, 0x9e,
	0xc4, 0x54, 0x51, 0x27, 0x76, 0x14, 0x7b, 0x06, 0x66, 0xc3, 0x86, 0x05, 0x2f, 0xc7, 0x2b, 0x21,
	0x21, 0xdb, 0xb1, 0x93, 0x26, 0x95, 0xe0, 0x01, 0x58, 0x35, 0xf7, 0xf8, 0xf6, 0x9e, 0xe3, 0x63,
	0xfb, 0x0c, 0xcc, 0x12, 0xce, 0x64, 0x49, 0x12, 0xb9, 0x28, 0x4a, 0x2e, 0x39, 0x1a, 0xea, 0x3f,
	0xe2, 0xf4, 0xe4, 0x9e, 0xf3, 0xfb, 0x35, 0xbd, 0xd4, 0xe5, 0x6a, 0xf3, 0xf9, 0x92, 0xb0, 0x9d,
	0x69, 0x89, 0xb7, 0x80, 0x6e, 0xaa, 0x7f, 0xba, 0x96, 0xb2, 0xcc, 0x56, 0x1b, 0x49, 0x05, 0x42,
	0xd0, 0x67, 0x24, 0xa7, 0xa1, 0x17, 0x79, 0xf3, 0x00, 0xeb, 0x6f, 0x74, 0x01, 0xfb, 0x42, 0x92,
	0x52, 0x2e, 0x65, 0x96, 0x53, 0x21, 0x49, 0x5e, 0x84, 0x7e, 0xe4, 0xcd, 0x7b, 0x78, 0xa6, 0xe1,
	0x3b, 0x8b, 0xa2, 0xe7, 0xb0, 0x47, 0x59, 0xda, 0x68, 0xeb, 0xe9, 0xb6, 0x29, 0x65, 0xa9, 0x6b,
	0x8a, 0x4b, 0x38, 0xb4, 0xbc, 0x1f, 0xb3, 0x7b, 0x46, 0xe4, 0xa6, 0xa4, 0xe8, 0x19, 0x4c, 0xf9,
	0x17, 0x46, 0xcb, 0x65, 0xb1, 0x59, 0x3d, 0xd0, 0x5d, 0x45, 0x3f, 0xd1, 0xd8, 0xad, 0x86, 0xd0,
	0x19, 0x04, 0xc2, 0xf6, 0x6b, 0xfe, 0x00, 0xd7, 0x80, 0x5a, 0x6d, 0xd3, 0xd6, 0x40, 0xfc, 0x0d,
	0x66, 0x6e, 0xaf, 0x42, 0xd0, 0x52, 0xa2, 0x19, 0xf8, 0x59, 0x5a, 0xd1, 0xf8, 0x59, 0xea, 0xf6,
	0xed, 0x37, 0xf6, 0x7d, 0x0c, 0x43, 0x92, 0xf3, 0x0d, 0x93, 0x7a, 0xa0, 0x8f, 0xab, 0x0a, 0xbd,
	0x86, 0x71, 0x4e, 0x25, 0x49, 0x89, 0x24, 0x61, 0x3f, 0xf2, 0xe6, 0x93, 0xab, 0xa3, 0x85, 0xf1,
	0x79, 0x61, 0x7d, 0x5e, 0x5c, 0xb3, 0x1d, 0x76, 0x5d, 0xf1, 0x77, 0x0f, 0x0e, 0x6f, 0xd7, 0x84,
	0xdd, 0x11, 0xf1, 0x70, 0xc3, 0x59, 0x9a, 0xc9, 0x8c, 0xb3, 0xa7, 0x34, 0xc8, 0x5d, 0xe1, 0x34,
	0xa8, 0x6f, 0xa7, 0xab, 0xd7, 0xd0, 0x75, 0x04, 0x83, 0x2d, 0x59, 0x6f, 0xa8, 0x26, 0x0f, 0xb0,
	0x29, 0x50, 0x04, 0x93, 0x94, 0x8a, 0xa4, 0xcc, 0x0a, 0x35, 0x3c, 0x1c, 0x18, 0x07, 0x1b, 0x50,
	0xfc, 0xc3, 0x87, 0xbe, 0x52, 0xf1, 0x57, 0xc4, 0x47, 0x30, 0x10, 0x92, 0x48, 0xcb, 0x6c, 0x0a,
	0x27, 0xa7, 0xdf, 0x90, 0xf3, 0x47, 0x62, 0xf4, 0x16, 0x82, 0xc4, 0xee, 0x3a, 0x1c, 0x46, 0xbd,
	0xf9, 0xe4, 0xea, 0xc4, 0x58, 0x25, 0x16, 0x1d, 0x5b, 0x70, 0xdd, 0xab, 0x44, 0xac, 0xe9, 0x96,
	0xae, 0xc3, 0x51, 0xe4, 0xcd, 0x07, 0xd8, 0x14, 0xea, 0x9a, 0xd9, 0xeb, 0xbe, 0xd4, 0xba, 0xc7,
	0x9a, 0x72, 0x6a, 0xc1, 0x3b, 0xa5, 0xff, 0x7f, 0x08, 0x18, 0xfd, 0x2a, 0x97, 0x92, 0x88, 0x87,
	0x30, 0x88, 0x7a, 0xf3, 0x00, 0x8f, 0x15, 0xa0, 0xb8, 0xb4, 0x13, 0xea, 0xe3, 0x9f, 0x13, 0x9f,
	0xea, 0x14, 0xb8, 0xe1, 0x79, 0xc1, 0x19, 0x65, 0x52, 0xa0, 0x18, 0x06, 0xc5, 0x9a, 0x30, 0x11,
	0x7a, 0x5a, 0xe2, 0xb4, 0x29, 0x11, 0x9b, 0x25, 0xd5, 0xa3, 0x26, 0x8a, 0xd0, 0x7f, 0xdc, 0xa3,
	0xc6, 0x62, 0xb3, 0x14, 0xff, 0xec, 0xc1, 0xd8, 0x8e, 0x47, 0x2f, 0x60, 0x96, 0x94, 0x94, 0x48,
	0xde, 0x7a, 0xe5, 0x7b, 0x15, 0x5a, 0xbd, 0xf3, 0x97, 0x70, 0xa0, 0x01, 0xda, 0x89, 0x9b, 0x7d,
	0x83, 0xd7, 0x79, 0x73, 0x06, 0x01, 0x2f, 0x68, 0x49, 0xb4, 0x9b, 0xe6, 0x74, 0x6a, 0x00, 0xbd,
	0x87, 0xff, 0x9c, 0x39, 0xc4, 0x25, 0x5c, 0xf5, 0x62, 0x4f, 0xad, 0xdc, 0x6e, 0x06, 0x62, 0x94,
	0x74, 0x73, 0xf1, 0x02, 0xf6, 0xdd, 0x30, 0x9d, 0x4a, 0x22, 0x1c, 0x68, 0x2b, 0x5d, 0xf2, 0x7e,
	0xd0, 0x28, 0x7a, 0xd7, 0x60, 0x75, 0xf1, 0x24, 0xda, 0x67, 0xdd, 0x49, 0xc0, 0x9a, 0xd4, 0x41,
	0x02, 0x5d, 0xc3, 0x41, 0xbd, 0x03, 0x9d, 0x5b, 0x22, 0x1c, 0xe9, 0x41, 0xc7, 0x1d, 0xf9, 0x7a,
	0x19, 0x3b, 0x91, 0xa6, 0x16, 0x8f, 0x4c, 0x48, 0xdc, 0x01, 0xeb, 0x7b, 0xf2, 0x84, 0x09, 0xf5,
	0x15, 0xa8, 0xf5, 0xd4, 0x58, 0xfc, 0xcb, 0x83, 0x3d, 0xdb, 0x7a, 0xab, 0x7f, 0x67, 0xda, 0xef,
	0xe7, 0x1c, 0x26, 0x8c, 0xa7, 0xd4, 0x1e, 0xb0, 0x79, 0x46, 0xa0, 0xa0, 0xea, 0x74, 0xcf, 0x61,
	0x92, 0x93, 0x8c, 0xd9, 0x06, 0x73, 0x68, 0xa0, 0xa0, 0xa7, 0x62, 0xbe, 0xdf, 0x8e, 0xf9, 0x63,
	0x18, 0x6e, 0xb9, 0xac, 0xdd, 0xaf, 0xaa, 0xc7, 0xf1, 0x3f, 0x6c, 0xc5, 0x3f, 0x0a, 0x61, 0xb4,
	0xa5, 0xa5, 0x50, 0xb7, 0x64, 0xa4, 0x27, 0xda, 0x12, 0xbd, 0x82, 0xb1, 0xdd, 0x67, 0xe5, 0xc9,
	0x41, 0xdb, 0x13, 0xec, 0x3a, 0x56, 0xe6, 0x57, 0xf5, 0xcd, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x58, 0xd7, 0x63, 0xc9, 0x6e, 0x07, 0x00, 0x00,
}
