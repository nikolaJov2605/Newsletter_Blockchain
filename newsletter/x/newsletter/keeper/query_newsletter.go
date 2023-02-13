package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"newsletter/x/newsletter/types"
)

func (k Keeper) NewsletterAll(goCtx context.Context, req *types.QueryAllNewsletterRequest) (*types.QueryAllNewsletterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var newsletters []types.Newsletter
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	newsletterStore := prefix.NewStore(store, types.KeyPrefix(types.NewsletterKeyPrefix))

	pageRes, err := query.Paginate(newsletterStore, req.Pagination, func(key []byte, value []byte) error {
		var newsletter types.Newsletter
		if err := k.cdc.Unmarshal(value, &newsletter); err != nil {
			return err
		}

		newsletters = append(newsletters, newsletter)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNewsletterResponse{Newsletter: newsletters, Pagination: pageRes}, nil
}

func (k Keeper) Newsletter(goCtx context.Context, req *types.QueryGetNewsletterRequest) (*types.QueryGetNewsletterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetNewsletter(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNewsletterResponse{Newsletter: val}, nil
}
