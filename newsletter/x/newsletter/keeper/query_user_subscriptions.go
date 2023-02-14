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

func (k Keeper) UserSubscriptions(goCtx context.Context, req *types.QueryUserSubscriptionsRequest) (*types.QueryUserSubscriptionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	newsletterStore := prefix.NewStore(store, types.KeyPrefix(types.NewsletterKeyPrefix))

	var retUser types.User

	_, err := query.Paginate(newsletterStore, req.Pagination, func(key []byte, value []byte) error {
		var newsletter types.Newsletter
		if err := k.cdc.Unmarshal(value, &newsletter); err != nil {
			return err
		}

		var i uint64

		for i = 0; i < newsletter.SubscriberCount; i++ {
			if newsletter.SubscriberList[i] == req.UserAddress {
				retUser.SubscriptionTitles = append(retUser.SubscriptionTitles, newsletter.Title)
			}
		}

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryUserSubscriptionsResponse{UserInfo: &retUser}, nil
}
