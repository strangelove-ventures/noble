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
	MinterControllerList []MinterController `protobuf:"bytes,9,rep,name=minterControllerList,proto3" json:"minterControllerList"`
	MintingDenom         *MintingDenom      `protobuf:"bytes,10,opt,name=mintingDenom,proto3" json:"mintingDenom,omitempty"`
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
<<<<<<< HEAD
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x3f, 0x8f, 0xd3, 0x30,
	0x18, 0xc6, 0x13, 0xee, 0x5a, 0xc0, 0x39, 0x09, 0xc9, 0xba, 0xc1, 0x17, 0xa4, 0x5c, 0x84, 0x6e,
	0xb8, 0x85, 0x44, 0x2a, 0x0b, 0x62, 0xa3, 0xad, 0xc4, 0x42, 0x55, 0x14, 0x98, 0x18, 0xa8, 0x92,
	0xd4, 0x84, 0xa8, 0x89, 0x5d, 0xd9, 0x6e, 0xa1, 0xdf, 0x82, 0x4f, 0xc1, 0x67, 0xe9, 0xd8, 0x91,
	0x09, 0xa1, 0xf6, 0x8b, 0xa0, 0xd8, 0x26, 0x8d, 0xaf, 0x69, 0xbb, 0x45, 0x7a, 0x7e, 0xcf, 0x9b,
	0xe7, 0xfd, 0x63, 0xe0, 0x0a, 0x3a, 0xc3, 0xe4, 0x6b, 0x9c, 0x0a, 0xca, 0x56, 0x61, 0x86, 0x09,
	0xe6, 0x39, 0x0f, 0xe6, 0x8c, 0x0a, 0x0a, 0x21, 0xa1, 0x49, 0x81, 0x83, 0x26, 0xe1, 0x5e, 0x67,
	0x34, 0xa3, 0x52, 0x0e, 0xab, 0x2f, 0x45, 0xba, 0x37, 0x46, 0x95, 0x79, 0xcc, 0xe2, 0x52, 0x17,
	0x71, 0x3d, 0x43, 0x4a, 0x8a, 0x38, 0x9d, 0x15, 0x39, 0x17, 0x78, 0x7a, 0xc4, 0xba, 0xe0, 0xb5,
	0xe4, 0x1b, 0x52, 0x19, 0x73, 0x81, 0xd9, 0xa4, 0xcc, 0x89, 0xc0, 0x4c, 0x13, 0x66, 0x7a, 0x25,
	0xf1, 0xe3, 0x85, 0xd9, 0x99, 0x4c, 0xff, 0x75, 0x64, 0xe8, 0xf4, 0x3b, 0xa9, 0x95, 0xbb, 0x96,
	0x1f, 0x4e, 0x52, 0x4a, 0x04, 0xa3, 0x45, 0x51, 0x53, 0xfe, 0x01, 0x95, 0x93, 0x6c, 0x32, 0xc5,
	0x84, 0x96, 0x8a, 0x78, 0xf1, 0xab, 0x03, 0xae, 0xde, 0xa9, 0x61, 0x7f, 0x14, 0xb1, 0xc0, 0xf0,
	0x35, 0xe8, 0xaa, 0xb1, 0x21, 0xdb, 0xb7, 0xef, 0x9d, 0x9e, 0x1b, 0x1c, 0x0e, 0x3f, 0xf8, 0x20,
	0x89, 0xfe, 0xe5, 0xfa, 0xcf, 0xad, 0x15, 0x69, 0x1e, 0x8e, 0xc1, 0xb3, 0xc6, 0x54, 0xdf, 0xe7,
	0x5c, 0xa0, 0x47, 0xfe, 0xc5, 0xbd, 0xd3, 0xbb, 0x6d, 0x2b, 0xd1, 0xdf, 0xa3, 0xba, 0xce, 0x43,
	0x37, 0xec, 0x55, 0x51, 0xaa, 0x35, 0xa0, 0x8b, 0x53, 0x51, 0x2a, 0x22, 0xd2, 0x24, 0x1c, 0x82,
	0x2b, 0xb5, 0x9f, 0x91, 0x1c, 0x09, 0xba, 0x94, 0x4e, 0xbf, 0xcd, 0x39, 0x6a, 0x70, 0x91, 0xe1,
	0x82, 0x03, 0xe0, 0xe8, 0x1d, 0xca, 0x36, 0x3a, 0xb2, 0x8d, 0xe7, 0xad, 0x45, 0x14, 0xa6, 0x5b,
	0x68, 0xba, 0xea, 0xf8, 0x0c, 0x75, 0xcf, 0xc4, 0x67, 0x3a, 0x3e, 0x83, 0x6f, 0x81, 0xd3, 0xb8,
	0x02, 0xf4, 0x58, 0x1a, 0xcf, 0xcc, 0x8f, 0x45, 0x4d, 0x0f, 0x0c, 0x41, 0x47, 0x1e, 0x0a, 0x7a,
	0x22, 0xcd, 0x37, 0x6d, 0xe6, 0x71, 0x05, 0x44, 0x8a, 0x83, 0x5f, 0xc0, 0xb5, 0x8a, 0x3d, 0xa8,
	0xcf, 0x47, 0x76, 0xfd, 0x54, 0x76, 0x7d, 0x77, 0xbc, 0xeb, 0x3d, 0xaf, 0xdb, 0x6f, 0xad, 0x23,
	0x57, 0xa2, 0x2e, 0x6f, 0x58, 0x1d, 0x1e, 0x02, 0x27, 0x56, 0xd2, 0xe0, 0x22, 0xc3, 0xd5, 0xff,
	0xb4, 0xde, 0x7a, 0xf6, 0x66, 0xeb, 0xd9, 0x7f, 0xb7, 0x9e, 0xfd, 0x73, 0xe7, 0x59, 0x9b, 0x9d,
	0x67, 0xfd, 0xde, 0x79, 0xd6, 0xe7, 0x37, 0x59, 0x2e, 0xbe, 0x2d, 0x92, 0x20, 0xa5, 0x65, 0xc8,
	0x05, 0x8b, 0x49, 0x86, 0x0b, 0xba, 0xc4, 0x2f, 0x97, 0x98, 0x88, 0x05, 0xc3, 0x3c, 0x94, 0x3f,
	0x0a, 0x7f, 0x84, 0xc6, 0x63, 0x10, 0xab, 0x39, 0xe6, 0x49, 0x57, 0xbe, 0x82, 0x57, 0xff, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x87, 0x27, 0x28, 0x63, 0x7e, 0x04, 0x00, 0x00,
