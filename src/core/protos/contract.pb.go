// Code generated by protoc-gen-go.
// source: contract.proto
// DO NOT EDIT!

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	contract.proto
	data.proto

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
	ContractData
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
	StartTimestamp string `protobuf:"bytes,2,opt,name=start_timestamp,json=startTimestamp" json:"start_timestamp,omitempty"`
	EndTimestamp   string `protobuf:"bytes,3,opt,name=end_timestamp,json=endTimestamp" json:"end_timestamp,omitempty"`
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

func (m *ContractAttributes) GetStartTimestamp() string {
	if m != nil {
		return m.StartTimestamp
	}
	return ""
}

func (m *ContractAttributes) GetEndTimestamp() string {
	if m != nil {
		return m.EndTimestamp
	}
	return ""
}

type ContractSignature struct {
	OwnerPubkey string `protobuf:"bytes,1,opt,name=owner_pubkey,json=ownerPubkey" json:"owner_pubkey,omitempty"`
	Signature   string `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	Timestamp   string `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
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

func (m *ContractSignature) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
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
	CreateTimestamp    string               `protobuf:"bytes,2,opt,name=create_timestamp,json=createTimestamp" json:"create_timestamp,omitempty"`
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

func (m *Contract) GetCreateTimestamp() string {
	if m != nil {
		return m.CreateTimestamp
	}
	return ""
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
	Timestamp  string    `protobuf:"bytes,6,opt,name=timestamp" json:"timestamp,omitempty"`
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

func (m *ContractProto) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
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
	// 692 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x55, 0xcd, 0x6e, 0xdb, 0x3c,
	0x10, 0x84, 0xe4, 0x5f, 0xad, 0x13, 0x27, 0xe1, 0x17, 0x04, 0x4a, 0xbe, 0x00, 0x71, 0x55, 0x14,
	0x71, 0x81, 0xc2, 0x29, 0xd2, 0x43, 0xcf, 0x41, 0x6e, 0xed, 0xa1, 0x81, 0x9a, 0x63, 0x01, 0x83,
	0x96, 0xd8, 0x40, 0x88, 0x45, 0x0a, 0x22, 0xed, 0xd6, 0x97, 0x5e, 0x7a, 0xe8, 0xcb, 0xf5, 0x95,
	0x0a, 0x14, 0x5c, 0x8a, 0x94, 0x22, 0x1b, 0x68, 0x1f, 0xa0, 0x27, 0x73, 0x87, 0xeb, 0xdd, 0xe1,
	0x70, 0x39, 0x82, 0x71, 0x22, 0xb8, 0x2a, 0x69, 0xa2, 0x66, 0x45, 0x29, 0x94, 0x20, 0x7d, 0xfc,
	0x91, 0x67, 0xa7, 0x0f, 0x42, 0x3c, 0x2c, 0xd9, 0x15, 0x86, 0x8b, 0xd5, 0xe7, 0x2b, 0xca, 0x37,
	0x26, 0x25, 0x5a, 0x03, 0xb9, 0xad, 0xfe, 0x74, 0xa3, 0x54, 0x99, 0x2d, 0x56, 0x8a, 0x49, 0x42,
	0xa0, 0xcb, 0x69, 0xce, 0x42, 0x6f, 0xe2, 0x4d, 0x83, 0x18, 0xd7, 0xe4, 0x12, 0x0e, 0xa4, 0xa2,
	0xa5, 0x9a, 0xab, 0x2c, 0x67, 0x52, 0xd1, 0xbc, 0x08, 0x7d, 0xdc, 0x1e, 0x23, 0x7c, 0x6f, 0x51,
	0xf2, 0x1c, 0xf6, 0x19, 0x4f, 0x1b, 0x69, 0x1d, 0x4c, 0xdb, 0x63, 0x3c, 0x75, 0x49, 0x51, 0x09,
	0x47, 0xb6, 0xef, 0xc7, 0xec, 0x81, 0x53, 0xb5, 0x2a, 0x19, 0x79, 0x06, 0x7b, 0xe2, 0x0b, 0x67,
	0xe5, 0xbc, 0x58, 0x2d, 0x1e, 0xd9, 0xa6, 0x6a, 0x3f, 0x42, 0xec, 0x0e, 0x21, 0x72, 0x0e, 0x81,
	0xb4, 0xf9, 0x55, 0xff, 0x1a, 0xd0, 0xbb, 0xed, 0xb6, 0x35, 0x10, 0x7d, 0x83, 0xb1, 0x3b, 0xab,
	0x94, 0xac, 0x54, 0x64, 0x0c, 0x7e, 0x96, 0x56, 0x6d, 0xfc, 0x2c, 0x75, 0xe7, 0xf6, 0x1b, 0xe7,
	0x3e, 0x81, 0x3e, 0xcd, 0xc5, 0x8a, 0x2b, 0x2c, 0xe8, 0xc7, 0x55, 0x44, 0x5e, 0xc3, 0x30, 0x67,
	0x8a, 0xa6, 0x54, 0xd1, 0xb0, 0x3b, 0xf1, 0xa6, 0xa3, 0xeb, 0xe3, 0x99, 0xd1, 0x79, 0x66, 0x75,
	0x9e, 0xdd, 0xf0, 0x4d, 0xec, 0xb2, 0xa2, 0xef, 0x1e, 0x1c, 0xdd, 0x2d, 0x29, 0xbf, 0xa7, 0xf2,
	0xf1, 0x56, 0xf0, 0x34, 0x53, 0x99, 0xe0, 0xbb, 0x38, 0xa8, 0x4d, 0xe1, 0x38, 0xe8, 0xb5, 0xe3,
	0xd5, 0x69, 0xf0, 0x3a, 0x86, 0xde, 0x9a, 0x2e, 0x57, 0x0c, 0x9b, 0x07, 0xb1, 0x09, 0xc8, 0x04,
	0x46, 0x29, 0x93, 0x49, 0x99, 0x15, 0xba, 0x78, 0xd8, 0x33, 0x0a, 0x36, 0xa0, 0xe8, 0x87, 0x0f,
	0x5d, 0xcd, 0xe2, 0xaf, 0x1a, 0x1f, 0x43, 0x4f, 0x2a, 0xaa, 0x6c, 0x67, 0x13, 0x38, 0x3a, 0xdd,
	0x06, 0x9d, 0x3f, 0x36, 0x26, 0x6f, 0x21, 0x48, 0xec, 0xa9, 0xc3, 0xfe, 0xa4, 0x33, 0x1d, 0x5d,
	0x9f, 0x1a, 0xa9, 0xe4, 0x6c, 0x4b, 0x96, 0xb8, 0xce, 0xd5, 0x24, 0x96, 0x6c, 0xcd, 0x96, 0xe1,
	0x60, 0xe2, 0x4d, 0x7b, 0xb1, 0x09, 0xf4, 0x98, 0xd9, 0x71, 0x9f, 0x23, 0xef, 0xa1, 0x19, 0x33,
	0x0b, 0xde, 0x6b, 0xfe, 0xff, 0x43, 0xc0, 0xd9, 0x57, 0x35, 0x57, 0x54, 0x3e, 0x86, 0xc1, 0xa4,
	0x33, 0x0d, 0xe2, 0xa1, 0x06, 0x74, 0x2f, 0x54, 0x42, 0x2f, 0xfe, 0x29, 0xf1, 0xa9, 0x76, 0x81,
	0x5b, 0x91, 0x17, 0x82, 0x33, 0xae, 0x24, 0x89, 0xa0, 0x57, 0x2c, 0x29, 0x97, 0xa1, 0x87, 0x14,
	0xf7, 0x9a, 0x14, 0x63, 0xb3, 0xa5, 0x73, 0x74, 0x45, 0x19, 0xfa, 0x4f, 0x73, 0x74, 0xd9, 0xd8,
	0x6c, 0x45, 0x3f, 0x3b, 0x30, 0xb4, 0xe5, 0xc9, 0x0b, 0x18, 0x27, 0x25, 0xa3, 0x4a, 0xb4, 0x5e,
	0xf9, 0x7e, 0x85, 0x56, 0xef, 0xfc, 0x25, 0x1c, 0x22, 0xc0, 0xb6, 0xec, 0xe6, 0xc0, 0xe0, 0xb5,
	0xdf, 0x9c, 0x43, 0x20, 0x0a, 0x56, 0x52, 0x54, 0xb3, 0x7a, 0xf4, 0x0e, 0x20, 0xef, 0xe1, 0x3f,
	0x27, 0x0e, 0x75, 0x0e, 0x57, 0xbd, 0xd8, 0x33, 0x4b, 0x77, 0xdb, 0x03, 0x63, 0x92, 0x6c, 0xfb,
	0xe2, 0x25, 0x1c, 0xb8, 0x62, 0xe8, 0x4a, 0x32, 0xec, 0xa1, 0x94, 0xce, 0x79, 0x3f, 0x20, 0x4a,
	0xde, 0x35, 0xba, 0x3a, 0x7b, 0x92, 0xed, 0xbb, 0xde, 0x72, 0xc0, 0xba, 0xa9, 0x83, 0x24, 0xb9,
	0x81, 0xc3, 0xfa, 0x04, 0xe8, 0x5b, 0x32, 0x1c, 0x60, 0xa1, 0x93, 0x2d, 0xfa, 0xb8, 0x1d, 0x3b,
	0x92, 0x26, 0x96, 0x4f, 0x44, 0x48, 0xdc, 0x05, 0xe3, 0x9c, 0xec, 0x10, 0xa1, 0x1e, 0x81, 0x9a,
	0x4f, 0x8d, 0x45, 0xbf, 0x3c, 0xd8, 0xb7, 0xa9, 0x77, 0xf8, 0x9d, 0x69, 0xbf, 0x9f, 0x0b, 0x18,
	0x71, 0x91, 0x32, 0x7b, 0xc1, 0xe6, 0xde, 0x40, 0x43, 0xd5, 0xed, 0x5e, 0xc0, 0x28, 0xa7, 0x19,
	0xb7, 0x09, 0xe6, 0xd2, 0x40, 0x43, 0xbb, 0x6c, 0xbe, 0xdb, 0xb6, 0xf9, 0x13, 0xe8, 0xaf, 0x85,
	0xaa, 0xd5, 0xaf, 0xa2, 0xa7, 0xf6, 0xdf, 0x6f, 0xd9, 0x3f, 0x09, 0x61, 0xb0, 0x66, 0xa5, 0xd4,
	0x53, 0x32, 0xc0, 0x3d, 0x1b, 0x92, 0x57, 0x30, 0xb4, 0xe7, 0xac, 0x34, 0x39, 0x6c, 0x6b, 0x12,
	0xbb, 0x8c, 0x85, 0xf9, 0xaa, 0xbe, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xcd, 0x97, 0x55,
	0x6e, 0x07, 0x00, 0x00,
}
