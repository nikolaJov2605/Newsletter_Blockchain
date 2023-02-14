package keeper

import (
	"context"

	"newsletter/x/newsletter/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NewslettersBasic(goCtx context.Context, req *types.QueryNewslettersBasicRequest) (*types.QueryNewslettersBasicResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var newsletters []types.NewslettersBasic

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	newsletterStore := prefix.NewStore(store, types.KeyPrefix(types.NewsletterKeyPrefix))

	pageRes, err := query.Paginate(newsletterStore, req.Pagination, func(key []byte, value []byte) error {
		var newsletter types.Newsletter
		var newsletterBasic types.NewslettersBasic
		if err := k.cdc.Unmarshal(value, &newsletter); err != nil {
			return err
		}
		newsletterBasic.Title = newsletter.Title
		newsletterBasic.SubscriberCount = newsletter.SubscriberCount
		newsletters = append(newsletters, newsletterBasic)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryNewslettersBasicResponse{Newsletters: newsletters, Pagination: pageRes}, nil
}
