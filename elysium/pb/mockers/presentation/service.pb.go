// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mockers/presentation/service.proto

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

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f918166c71d7ae4, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f918166c71d7ae4, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "presentation.api.Request")
	proto.RegisterType((*Response)(nil), "presentation.api.Response")
}

func init() {
	proto.RegisterFile("mockers/presentation/service.proto", fileDescriptor_8f918166c71d7ae4)
}

var fileDescriptor_8f918166c71d7ae4 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0xd4, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0x07, 0x70, 0x53, 0xd6, 0xdd, 0xee, 0xc3, 0xb6, 0xdb, 0x69, 0xb5, 0x69, 0x90, 0x50, 0x72,
	0xea, 0x29, 0x83, 0x8a, 0x20, 0x78, 0x13, 0xac, 0x27, 0xa1, 0x26, 0xfe, 0xa0, 0x8a, 0x94, 0xe9,
	0xec, 0xeb, 0x3a, 0x98, 0x64, 0xa6, 0x33, 0x93, 0x68, 0xaf, 0xfe, 0x0b, 0x5e, 0xfc, 0x7b, 0x3c,
	0x79, 0x14, 0xfc, 0x07, 0x64, 0xf5, 0x0f, 0x91, 0x9d, 0x64, 0x96, 0x2a, 0x9e, 0x76, 0xf7, 0x36,
	0x8f, 0x6f, 0xf8, 0xe4, 0xbd, 0xc3, 0x7b, 0x90, 0x94, 0x92, 0xbf, 0x47, 0x6d, 0xa8, 0xd2, 0x68,
	0xb0, 0xb2, 0xcc, 0x0a, 0x59, 0x51, 0x83, 0xba, 0x11, 0x1c, 0x53, 0xa5, 0xa5, 0x95, 0x64, 0x74,
	0x35, 0x4b, 0x99, 0x12, 0xd1, 0xed, 0x89, 0x94, 0x93, 0x02, 0x29, 0x53, 0x82, 0xb2, 0xaa, 0x92,
	0x6d, 0x64, 0xda, 0xef, 0xa3, 0xb8, 0x4b, 0x5d, 0x75, 0x56, 0x9f, 0xd3, 0x0f, 0x9a, 0x29, 0x85,
	0xda, 0xe7, 0x7b, 0x0d, 0x2b, 0xc4, 0x98, 0x59, 0xa4, 0xfe, 0xd1, 0x06, 0xc9, 0x10, 0x06, 0x19,
	0x5e, 0xd4, 0x68, 0x6c, 0xf2, 0x00, 0xd6, 0x33, 0x34, 0x4a, 0x56, 0x06, 0x09, 0x81, 0x1e, 0x97,
	0x63, 0x0c, 0x83, 0x83, 0xe0, 0x70, 0x23, 0x73, 0x6f, 0x12, 0xc2, 0xa0, 0x44, 0x63, 0xd8, 0x04,
	0xc3, 0xb5, 0x83, 0xe0, 0x70, 0x98, 0xf9, 0xf2, 0xee, 0xd7, 0x3e, 0xec, 0x1c, 0x5f, 0x69, 0x38,
	0x6f, 0x67, 0x21, 0x39, 0xf4, 0x1e, 0xf3, 0x77, 0x92, 0xec, 0xa7, 0xff, 0x8e, 0x93, 0x76, 0x3f,
	0x8d, 0xa2, 0xff, 0x45, 0x6d, 0x13, 0xc9, 0xee, 0xa7, 0x1f, 0xbf, 0x3f, 0xaf, 0x6d, 0x92, 0x1b,
	0x6e, 0xe8, 0xe6, 0x0e, 0xc5, 0x19, 0x96, 0x43, 0xef, 0x88, 0xd5, 0x1f, 0x57, 0x86, 0x9e, 0xcf,
	0xb0, 0x57, 0xd0, 0x3f, 0x12, 0xdc, 0x8a, 0x85, 0x7b, 0xbd, 0xe5, 0xd8, 0x11, 0xd9, 0x9c, 0xb3,
	0x2d, 0xf7, 0x02, 0xae, 0x3f, 0x15, 0xa5, 0xe0, 0x8b, 0xba, 0x37, 0x9d, 0xbb, 0x45, 0x36, 0xbc,
	0x5b, 0x3a, 0xed, 0x04, 0x06, 0x2f, 0x85, 0xb6, 0x35, 0x2b, 0x16, 0x85, 0xf7, 0x1c, 0xbc, 0x4d,
	0xb6, 0x3c, 0xdc, 0x74, 0xde, 0x09, 0x0c, 0x8e, 0xb5, 0x28, 0x99, 0xbe, 0x5c, 0x19, 0xad, 0x3a,
	0xef, 0x2d, 0x0c, 0x73, 0xe4, 0xb2, 0x1a, 0x2f, 0x81, 0xef, 0x3b, 0x7c, 0x87, 0x6c, 0x7b, 0xdc,
	0xcc, 0xc5, 0x37, 0xb0, 0xfe, 0x1c, 0xb5, 0x15, 0x4b, 0xe8, 0xa1, 0xd3, 0x09, 0x19, 0x79, 0xdd,
	0x7a, 0xf0, 0x14, 0xe0, 0x59, 0xcd, 0x2c, 0xea, 0x6a, 0x09, 0x3e, 0x72, 0xfc, 0x2e, 0x21, 0x9e,
	0xbf, 0x98, 0x93, 0x8f, 0x9e, 0x7c, 0x9b, 0xc6, 0xc1, 0xf7, 0x69, 0x1c, 0xfc, 0x9c, 0xc6, 0xc1,
	0x97, 0x5f, 0xf1, 0xb5, 0xd7, 0xf7, 0xb1, 0xb8, 0x34, 0xa2, 0x2e, 0x53, 0x2e, 0x4b, 0xca, 0x94,
	0x2a, 0x04, 0x6f, 0x97, 0xfe, 0xaf, 0xcb, 0x71, 0xda, 0x5d, 0x8e, 0x19, 0xf9, 0x90, 0x29, 0x71,
	0xd6, 0x77, 0x9b, 0x7d, 0xef, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x50, 0xde, 0x2a, 0x68,
	0x04, 0x00, 0x00,
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *Request) Size() (n int) {
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

func (m *Response) Size() (n int) {
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

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *Response) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
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