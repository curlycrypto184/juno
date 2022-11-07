package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/CosmosContracts/juno/v11/x/globalfee/types"
)

// NewTxCmd returns a root CLI command handler for all x/exp transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "GlobalFee transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewOverrideCmd(),
		DEBUGAddAccountViaGovernance(),
	)
	return txCmd
}

func NewOverrideCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "override [coins]",
		Short: `Changes the global fees for a given denom for a temporary length of time (in hours)`,
		Long:  `...`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			coins, err := sdk.ParseCoinsNormalized(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgTemporarilyOverrideFees(from, 1, coins)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func DEBUGAddAccountViaGovernance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gov [acc] [number_of_times]",
		Short: `DEBUGGING, acting like gov`,
		Long:  `...`,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			times, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			// content := types.NewAddOverrideAccountProposal(title, description, address, amount)
			// v := govtypes.NewMsgSubmitProposal(&content, deposit, from)

			msg := types.NewMsgAddOverride(addr, times)
			if msg.ValidateBasic() != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
