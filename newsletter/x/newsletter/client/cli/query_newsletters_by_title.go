package cli

import (
	"strconv"

	"newsletter/x/newsletter/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdNewslettersByTitle() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "newsletters-by-title [title]",
		Short: "Query newsletters-by-title",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argTitle := args[0]

			params := &types.QueryNewslettersByTitleRequest{
				Title: argTitle,
			}

			res, err := queryClient.NewslettersByTitle(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
