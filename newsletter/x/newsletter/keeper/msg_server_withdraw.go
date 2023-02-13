package keeper

import (
	"context"

	"newsletter/x/newsletter/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	newsletter, found := k.Keeper.GetNewsletter(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Newsletter doesn't exist")
	}

	if newsletter.Creator != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "You don't have permission to withdraw funds for this newsletter")
	}

	toWithdraw := len(newsletter.SubscriberList) - int(newsletter.WithdrawCount)

	if toWithdraw == 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "No subscribers")
	}

	price, _ := sdk.ParseCoinsNormalized(newsletter.Price)
	toWithdrawInt := sdk.NewIntFromUint64(newsletter.SubscriberCount)

	coinsToWithdraw := price.MulInt(toWithdrawInt)

	owner, _ := sdk.AccAddressFromBech32(newsletter.Creator)

	_ = coinsToWithdraw
	_ = owner

	/*err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, coinsToWithdraw)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "There is no enough funds on wallet")
	}*/

	newsletter.WithdrawCount += toWithdrawInt.Uint64()

	k.SetNewsletter(ctx, newsletter)

	return &types.MsgWithdrawResponse{}, nil
}
