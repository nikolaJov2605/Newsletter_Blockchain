package keeper

import (
	"context"

	"newsletter/x/newsletter/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Subscribe(goCtx context.Context, msg *types.MsgSubscribe) (*types.MsgSubscribeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	newsletter, found := k.Keeper.GetNewsletter(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Newsletter doesn't exist")
	}

	if newsletter.Creator == msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Can't subscribe to your own newsletter")
	}

	subscriptionPrice, _ := sdk.ParseCoinsNormalized(newsletter.Price)

	buyer, _ := sdk.AccAddressFromBech32(msg.Creator)

	_ = subscriptionPrice
	_ = buyer

	/*err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, buyer, types.ModuleName, subscriptionPrice)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "There is no enough funds on wallet")
	}*/

	newsletter.SubscriberList = append(newsletter.SubscriberList, msg.Creator)
	newsletter.SubscriberCount = uint64(len(newsletter.SubscriberList))

	k.SetNewsletter(ctx, newsletter)

	return &types.MsgSubscribeResponse{}, nil
}
