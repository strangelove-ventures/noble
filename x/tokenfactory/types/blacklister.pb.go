// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tokenfactory/blacklister.proto

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

type Blacklister struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *Blacklister) Reset()         { *m = Blacklister{} }
func (m *Blacklister) String() string { return proto.CompactTextString(m) }
func (*Blacklister) ProtoMessage()    {}
func (*Blacklister) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4e04641cbe52423, []int{0}
}
func (m *Blacklister) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Blacklister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Blacklister.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Blacklister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blacklister.Merge(m, src)
}
func (m *Blacklister) XXX_Size() int {
	return m.Size()
}
func (m *Blacklister) XXX_DiscardUnknown() {
	xxx_messageInfo_Blacklister.DiscardUnknown(m)
}

var xxx_messageInfo_Blacklister proto.InternalMessageInfo

func (m *Blacklister) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*Blacklister)(nil), "noble.tokenfactory.Blacklister")
}

func init() { proto.RegisterFile("tokenfactory/blacklister.proto", fileDescriptor_c4e04641cbe52423) }

var fileDescriptor_c4e04641cbe52423 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0xc9, 0xcf, 0x4e,
	0xcd, 0x4b, 0x4b, 0x4c, 0x2e, 0xc9, 0x2f, 0xaa, 0xd4, 0x4f, 0xca, 0x49, 0x4c, 0xce, 0xce, 0xc9,
	0x2c, 0x2e, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xca, 0xcb, 0x4f, 0xca,
	0x49, 0xd5, 0x43, 0x56, 0xa5, 0xa4, 0xce, 0xc5, 0xed, 0x84, 0x50, 0x28, 0x24, 0xc1, 0xc5, 0x9e,
	0x98, 0x92, 0x52, 0x94, 0x5a, 0x5c, 0x2c, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe3, 0x3a,
<<<<<<< HEAD
	0x85, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e,
	0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x6d, 0x7a, 0x66, 0x49,
	0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x71, 0x49, 0x51, 0x62, 0x5e, 0x7a, 0x6a, 0x4e,
	0x7e, 0x59, 0xaa, 0x6e, 0x59, 0x6a, 0x5e, 0x49, 0x69, 0x51, 0x6a, 0xb1, 0x3e, 0xd8, 0x5a, 0xfd,
	0x32, 0x13, 0xfd, 0x0a, 0x7d, 0x14, 0x17, 0x96, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x1d,
	0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x52, 0x60, 0x22, 0x0c, 0xbe, 0x00, 0x00, 0x00,
=======
	0xf9, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e,
	0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x69, 0x7a, 0x66, 0x49,
	0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xd8, 0x06, 0xdd, 0xc4, 0xe2, 0xe2, 0xd4, 0x92,
	0x62, 0x08, 0x47, 0xbf, 0xcc, 0x54, 0xbf, 0x42, 0x1f, 0xc5, 0x65, 0x25, 0x95, 0x05, 0xa9, 0xc5,
	0x49, 0x6c, 0x60, 0x47, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xd6, 0xfa, 0xda, 0xb6,
	0x00, 0x00, 0x00,
>>>>>>> a4ad980 (chore: rename module path (#283))
}

func (m *Blacklister) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Blacklister) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Blacklister) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintBlacklister(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlacklister(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlacklister(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Blacklister) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovBlacklister(uint64(l))
	}
	return n
}

func sovBlacklister(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlacklister(x uint64) (n int) {
	return sovBlacklister(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Blacklister) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlacklister
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
			return fmt.Errorf("proto: Blacklister: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Blacklister: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlacklister
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
				return ErrInvalidLengthBlacklister
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBlacklister
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlacklister(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlacklister
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
func skipBlacklister(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlacklister
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
					return 0, ErrIntOverflowBlacklister
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
					return 0, ErrIntOverflowBlacklister
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
				return 0, ErrInvalidLengthBlacklister
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlacklister
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlacklister
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlacklister        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlacklister          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlacklister = fmt.Errorf("proto: unexpected end of group")
)
