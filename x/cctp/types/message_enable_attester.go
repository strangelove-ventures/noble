package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEnableAttester = "add_public_key"

var _ sdk.Msg = &MsgEnableAttester{}

func NewMsgEnableAttester(from string, publicKey []byte) *MsgEnableAttester {
	return &MsgEnableAttester{
		From:     from,
		Attester: publicKey,
	}
}

func (msg *MsgEnableAttester) Route() string {
	return RouterKey
}

func (msg *MsgEnableAttester) Type() string {
	return TypeMsgEnableAttester
}

func (msg *MsgEnableAttester) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgEnableAttester) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnableAttester) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}
	if len(msg.Attester) == 0 {
		return sdkerrors.Wrapf(ErrMalformedField, "Public key cannot be empty or nil")
	}
	return nil
}
