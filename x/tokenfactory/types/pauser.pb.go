// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tokenfactory/pauser.proto

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

type Pauser struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *Pauser) Reset()         { *m = Pauser{} }
func (m *Pauser) String() string { return proto.CompactTextString(m) }
func (*Pauser) ProtoMessage()    {}
func (*Pauser) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e2f3360489cd3e, []int{0}
}
func (m *Pauser) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pauser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pauser.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pauser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pauser.Merge(m, src)
}
func (m *Pauser) XXX_Size() int {
	return m.Size()
}
func (m *Pauser) XXX_DiscardUnknown() {
	xxx_messageInfo_Pauser.DiscardUnknown(m)
}

var xxx_messageInfo_Pauser proto.InternalMessageInfo

func (m *Pauser) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*Pauser)(nil), "noble.tokenfactory.Pauser")
}

func init() { proto.RegisterFile("tokenfactory/pauser.proto", fileDescriptor_67e2f3360489cd3e) }

var fileDescriptor_67e2f3360489cd3e = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0xc9, 0xcf, 0x4e,
	0xcd, 0x4b, 0x4b, 0x4c, 0x2e, 0xc9, 0x2f, 0xaa, 0xd4, 0x2f, 0x48, 0x2c, 0x2d, 0x4e, 0x2d, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcc, 0x48, 0x2d, 0xca, 0xd7, 0x43, 0x96, 0x57, 0x52,
	0xe2, 0x62, 0x0b, 0x00, 0x2b, 0x11, 0x92, 0xe0, 0x62, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e,
	0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x9d, 0x82, 0x4f, 0x3c, 0x92, 0x63, 0xbc,
	0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63,
	0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x32, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f,
	0x57, 0xbf, 0xb8, 0xa4, 0x28, 0x31, 0x2f, 0x3d, 0x35, 0x27, 0xbf, 0x2c, 0x55, 0xb7, 0x2c, 0x35,
	0xaf, 0xa4, 0xb4, 0x28, 0xb5, 0x58, 0x1f, 0x64, 0xa1, 0x7e, 0x85, 0x3e, 0x8a, 0x93, 0x4a, 0x2a,
	0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x4e, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xcb,
	0x78, 0x61, 0xaf, 0x00, 0x00, 0x00,
}

func (m *Pauser) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pauser) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pauser) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPauser(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPauser(dAtA []byte, offset int, v uint64) int {
	offset -= sovPauser(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pauser) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPauser(uint64(l))
	}
	return n
}

func sovPauser(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPauser(x uint64) (n int) {
	return sovPauser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pauser) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPauser
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
			return fmt.Errorf("proto: Pauser: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pauser: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPauser
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
				return ErrInvalidLengthPauser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPauser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPauser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPauser
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
func skipPauser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPauser
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
					return 0, ErrIntOverflowPauser
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
					return 0, ErrIntOverflowPauser
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
				return 0, ErrInvalidLengthPauser
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPauser
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPauser
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPauser        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPauser          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPauser = fmt.Errorf("proto: unexpected end of group")
)
