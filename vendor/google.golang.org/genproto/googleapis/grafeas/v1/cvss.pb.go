// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grafeas/v1/cvss.proto

package grafeas

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type CVSSv3_AttackVector int32

const (
	CVSSv3_ATTACK_VECTOR_UNSPECIFIED CVSSv3_AttackVector = 0
	CVSSv3_ATTACK_VECTOR_NETWORK     CVSSv3_AttackVector = 1
	CVSSv3_ATTACK_VECTOR_ADJACENT    CVSSv3_AttackVector = 2
	CVSSv3_ATTACK_VECTOR_LOCAL       CVSSv3_AttackVector = 3
	CVSSv3_ATTACK_VECTOR_PHYSICAL    CVSSv3_AttackVector = 4
)

var CVSSv3_AttackVector_name = map[int32]string{
	0: "ATTACK_VECTOR_UNSPECIFIED",
	1: "ATTACK_VECTOR_NETWORK",
	2: "ATTACK_VECTOR_ADJACENT",
	3: "ATTACK_VECTOR_LOCAL",
	4: "ATTACK_VECTOR_PHYSICAL",
}

var CVSSv3_AttackVector_value = map[string]int32{
	"ATTACK_VECTOR_UNSPECIFIED": 0,
	"ATTACK_VECTOR_NETWORK":     1,
	"ATTACK_VECTOR_ADJACENT":    2,
	"ATTACK_VECTOR_LOCAL":       3,
	"ATTACK_VECTOR_PHYSICAL":    4,
}

func (x CVSSv3_AttackVector) String() string {
	return proto.EnumName(CVSSv3_AttackVector_name, int32(x))
}

func (CVSSv3_AttackVector) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_20c164700ea8df83, []int{0, 0}
}

type CVSSv3_AttackComplexity int32

const (
	CVSSv3_ATTACK_COMPLEXITY_UNSPECIFIED CVSSv3_AttackComplexity = 0
	CVSSv3_ATTACK_COMPLEXITY_LOW         CVSSv3_AttackComplexity = 1
	CVSSv3_ATTACK_COMPLEXITY_HIGH        CVSSv3_AttackComplexity = 2
)

var CVSSv3_AttackComplexity_name = map[int32]string{
	0: "ATTACK_COMPLEXITY_UNSPECIFIED",
	1: "ATTACK_COMPLEXITY_LOW",
	2: "ATTACK_COMPLEXITY_HIGH",
}

var CVSSv3_AttackComplexity_value = map[string]int32{
	"ATTACK_COMPLEXITY_UNSPECIFIED": 0,
	"ATTACK_COMPLEXITY_LOW":         1,
	"ATTACK_COMPLEXITY_HIGH":        2,
}

func (x CVSSv3_AttackComplexity) String() string {
	return proto.EnumName(CVSSv3_AttackComplexity_name, int32(x))
}

func (CVSSv3_AttackComplexity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_20c164700ea8df83, []int{0, 1}
}

type CVSSv3_PrivilegesRequired int32

const (
	CVSSv3_PRIVILEGES_REQUIRED_UNSPECIFIED CVSSv3_PrivilegesRequired = 0
	CVSSv3_PRIVILEGES_REQUIRED_NONE        CVSSv3_PrivilegesRequired = 1
	CVSSv3_PRIVILEGES_REQUIRED_LOW         CVSSv3_PrivilegesRequired = 2
	CVSSv3_PRIVILEGES_REQUIRED_HIGH        CVSSv3_PrivilegesRequired = 3
)

var CVSSv3_PrivilegesRequired_name = map[int32]string{
	0: "PRIVILEGES_REQUIRED_UNSPECIFIED",
	1: "PRIVILEGES_REQUIRED_NONE",
	2: "PRIVILEGES_REQUIRED_LOW",
	3: "PRIVILEGES_REQUIRED_HIGH",
}

var CVSSv3_PrivilegesRequired_value = map[string]int32{
	"PRIVILEGES_REQUIRED_UNSPECIFIED": 0,
	"PRIVILEGES_REQUIRED_NONE":        1,
	"PRIVILEGES_REQUIRED_LOW":         2,
	"PRIVILEGES_REQUIRED_HIGH":        3,
}

