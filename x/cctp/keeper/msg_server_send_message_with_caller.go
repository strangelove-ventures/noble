package keeper

import (
	"bytes"
	"context"
	"github.com/strangelove-ventures/noble/x/cctp"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/noble/x/cctp/types"
)

func (k msgServer) SendMessageWithCaller(goCtx context.Context, msg *types.MsgSendMessageWithCaller) (*types.MsgSendMessageWithCallerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	paused, found := k.GetSendingAndReceivingMessagesPaused(ctx)
	if found && paused.Paused {
		return nil, sdkerrors.Wrap(types.ErrReceiveMessage, "sending and receiving messages is paused")
	}

	// check if message body is too long, ignore if max length not found
	max, found := k.GetMaxMessageBodySize(ctx)
	if found && uint64(len(msg.MessageBody)) > max.Amount {
		return nil, sdkerrors.Wrap(types.ErrSendMessageWithCaller, "message body exceeds max size")
	}

	if msg.Recipient == nil || len(msg.Recipient) == 0 {
		return nil, sdkerrors.Wrap(types.ErrSendMessageWithCaller, "recipient must not be nil or empty")
	}

	emptyByteArr := make([]byte, cctp.Bytes32Len)
	if bytes.Equal(msg.DestinationCaller, emptyByteArr) {
		return nil, sdkerrors.Wrap(types.ErrSendMessageWithCaller, "destination caller must be nonzero")
	}

	event := types.MessageSent{
		Message: msg.MessageBody,
	}
	err := ctx.EventManager().EmitTypedEvent(&event)

	return &types.MsgSendMessageWithCallerResponse{}, err
}