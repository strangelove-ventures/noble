package types

import (
	"testing"

	"github.com/strangelove-ventures/noble/testutil/sample"
	"github.com/stretchr/testify/require"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func TestMsgRemoveMinter_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRemoveMinter
		err  error
	}{
		{
			name: "invalid from",
			msg: MsgRemoveMinter{
				From:    "invalid_address",
				Address: sample.AccAddress(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid address",
			msg: MsgRemoveMinter{
				From:    sample.AccAddress(),
				Address: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid address and from",
			msg: MsgRemoveMinter{
				From:    sample.AccAddress(),
				Address: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
