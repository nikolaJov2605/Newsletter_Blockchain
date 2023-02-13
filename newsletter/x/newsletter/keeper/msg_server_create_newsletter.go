package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"newsletter/x/newsletter/types"
)

func (k msgServer) CreateNewsletter(goCtx context.Context, msg *types.MsgCreateNewsletter) (*types.MsgCreateNewsletterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateNewsletterResponse{}, nil
}
