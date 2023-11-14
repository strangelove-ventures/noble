// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tokenfactory/blacklisted.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Blacklisted struct {
	AddressBz []byte `protobuf:"bytes,1,opt,name=addressBz,proto3" json:"addressBz,omitempty"`
}

func (m *Blacklisted) Reset()         { *m = Blacklisted{} }
func (m *Blacklisted) String() string { return proto.CompactTextString(m) }
func (*Blacklisted) ProtoMessage()    {}
func (*Blacklisted) Descriptor() ([]byte, []int) {
	return fileDescriptor_43ff59c42df01ab4, []int{0}
}
func (m *Blacklisted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Blacklisted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Blacklisted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Blacklisted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blacklisted.Merge(m, src)
}
func (m *Blacklisted) XXX_Size() int {
	return m.Size()
}
func (m *Blacklisted) XXX_DiscardUnknown() {
	xxx_messageInfo_Blacklisted.DiscardUnknown(m)
}

var xxx_messageInfo_Blacklisted proto.InternalMessageInfo

func (m *Blacklisted) GetAddressBz() []byte {
	if m != nil {
		return m.AddressBz
	}
	return nil
}

func init() {
	proto.RegisterType((*Blacklisted)(nil), "noble.tokenfactory.Blacklisted")
}

func init() { proto.RegisterFile("tokenfactory/blacklisted.proto", fileDescriptor_43ff59c42df01ab4) }

var fileDescriptor_43ff59c42df01ab4 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0xc9, 0xcf, 0x4e,
	0xcd, 0x4b, 0x4b, 0x4c, 0x2e, 0xc9, 0x2f, 0xaa, 0xd4, 0x4f, 0xca, 0x49, 0x4c, 0xce, 0xce, 0xc9,
	0x2c, 0x2e, 0x49, 0x4d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xca, 0xcb, 0x4f, 0xca,
	0x49, 0xd5, 0x43, 0x56, 0xa5, 0xa4, 0xcd, 0xc5, 0xed, 0x84, 0x50, 0x28, 0x24, 0xc3, 0xc5, 0x99,
	0x98, 0x92, 0x52, 0x94, 0x5a, 0x5c, 0xec, 0x54, 0x25, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x13, 0x84,
	0x10, 0x70, 0x0a, 0x3f, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18,
	0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xdb, 0xf4,
	0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xe2, 0x92, 0xa2, 0xc4, 0xbc, 0xf4,
	0xd4, 0x9c, 0xfc, 0xb2, 0x54, 0xdd, 0xb2, 0xd4, 0xbc, 0x92, 0xd2, 0xa2, 0xd4, 0x62, 0x7d, 0xb0,
	0xd5, 0xfa, 0x65, 0xa6, 0xfa, 0x15, 0xfa, 0x28, 0xae, 0x2c, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62,
	0x03, 0x3b, 0xd0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x29, 0xdd, 0x40, 0x60, 0xc2, 0x00, 0x00,
	0x00,
}

func (m *Blacklisted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Blacklisted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Blacklisted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AddressBz) > 0 {
		i -= len(m.AddressBz)
		copy(dAtA[i:], m.AddressBz)
		i = encodeVarintBlacklisted(dAtA, i, uint64(len(m.AddressBz)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlacklisted(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlacklisted(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Blacklisted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AddressBz)
	if l > 0 {
		n += 1 + l + sovBlacklisted(uint64(l))
	}
	return n
}

func sovBlacklisted(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlacklisted(x uint64) (n int) {
	return sovBlacklisted(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Blacklisted) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlacklisted
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
			return fmt.Errorf("proto: Blacklisted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Blacklisted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressBz", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlacklisted
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlacklisted
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlacklisted
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AddressBz = append(m.AddressBz[:0], dAtA[iNdEx:postIndex]...)
			if m.AddressBz == nil {
				m.AddressBz = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlacklisted(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlacklisted
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBlacklisted(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlacklisted
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
					return 0, ErrIntOverflowBlacklisted
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
					return 0, ErrIntOverflowBlacklisted
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
				return 0, ErrInvalidLengthBlacklisted
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlacklisted
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlacklisted
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlacklisted        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlacklisted          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlacklisted = fmt.Errorf("proto: unexpected end of group")
)
