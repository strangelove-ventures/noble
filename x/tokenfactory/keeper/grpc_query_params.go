package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
<<<<<<< HEAD
	"github.com/strangelove-ventures/noble/v4/x/tokenfactory/types"
=======
	"github.com/noble-assets/noble/v5/x/tokenfactory/types"
>>>>>>> a4ad980 (chore: rename module path (#283))
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}
