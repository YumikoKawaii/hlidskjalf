// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pawn/service.proto

package api

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	io "io"
	math "math"
	math_bits "math/bits"
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

// GreetRequest ...
type GreetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf98b6ac71b62bab, []int{0}
}
func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return m.Size()
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

// GreetResponse ...
type GreetResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResponse) Reset()         { *m = GreetResponse{} }
func (m *GreetResponse) String() string { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()    {}
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf98b6ac71b62bab, []int{1}
}
func (m *GreetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GreetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResponse.Merge(m, src)
}
func (m *GreetResponse) XXX_Size() int {
	return m.Size()
}
func (m *GreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResponse proto.InternalMessageInfo

func (m *GreetResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GreetResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetStudentsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStudentsRequest) Reset()         { *m = GetStudentsRequest{} }
func (m *GetStudentsRequest) String() string { return proto.CompactTextString(m) }
func (*GetStudentsRequest) ProtoMessage()    {}
func (*GetStudentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf98b6ac71b62bab, []int{2}
}
func (m *GetStudentsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetStudentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetStudentsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetStudentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStudentsRequest.Merge(m, src)
}
func (m *GetStudentsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetStudentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStudentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStudentsRequest proto.InternalMessageInfo

type GetStudentsResponse struct {
	Code                 uint32     `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Students             []*Student `protobuf:"bytes,3,rep,name=students,proto3" json:"students,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetStudentsResponse) Reset()         { *m = GetStudentsResponse{} }
func (m *GetStudentsResponse) String() string { return proto.CompactTextString(m) }
func (*GetStudentsResponse) ProtoMessage()    {}
func (*GetStudentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf98b6ac71b62bab, []int{3}
}
func (m *GetStudentsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetStudentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetStudentsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetStudentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStudentsResponse.Merge(m, src)
}
func (m *GetStudentsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetStudentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStudentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStudentsResponse proto.InternalMessageInfo

func (m *GetStudentsResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetStudentsResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetStudentsResponse) GetStudents() []*Student {
	if m != nil {
		return m.Students
	}
	return nil
}

type Student struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName             string   `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Age                  uint32   `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Sex                  string   `protobuf:"bytes,4,opt,name=sex,proto3" json:"sex,omitempty"`
	Major                string   `protobuf:"bytes,5,opt,name=major,proto3" json:"major,omitempty"`
	Level                string   `protobuf:"bytes,6,opt,name=level,proto3" json:"level,omitempty"`
	Gpa                  float32  `protobuf:"fixed32,7,opt,name=gpa,proto3" json:"gpa,omitempty"`
	Hobbies              []string `protobuf:"bytes,8,rep,name=hobbies,proto3" json:"hobbies,omitempty"`
	Country              string   `protobuf:"bytes,9,opt,name=country,proto3" json:"country,omitempty"`
	Province             string   `protobuf:"bytes,10,opt,name=province,proto3" json:"province,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Student) Reset()         { *m = Student{} }
func (m *Student) String() string { return proto.CompactTextString(m) }
func (*Student) ProtoMessage()    {}
func (*Student) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf98b6ac71b62bab, []int{4}
}
func (m *Student) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Student) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Student.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Student) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Student.Merge(m, src)
}
func (m *Student) XXX_Size() int {
	return m.Size()
}
func (m *Student) XXX_DiscardUnknown() {
	xxx_messageInfo_Student.DiscardUnknown(m)
}

var xxx_messageInfo_Student proto.InternalMessageInfo

func (m *Student) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Student) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *Student) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Student) GetSex() string {
	if m != nil {
		return m.Sex
	}
	return ""
}

func (m *Student) GetMajor() string {
	if m != nil {
		return m.Major
	}
	return ""
}

func (m *Student) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *Student) GetGpa() float32 {
	if m != nil {
		return m.Gpa
	}
	return 0
}

func (m *Student) GetHobbies() []string {
	if m != nil {
		return m.Hobbies
	}
	return nil
}

func (m *Student) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Student) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func init() {
	proto.RegisterType((*GreetRequest)(nil), "pawn.api.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "pawn.api.GreetResponse")
	proto.RegisterType((*GetStudentsRequest)(nil), "pawn.api.GetStudentsRequest")
	proto.RegisterType((*GetStudentsResponse)(nil), "pawn.api.GetStudentsResponse")
	proto.RegisterType((*Student)(nil), "pawn.api.Student")
}

func init() { proto.RegisterFile("pawn/service.proto", fileDescriptor_cf98b6ac71b62bab) }

var fileDescriptor_cf98b6ac71b62bab = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x49, 0xd3, 0x38, 0x5b, 0x52, 0xca, 0x50, 0xe8, 0xca, 0x94, 0xc8, 0xb2, 0x84, 0x94,
	0x0b, 0xb6, 0x08, 0x07, 0x0e, 0x88, 0x0b, 0x97, 0xde, 0x00, 0xb9, 0x37, 0x2e, 0xd5, 0xc6, 0x9e,
	0x9a, 0x45, 0xf6, 0xee, 0xe2, 0x5d, 0x1b, 0xca, 0x91, 0x5f, 0xe0, 0xc2, 0x27, 0x71, 0x44, 0xe2,
	0x07, 0x50, 0xe0, 0xc0, 0x67, 0xa0, 0x5d, 0xdb, 0x10, 0x54, 0x10, 0xb7, 0x79, 0x6f, 0x26, 0xef,
	0xed, 0x64, 0x9e, 0x09, 0x28, 0xf6, 0x46, 0x24, 0x1a, 0xeb, 0x96, 0x67, 0x18, 0xab, 0x5a, 0x1a,
	0x09, 0xbe, 0xe5, 0x62, 0xa6, 0x78, 0x70, 0x5c, 0x48, 0x59, 0x94, 0x98, 0x30, 0xc5, 0x13, 0x26,
	0x84, 0x34, 0xcc, 0x70, 0x29, 0x74, 0x37, 0x17, 0x1c, 0xb5, 0xac, 0xe4, 0x39, 0x33, 0x98, 0x0c,
	0x45, 0xd7, 0x88, 0xf6, 0xc9, 0xd5, 0x93, 0x1a, 0xd1, 0xa4, 0xf8, 0xba, 0x41, 0x6d, 0xa2, 0xc7,
	0x64, 0xde, 0x63, 0xad, 0xa4, 0xd0, 0x08, 0x40, 0x76, 0x32, 0x99, 0x23, 0xf5, 0x42, 0x6f, 0x39,
	0x49, 0x5d, 0x0d, 0x94, 0x4c, 0x2b, 0xd4, 0x9a, 0x15, 0x48, 0x47, 0xa1, 0xb7, 0x9c, 0xa5, 0x03,
	0x8c, 0x0e, 0x09, 0x9c, 0xa0, 0x39, 0x35, 0x4d, 0x8e, 0xc2, 0xe8, 0x41, 0xb4, 0x26, 0x37, 0xfe,
	0x60, 0xff, 0x22, 0x3d, 0xff, 0x9f, 0x34, 0xdc, 0x23, 0xbe, 0xee, 0x15, 0xe8, 0x38, 0x1c, 0x2f,
	0xf7, 0x56, 0xd7, 0xe3, 0x61, 0xfb, 0xb8, 0xd7, 0x4e, 0x7f, 0x8d, 0x44, 0x3f, 0x3c, 0x32, 0xed,
	0x59, 0xd8, 0x27, 0x23, 0x9e, 0x3b, 0x9b, 0x59, 0x3a, 0xe2, 0x39, 0xdc, 0x26, 0xb3, 0xf3, 0xa6,
	0x2c, 0xcf, 0x04, 0xab, 0x06, 0x1b, 0xdf, 0x12, 0x4f, 0x59, 0x85, 0x70, 0x40, 0xc6, 0xd6, 0x7d,
	0xec, 0x1e, 0x65, 0x4b, 0xcb, 0x68, 0x7c, 0x4b, 0x77, 0xdc, 0xa0, 0x2d, 0xe1, 0x90, 0x4c, 0x2a,
	0xf6, 0x4a, 0xd6, 0x74, 0xe2, 0xb8, 0x0e, 0x58, 0xb6, 0xc4, 0x16, 0x4b, 0xba, 0xdb, 0xb1, 0x0e,
	0xd8, 0x5f, 0x17, 0x8a, 0xd1, 0x69, 0xe8, 0x2d, 0x47, 0xa9, 0x2d, 0xed, 0x8e, 0x2f, 0xe5, 0x7a,
	0xcd, 0x51, 0x53, 0x3f, 0x1c, 0xdb, 0x1d, 0x7b, 0x68, 0x3b, 0x99, 0x6c, 0x84, 0xa9, 0x2f, 0xe8,
	0xac, 0xdb, 0xbe, 0x87, 0x10, 0x10, 0x5f, 0xd5, 0xb2, 0xe5, 0x22, 0x43, 0x4a, 0xba, 0x17, 0x0f,
	0x78, 0x75, 0xd6, 0xdf, 0xf0, 0xb4, 0x8b, 0x06, 0x3c, 0x23, 0x13, 0x87, 0xe1, 0xd6, 0xef, 0x3f,
	0x68, 0xfb, 0xc8, 0xc1, 0xd1, 0x25, 0xbe, 0xbb, 0x48, 0x74, 0xf3, 0xfd, 0x97, 0xef, 0x1f, 0x46,
	0xd7, 0x60, 0xee, 0x62, 0xd4, 0xde, 0x4f, 0x0a, 0xdb, 0x5e, 0xbd, 0x23, 0xf0, 0x1c, 0xeb, 0x73,
	0x59, 0x57, 0x4c, 0x64, 0x38, 0xd8, 0xe4, 0x64, 0x6f, 0xeb, 0xaa, 0x70, 0xbc, 0x25, 0x7a, 0x29,
	0x02, 0xc1, 0x9d, 0x7f, 0x74, 0x7b, 0x63, 0xea, 0x8c, 0x01, 0x0e, 0x06, 0xe3, 0xe1, 0x8e, 0x4f,
	0x1e, 0x7e, 0xda, 0x2c, 0xbc, 0xcf, 0x9b, 0x85, 0xf7, 0x75, 0xb3, 0xf0, 0x3e, 0x7e, 0x5b, 0x5c,
	0x79, 0x71, 0x17, 0xcb, 0x0b, 0xcd, 0x9b, 0x2a, 0xce, 0x64, 0x95, 0x30, 0xa5, 0x4a, 0x9e, 0x75,
	0x59, 0x4f, 0xdc, 0x07, 0xc2, 0x14, 0x7f, 0xc4, 0x14, 0x5f, 0xef, 0xba, 0x80, 0x3f, 0xf8, 0x19,
	0x00, 0x00, 0xff, 0xff, 0xc3, 0x48, 0x12, 0x28, 0x37, 0x03, 0x00, 0x00,
}

func (m *GreetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GreetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GreetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *GreetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GreetResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GreetResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintService(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetStudentsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetStudentsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetStudentsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *GetStudentsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetStudentsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetStudentsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Students) > 0 {
		for iNdEx := len(m.Students) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Students[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintService(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Student) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Student) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Student) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Province) > 0 {
		i -= len(m.Province)
		copy(dAtA[i:], m.Province)
		i = encodeVarintService(dAtA, i, uint64(len(m.Province)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Country) > 0 {
		i -= len(m.Country)
		copy(dAtA[i:], m.Country)
		i = encodeVarintService(dAtA, i, uint64(len(m.Country)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Hobbies) > 0 {
		for iNdEx := len(m.Hobbies) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Hobbies[iNdEx])
			copy(dAtA[i:], m.Hobbies[iNdEx])
			i = encodeVarintService(dAtA, i, uint64(len(m.Hobbies[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	if m.Gpa != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Gpa))))
		i--
		dAtA[i] = 0x3d
	}
	if len(m.Level) > 0 {
		i -= len(m.Level)
		copy(dAtA[i:], m.Level)
		i = encodeVarintService(dAtA, i, uint64(len(m.Level)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Major) > 0 {
		i -= len(m.Major)
		copy(dAtA[i:], m.Major)
		i = encodeVarintService(dAtA, i, uint64(len(m.Major)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Sex) > 0 {
		i -= len(m.Sex)
		copy(dAtA[i:], m.Sex)
		i = encodeVarintService(dAtA, i, uint64(len(m.Sex)))
		i--
		dAtA[i] = 0x22
	}
	if m.Age != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.Age))
		i--
		dAtA[i] = 0x18
	}
	if len(m.FullName) > 0 {
		i -= len(m.FullName)
		copy(dAtA[i:], m.FullName)
		i = encodeVarintService(dAtA, i, uint64(len(m.FullName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintService(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	offset -= sovService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GreetRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GreetResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovService(uint64(m.Code))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetStudentsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetStudentsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovService(uint64(m.Code))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	if len(m.Students) > 0 {
		for _, e := range m.Students {
			l = e.Size()
			n += 1 + l + sovService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Student) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.FullName)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	if m.Age != 0 {
		n += 1 + sovService(uint64(m.Age))
	}
	l = len(m.Sex)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.Major)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.Level)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	if m.Gpa != 0 {
		n += 5
	}
	if len(m.Hobbies) > 0 {
		for _, s := range m.Hobbies {
			l = len(s)
			n += 1 + l + sovService(uint64(l))
		}
	}
	l = len(m.Country)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.Province)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GreetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GreetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GreetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GreetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GreetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GreetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetStudentsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetStudentsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetStudentsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetStudentsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetStudentsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetStudentsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Students", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Students = append(m.Students, &Student{})
			if err := m.Students[len(m.Students)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Student) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Student: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Student: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FullName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FullName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Age", wireType)
			}
			m.Age = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Age |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Major", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Major = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Level = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gpa", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Gpa = float32(math.Float32frombits(v))
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hobbies", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hobbies = append(m.Hobbies, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Country", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Country = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Province", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Province = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupService = fmt.Errorf("proto: unexpected end of group")
)
