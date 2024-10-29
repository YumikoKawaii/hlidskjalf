// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mockers/echo/service.proto

package api

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/types"
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

type EchoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6d45a56d7927e39, []int{0}
}
func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(m, src)
}
func (m *EchoRequest) XXX_Size() int {
	return m.Size()
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

type Stat struct {
	First                string   `protobuf:"bytes,1,opt,name=first,proto3" json:"first,omitempty"`
	Second               string   `protobuf:"bytes,2,opt,name=second,proto3" json:"second,omitempty"`
	Third                string   `protobuf:"bytes,3,opt,name=third,proto3" json:"third,omitempty"`
	Fourth               string   `protobuf:"bytes,4,opt,name=fourth,proto3" json:"fourth,omitempty"`
	Fifth                string   `protobuf:"bytes,5,opt,name=fifth,proto3" json:"fifth,omitempty"`
	Sixth                string   `protobuf:"bytes,6,opt,name=sixth,proto3" json:"sixth,omitempty"`
	Seventh              string   `protobuf:"bytes,7,opt,name=seventh,proto3" json:"seventh,omitempty"`
	Eighth               string   `protobuf:"bytes,8,opt,name=eighth,proto3" json:"eighth,omitempty"`
	Ninth                string   `protobuf:"bytes,9,opt,name=ninth,proto3" json:"ninth,omitempty"`
	Tenth                string   `protobuf:"bytes,10,opt,name=tenth,proto3" json:"tenth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stat) Reset()         { *m = Stat{} }
func (m *Stat) String() string { return proto.CompactTextString(m) }
func (*Stat) ProtoMessage()    {}
func (*Stat) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6d45a56d7927e39, []int{1}
}
func (m *Stat) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Stat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Stat.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Stat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stat.Merge(m, src)
}
func (m *Stat) XXX_Size() int {
	return m.Size()
}
func (m *Stat) XXX_DiscardUnknown() {
	xxx_messageInfo_Stat.DiscardUnknown(m)
}

var xxx_messageInfo_Stat proto.InternalMessageInfo

func (m *Stat) GetFirst() string {
	if m != nil {
		return m.First
	}
	return ""
}

func (m *Stat) GetSecond() string {
	if m != nil {
		return m.Second
	}
	return ""
}

func (m *Stat) GetThird() string {
	if m != nil {
		return m.Third
	}
	return ""
}

func (m *Stat) GetFourth() string {
	if m != nil {
		return m.Fourth
	}
	return ""
}

func (m *Stat) GetFifth() string {
	if m != nil {
		return m.Fifth
	}
	return ""
}

func (m *Stat) GetSixth() string {
	if m != nil {
		return m.Sixth
	}
	return ""
}

func (m *Stat) GetSeventh() string {
	if m != nil {
		return m.Seventh
	}
	return ""
}

func (m *Stat) GetEighth() string {
	if m != nil {
		return m.Eighth
	}
	return ""
}

func (m *Stat) GetNinth() string {
	if m != nil {
		return m.Ninth
	}
	return ""
}

func (m *Stat) GetTenth() string {
	if m != nil {
		return m.Tenth
	}
	return ""
}

type EchoResponse struct {
	Code                 uint32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string             `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *EchoResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6d45a56d7927e39, []int{2}
}
func (m *EchoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EchoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResponse.Merge(m, src)
}
func (m *EchoResponse) XXX_Size() int {
	return m.Size()
}
func (m *EchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResponse proto.InternalMessageInfo

func (m *EchoResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *EchoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *EchoResponse) GetData() *EchoResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type EchoResponse_Data struct {
	Stats                []*Stat  `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoResponse_Data) Reset()         { *m = EchoResponse_Data{} }
func (m *EchoResponse_Data) String() string { return proto.CompactTextString(m) }
func (*EchoResponse_Data) ProtoMessage()    {}
func (*EchoResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6d45a56d7927e39, []int{2, 0}
}
func (m *EchoResponse_Data) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EchoResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EchoResponse_Data.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EchoResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResponse_Data.Merge(m, src)
}
func (m *EchoResponse_Data) XXX_Size() int {
	return m.Size()
}
func (m *EchoResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResponse_Data proto.InternalMessageInfo

func (m *EchoResponse_Data) GetStats() []*Stat {
	if m != nil {
		return m.Stats
	}
	return nil
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "echo.api.EchoRequest")
	proto.RegisterType((*Stat)(nil), "echo.api.Stat")
	proto.RegisterType((*EchoResponse)(nil), "echo.api.EchoResponse")
	proto.RegisterType((*EchoResponse_Data)(nil), "echo.api.EchoResponse.Data")
}

func init() { proto.RegisterFile("mockers/echo/service.proto", fileDescriptor_a6d45a56d7927e39) }

var fileDescriptor_a6d45a56d7927e39 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xb1, 0x6e, 0xd4, 0x40,
	0x10, 0x86, 0x71, 0xe2, 0xbb, 0x24, 0x7b, 0x49, 0x8a, 0x55, 0x08, 0x2b, 0x83, 0xac, 0xe8, 0x44,
	0x91, 0x02, 0x6c, 0x71, 0x94, 0x14, 0x48, 0x08, 0x2a, 0x3a, 0xa7, 0x4b, 0x83, 0x36, 0xf6, 0xf8,
	0x76, 0xc5, 0x9d, 0x77, 0xf1, 0xce, 0x1d, 0xd0, 0xf2, 0x0a, 0x34, 0x48, 0xbc, 0x10, 0x25, 0x12,
	0x2f, 0x80, 0x0e, 0x4a, 0x1e, 0x02, 0xed, 0x8c, 0x2d, 0x10, 0xa2, 0xf3, 0x37, 0xf3, 0xef, 0xbf,
	0x3b, 0x9e, 0x5f, 0x64, 0x6b, 0x57, 0xbf, 0x86, 0x3e, 0x94, 0x50, 0x1b, 0x57, 0x06, 0xe8, 0xb7,
	0xb6, 0x86, 0xc2, 0xf7, 0x0e, 0x9d, 0x3c, 0x8c, 0xb5, 0x42, 0x7b, 0x9b, 0xdd, 0x5b, 0x3a, 0xb7,
	0x5c, 0x41, 0xa9, 0xbd, 0x2d, 0x75, 0xd7, 0x39, 0xd4, 0x68, 0x5d, 0x17, 0x58, 0x97, 0xe5, 0x43,
	0x97, 0xe8, 0x66, 0xd3, 0x96, 0x6f, 0x7b, 0xed, 0x3d, 0xf4, 0x63, 0xff, 0xce, 0x56, 0xaf, 0x6c,
	0xa3, 0x11, 0xca, 0xf1, 0x83, 0x1b, 0xf3, 0x13, 0x31, 0x7b, 0x51, 0x1b, 0x57, 0xc1, 0x9b, 0x0d,
	0x04, 0x9c, 0xff, 0x4a, 0x44, 0x7a, 0x85, 0x1a, 0xe5, 0x99, 0x98, 0xb4, 0xb6, 0x0f, 0xa8, 0x92,
	0x8b, 0xe4, 0xf2, 0xa8, 0x62, 0x90, 0xe7, 0x62, 0x1a, 0xa0, 0x76, 0x5d, 0xa3, 0xf6, 0xa8, 0x3c,
	0x50, 0x54, 0xa3, 0xb1, 0x7d, 0xa3, 0xf6, 0x59, 0x4d, 0x10, 0xd5, 0xad, 0xdb, 0xf4, 0x68, 0x54,
	0xca, 0x6a, 0x26, 0xf6, 0x6e, 0xd1, 0xa8, 0xc9, 0xe8, 0xdd, 0x72, 0x35, 0xd8, 0x77, 0x68, 0xd4,
	0x94, 0xab, 0x04, 0x52, 0x89, 0x83, 0x00, 0x5b, 0xe8, 0xd0, 0xa8, 0x03, 0xaa, 0x8f, 0x18, 0xdd,
	0xc1, 0x2e, 0x0d, 0x1a, 0x75, 0xc8, 0xee, 0x4c, 0xd1, 0xa7, 0xb3, 0x51, 0x7f, 0xc4, 0x3e, 0x04,
	0xf4, 0x42, 0x72, 0x11, 0xc3, 0x0b, 0x23, 0xcc, 0x3f, 0x27, 0xe2, 0x98, 0xc7, 0x0f, 0xde, 0x75,
	0x01, 0xa4, 0x14, 0x69, 0xed, 0x1a, 0xa0, 0xa9, 0x4f, 0x2a, 0xfa, 0x8e, 0x4f, 0x58, 0x43, 0x08,
	0x7a, 0x09, 0xc3, 0xd4, 0x23, 0xca, 0x52, 0xa4, 0x8d, 0x46, 0x4d, 0x53, 0xcf, 0x16, 0x77, 0x8b,
	0x71, 0x59, 0xc5, 0xdf, 0x9e, 0xc5, 0x73, 0x8d, 0xba, 0x22, 0x61, 0xf6, 0x40, 0xa4, 0x91, 0xe4,
	0x7d, 0x31, 0x09, 0xa8, 0x31, 0xa8, 0xe4, 0x62, 0xff, 0x72, 0xb6, 0x38, 0xfd, 0x73, 0x32, 0xfe,
	0xfc, 0x8a, 0x9b, 0x8b, 0x6b, 0xde, 0xcd, 0x15, 0x27, 0x42, 0xbe, 0x14, 0x69, 0x44, 0x79, 0xfb,
	0xdf, 0x7b, 0x68, 0x75, 0xd9, 0xf9, 0xff, 0xaf, 0x9f, 0x9f, 0x7d, 0xf8, 0xf6, 0xf3, 0xe3, 0xde,
	0xa9, 0x3c, 0xa6, 0xe8, 0x6c, 0x1f, 0x51, 0xcc, 0x9e, 0x3d, 0xfd, 0xb2, 0xcb, 0x93, 0xaf, 0xbb,
	0x3c, 0xf9, 0xbe, 0xcb, 0x93, 0x4f, 0x3f, 0xf2, 0x5b, 0xd7, 0x0f, 0x61, 0xf5, 0x3e, 0xd8, 0xcd,
	0xba, 0xa8, 0xdd, 0xba, 0xd4, 0xde, 0xaf, 0x6c, 0xcd, 0x11, 0xa3, 0x03, 0xaf, 0x86, 0x5c, 0x46,
	0x93, 0x27, 0xda, 0xdb, 0x9b, 0x29, 0xe5, 0xe7, 0xf1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16,
	0x56, 0x1e, 0x90, 0xbe, 0x02, 0x00, 0x00,
}

func (m *EchoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EchoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EchoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *Stat) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Stat) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Stat) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Tenth) > 0 {
		i -= len(m.Tenth)
		copy(dAtA[i:], m.Tenth)
		i = encodeVarintService(dAtA, i, uint64(len(m.Tenth)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Ninth) > 0 {
		i -= len(m.Ninth)
		copy(dAtA[i:], m.Ninth)
		i = encodeVarintService(dAtA, i, uint64(len(m.Ninth)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Eighth) > 0 {
		i -= len(m.Eighth)
		copy(dAtA[i:], m.Eighth)
		i = encodeVarintService(dAtA, i, uint64(len(m.Eighth)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Seventh) > 0 {
		i -= len(m.Seventh)
		copy(dAtA[i:], m.Seventh)
		i = encodeVarintService(dAtA, i, uint64(len(m.Seventh)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Sixth) > 0 {
		i -= len(m.Sixth)
		copy(dAtA[i:], m.Sixth)
		i = encodeVarintService(dAtA, i, uint64(len(m.Sixth)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Fifth) > 0 {
		i -= len(m.Fifth)
		copy(dAtA[i:], m.Fifth)
		i = encodeVarintService(dAtA, i, uint64(len(m.Fifth)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Fourth) > 0 {
		i -= len(m.Fourth)
		copy(dAtA[i:], m.Fourth)
		i = encodeVarintService(dAtA, i, uint64(len(m.Fourth)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Third) > 0 {
		i -= len(m.Third)
		copy(dAtA[i:], m.Third)
		i = encodeVarintService(dAtA, i, uint64(len(m.Third)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Second) > 0 {
		i -= len(m.Second)
		copy(dAtA[i:], m.Second)
		i = encodeVarintService(dAtA, i, uint64(len(m.Second)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.First) > 0 {
		i -= len(m.First)
		copy(dAtA[i:], m.First)
		i = encodeVarintService(dAtA, i, uint64(len(m.First)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EchoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EchoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EchoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
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

func (m *EchoResponse_Data) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EchoResponse_Data) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EchoResponse_Data) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Stats) > 0 {
		for iNdEx := len(m.Stats) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Stats[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *EchoRequest) Size() (n int) {
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

func (m *Stat) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.First)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.Second)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.Third)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.Fourth)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.Fifth)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.Sixth)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.Seventh)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.Eighth)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.Ninth)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.Tenth)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *EchoResponse) Size() (n int) {
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
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *EchoResponse_Data) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Stats) > 0 {
		for _, e := range m.Stats {
			l = e.Size()
			n += 1 + l + sovService(uint64(l))
		}
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
func (m *EchoRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EchoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EchoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *Stat) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Stat: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Stat: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field First", wireType)
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
			m.First = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Second", wireType)
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
			m.Second = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Third", wireType)
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
			m.Third = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fourth", wireType)
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
			m.Fourth = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fifth", wireType)
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
			m.Fifth = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sixth", wireType)
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
			m.Sixth = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seventh", wireType)
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
			m.Seventh = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Eighth", wireType)
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
			m.Eighth = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ninth", wireType)
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
			m.Ninth = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tenth", wireType)
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
			m.Tenth = string(dAtA[iNdEx:postIndex])
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
func (m *EchoResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EchoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EchoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
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
			if m.Data == nil {
				m.Data = &EchoResponse_Data{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *EchoResponse_Data) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Data: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Data: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
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
			m.Stats = append(m.Stats, &Stat{})
			if err := m.Stats[len(m.Stats)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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