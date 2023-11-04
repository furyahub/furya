package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/furyahub/furya/x/feeburner/types"
	"github.com/spf13/cobra"
)

func CmdQueryTotalBurnedFuryasAmount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "total-burned-furyas-amount",
		Short: "shows total amount of burned furyas",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.TotalBurnedFuryasAmount(context.Background(), &types.QueryTotalBurnedFuryasAmountRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
