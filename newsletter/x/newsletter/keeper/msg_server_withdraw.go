package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"newsletter/x/newsletter/types"
)

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgWithdrawResponse{}, nil
}