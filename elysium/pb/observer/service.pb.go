// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: observer/service.proto

package api

import (
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

type EntryRequest struct {
	CurrentSession       string   `protobuf:"bytes,1,opt,name=current_session,json=currentSession,proto3" json:"current_session,omitempty"`
	PreviousSession      string   `protobuf:"bytes,2,opt,name=previous_session,json=previousSession,proto3" json:"previous_session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntryRequest) Reset()         { *m = EntryRequest{} }
func (m *EntryRequest) String() string { return proto.CompactTextString(m) }
func (*EntryRequest) ProtoMessage()    {}
func (*EntryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f54fe327b4fa0fce, []int{0}
}
func (m *EntryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EntryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EntryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EntryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntryRequest.Merge(m, src)
}
func (m *EntryRequest) XXX_Size() int {
	return m.Size()
}
func (m *EntryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EntryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EntryRequest proto.InternalMessageInfo

func (m *EntryRequest) GetCurrentSession() string {
	if m != nil {
		return m.CurrentSession
	}
	return ""
}

func (m *EntryRequest) GetPreviousSession() string {
	if m != nil {
		return m.PreviousSession
	}
	return ""
}

type EntryResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntryResponse) Reset()         { *m = EntryResponse{} }
func (m *EntryResponse) String() string { return proto.CompactTextString(m) }
func (*EntryResponse) ProtoMessage()    {}
func (*EntryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f54fe327b4fa0fce, []int{1}
}
func (m *EntryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EntryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EntryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EntryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntryResponse.Merge(m, src)
}
func (m *EntryResponse) XXX_Size() int {
	return m.Size()
}
func (m *EntryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EntryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EntryResponse proto.InternalMessageInfo

func (m *EntryResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *EntryResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ViewStreamRequest struct {
	Session              string   `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewStreamRequest) Reset()         { *m = ViewStreamRequest{} }
func (m *ViewStreamRequest) String() string { return proto.CompactTextString(m) }
func (*ViewStreamRequest) ProtoMessage()    {}
func (*ViewStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f54fe327b4fa0fce, []int{2}
}
func (m *ViewStreamRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ViewStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ViewStreamRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ViewStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewStreamRequest.Merge(m, src)
}
func (m *ViewStreamRequest) XXX_Size() int {
	return m.Size()
}
func (m *ViewStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ViewStreamRequest proto.InternalMessageInfo

func (m *ViewStreamRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

type ViewStreamResponse struct {
	View                 int64    `protobuf:"varint,1,opt,name=view,proto3" json:"view,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewStreamResponse) Reset()         { *m = ViewStreamResponse{} }
func (m *ViewStreamResponse) String() string { return proto.CompactTextString(m) }
func (*ViewStreamResponse) ProtoMessage()    {}
func (*ViewStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f54fe327b4fa0fce, []int{3}
}
func (m *ViewStreamResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ViewStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ViewStreamResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ViewStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewStreamResponse.Merge(m, src)
}
func (m *ViewStreamResponse) XXX_Size() int {
	return m.Size()
}
func (m *ViewStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ViewStreamResponse proto.InternalMessageInfo

func (m *ViewStreamResponse) GetView() int64 {
	if m != nil {
		return m.View
	}
	return 0
}

func init() {
	proto.RegisterType((*EntryRequest)(nil), "observer.api.EntryRequest")
	proto.RegisterType((*EntryResponse)(nil), "observer.api.EntryResponse")
	proto.RegisterType((*ViewStreamRequest)(nil), "observer.api.ViewStreamRequest")
	proto.RegisterType((*ViewStreamResponse)(nil), "observer.api.ViewStreamResponse")
}

func init() { proto.RegisterFile("observer/service.proto", fileDescriptor_f54fe327b4fa0fce) }

var fileDescriptor_f54fe327b4fa0fce = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xed, 0xba, 0x75, 0xdd, 0x2e, 0x76, 0xdd, 0xee, 0xc1, 0x15, 0x6a, 0x51, 0x8d, 0x2e, 0xb5,
	0x0b, 0xd5, 0xb6, 0xc9, 0x2d, 0x26, 0x97, 0x40, 0xce, 0x01, 0x19, 0x72, 0xf0, 0x25, 0xac, 0xe5,
	0xc1, 0x2c, 0x48, 0xda, 0xf5, 0xee, 0x4a, 0xc6, 0x84, 0x5c, 0xf2, 0x17, 0x72, 0xc9, 0x4f, 0xca,
	0x31, 0x90, 0x53, 0x6e, 0xc1, 0xc9, 0x0f, 0x09, 0x5e, 0x49, 0xfe, 0x20, 0xc9, 0x49, 0xb3, 0x6f,
	0x1e, 0xf3, 0xde, 0x3c, 0x0d, 0xee, 0x88, 0xb1, 0x06, 0x95, 0x83, 0xa2, 0xab, 0x0f, 0x8f, 0x20,
	0x90, 0x4a, 0x18, 0x41, 0x9a, 0x15, 0x1e, 0x30, 0xc9, 0xdd, 0x9f, 0x53, 0x21, 0xa6, 0x31, 0x50,
	0x26, 0x39, 0x65, 0x69, 0x2a, 0x0c, 0x33, 0x5c, 0xa4, 0xba, 0xe0, 0xba, 0xdf, 0x73, 0x16, 0xf3,
	0x09, 0x33, 0x40, 0xab, 0xa2, 0x68, 0xf8, 0x63, 0xdc, 0x3c, 0x4e, 0x8d, 0x5a, 0x84, 0x30, 0xcb,
	0x40, 0x1b, 0xf2, 0x1b, 0xb7, 0xa3, 0x4c, 0x29, 0x48, 0xcd, 0x99, 0x06, 0xad, 0xb9, 0x48, 0x1d,
	0xd4, 0x45, 0xbd, 0xcf, 0xe1, 0x97, 0x12, 0x1e, 0x16, 0x28, 0xe9, 0xe3, 0xaf, 0x52, 0x41, 0xce,
	0x45, 0xa6, 0xd7, 0xcc, 0x9a, 0x65, 0xb6, 0x2b, 0xbc, 0xa4, 0xfa, 0x87, 0xb8, 0x55, 0x6a, 0x68,
	0x29, 0x52, 0x0d, 0x84, 0xe0, 0x0f, 0x91, 0x98, 0x80, 0x9d, 0x5c, 0x0f, 0x6d, 0x4d, 0x1c, 0xdc,
	0x48, 0x40, 0x6b, 0x36, 0x85, 0x72, 0x4c, 0xf5, 0xf4, 0xff, 0xe2, 0x6f, 0xa7, 0x1c, 0xe6, 0x43,
	0xa3, 0x80, 0x25, 0x95, 0x4f, 0x07, 0x37, 0x76, 0xfd, 0x55, 0x4f, 0xbf, 0x87, 0xc9, 0x36, 0x7d,
	0x23, 0x99, 0x73, 0x98, 0x5b, 0xf2, 0xfb, 0xd0, 0xd6, 0x7b, 0xf7, 0x08, 0x7f, 0x3a, 0x29, 0x33,
	0x24, 0x23, 0x5c, 0xb7, 0x26, 0x89, 0x1b, 0x6c, 0xe7, 0x1a, 0x6c, 0xa7, 0xe3, 0xfe, 0x78, 0xb5,
	0x57, 0x48, 0xf8, 0xce, 0xe5, 0xdd, 0xd3, 0x55, 0x8d, 0xf8, 0x2d, 0xfb, 0x0f, 0xf2, 0xff, 0x14,
	0x56, 0xed, 0x03, 0xf4, 0x87, 0xcc, 0x30, 0xde, 0x58, 0x22, 0xbf, 0x76, 0x87, 0xbc, 0xd8, 0xcd,
	0xed, 0xbe, 0x4d, 0x28, 0xa5, 0x3c, 0x2b, 0xe5, 0x90, 0x4e, 0x25, 0xb5, 0xda, 0x87, 0x9e, 0x97,
	0x11, 0x5c, 0xfc, 0x43, 0x47, 0x83, 0x9b, 0xa5, 0x87, 0x6e, 0x97, 0x1e, 0x7a, 0x58, 0x7a, 0xe8,
	0xfa, 0xd1, 0x7b, 0x37, 0xea, 0x43, 0xbc, 0xd0, 0x3c, 0x4b, 0x82, 0x48, 0x24, 0x94, 0x49, 0x19,
	0xf3, 0xa8, 0x38, 0x11, 0xba, 0xbe, 0x2f, 0x26, 0xf9, 0x80, 0x49, 0x3e, 0xfe, 0x68, 0x6f, 0x63,
	0xff, 0x39, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xd8, 0x7b, 0x6b, 0x7a, 0x02, 0x00, 0x00,
}

func (m *EntryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EntryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.PreviousSession) > 0 {
		i -= len(m.PreviousSession)
		copy(dAtA[i:], m.PreviousSession)
		i = encodeVarintService(dAtA, i, uint64(len(m.PreviousSession)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CurrentSession) > 0 {
		i -= len(m.CurrentSession)
		copy(dAtA[i:], m.CurrentSession)
		i = encodeVarintService(dAtA, i, uint64(len(m.CurrentSession)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EntryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EntryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *ViewStreamRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ViewStreamRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ViewStreamRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Session) > 0 {
		i -= len(m.Session)
		copy(dAtA[i:], m.Session)
		i = encodeVarintService(dAtA, i, uint64(len(m.Session)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ViewStreamResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ViewStreamResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ViewStreamResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.View != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.View))
		i--
		dAtA[i] = 0x8
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
func (m *EntryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CurrentSession)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.PreviousSession)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *EntryResponse) Size() (n int) {
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

func (m *ViewStreamRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Session)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ViewStreamResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.View != 0 {
		n += 1 + sovService(uint64(m.View))
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
func (m *EntryRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EntryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EntryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentSession", wireType)
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
			m.CurrentSession = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousSession", wireType)
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
			m.PreviousSession = string(dAtA[iNdEx:postIndex])
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
func (m *EntryResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EntryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EntryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *ViewStreamRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ViewStreamRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ViewStreamRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Session", wireType)
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
			m.Session = string(dAtA[iNdEx:postIndex])
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
func (m *ViewStreamResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ViewStreamResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ViewStreamResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field View", wireType)
			}
			m.View = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.View |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
