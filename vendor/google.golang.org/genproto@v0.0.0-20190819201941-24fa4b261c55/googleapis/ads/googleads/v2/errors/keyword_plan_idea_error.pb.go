// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/keyword_plan_idea_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum describing possible errors from KeywordPlanIdeaService.
type KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError int32

const (
	// Enum unspecified.
	KeywordPlanIdeaErrorEnum_UNSPECIFIED KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError = 0
	// The received error code is not known in this version.
	KeywordPlanIdeaErrorEnum_UNKNOWN KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError = 1
	// Error when crawling the input URL.
	KeywordPlanIdeaErrorEnum_URL_CRAWL_ERROR KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError = 2
	// The input has an invalid value.
	KeywordPlanIdeaErrorEnum_INVALID_VALUE KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError = 3
)

var KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "URL_CRAWL_ERROR",
	3: "INVALID_VALUE",
}

var KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError_value = map[string]int32{
	"UNSPECIFIED":     0,
	"UNKNOWN":         1,
	"URL_CRAWL_ERROR": 2,
	"INVALID_VALUE":   3,
}

func (x KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError) String() string {
	return proto.EnumName(KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError_name, int32(x))
}

func (KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b52051ec051c7273, []int{0, 0}
}

// Container for enum describing possible errors from KeywordPlanIdeaService.
type KeywordPlanIdeaErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordPlanIdeaErrorEnum) Reset()         { *m = KeywordPlanIdeaErrorEnum{} }
func (m *KeywordPlanIdeaErrorEnum) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanIdeaErrorEnum) ProtoMessage()    {}
func (*KeywordPlanIdeaErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52051ec051c7273, []int{0}
}

func (m *KeywordPlanIdeaErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanIdeaErrorEnum.Unmarshal(m, b)
}
func (m *KeywordPlanIdeaErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanIdeaErrorEnum.Marshal(b, m, deterministic)
}
func (m *KeywordPlanIdeaErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanIdeaErrorEnum.Merge(m, src)
}
func (m *KeywordPlanIdeaErrorEnum) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanIdeaErrorEnum.Size(m)
}
func (m *KeywordPlanIdeaErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanIdeaErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanIdeaErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError", KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError_name, KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError_value)
	proto.RegisterType((*KeywordPlanIdeaErrorEnum)(nil), "google.ads.googleads.v2.errors.KeywordPlanIdeaErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/keyword_plan_idea_error.proto", fileDescriptor_b52051ec051c7273)
}

var fileDescriptor_b52051ec051c7273 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4a, 0xf3, 0x30,
	0x1c, 0xc5, 0xbf, 0x75, 0xf0, 0x09, 0x19, 0xb2, 0x5a, 0xbd, 0x50, 0x91, 0x5d, 0xf4, 0x01, 0x52,
	0xa8, 0x77, 0xd1, 0x9b, 0x6c, 0xab, 0xa3, 0xac, 0x74, 0xa5, 0xd2, 0x0e, 0xa4, 0x50, 0xa2, 0x09,
	0xa1, 0xd8, 0x25, 0x25, 0x99, 0x53, 0x5f, 0xc7, 0x4b, 0x1f, 0xc5, 0x47, 0xf1, 0xda, 0x07, 0x90,
	0x36, 0x6e, 0x57, 0xd3, 0xab, 0x1c, 0xf2, 0x3f, 0xbf, 0x93, 0x93, 0x3f, 0xb8, 0xe6, 0x52, 0xf2,
	0x9a, 0x79, 0x84, 0x6a, 0xcf, 0xc8, 0x56, 0x6d, 0x7c, 0x8f, 0x29, 0x25, 0x95, 0xf6, 0x1e, 0xd9,
	0xeb, 0xb3, 0x54, 0xb4, 0x6c, 0x6a, 0x22, 0xca, 0x8a, 0x32, 0x52, 0x76, 0x03, 0xd8, 0x28, 0xb9,
	0x96, 0xce, 0xc8, 0x20, 0x90, 0x50, 0x0d, 0x77, 0x34, 0xdc, 0xf8, 0xd0, 0xd0, 0xe7, 0x17, 0xdb,
	0xf4, 0xa6, 0xf2, 0x88, 0x10, 0x72, 0x4d, 0xd6, 0x95, 0x14, 0xda, 0xd0, 0xee, 0x0b, 0x38, 0x9d,
	0x9b, 0xf8, 0xa4, 0x26, 0x22, 0xa4, 0x8c, 0x04, 0x2d, 0x16, 0x88, 0xa7, 0x95, 0x5b, 0x80, 0x93,
	0x7d, 0x33, 0x67, 0x08, 0x06, 0x59, 0x7c, 0x9b, 0x04, 0x93, 0xf0, 0x26, 0x0c, 0xa6, 0xf6, 0x3f,
	0x67, 0x00, 0x0e, 0xb2, 0x78, 0x1e, 0x2f, 0x96, 0xb1, 0xdd, 0x73, 0x8e, 0xc1, 0x30, 0x4b, 0xa3,
	0x72, 0x92, 0xe2, 0x65, 0x54, 0x06, 0x69, 0xba, 0x48, 0x6d, 0xcb, 0x39, 0x02, 0x87, 0x61, 0x9c,
	0xe3, 0x28, 0x9c, 0x96, 0x39, 0x8e, 0xb2, 0xc0, 0xee, 0x8f, 0xbf, 0x7a, 0xc0, 0x7d, 0x90, 0x2b,
	0xf8, 0x77, 0xfd, 0xf1, 0xd9, 0xbe, 0x0a, 0x49, 0xdb, 0x3d, 0xe9, 0xdd, 0x4d, 0x7f, 0x60, 0x2e,
	0x6b, 0x22, 0x38, 0x94, 0x8a, 0x7b, 0x9c, 0x89, 0xee, 0x67, 0xdb, 0x4d, 0x36, 0x95, 0xfe, 0x6d,
	0xb1, 0x57, 0xe6, 0x78, 0xb3, 0xfa, 0x33, 0x8c, 0xdf, 0xad, 0xd1, 0xcc, 0x84, 0x61, 0xaa, 0xa1,
	0x91, 0xad, 0xca, 0x7d, 0xd8, 0x3d, 0xa9, 0x3f, 0xb6, 0x86, 0x02, 0x53, 0x5d, 0xec, 0x0c, 0x45,
	0xee, 0x17, 0xc6, 0xf0, 0x69, 0xb9, 0xe6, 0x16, 0x21, 0x4c, 0x35, 0x42, 0x3b, 0x0b, 0x42, 0xb9,
	0x8f, 0x90, 0x31, 0xdd, 0xff, 0xef, 0xda, 0x5d, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x66, 0x3f,
	0x4b, 0xfb, 0xf5, 0x01, 0x00, 0x00,
}
