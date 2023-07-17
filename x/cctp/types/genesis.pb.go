// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cctp/v1/genesis.proto

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

// GenesisState defines the cctp module's genesis state.
type GenesisState struct {
	Params                            Params                             `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Authority                         *Authority                         `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
	AttesterList                      []Attester                         `protobuf:"bytes,3,rep,name=attester_list,json=attesterList,proto3" json:"attester_list"`
	MinterAllowanceList               []MinterAllowances                 `protobuf:"bytes,4,rep,name=minter_allowance_list,json=minterAllowanceList,proto3" json:"minter_allowance_list"`
	PerMessageBurnLimit               *PerMessageBurnLimit               `protobuf:"bytes,5,opt,name=per_message_burn_limit,json=perMessageBurnLimit,proto3" json:"per_message_burn_limit,omitempty"`
	BurningAndMintingPaused           *BurningAndMintingPaused           `protobuf:"bytes,6,opt,name=burning_and_minting_paused,json=burningAndMintingPaused,proto3" json:"burning_and_minting_paused,omitempty"`
	SendingAndReceivingMessagesPaused *SendingAndReceivingMessagesPaused `protobuf:"bytes,7,opt,name=sending_and_receiving_messages_paused,json=sendingAndReceivingMessagesPaused,proto3" json:"sending_and_receiving_messages_paused,omitempty"`
	MaxMessageBodySize                *MaxMessageBodySize                `protobuf:"bytes,8,opt,name=max_message_body_size,json=maxMessageBodySize,proto3" json:"max_message_body_size,omitempty"`
	Nonce                             *Nonce                             `protobuf:"bytes,9,opt,name=nonce,proto3" json:"nonce,omitempty"`
	SignatureThreshold                *SignatureThreshold                `protobuf:"bytes,10,opt,name=signature_threshold,json=signatureThreshold,proto3" json:"signature_threshold,omitempty"`
	TokenPairList                     []TokenPairs                       `protobuf:"bytes,11,rep,name=token_pair_list,json=tokenPairList,proto3" json:"token_pair_list"`
	UsedNoncesList                    []Nonce                            `protobuf:"bytes,12,rep,name=used_nonces_list,json=usedNoncesList,proto3" json:"used_nonces_list"`
	TokenMessengerList                []TokenMessenger                   `protobuf:"bytes,13,rep,name=token_messenger_list,json=tokenMessengerList,proto3" json:"token_messenger_list"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_6890c4b88479c059, []int{0}
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

func (m *GenesisState) GetAuthority() *Authority {
	if m != nil {
		return m.Authority
	}
	return nil
}

func (m *GenesisState) GetAttesterList() []Attester {
	if m != nil {
		return m.AttesterList
	}
	return nil
}

func (m *GenesisState) GetMinterAllowanceList() []MinterAllowances {
	if m != nil {
		return m.MinterAllowanceList
	}
	return nil
}

func (m *GenesisState) GetPerMessageBurnLimit() *PerMessageBurnLimit {
	if m != nil {
		return m.PerMessageBurnLimit
	}
	return nil
}

func (m *GenesisState) GetBurningAndMintingPaused() *BurningAndMintingPaused {
	if m != nil {
		return m.BurningAndMintingPaused
	}
	return nil
}

func (m *GenesisState) GetSendingAndReceivingMessagesPaused() *SendingAndReceivingMessagesPaused {
	if m != nil {
		return m.SendingAndReceivingMessagesPaused
	}
	return nil
}

func (m *GenesisState) GetMaxMessageBodySize() *MaxMessageBodySize {
	if m != nil {
		return m.MaxMessageBodySize
	}
	return nil
}

func (m *GenesisState) GetNonce() *Nonce {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *GenesisState) GetSignatureThreshold() *SignatureThreshold {
	if m != nil {
		return m.SignatureThreshold
	}
	return nil
}

func (m *GenesisState) GetTokenPairList() []TokenPairs {
	if m != nil {
		return m.TokenPairList
	}
	return nil
}

func (m *GenesisState) GetUsedNoncesList() []Nonce {
	if m != nil {
		return m.UsedNoncesList
	}
	return nil
}

func (m *GenesisState) GetTokenMessengerList() []TokenMessenger {
	if m != nil {
		return m.TokenMessengerList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "cctp.v1.GenesisState")
}

func init() { proto.RegisterFile("cctp/v1/genesis.proto", fileDescriptor_6890c4b88479c059) }

var fileDescriptor_6890c4b88479c059 = []byte{
	// 670 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x4d, 0x6f, 0x13, 0x3d,
	0x10, 0xc7, 0x93, 0xa7, 0x6f, 0x4f, 0xdd, 0x37, 0x70, 0xfa, 0xb2, 0x0d, 0x90, 0xb6, 0x50, 0xa4,
	0x0a, 0xa9, 0x09, 0x6d, 0xaf, 0x08, 0x29, 0xb9, 0x70, 0x49, 0x4a, 0xd9, 0xf4, 0x84, 0x84, 0x56,
	0x4e, 0x76, 0xb4, 0xb1, 0xd8, 0xb5, 0x57, 0xb6, 0x37, 0x34, 0x15, 0x1f, 0x82, 0x8f, 0x55, 0x71,
	0xea, 0x91, 0x13, 0x42, 0xed, 0x17, 0x41, 0xeb, 0xb5, 0x1d, 0x25, 0x0d, 0x82, 0xdb, 0xe6, 0x3f,
	0xbf, 0xf9, 0x8f, 0xed, 0x99, 0x09, 0xda, 0xea, 0xf7, 0x55, 0xda, 0x18, 0x9e, 0x34, 0x22, 0x60,
	0x20, 0xa9, 0xac, 0xa7, 0x82, 0x2b, 0x8e, 0x97, 0x72, 0xb9, 0x3e, 0x3c, 0xa9, 0xee, 0xd8, 0x38,
	0xc9, 0xd4, 0x80, 0x0b, 0xaa, 0x46, 0x05, 0x51, 0x3d, 0xb2, 0x81, 0x5e, 0x26, 0x18, 0x65, 0x51,
	0x40, 0x58, 0x18, 0x24, 0x94, 0xa9, 0xfc, 0x3b, 0x25, 0x99, 0x84, 0xd0, 0x90, 0x2f, 0x2c, 0x99,
	0x90, 0xab, 0x20, 0x01, 0x29, 0x49, 0x04, 0x41, 0x8f, 0x87, 0xa3, 0x40, 0xd2, 0x6b, 0x30, 0xd0,
	0x9e, 0x83, 0x28, 0x53, 0x20, 0x02, 0x12, 0xc7, 0xfc, 0x0b, 0x61, 0x7d, 0x30, 0x27, 0xaa, 0x56,
	0x2c, 0xc0, 0x38, 0xeb, 0xdb, 0xac, 0x4d, 0x2b, 0xa6, 0x44, 0x90, 0xc4, 0xa2, 0x87, 0x4e, 0x05,
	0x31, 0x2e, 0x98, 0x09, 0x16, 0xc4, 0x34, 0xa1, 0xca, 0x50, 0xdb, 0xee, 0x66, 0x4a, 0x81, 0x54,
	0x20, 0x8c, 0x7e, 0x66, 0x75, 0x09, 0x2c, 0xb4, 0x17, 0x13, 0xd0, 0x07, 0x3a, 0xcc, 0x7f, 0x19,
	0x3f, 0x39, 0x79, 0xc7, 0x03, 0x97, 0x44, 0x23, 0x46, 0x54, 0x26, 0x20, 0x50, 0x03, 0x01, 0x72,
	0xc0, 0x63, 0x8b, 0xec, 0x5a, 0x44, 0xf1, 0xcf, 0xc0, 0x82, 0x94, 0x50, 0x61, 0x0f, 0xfc, 0x6c,
	0x32, 0x94, 0x97, 0x00, 0x16, 0xb9, 0x13, 0x6d, 0x46, 0x3c, 0xe2, 0xfa, 0xb3, 0x91, 0x7f, 0x15,
	0xea, 0xf3, 0xef, 0x4b, 0x68, 0xf5, 0x5d, 0xd1, 0xb4, 0xae, 0x22, 0x0a, 0xf0, 0x31, 0x5a, 0x2c,
	0x9e, 0xc1, 0x2b, 0xef, 0x97, 0x8f, 0x56, 0x4e, 0x37, 0xea, 0xa6, 0x89, 0xf5, 0x0b, 0x2d, 0xb7,
	0xe6, 0x6f, 0x7e, 0xee, 0x95, 0x7c, 0x03, 0xe1, 0xd7, 0x68, 0xd9, 0xf5, 0xd4, 0xfb, 0x4f, 0x67,
	0x60, 0x97, 0xd1, 0xb4, 0x11, 0x7f, 0x0c, 0xe1, 0x37, 0x68, 0xcd, 0xbe, 0x55, 0x10, 0x53, 0xa9,
	0xbc, 0xb9, 0xfd, 0xb9, 0xa3, 0x95, 0xd3, 0xc7, 0xe3, 0x2c, 0x13, 0x35, 0x95, 0x56, 0x2d, 0xdd,
	0xa6, 0x52, 0xe1, 0x2e, 0xda, 0x9a, 0xee, 0x6d, 0xe1, 0x32, 0xaf, 0x5d, 0x76, 0x9d, 0x4b, 0x47,
	0x53, 0x4d, 0x37, 0x00, 0xc6, 0xad, 0x92, 0x4c, 0xea, 0xda, 0xf4, 0x03, 0xda, 0x9e, 0xdd, 0x64,
	0x6f, 0x41, 0xdf, 0xe8, 0xe9, 0xf8, 0x0d, 0x40, 0x74, 0x0a, 0xaa, 0x95, 0x09, 0xd6, 0xce, 0x19,
	0xbf, 0x92, 0x3e, 0x14, 0xf1, 0x27, 0x54, 0xfd, 0xf3, 0x48, 0x7b, 0x8b, 0xda, 0x76, 0xdf, 0xd9,
	0xb6, 0x0a, 0xb4, 0xc9, 0xc2, 0x4e, 0x01, 0x5e, 0x68, 0xce, 0xdf, 0xe9, 0xcd, 0x0e, 0xe0, 0xaf,
	0xe8, 0xe5, 0x3f, 0x0d, 0x96, 0xb7, 0xa4, 0x2b, 0xbd, 0x72, 0x95, 0xba, 0x45, 0x56, 0x93, 0x85,
	0xbe, 0xcd, 0x31, 0x67, 0x97, 0xa6, 0xe6, 0x81, 0xfc, 0x1b, 0x82, 0xcf, 0xd1, 0xd6, 0xcc, 0x2d,
	0xf4, 0xfe, 0xd7, 0xd5, 0x9e, 0x8c, 0x9b, 0x40, 0xae, 0xec, 0xcb, 0xf0, 0x70, 0xd4, 0xa5, 0xd7,
	0xe0, 0xe3, 0xe4, 0x81, 0x86, 0x0f, 0xd1, 0x82, 0xde, 0x47, 0x6f, 0x59, 0xe7, 0xaf, 0xbb, 0xfc,
	0xf3, 0x5c, 0xf5, 0x8b, 0x20, 0x6e, 0xa3, 0xca, 0x8c, 0xbd, 0xf0, 0xd0, 0x54, 0xcd, 0xae, 0x65,
	0x2e, 0x2d, 0xe2, 0x63, 0xf9, 0x40, 0xc3, 0x4d, 0xb4, 0x31, 0x5e, 0xa1, 0x62, 0x84, 0x56, 0xf4,
	0x08, 0x55, 0x9c, 0xd3, 0x65, 0x1e, 0xbf, 0xc8, 0x37, 0xcc, 0x0c, 0xcf, 0x9a, 0xb2, 0x8a, 0x1e,
	0x9b, 0xb7, 0xe8, 0x51, 0xfe, 0x1c, 0x81, 0x3e, 0x9e, 0x2c, 0x3c, 0x56, 0xb5, 0xc7, 0xd4, 0x0d,
	0x4c, 0xfa, 0x7a, 0x4e, 0x6b, 0x41, 0xea, 0xfc, 0xf7, 0x68, 0x73, 0x6a, 0x55, 0x0b, 0x8f, 0x35,
	0xed, 0xb1, 0x33, 0x79, 0x8e, 0x8e, 0x65, 0x8c, 0x19, 0x56, 0x13, 0x6a, 0x6e, 0xd8, 0x6a, 0xdf,
	0xdc, 0xd5, 0xca, 0xb7, 0x77, 0xb5, 0xf2, 0xaf, 0xbb, 0x5a, 0xf9, 0xdb, 0x7d, 0xad, 0x74, 0x7b,
	0x5f, 0x2b, 0xfd, 0xb8, 0xaf, 0x95, 0x3e, 0x9e, 0x46, 0x54, 0x0d, 0xb2, 0x5e, 0xbd, 0xcf, 0x93,
	0x86, 0x54, 0x82, 0xb0, 0x08, 0x62, 0x3e, 0x84, 0xe3, 0x21, 0xb0, 0xfc, 0x59, 0x64, 0x83, 0xf1,
	0x5e, 0x0c, 0x8d, 0xab, 0x86, 0xfe, 0x0f, 0x51, 0xa3, 0x14, 0x64, 0x6f, 0x51, 0xff, 0x43, 0x9c,
	0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x8b, 0x64, 0x8c, 0xdd, 0x05, 0x00, 0x00,
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
	if len(m.TokenMessengerList) > 0 {
		for iNdEx := len(m.TokenMessengerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenMessengerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x6a
		}
	}
	if len(m.UsedNoncesList) > 0 {
		for iNdEx := len(m.UsedNoncesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UsedNoncesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x62
		}
	}
	if len(m.TokenPairList) > 0 {
		for iNdEx := len(m.TokenPairList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenPairList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if m.SignatureThreshold != nil {
		{
			size, err := m.SignatureThreshold.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if m.Nonce != nil {
		{
			size, err := m.Nonce.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.MaxMessageBodySize != nil {
		{
			size, err := m.MaxMessageBodySize.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.SendingAndReceivingMessagesPaused != nil {
		{
			size, err := m.SendingAndReceivingMessagesPaused.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.BurningAndMintingPaused != nil {
		{
			size, err := m.BurningAndMintingPaused.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.PerMessageBurnLimit != nil {
		{
			size, err := m.PerMessageBurnLimit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.MinterAllowanceList) > 0 {
		for iNdEx := len(m.MinterAllowanceList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinterAllowanceList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.AttesterList) > 0 {
		for iNdEx := len(m.AttesterList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AttesterList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Authority != nil {
		{
			size, err := m.Authority.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
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
	if m.Authority != nil {
		l = m.Authority.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.AttesterList) > 0 {
		for _, e := range m.AttesterList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MinterAllowanceList) > 0 {
		for _, e := range m.MinterAllowanceList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.PerMessageBurnLimit != nil {
		l = m.PerMessageBurnLimit.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.BurningAndMintingPaused != nil {
		l = m.BurningAndMintingPaused.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.SendingAndReceivingMessagesPaused != nil {
		l = m.SendingAndReceivingMessagesPaused.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.MaxMessageBodySize != nil {
		l = m.MaxMessageBodySize.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Nonce != nil {
		l = m.Nonce.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.SignatureThreshold != nil {
		l = m.SignatureThreshold.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.TokenPairList) > 0 {
		for _, e := range m.TokenPairList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UsedNoncesList) > 0 {
		for _, e := range m.UsedNoncesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TokenMessengerList) > 0 {
		for _, e := range m.TokenMessengerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
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
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
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
			if m.Authority == nil {
				m.Authority = &Authority{}
			}
			if err := m.Authority.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttesterList", wireType)
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
			m.AttesterList = append(m.AttesterList, Attester{})
			if err := m.AttesterList[len(m.AttesterList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinterAllowanceList", wireType)
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
			m.MinterAllowanceList = append(m.MinterAllowanceList, MinterAllowances{})
			if err := m.MinterAllowanceList[len(m.MinterAllowanceList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PerMessageBurnLimit", wireType)
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
			if m.PerMessageBurnLimit == nil {
				m.PerMessageBurnLimit = &PerMessageBurnLimit{}
			}
			if err := m.PerMessageBurnLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurningAndMintingPaused", wireType)
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
			if m.BurningAndMintingPaused == nil {
				m.BurningAndMintingPaused = &BurningAndMintingPaused{}
			}
			if err := m.BurningAndMintingPaused.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendingAndReceivingMessagesPaused", wireType)
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
			if m.SendingAndReceivingMessagesPaused == nil {
				m.SendingAndReceivingMessagesPaused = &SendingAndReceivingMessagesPaused{}
			}
			if err := m.SendingAndReceivingMessagesPaused.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMessageBodySize", wireType)
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
			if m.MaxMessageBodySize == nil {
				m.MaxMessageBodySize = &MaxMessageBodySize{}
			}
			if err := m.MaxMessageBodySize.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
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
			if m.Nonce == nil {
				m.Nonce = &Nonce{}
			}
			if err := m.Nonce.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignatureThreshold", wireType)
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
			if m.SignatureThreshold == nil {
				m.SignatureThreshold = &SignatureThreshold{}
			}
			if err := m.SignatureThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenPairList", wireType)
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
			m.TokenPairList = append(m.TokenPairList, TokenPairs{})
			if err := m.TokenPairList[len(m.TokenPairList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsedNoncesList", wireType)
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
			m.UsedNoncesList = append(m.UsedNoncesList, Nonce{})
			if err := m.UsedNoncesList[len(m.UsedNoncesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenMessengerList", wireType)
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
			m.TokenMessengerList = append(m.TokenMessengerList, TokenMessenger{})
			if err := m.TokenMessengerList[len(m.TokenMessengerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
