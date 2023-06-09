// Code generated by protoc-gen-go. DO NOT EDIT.
// source: person.proto

package communication

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Person struct {
	Id                   *int32   `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name                 *string  `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Family               *string  `protobuf:"bytes,3,opt,name=family" json:"family,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Person) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Person) GetFamily() string {
	if m != nil && m.Family != nil {
		return *m.Family
	}
	return ""
}

func init() {
	proto.RegisterType((*Person)(nil), "communication.Person")
}

func init() { proto.RegisterFile("person.proto", fileDescriptor_4c9e10cf24b1156d) }

var fileDescriptor_4c9e10cf24b1156d = []byte{
	// 102 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x48, 0x2d, 0x2a,
	0xce, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4d, 0xce, 0xcf, 0xcd, 0x2d, 0xcd,
	0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x53, 0x32, 0xe2, 0x62, 0x0b, 0x00, 0x4b, 0x0b, 0x71,
	0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x69, 0xb0, 0x0a, 0xf1, 0x70, 0xb1, 0xe4, 0x25,
	0xe6, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x69, 0x70, 0x0a, 0xf1, 0x71, 0xb1, 0xa5, 0x25, 0xe6, 0x66,
	0xe6, 0x54, 0x4a, 0x30, 0x2b, 0x30, 0x6a, 0x70, 0x02, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x51,
	0x60, 0x05, 0x51, 0x00, 0x00, 0x00,
}