=======
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x41, 0x8e, 0xd3, 0x30,
	0x18, 0x85, 0x13, 0x66, 0x5a, 0xc0, 0x19, 0x09, 0xc9, 0x9a, 0x85, 0x27, 0x48, 0x99, 0x08, 0xcd,
	0x62, 0x36, 0x24, 0x52, 0xd1, 0x48, 0x6c, 0x69, 0x2b, 0xb1, 0xa1, 0x2a, 0x0a, 0x3b, 0x16, 0x54,
	0x49, 0x6a, 0x42, 0xd4, 0xc4, 0xae, 0x6c, 0x17, 0xe8, 0x2d, 0x38, 0x05, 0x67, 0xe9, 0xb2, 0x4b,
	0x56, 0x08, 0xb5, 0x17, 0x41, 0xb1, 0xad, 0x34, 0x06, 0xb7, 0xdd, 0xcd, 0xe8, 0x7d, 0xef, 0xf5,
	0x7f, 0xbf, 0xff, 0x00, 0x5f, 0xd0, 0x05, 0x26, 0x9f, 0xd3, 0x5c, 0x50, 0xb6, 0x8e, 0x0b, 0x4c,
	0x30, 0x2f, 0x79, 0xb4, 0x64, 0x54, 0x50, 0x08, 0x09, 0xcd, 0x2a, 0x1c, 0x75, 0x09, 0xff, 0xba,
	0xa0, 0x05, 0x95, 0x72, 0xdc, 0xfc, 0xa5, 0x48, 0x3f, 0x30, 0x52, 0xb2, 0x2a, 0xcd, 0x17, 0x55,
	0xc9, 0x05, 0x9e, 0x9f, 0xd1, 0x99, 0xd6, 0x43, 0x43, 0xaf, 0xd3, 0x46, 0x9a, 0xd5, 0x25, 0x39,
	0x10, 0x77, 0x26, 0x21, 0xa5, 0x59, 0x4e, 0x89, 0x60, 0xb4, 0xaa, 0x5a, 0xca, 0xb7, 0x50, 0xdc,
	0xfe, 0x1b, 0x25, 0x11, 0x25, 0x29, 0x66, 0x73, 0x4c, 0x68, 0xad, 0x09, 0x64, 0x10, 0xf4, 0x1b,
	0x69, 0x73, 0x6f, 0x0c, 0x65, 0x99, 0xb2, 0xb4, 0xe6, 0x47, 0xa4, 0x15, 0x6f, 0x5b, 0x5b, 0x24,
	0x1d, 0xf8, 0xe2, 0x67, 0x0f, 0x5c, 0xbd, 0x55, 0xcb, 0xfe, 0x20, 0x52, 0x81, 0xe1, 0x6b, 0xd0,
	0x57, 0xb1, 0xc8, 0x0d, 0xdd, 0x7b, 0x6f, 0xe0, 0x47, 0xff, 0x2f, 0x3f, 0x7a, 0x2f, 0x89, 0xe1,
	0xe5, 0xe6, 0xf7, 0xad, 0x93, 0x68, 0x1e, 0x4e, 0xc1, 0xb3, 0xce, 0xc2, 0xdf, 0x95, 0x5c, 0xa0,
	0x47, 0xe1, 0xc5, 0xbd, 0x37, 0xb8, 0xb5, 0x45, 0x0c, 0x0f, 0xa8, 0xce, 0xf9, 0xd7, 0x0d, 0x07,
	0xcd, 0x28, 0x4d, 0x0d, 0x74, 0x71, 0x6a, 0x94, 0x86, 0x48, 0x34, 0x09, 0xc7, 0xe0, 0x4a, 0xbd,
	0xda, 0x44, 0xee, 0x1c, 0x5d, 0x4a, 0x67, 0x68, 0x73, 0x4e, 0x3a, 0x5c, 0x62, 0xb8, 0xe0, 0x08,
	0x78, 0xfa, 0xcd, 0x64, 0x8d, 0x9e, 0xac, 0xf1, 0xdc, 0x1a, 0xa2, 0x30, 0x5d, 0xa1, 0xeb, 0x6a,
	0xc7, 0x67, 0xa8, 0x7f, 0x66, 0x7c, 0xa6, 0xc7, 0x67, 0xf0, 0x0d, 0xf0, 0x3a, 0x47, 0x89, 0x1e,
	0x4b, 0xe3, 0x99, 0xfd, 0xb1, 0xa4, 0xeb, 0x81, 0x31, 0xe8, 0xc9, 0x8b, 0x41, 0x4f, 0xa4, 0xf9,
	0xc6, 0x66, 0x9e, 0x36, 0x40, 0xa2, 0x38, 0xf8, 0x09, 0x5c, 0xab, 0xb1, 0x47, 0xed, 0x15, 0xcb,
	0xd6, 0x4f, 0x65, 0xeb, 0xbb, 0xe3, 0xad, 0x0f, 0xbc, 0xae, 0x6f, 0xcd, 0x91, 0x4f, 0xa2, 0x8e,
	0x7c, 0xdc, 0xdc, 0x38, 0x02, 0x27, 0x9e, 0xa4, 0xc3, 0x25, 0x86, 0x6b, 0x38, 0xdd, 0xec, 0x02,
	0x77, 0xbb, 0x0b, 0xdc, 0x3f, 0xbb, 0xc0, 0xfd, 0xb1, 0x0f, 0x9c, 0xed, 0x3e, 0x70, 0x7e, 0xed,
	0x03, 0xe7, 0xe3, 0x43, 0x51, 0x8a, 0x2f, 0xab, 0x2c, 0xca, 0x69, 0x1d, 0xcb, 0xcc, 0x97, 0x29,
	0xe7, 0x58, 0x70, 0xf5, 0x4f, 0xfc, 0xf5, 0x21, 0xfe, 0x1e, 0x1b, 0x1f, 0x80, 0x58, 0x2f, 0x31,
	0xcf, 0xfa, 0xf2, 0x03, 0x78, 0xf5, 0x37, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x1f, 0x15, 0x34, 0x79,
	0x04, 0x00, 0x00,
>>>>>>> a4ad980 (chore: rename module path (#283))
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
		dAtA[i] = 0x52
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
			dAtA[i] = 0x4a
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
		case 9:
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
		case 10:
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
