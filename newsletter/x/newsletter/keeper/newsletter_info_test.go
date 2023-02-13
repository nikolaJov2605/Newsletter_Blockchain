package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "newsletter/testutil/keeper"
	"newsletter/testutil/nullify"
	"newsletter/x/newsletter/keeper"
	"newsletter/x/newsletter/types"
)

func createTestNewsletterInfo(keeper *keeper.Keeper, ctx sdk.Context) types.NewsletterInfo {
	item := types.NewsletterInfo{}
	keeper.SetNewsletterInfo(ctx, item)
	return item
}

func TestNewsletterInfoGet(t *testing.T) {
	keeper, ctx := keepertest.NewsletterKeeper(t)
	item := createTestNewsletterInfo(keeper, ctx)
	rst, found := keeper.GetNewsletterInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestNewsletterInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.NewsletterKeeper(t)
	createTestNewsletterInfo(keeper, ctx)
	keeper.RemoveNewsletterInfo(ctx)
	_, found := keeper.GetNewsletterInfo(ctx)
	require.False(t, found)
}
