package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "newsletter/testutil/keeper"
	"newsletter/x/newsletter/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NewsletterKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