func (x CVSSv3_PrivilegesRequired) String() string {
	return proto.EnumName(CVSSv3_PrivilegesRequired_name, int32(x))
}

func (CVSSv3_PrivilegesRequired) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_20c164700ea8df83, []int{0, 2}
}

type CVSSv3_UserInteraction int32

const (
	CVSSv3_USER_INTERACTION_UNSPECIFIED CVSSv3_UserInteraction = 0
	CVSSv3_USER_INTERACTION_NONE        CVSSv3_UserInteraction = 1
	CVSSv3_USER_INTERACTION_REQUIRED    CVSSv3_UserInteraction = 2
)

var CVSSv3_UserInteraction_name = map[int32]string{
	0: "USER_INTERACTION_UNSPECIFIED",
	1: "USER_INTERACTION_NONE",
	2: "USER_INTERACTION_REQUIRED",
}

var CVSSv3_UserInteraction_value = map[string]int32{
	"USER_INTERACTION_UNSPECIFIED": 0,
	"USER_INTERACTION_NONE":        1,
	"USER_INTERACTION_REQUIRED":    2,
}

func (x CVSSv3_UserInteraction) String() string {
	return proto.EnumName(CVSSv3_UserInteraction_name, int32(x))
}

func (CVSSv3_UserInteraction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_20c164700ea8df83, []int{0, 3}
}

type CVSSv3_Scope int32

const (
	CVSSv3_SCOPE_UNSPECIFIED CVSSv3_Scope = 0
	CVSSv3_SCOPE_UNCHANGED   CVSSv3_Scope = 1
	CVSSv3_SCOPE_CHANGED     CVSSv3_Scope = 2
)

var CVSSv3_Scope_name = map[int32]string{
	0: "SCOPE_UNSPECIFIED",
	1: "SCOPE_UNCHANGED",
	2: "SCOPE_CHANGED",
}

var CVSSv3_Scope_value = map[string]int32{
	"SCOPE_UNSPECIFIED": 0,
	"SCOPE_UNCHANGED":   1,
	"SCOPE_CHANGED":     2,
}

func (x CVSSv3_Scope) String() string {
	return proto.EnumName(CVSSv3_Scope_name, int32(x))
}

func (CVSSv3_Scope) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_20c164700ea8df83, []int{0, 4}
}

type CVSSv3_Impact int32

const (
	CVSSv3_IMPACT_UNSPECIFIED CVSSv3_Impact = 0
	CVSSv3_IMPACT_HIGH        CVSSv3_Impact = 1
	CVSSv3_IMPACT_LOW         CVSSv3_Impact = 2
	CVSSv3_IMPACT_NONE        CVSSv3_Impact = 3
)

var CVSSv3_Impact_name = map[int32]string{
	0: "IMPACT_UNSPECIFIED",
	1: "IMPACT_HIGH",
	2: "IMPACT_LOW",
	3: "IMPACT_NONE",
}

var CVSSv3_Impact_value = map[string]int32{
	"IMPACT_UNSPECIFIED": 0,
	"IMPACT_HIGH":        1,
	"IMPACT_LOW":         2,
	"IMPACT_NONE":        3,
}

func (x CVSSv3_Impact) String() string {
	return proto.EnumName(CVSSv3_Impact_name, int32(x))
}

func (CVSSv3_Impact) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_20c164700ea8df83, []int{0, 5}
}

