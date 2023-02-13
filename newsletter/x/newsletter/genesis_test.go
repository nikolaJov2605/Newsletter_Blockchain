package newsletter_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "newsletter/testutil/keeper"
	"newsletter/testutil/nullify"
	"newsletter/x/newsletter"
	"newsletter/x/newsletter/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		NewsletterInfo: &types.NewsletterInfo{
			NextId: 65,
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
	// this line is used by starport scaffolding # genesis/test/assert
}
