package newsletter_test

import (
	"testing"

	keepertest "newsletter/testutil/keeper"
	"newsletter/testutil/nullify"
	"newsletter/x/newsletter"
	"newsletter/x/newsletter/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		NewsletterInfo: types.NewsletterInfo{
			NextId: 65,
		},
		NewsletterList: []types.Newsletter{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NewsletterKeeper(t)
	newsletter.InitGenesis(ctx, *k, genesisState)
	got := newsletter.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.NewsletterInfo, got.NewsletterInfo)
	require.ElementsMatch(t, genesisState.NewsletterList, got.NewsletterList)
	// this line is used by starport scaffolding # genesis/test/assert
}
