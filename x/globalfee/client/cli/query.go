package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/CosmosContracts/juno/v11/x/globalfee/types"
)

func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the global fee module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	queryCmd.AddCommand(
		GetCmdShowMinimumGasPrices(),
		GetCmdShowWhitelistOverrideAccounts(),
	)
	return queryCmd
}

func GetCmdShowMinimumGasPrices() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "minimum-gas-prices",
		Short:   "Show minimum gas prices",
		Long:    "Show all minimum gas prices",
		Aliases: []string{"min"},
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.MinimumGasPrices(cmd.Context(), &types.QueryMinimumGasPricesRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdShowWhitelistOverrideAccounts() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "whitelist",
		Short:   "Show accounts which can override fees temporarily",
		Long:    "Show accounts which can override fees temporarily",
		Aliases: []string{"w"},
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.WhiteList(cmd.Context(), &types.QueryWhiteListRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
