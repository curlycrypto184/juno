package keeper

import (
	"fmt"

	"github.com/CosmosContracts/juno/v11/x/globalfee/types"
	codec "github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper for GlobalFeeModule.
type GlobalFeeKeeper struct {
	cdc codec.BinaryCodec

	storeKey   storetypes.StoreKey
	paramSpace paramtypes.Subspace
}

func NewKeeper(
	key storetypes.StoreKey,
	cdc codec.BinaryCodec,
	paramSpace paramtypes.Subspace,
) GlobalFeeKeeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return GlobalFeeKeeper{
		cdc:        cdc,
		storeKey:   key,
		paramSpace: paramSpace,
	}
}

// Logger returns the application logger, scoped to the associated module
func (k GlobalFeeKeeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// verify the account is allowed to override fees (as signaled by governance)
func (k GlobalFeeKeeper) verifyAccountCanOverride(ctx sdk.Context, addr sdk.AccAddress) error {
	whiteList := k.GetWhiteList(ctx)

	for _, accountRecord := range whiteList {
		if addr.String() == accountRecord.Account {
			return nil
		}
	}
	return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "account is not in whitelist")
}

func (k GlobalFeeKeeper) GetOverrideFees(ctx sdk.Context) types.MsgTemporarilyOverrideFees {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.OverrideFees)
	var overrideFeesInfo types.MsgTemporarilyOverrideFees
	k.cdc.MustUnmarshal(bz, &overrideFeesInfo)
	return overrideFeesInfo
}

func (k GlobalFeeKeeper) SetOverrideFees(ctx sdk.Context, overrideFeesInfo types.MsgTemporarilyOverrideFees) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&overrideFeesInfo)
	store.Set(types.OverrideFees, bz)
}

func (k GlobalFeeKeeper) RemoveOverrideFees(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.OverrideFees)
}
