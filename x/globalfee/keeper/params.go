package keeper

import (
	"github.com/CosmosContracts/juno/v11/x/globalfee/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams gets the global fee module's parameters.
func (k GlobalFeeKeeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the global fee module's parameters.
func (k GlobalFeeKeeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}
