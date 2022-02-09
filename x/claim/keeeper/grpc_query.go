package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmosquad-labs/squad/x/claim/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	// ctx := sdk.UnwrapSDKContext(c)

	// return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
	return &types.QueryParamsResponse{}, nil
}
