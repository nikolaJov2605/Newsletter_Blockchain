package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "newsletter/testutil/keeper"
	"newsletter/testutil/nullify"
	"newsletter/x/newsletter/keeper"
	"newsletter/x/newsletter/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNNewsletter(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Newsletter {
	items := make([]types.Newsletter, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetNewsletter(ctx, items[i])
	}
	return items
}

func TestNewsletterGet(t *testing.T) {
	keeper, ctx := keepertest.NewsletterKeeper(t)
	items := createNNewsletter(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNewsletter(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNewsletterRemove(t *testing.T) {
	keeper, ctx := keepertest.NewsletterKeeper(t)
	items := createNNewsletter(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNewsletter(ctx,
			item.Index,
		)
		_, found := keeper.GetNewsletter(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestNewsletterGetAll(t *testing.T) {
	keeper, ctx := keepertest.NewsletterKeeper(t)
	items := createNNewsletter(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNewsletter(ctx)),
	)
}
