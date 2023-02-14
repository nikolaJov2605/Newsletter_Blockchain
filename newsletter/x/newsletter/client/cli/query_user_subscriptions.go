package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"newsletter/x/newsletter/types"
)

var _ = strconv.Itoa(0)

func CmdUserSubscriptions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user-subscriptions [user-address]",
		Short: "Query user-subscriptions",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqUserAddress := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryUserSubscriptionsRequest{

				UserAddress: reqUserAddress,
			}

			res, err := queryClient.UserSubscriptions(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
