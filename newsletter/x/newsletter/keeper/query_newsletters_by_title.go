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

func (k Keeper) NewslettersByTitle(goCtx context.Context, req *types.QueryNewslettersByTitleRequest) (*types.QueryNewslettersByTitleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	newsletterStore := prefix.NewStore(store, types.KeyPrefix(types.NewsletterKeyPrefix))

	var retNewsletter types.Newsletter

	_, err := query.Paginate(newsletterStore, req.Pagination, func(key []byte, value []byte) error {
		var newsletter types.Newsletter
		if err := k.cdc.Unmarshal(value, &newsletter); err != nil {
			return err
		}

		if newsletter.Title == req.Title {
			retNewsletter = newsletter
		}

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryNewslettersByTitleResponse{Newsletter: retNewsletter}, nil
}
