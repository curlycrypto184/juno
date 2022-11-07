package globalfee

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/CosmosContracts/juno/v11/x/globalfee/keeper"
	"github.com/CosmosContracts/juno/v11/x/globalfee/types"
)

// EndBlocker is called at the end of every block
// func EndBlocker(ctx sdk.Context, k keeper.GlobalFeeKeeper) error {
func EndBlocker(ctx sdk.Context, k keeper.GlobalFeeKeeper) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	current_time := ctx.BlockTime()
	endTimeSeconds := uint64(k.GetOverrideFees(ctx).EndTime)

	if endTimeSeconds == 0 {
		return nil
	}

	if endTimeSeconds-uint64(current_time.Second()) <= 0 {
		// remove the override
		k.RemoveOverrideFees(ctx)
	}

	return nil
}
