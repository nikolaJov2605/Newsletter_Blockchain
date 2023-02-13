package keeper

import (
	"context"
	"strconv"

	"newsletter/x/newsletter/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateNewsletter(goCtx context.Context, msg *types.MsgCreateNewsletter) (*types.MsgCreateNewsletterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	nInfo, found := k.Keeper.GetNewsletterInfo(ctx)
	if !found {
		panic("Info not found")
	}

	newIndex := strconv.FormatUint(nInfo.NextId, 10)

	newNewsletter := types.Newsletter{
		Index:          newIndex,
		Title:          msg.Title,
		Description:    msg.Description,
		Price:          msg.Price,
		SubscriberList: make([]string, 0),
	}

	k.Keeper.SetNewsletter(ctx, newNewsletter)

	nInfo.NextId++
	k.Keeper.SetNewsletterInfo(ctx, nInfo)

	return &types.MsgCreateNewsletterResponse{NewsletterIndex: newIndex}, nil
}
