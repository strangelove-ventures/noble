package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/tokenfactory module sentinel errors
var (
	ErrUnauthorized       = sdkerrors.Register(ModuleName, 2, "unauthorized")
	ErrUserNotFound       = sdkerrors.Register(ModuleName, 3, "user not found")
	ErrMint               = sdkerrors.Register(ModuleName, 4, "tokens can not be minted")
	ErrSendCoinsToAccount = sdkerrors.Register(ModuleName, 5, "can't send tokens to account")
	ErrBurn               = sdkerrors.Register(ModuleName, 6, "tokens can not be burned")
	ErrPaused             = sdkerrors.Register(ModuleName, 7, "the chain is paused")
	ErrInvalidPubkey      = sdkerrors.Register(ModuleName, 8, "pubkey bytes are invalid")
	ErrUserBlacklisted    = sdkerrors.Register(ModuleName, 9, "user is already blacklisted")
)
