package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "newsletter/testutil/keeper"
	"newsletter/testutil/nullify"
	"newsletter/x/newsletter/types"
)

func TestNewsletterInfoQuery(t *testing.T) {
	keeper, ctx := keepertest.NewsletterKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestNewsletterInfo(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetNewsletterInfoRequest
		response *types.QueryGetNewsletterInfoResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetNewsletterInfoRequest{},
			response: &types.QueryGetNewsletterInfoResponse{NewsletterInfo: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.NewsletterInfo(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
