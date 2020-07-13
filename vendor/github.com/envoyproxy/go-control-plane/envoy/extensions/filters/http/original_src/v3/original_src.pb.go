// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/original_src/v3/original_src.proto

package envoy_extensions_filters_http_original_src_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type OriginalSrc struct {
	Mark                 uint32   `protobuf:"varint,1,opt,name=mark,proto3" json:"mark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OriginalSrc) Reset()         { *m = OriginalSrc{} }
func (m *OriginalSrc) String() string { return proto.CompactTextString(m) }
func (*OriginalSrc) ProtoMessage()    {}
func (*OriginalSrc) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac73f7eba2c8a298, []int{0}
}

func (m *OriginalSrc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OriginalSrc.Unmarshal(m, b)
}
func (m *OriginalSrc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OriginalSrc.Marshal(b, m, deterministic)
}
func (m *OriginalSrc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OriginalSrc.Merge(m, src)
}
func (m *OriginalSrc) XXX_Size() int {
	return xxx_messageInfo_OriginalSrc.Size(m)
}
func (m *OriginalSrc) XXX_DiscardUnknown() {
	xxx_messageInfo_OriginalSrc.DiscardUnknown(m)
}

var xxx_messageInfo_OriginalSrc proto.InternalMessageInfo

func (m *OriginalSrc) GetMark() uint32 {
	if m != nil {
		return m.Mark
	}
	return 0
}

func init() {
	proto.RegisterType((*OriginalSrc)(nil), "envoy.extensions.filters.http.original_src.v3.OriginalSrc")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/original_src/v3/original_src.proto", fileDescriptor_ac73f7eba2c8a298)
}

var fileDescriptor_ac73f7eba2c8a298 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x48, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x4f, 0xcb, 0xcc,
	0x29, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0x28, 0x29, 0x29, 0xd0, 0xcf, 0x2f, 0xca, 0x4c, 0xcf, 0xcc,
	0x4b, 0xcc, 0x89, 0x2f, 0x2e, 0x4a, 0xd6, 0x2f, 0x33, 0x46, 0xe1, 0xeb, 0x15, 0x14, 0xe5, 0x97,
	0xe4, 0x0b, 0xe9, 0x82, 0x4d, 0xd0, 0x43, 0x98, 0xa0, 0x07, 0x35, 0x41, 0x0f, 0x64, 0x82, 0x1e,
	0x8a, 0x8e, 0x32, 0x63, 0x29, 0xc5, 0xd2, 0x94, 0x82, 0x44, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92,
	0xc4, 0x12, 0xb0, 0x85, 0x65, 0xa9, 0x45, 0x20, 0x7d, 0x99, 0x79, 0xe9, 0x10, 0x13, 0xa5, 0xc4,
	0xcb, 0x12, 0x73, 0x32, 0x53, 0x12, 0x4b, 0x52, 0xf5, 0x61, 0x0c, 0x88, 0x84, 0x52, 0x0a, 0x17,
	0xb7, 0x3f, 0xd4, 0xb8, 0xe0, 0xa2, 0x64, 0x21, 0x21, 0x2e, 0x96, 0xdc, 0xc4, 0xa2, 0x6c, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xde, 0x20, 0x30, 0xdb, 0xca, 0x71, 0xd6, 0xd1, 0x0e, 0x39, 0x1b, 0x2e,
	0x2b, 0x88, 0xa3, 0x92, 0xf3, 0xf3, 0xd2, 0x32, 0xd3, 0xa1, 0x0e, 0xc2, 0xe6, 0x1e, 0xa3, 0xc4,
	0x9c, 0x82, 0x8c, 0x44, 0x43, 0x3d, 0x24, 0x63, 0x9d, 0x02, 0xb9, 0xac, 0x33, 0xf3, 0xf5, 0xc0,
	0x06, 0x14, 0x14, 0xe5, 0x57, 0x54, 0xea, 0x91, 0xe4, 0x41, 0x27, 0x01, 0x24, 0xb3, 0x02, 0x40,
	0xce, 0x0e, 0x60, 0x4c, 0x62, 0x03, 0xbb, 0xdf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x97,
	0x74, 0xa0, 0x6e, 0x01, 0x00, 0x00,
}
