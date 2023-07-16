package keeper

import (
	"context"

	"github.com/strangelove-ventures/noble/x/tokenfactory/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Paused(c context.Context, req *types.QueryGetPausedRequest) (*types.QueryGetPausedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val := k.GetPaused(ctx)

	return &types.QueryGetPausedResponse{Paused: val}, nil
}