// Common Vulnerability Scoring System version 3.
// For details, see https://www.first.org/cvss/specification-document
type CVSSv3 struct {
	// The base score is a function of the base metric scores.
	BaseScore           float32 `protobuf:"fixed32,1,opt,name=base_score,json=baseScore,proto3" json:"base_score,omitempty"`
	ExploitabilityScore float32 `protobuf:"fixed32,2,opt,name=exploitability_score,json=exploitabilityScore,proto3" json:"exploitability_score,omitempty"`
	ImpactScore         float32 `protobuf:"fixed32,3,opt,name=impact_score,json=impactScore,proto3" json:"impact_score,omitempty"`
	// Base Metrics
	// Represents the intrinsic characteristics of a vulnerability that are
	// constant over time and across user environments.
	AttackVector          CVSSv3_AttackVector       `protobuf:"varint,5,opt,name=attack_vector,json=attackVector,proto3,enum=grafeas.v1.CVSSv3_AttackVector" json:"attack_vector,omitempty"`
	AttackComplexity      CVSSv3_AttackComplexity   `protobuf:"varint,6,opt,name=attack_complexity,json=attackComplexity,proto3,enum=grafeas.v1.CVSSv3_AttackComplexity" json:"attack_complexity,omitempty"`
	PrivilegesRequired    CVSSv3_PrivilegesRequired `protobuf:"varint,7,opt,name=privileges_required,json=privilegesRequired,proto3,enum=grafeas.v1.CVSSv3_PrivilegesRequired" json:"privileges_required,omitempty"`
	UserInteraction       CVSSv3_UserInteraction    `protobuf:"varint,8,opt,name=user_interaction,json=userInteraction,proto3,enum=grafeas.v1.CVSSv3_UserInteraction" json:"user_interaction,omitempty"`
	Scope                 CVSSv3_Scope              `protobuf:"varint,9,opt,name=scope,proto3,enum=grafeas.v1.CVSSv3_Scope" json:"scope,omitempty"`
	ConfidentialityImpact CVSSv3_Impact             `protobuf:"varint,10,opt,name=confidentiality_impact,json=confidentialityImpact,proto3,enum=grafeas.v1.CVSSv3_Impact" json:"confidentiality_impact,omitempty"`
	IntegrityImpact       CVSSv3_Impact             `protobuf:"varint,11,opt,name=integrity_impact,json=integrityImpact,proto3,enum=grafeas.v1.CVSSv3_Impact" json:"integrity_impact,omitempty"`
	AvailabilityImpact    CVSSv3_Impact             `protobuf:"varint,12,opt,name=availability_impact,json=availabilityImpact,proto3,enum=grafeas.v1.CVSSv3_Impact" json:"availability_impact,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                  `json:"-"`
	XXX_unrecognized      []byte                    `json:"-"`
	XXX_sizecache         int32                     `json:"-"`
}

func (m *CVSSv3) Reset()         { *m = CVSSv3{} }
func (m *CVSSv3) String() string { return proto.CompactTextString(m) }
func (*CVSSv3) ProtoMessage()    {}
func (*CVSSv3) Descriptor() ([]byte, []int) {
	return fileDescriptor_20c164700ea8df83, []int{0}
}

func (m *CVSSv3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CVSSv3.Unmarshal(m, b)
}
func (m *CVSSv3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CVSSv3.Marshal(b, m, deterministic)
}
func (m *CVSSv3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CVSSv3.Merge(m, src)
}
func (m *CVSSv3) XXX_Size() int {
	return xxx_messageInfo_CVSSv3.Size(m)
}
func (m *CVSSv3) XXX_DiscardUnknown() {
	xxx_messageInfo_CVSSv3.DiscardUnknown(m)
}

var xxx_messageInfo_CVSSv3 proto.InternalMessageInfo

func (m *CVSSv3) GetBaseScore() float32 {
	if m != nil {
		return m.BaseScore
	}
	return 0
}

func (m *CVSSv3) GetExploitabilityScore() float32 {
	if m != nil {
		return m.ExploitabilityScore
	}
	return 0
}

func (m *CVSSv3) GetImpactScore() float32 {
	if m != nil {
		return m.ImpactScore
	}
	return 0
}

func (m *CVSSv3) GetAttackVector() CVSSv3_AttackVector {
	if m != nil {
		return m.AttackVector
	}
	return CVSSv3_ATTACK_VECTOR_UNSPECIFIED
}

func (m *CVSSv3) GetAttackComplexity() CVSSv3_AttackComplexity {
	if m != nil {
		return m.AttackComplexity
	}
	return CVSSv3_ATTACK_COMPLEXITY_UNSPECIFIED
}

func (m *CVSSv3) GetPrivilegesRequired() CVSSv3_PrivilegesRequired {
	if m != nil {
		return m.PrivilegesRequired
	}
	return CVSSv3_PRIVILEGES_REQUIRED_UNSPECIFIED
}

func (m *CVSSv3) GetUserInteraction() CVSSv3_UserInteraction {
	if m != nil {
		return m.UserInteraction
	}
	return CVSSv3_USER_INTERACTION_UNSPECIFIED
}

func (m *CVSSv3) GetScope() CVSSv3_Scope {
	if m != nil {
		return m.Scope
	}
	return CVSSv3_SCOPE_UNSPECIFIED
}

func (m *CVSSv3) GetConfidentialityImpact() CVSSv3_Impact {
	if m != nil {
		return m.ConfidentialityImpact
	}
	return CVSSv3_IMPACT_UNSPECIFIED
}

func (m *CVSSv3) GetIntegrityImpact() CVSSv3_Impact {
	if m != nil {
		return m.IntegrityImpact
	}
	return CVSSv3_IMPACT_UNSPECIFIED
}

func (m *CVSSv3) GetAvailabilityImpact() CVSSv3_Impact {
	if m != nil {
		return m.AvailabilityImpact
	}
	return CVSSv3_IMPACT_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("grafeas.v1.CVSSv3_AttackVector", CVSSv3_AttackVector_name, CVSSv3_AttackVector_value)
	proto.RegisterEnum("grafeas.v1.CVSSv3_AttackComplexity", CVSSv3_AttackComplexity_name, CVSSv3_AttackComplexity_value)
	proto.RegisterEnum("grafeas.v1.CVSSv3_PrivilegesRequired", CVSSv3_PrivilegesRequired_name, CVSSv3_PrivilegesRequired_value)
	proto.RegisterEnum("grafeas.v1.CVSSv3_UserInteraction", CVSSv3_UserInteraction_name, CVSSv3_UserInteraction_value)
	proto.RegisterEnum("grafeas.v1.CVSSv3_Scope", CVSSv3_Scope_name, CVSSv3_Scope_value)
	proto.RegisterEnum("grafeas.v1.CVSSv3_Impact", CVSSv3_Impact_name, CVSSv3_Impact_value)
	proto.RegisterType((*CVSSv3)(nil), "grafeas.v1.CVSSv3")
}

func init() { proto.RegisterFile("grafeas/v1/cvss.proto", fileDescriptor_20c164700ea8df83) }

var fileDescriptor_20c164700ea8df83 = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xcf, 0x53, 0x9b, 0x5a,
	0x14, 0xc7, 0x1f, 0xe4, 0x99, 0xf7, 0x3c, 0x46, 0x83, 0x37, 0x2f, 0x4a, 0x5e, 0xcd, 0xa8, 0x38,
	0x9d, 0x71, 0x45, 0xc6, 0xba, 0xe9, 0x4c, 0x57, 0x94, 0x60, 0x82, 0x46, 0xc0, 0x0b, 0x89, 0xb5,
	0x1b, 0x06, 0xf1, 0xca, 0xdc, 0x29, 0x06, 0x0a, 0x98, 0xd1, 0x7f, 0xa3, 0xbb, 0x6e, 0xfb, 0x67,
	0x76, 0xd5, 0xc9, 0x85, 0x98, 0x9f, 0x8e, 0xbb, 0xe4, 0x7b, 0x3e, 0xe7, 0xc3, 0x01, 0x0e, 0x17,
	0xea, 0x41, 0xe2, 0xdd, 0x13, 0x2f, 0x6d, 0x8d, 0x4e, 0x5a, 0xfe, 0x28, 0x4d, 0xe5, 0x38, 0x89,
	0xb2, 0x08, 0x41, 0x11, 0xcb, 0xa3, 0x13, 0xe9, 0x37, 0x40, 0x59, 0x1d, 0xd8, 0xf6, 0xe8, 0x14,
	0x35, 0x01, 0x6e, 0xbd, 0x94, 0xb8, 0xa9, 0x1f, 0x25, 0x44, 0xe4, 0x0e, 0xb8, 0x63, 0x1e, 0xaf,
	0x8f, 0x13, 0x7b, 0x1c, 0xa0, 0x13, 0xf8, 0x8f, 0x3c, 0xc5, 0x61, 0x44, 0x33, 0xef, 0x96, 0x86,
	0x34, 0x7b, 0x2e, 0x40, 0x9e, 0x81, 0xb5, 0xf9, 0x5a, 0xde, 0x72, 0x08, 0x15, 0xfa, 0x10, 0x7b,
	0x7e, 0x56, 0xa0, 0x25, 0x86, 0x6e, 0xe4, 0x59, 0x8e, 0xb4, 0x61, 0xd3, 0xcb, 0x32, 0xcf, 0xff,
	0xe6, 0x8e, 0x88, 0x9f, 0x45, 0x89, 0xb8, 0x76, 0xc0, 0x1d, 0x6f, 0x7d, 0xd8, 0x97, 0xa7, 0x33,
	0xca, 0xf9, 0x7c, 0xb2, 0xc2, 0xb8, 0x01, 0xc3, 0x70, 0xc5, 0x9b, 0xf9, 0x87, 0x2c, 0xd8, 0x2e,
	0x2c, 0x7e, 0xf4, 0x10, 0x87, 0xe4, 0x89, 0x66, 0xcf, 0x62, 0x99, 0x99, 0x8e, 0x5e, 0x35, 0xa9,
	0x2f, 0x28, 0x16, 0xbc, 0x85, 0x04, 0x0d, 0xa0, 0x16, 0x27, 0x74, 0x44, 0x43, 0x12, 0x90, 0xd4,
	0x4d, 0xc8, 0xf7, 0x47, 0x9a, 0x90, 0x3b, 0xf1, 0x1f, 0xe6, 0x7c, 0xbf, 0xc2, 0x69, 0xbd, 0xd0,
	0xb8, 0x80, 0x31, 0x8a, 0x97, 0x32, 0x74, 0x09, 0xc2, 0x63, 0x4a, 0x12, 0x97, 0x0e, 0x33, 0x92,
	0x78, 0x7e, 0x46, 0xa3, 0xa1, 0xf8, 0x2f, 0x93, 0x4a, 0x2b, 0xa4, 0xfd, 0x94, 0x24, 0xfa, 0x94,
	0xc4, 0xd5, 0xc7, 0xf9, 0x00, 0xc9, 0xb0, 0x96, 0xfa, 0x51, 0x4c, 0xc4, 0x75, 0xe6, 0x10, 0x57,
	0x38, 0xec, 0x71, 0x1d, 0xe7, 0x18, 0xb2, 0x60, 0xc7, 0x8f, 0x86, 0xf7, 0xf4, 0x8e, 0x0c, 0x33,
	0xea, 0xb1, 0xb7, 0x98, 0xbf, 0x0d, 0x11, 0x98, 0xa0, 0xb1, 0x42, 0xa0, 0x33, 0x00, 0xd7, 0x17,
	0x1a, 0xf3, 0x18, 0xb5, 0x41, 0x18, 0xdf, 0x4b, 0x90, 0xcc, 0xb8, 0x36, 0xde, 0x72, 0x55, 0x5f,
	0x5a, 0x0a, 0xcb, 0x39, 0xd4, 0xbc, 0x91, 0x47, 0xc3, 0xc9, 0x6a, 0x15, 0xa2, 0xca, 0x5b, 0x22,
	0x34, 0xdb, 0x95, 0x67, 0xd2, 0x4f, 0x0e, 0x2a, 0xb3, 0xbb, 0x82, 0x9a, 0xd0, 0x50, 0x1c, 0x47,
	0x51, 0x2f, 0xdc, 0x81, 0xa6, 0x3a, 0x26, 0x76, 0xfb, 0x86, 0x6d, 0x69, 0xaa, 0x7e, 0xa6, 0x6b,
	0x6d, 0xe1, 0x2f, 0xd4, 0x80, 0xfa, 0x7c, 0xd9, 0xd0, 0x9c, 0x6b, 0x13, 0x5f, 0x08, 0x1c, 0xfa,
	0x1f, 0x76, 0xe6, 0x4b, 0x4a, 0xfb, 0x5c, 0x51, 0x35, 0xc3, 0x11, 0x78, 0xb4, 0x0b, 0xb5, 0xf9,
	0x5a, 0xcf, 0x54, 0x95, 0x9e, 0x50, 0x5a, 0x6e, 0xb2, 0xba, 0x37, 0xb6, 0x3e, 0xae, 0xfd, 0x2d,
	0x85, 0x20, 0x2c, 0x2e, 0x1f, 0x3a, 0x84, 0x66, 0xc1, 0xab, 0xe6, 0xa5, 0xd5, 0xd3, 0xbe, 0xe8,
	0xce, 0xcd, 0xab, 0x23, 0xce, 0x20, 0x3d, 0xf3, 0x7a, 0x6e, 0xc4, 0x99, 0x52, 0x57, 0xef, 0x74,
	0x05, 0x5e, 0xfa, 0xc1, 0x01, 0x5a, 0xde, 0x4b, 0x74, 0x04, 0xfb, 0x16, 0xd6, 0x07, 0x7a, 0x4f,
	0xeb, 0x68, 0xb6, 0x8b, 0xb5, 0xab, 0xbe, 0x8e, 0xb5, 0xf6, 0xc2, 0x25, 0xf7, 0x40, 0x5c, 0x05,
	0x19, 0xa6, 0xa1, 0x09, 0x1c, 0x7a, 0x07, 0xbb, 0xab, 0xaa, 0xe3, 0x91, 0xf8, 0xd7, 0x5a, 0xd9,
	0x50, 0x25, 0xe9, 0x01, 0xaa, 0x0b, 0x6b, 0x8d, 0x0e, 0x60, 0xaf, 0x6f, 0x6b, 0xd8, 0xd5, 0x0d,
	0x47, 0xc3, 0x8a, 0xea, 0xe8, 0xa6, 0xb1, 0xfc, 0x00, 0x96, 0x88, 0x62, 0x94, 0x26, 0x34, 0x96,
	0x4a, 0x93, 0x6b, 0x0a, 0xbc, 0x74, 0x06, 0x6b, 0xec, 0x0b, 0x40, 0x75, 0xd8, 0xb6, 0x55, 0xd3,
	0xd2, 0x16, 0xcc, 0x35, 0xa8, 0x4e, 0x62, 0xb5, 0xab, 0x18, 0x1d, 0xad, 0x2d, 0x70, 0x68, 0x1b,
	0x36, 0xf3, 0x70, 0x12, 0xf1, 0x12, 0x86, 0x72, 0xb1, 0xab, 0x3b, 0x80, 0xf4, 0x4b, 0x4b, 0x51,
	0x9d, 0x05, 0x53, 0x15, 0x36, 0x8a, 0x9c, 0xdd, 0x29, 0x87, 0xb6, 0x00, 0x8a, 0x20, 0x7f, 0x2e,
	0x53, 0x80, 0x8d, 0x5e, 0xfa, 0x7c, 0x05, 0x9b, 0x34, 0x9a, 0x59, 0x6e, 0x8b, 0xfb, 0xfa, 0x31,
	0x88, 0xa2, 0x20, 0x24, 0x72, 0x10, 0x85, 0xde, 0x30, 0x90, 0xa3, 0x24, 0x68, 0x05, 0x64, 0xc8,
	0xce, 0xed, 0x56, 0x5e, 0xf2, 0x62, 0x9a, 0xb6, 0xa6, 0x07, 0xfb, 0xa7, 0xe2, 0xe7, 0x2f, 0xbe,
	0xd4, 0xc1, 0xca, 0x6d, 0x99, 0xa1, 0xa7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x76, 0x6f, 0xf4,
	0x24, 0xfb, 0x05, 0x00, 0x00,
}
