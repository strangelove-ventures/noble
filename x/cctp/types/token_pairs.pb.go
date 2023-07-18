// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cctp/token_pairs.proto

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

type TokenPairs struct {
	RemoteDomain uint32 `protobuf:"varint,1,opt,name=remoteDomain,proto3" json:"remoteDomain,omitempty"`
	RemoteToken  string `protobuf:"bytes,2,opt,name=remoteToken,proto3" json:"remoteToken,omitempty"`
	LocalToken   string `protobuf:"bytes,3,opt,name=localToken,proto3" json:"localToken,omitempty"`
}

func (m *TokenPairs) Reset()         { *m = TokenPairs{} }
func (m *TokenPairs) String() string { return proto.CompactTextString(m) }
func (*TokenPairs) ProtoMessage()    {}
func (*TokenPairs) Descriptor() ([]byte, []int) {
	return fileDescriptor_71c9ea03026a4cec, []int{0}
}
func (m *TokenPairs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenPairs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenPairs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenPairs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenPairs.Merge(m, src)
}
func (m *TokenPairs) XXX_Size() int {
	return m.Size()
}
func (m *TokenPairs) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenPairs.DiscardUnknown(m)
}

var xxx_messageInfo_TokenPairs proto.InternalMessageInfo

func (m *TokenPairs) GetRemoteDomain() uint32 {
	if m != nil {
		return m.RemoteDomain
	}
	return 0
}

func (m *TokenPairs) GetRemoteToken() string {
	if m != nil {
		return m.RemoteToken
	}
	return ""
}

func (m *TokenPairs) GetLocalToken() string {
	if m != nil {
		return m.LocalToken
	}
	return ""
}

func init() {
	proto.RegisterType((*TokenPairs)(nil), "noble.cctp.TokenPairs")
}

func init() { proto.RegisterFile("cctp/token_pairs.proto", fileDescriptor_71c9ea03026a4cec) }

var fileDescriptor_71c9ea03026a4cec = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x4e, 0x2e, 0x29,
	0xd0, 0x2f, 0xc9, 0xcf, 0x4e, 0xcd, 0x8b, 0x2f, 0x48, 0xcc, 0x2c, 0x2a, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0xca, 0xcb, 0x4f, 0xca, 0x49, 0xd5, 0x03, 0xc9, 0x2a, 0x15, 0x71, 0x71,
	0x85, 0x80, 0x14, 0x04, 0x80, 0xe4, 0x85, 0x94, 0xb8, 0x78, 0x8a, 0x52, 0x73, 0xf3, 0x4b, 0x52,
	0x5d, 0xf2, 0x73, 0x13, 0x33, 0xf3, 0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x83, 0x50, 0xc4, 0x84,
	0x14, 0xb8, 0xb8, 0x21, 0x7c, 0xb0, 0x3e, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x64, 0x21,
	0x21, 0x39, 0x2e, 0xae, 0x9c, 0xfc, 0xe4, 0xc4, 0x1c, 0x88, 0x02, 0x66, 0xb0, 0x02, 0x24, 0x11,
	0x27, 0x9f, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2,
	0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x32, 0x4a, 0xcf, 0x2c,
	0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0x2e, 0x29, 0x4a, 0xcc, 0x4b, 0x4f, 0xcd,
	0xc9, 0x2f, 0x4b, 0xd5, 0x2d, 0x4b, 0xcd, 0x2b, 0x29, 0x2d, 0x4a, 0x2d, 0xd6, 0x07, 0xbb, 0x5c,
	0xbf, 0x42, 0x1f, 0xe2, 0xb3, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xa7, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x26, 0x94, 0xa1, 0xf2, 0xee, 0x00, 0x00, 0x00,
}

func (m *TokenPairs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenPairs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenPairs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LocalToken) > 0 {
		i -= len(m.LocalToken)
		copy(dAtA[i:], m.LocalToken)
		i = encodeVarintTokenPairs(dAtA, i, uint64(len(m.LocalToken)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RemoteToken) > 0 {
		i -= len(m.RemoteToken)
		copy(dAtA[i:], m.RemoteToken)
		i = encodeVarintTokenPairs(dAtA, i, uint64(len(m.RemoteToken)))
		i--
		dAtA[i] = 0x12
	}
	if m.RemoteDomain != 0 {
		i = encodeVarintTokenPairs(dAtA, i, uint64(m.RemoteDomain))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTokenPairs(dAtA []byte, offset int, v uint64) int {
	offset -= sovTokenPairs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenPairs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RemoteDomain != 0 {
		n += 1 + sovTokenPairs(uint64(m.RemoteDomain))
	}
	l = len(m.RemoteToken)
	if l > 0 {
		n += 1 + l + sovTokenPairs(uint64(l))
	}
	l = len(m.LocalToken)
	if l > 0 {
		n += 1 + l + sovTokenPairs(uint64(l))
	}
	return n
}

func sovTokenPairs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTokenPairs(x uint64) (n int) {
	return sovTokenPairs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenPairs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenPairs
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
			return fmt.Errorf("proto: TokenPairs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenPairs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteDomain", wireType)
			}
			m.RemoteDomain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenPairs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RemoteDomain |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenPairs
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
				return ErrInvalidLengthTokenPairs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenPairs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RemoteToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenPairs
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
				return ErrInvalidLengthTokenPairs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenPairs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LocalToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTokenPairs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTokenPairs
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
func skipTokenPairs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTokenPairs
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
					return 0, ErrIntOverflowTokenPairs
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
					return 0, ErrIntOverflowTokenPairs
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
				return 0, ErrInvalidLengthTokenPairs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTokenPairs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTokenPairs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTokenPairs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTokenPairs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTokenPairs = fmt.Errorf("proto: unexpected end of group")
)
