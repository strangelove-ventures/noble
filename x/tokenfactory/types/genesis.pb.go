// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tokenfactory/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// GenesisState defines the tokenfactory module's genesis state.
type GenesisState struct {
	Params               Params             `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	BlacklistedList      []Blacklisted      `protobuf:"bytes,2,rep,name=blacklistedList,proto3" json:"blacklistedList"`
	Paused               *Paused            `protobuf:"bytes,3,opt,name=paused,proto3" json:"paused,omitempty"`
	MasterMinter         *MasterMinter      `protobuf:"bytes,4,opt,name=masterMinter,proto3" json:"masterMinter,omitempty"`
	MintersList          []Minters          `protobuf:"bytes,5,rep,name=mintersList,proto3" json:"mintersList"`
	Pauser               *Pauser            `protobuf:"bytes,6,opt,name=pauser,proto3" json:"pauser,omitempty"`
	Blacklister          *Blacklister       `protobuf:"bytes,7,opt,name=blacklister,proto3" json:"blacklister,omitempty"`
	Owner                *Owner             `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
	MinterControllerList []MinterController `protobuf:"bytes,10,rep,name=minterControllerList,proto3" json:"minterControllerList"`
	MintingDenom         *MintingDenom      `protobuf:"bytes,11,opt,name=mintingDenom,proto3" json:"mintingDenom,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_415d5acd9b7bd461, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetBlacklistedList() []Blacklisted {
	if m != nil {
		return m.BlacklistedList
	}
	return nil
}

func (m *GenesisState) GetPaused() *Paused {
	if m != nil {
		return m.Paused
	}
	return nil
}

func (m *GenesisState) GetMasterMinter() *MasterMinter {
	if m != nil {
		return m.MasterMinter
	}
	return nil
}

func (m *GenesisState) GetMintersList() []Minters {
	if m != nil {
		return m.MintersList
	}
	return nil
}

func (m *GenesisState) GetPauser() *Pauser {
	if m != nil {
		return m.Pauser
	}
	return nil
}

func (m *GenesisState) GetBlacklister() *Blacklister {
	if m != nil {
		return m.Blacklister
	}
	return nil
}

func (m *GenesisState) GetOwner() *Owner {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *GenesisState) GetMinterControllerList() []MinterController {
	if m != nil {
		return m.MinterControllerList
	}
	return nil
}

func (m *GenesisState) GetMintingDenom() *MintingDenom {
	if m != nil {
		return m.MintingDenom
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "noble.tokenfactory.GenesisState")
}

func init() { proto.RegisterFile("tokenfactory/genesis.proto", fileDescriptor_415d5acd9b7bd461) }

var fileDescriptor_415d5acd9b7bd461 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x3f, 0x8f, 0xd3, 0x30,
	0x18, 0xc6, 0x13, 0xee, 0x5a, 0x90, 0x73, 0x12, 0x92, 0x75, 0x83, 0x2f, 0x48, 0xb9, 0x08, 0xdd,
	0xd0, 0x85, 0x44, 0x2a, 0x0b, 0x62, 0xa3, 0xad, 0xc4, 0x42, 0x55, 0x14, 0x98, 0x18, 0xa8, 0x92,
	0xd4, 0x84, 0xa8, 0x89, 0x5d, 0xd9, 0x6e, 0xa1, 0xdf, 0x82, 0x4f, 0xc1, 0x67, 0xe9, 0xd8, 0x91,
	0x09, 0xa1, 0xf6, 0x8b, 0xa0, 0xd8, 0x26, 0x8d, 0x69, 0xda, 0x6e, 0x91, 0x9e, 0xdf, 0xf3, 0xe6,
	0x79, 0xff, 0x18, 0xb8, 0x82, 0xce, 0x31, 0xf9, 0x12, 0xa7, 0x82, 0xb2, 0x75, 0x98, 0x61, 0x82,
	0x79, 0xce, 0x83, 0x05, 0xa3, 0x82, 0x42, 0x48, 0x68, 0x52, 0xe0, 0xa0, 0x49, 0xb8, 0xb7, 0x19,
	0xcd, 0xa8, 0x94, 0xc3, 0xea, 0x4b, 0x91, 0xee, 0x9d, 0x51, 0x65, 0x11, 0xb3, 0xb8, 0xd4, 0x45,
	0x5c, 0xcf, 0x90, 0x92, 0x22, 0x4e, 0xe7, 0x45, 0xce, 0x05, 0x9e, 0x9d, 0xb0, 0x2e, 0x79, 0x2d,
	0xf9, 0x86, 0x54, 0xc6, 0x5c, 0x60, 0x36, 0x2d, 0x73, 0x22, 0x30, 0xd3, 0x84, 0x99, 0x5e, 0x49,
	0xfc, 0x74, 0x61, 0x76, 0x21, 0xd3, 0x3f, 0x1d, 0x19, 0x3a, 0xfd, 0x46, 0x6a, 0xe5, 0xa1, 0xe5,
	0x87, 0xd3, 0x94, 0x12, 0xc1, 0x68, 0x51, 0xd4, 0x94, 0x7f, 0x44, 0xe5, 0x24, 0x9b, 0xce, 0x30,
	0xa1, 0xa5, 0x22, 0x9e, 0xff, 0xec, 0x80, 0x9b, 0xb7, 0x6a, 0xd8, 0x1f, 0x44, 0x2c, 0x30, 0x7c,
	0x05, 0xba, 0x6a, 0x6c, 0xc8, 0xf6, 0xed, 0x9e, 0xd3, 0x77, 0x83, 0xe3, 0xe1, 0x07, 0xef, 0x25,
	0x31, 0xb8, 0xde, 0xfc, 0xbe, 0xb7, 0x22, 0xcd, 0xc3, 0x09, 0x78, 0xda, 0x98, 0xea, 0xbb, 0x9c,
	0x0b, 0xf4, 0xc8, 0xbf, 0xea, 0x39, 0xfd, 0xfb, 0xb6, 0x12, 0x83, 0x03, 0xaa, 0xeb, 0xfc, 0xef,
	0x86, 0xfd, 0x2a, 0x4a, 0xb5, 0x06, 0x74, 0x75, 0x2e, 0x4a, 0x45, 0x44, 0x9a, 0x84, 0x23, 0x70,
	0xa3, 0xf6, 0x33, 0x96, 0x23, 0x41, 0xd7, 0xd2, 0xe9, 0xb7, 0x39, 0xc7, 0x0d, 0x2e, 0x32, 0x5c,
	0x70, 0x08, 0x1c, 0xbd, 0x43, 0xd9, 0x46, 0x47, 0xb6, 0xf1, 0xac, 0xb5, 0x88, 0xc2, 0x74, 0x0b,
	0x4d, 0x57, 0x1d, 0x9f, 0xa1, 0xee, 0x85, 0xf8, 0x4c, 0xc7, 0x67, 0xf0, 0x0d, 0x70, 0x1a, 0x57,
	0x80, 0x1e, 0x4b, 0xe3, 0x85, 0xf9, 0xb1, 0xa8, 0xe9, 0x81, 0x21, 0xe8, 0xc8, 0x43, 0x41, 0x4f,
	0xa4, 0xf9, 0xae, 0xcd, 0x3c, 0xa9, 0x80, 0x48, 0x71, 0xf0, 0x33, 0xb8, 0x55, 0xb1, 0x87, 0xf5,
	0xf9, 0xc8, 0xae, 0x81, 0xec, 0xfa, 0xe1, 0x74, 0xd7, 0x07, 0x5e, 0xb7, 0xdf, 0x5a, 0x47, 0xae,
	0x44, 0x5d, 0xde, 0xa8, 0x3a, 0x3c, 0xe4, 0x9c, 0x59, 0x49, 0x83, 0x8b, 0x0c, 0xd7, 0xe0, 0xe3,
	0x66, 0xe7, 0xd9, 0xdb, 0x9d, 0x67, 0xff, 0xd9, 0x79, 0xf6, 0x8f, 0xbd, 0x67, 0x6d, 0xf7, 0x9e,
	0xf5, 0x6b, 0xef, 0x59, 0x9f, 0x5e, 0x67, 0xb9, 0xf8, 0xba, 0x4c, 0x82, 0x94, 0x96, 0x21, 0x17,
	0x2c, 0x26, 0x19, 0x2e, 0xe8, 0x0a, 0xbf, 0x58, 0x61, 0x22, 0x96, 0x0c, 0xf3, 0x50, 0xfe, 0x28,
	0xfc, 0x1e, 0x1a, 0x8f, 0x41, 0xac, 0x17, 0x98, 0x27, 0x5d, 0xf9, 0x0a, 0x5e, 0xfe, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0x0b, 0xbc, 0x21, 0xb3, 0x7e, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MintingDenom != nil {
		{
			size, err := m.MintingDenom.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	if len(m.MinterControllerList) > 0 {
		for iNdEx := len(m.MinterControllerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinterControllerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if m.Owner != nil {
		{
			size, err := m.Owner.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.Blacklister != nil {
		{
			size, err := m.Blacklister.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Pauser != nil {
		{
			size, err := m.Pauser.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.MintersList) > 0 {
		for iNdEx := len(m.MintersList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MintersList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.MasterMinter != nil {
		{
			size, err := m.MasterMinter.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Paused != nil {
		{
			size, err := m.Paused.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BlacklistedList) > 0 {
		for iNdEx := len(m.BlacklistedList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BlacklistedList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.BlacklistedList) > 0 {
		for _, e := range m.BlacklistedList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Paused != nil {
		l = m.Paused.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.MasterMinter != nil {
		l = m.MasterMinter.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.MintersList) > 0 {
		for _, e := range m.MintersList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Pauser != nil {
		l = m.Pauser.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Blacklister != nil {
		l = m.Blacklister.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Owner != nil {
		l = m.Owner.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.MinterControllerList) > 0 {
		for _, e := range m.MinterControllerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.MintingDenom != nil {
		l = m.MintingDenom.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlacklistedList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlacklistedList = append(m.BlacklistedList, Blacklisted{})
			if err := m.BlacklistedList[len(m.BlacklistedList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paused", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Paused == nil {
				m.Paused = &Paused{}
			}
			if err := m.Paused.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MasterMinter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MasterMinter == nil {
				m.MasterMinter = &MasterMinter{}
			}
			if err := m.MasterMinter.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintersList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintersList = append(m.MintersList, Minters{})
			if err := m.MintersList[len(m.MintersList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pauser", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pauser == nil {
				m.Pauser = &Pauser{}
			}
			if err := m.Pauser.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blacklister", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Blacklister == nil {
				m.Blacklister = &Blacklister{}
			}
			if err := m.Blacklister.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Owner == nil {
				m.Owner = &Owner{}
			}
			if err := m.Owner.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinterControllerList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinterControllerList = append(m.MinterControllerList, MinterController{})
			if err := m.MinterControllerList[len(m.MinterControllerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintingDenom", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MintingDenom == nil {
				m.MintingDenom = &MintingDenom{}
			}
			if err := m.MintingDenom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
