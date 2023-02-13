package newsletter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"newsletter/x/newsletter/keeper"
	"newsletter/x/newsletter/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.NewsletterInfo != nil {
		k.SetNewsletterInfo(ctx, *genState.NewsletterInfo)
	}
	// Set all the newsletter
	for _, elem := range genState.NewsletterList {
		k.SetNewsletter(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all newsletterInfo
	newsletterInfo, found := k.GetNewsletterInfo(ctx)
	if found {
		genesis.NewsletterInfo = &newsletterInfo
	}
	genesis.NewsletterList = k.GetAllNewsletter(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
