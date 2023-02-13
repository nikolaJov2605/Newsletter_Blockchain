package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"newsletter/x/newsletter/types"
)

func (k Keeper) NewsletterInfo(goCtx context.Context, req *types.QueryGetNewsletterInfoRequest) (*types.QueryGetNewsletterInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetNewsletterInfo(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNewsletterInfoResponse{NewsletterInfo: val}, nil
}
