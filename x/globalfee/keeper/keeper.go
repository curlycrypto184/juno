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
type Keeper struct {
	cdc codec.BinaryCodec

	storeKey   storetypes.StoreKey
	paramSpace paramtypes.Subspace
}

func NewKeeper(
	cdc codec.BinaryCodec,
	key storetypes.StoreKey,
	paramSpace paramtypes.Subspace,
) Keeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:        cdc,
		storeKey:   key,
		paramSpace: paramSpace,
	}
}

// Logger returns the application logger, scoped to the associated module
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// verify the account is allowed to override fees (as signaled by governance)
func (k Keeper) verifyAccountCanOverride(ctx sdk.Context, addr sdk.AccAddress) error {
	whiteList := k.GetWhiteList(ctx)

	for _, accountRecord := range whiteList {
		if addr.String() == accountRecord.Account {
			return nil
		}
	}
	return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "account is not in whitelist")
}

// FEES
func (k Keeper) GetOverrideFees(ctx sdk.Context) types.MsgTemporarilyOverrideFees {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ParamStoreOverrideFees)
	var overrideFeesInfo types.MsgTemporarilyOverrideFees
	k.cdc.MustUnmarshal(bz, &overrideFeesInfo)
	return overrideFeesInfo
}

func (k Keeper) SetOverrideFees(ctx sdk.Context, overrideFeesInfo types.MsgTemporarilyOverrideFees) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&overrideFeesInfo)
	store.Set(types.ParamStoreOverrideFees, bz)
}

func (k Keeper) RemoveOverrideFees(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.ParamStoreOverrideFees)
}

// ACCOUNTS (set by governance)
func (k Keeper) GetAccountAllowed(ctx sdk.Context) types.MsgTemporarilyOverrideFees {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ParamStoreOverrideFees)
	var overrideFeesInfo types.MsgTemporarilyOverrideFees
	k.cdc.MustUnmarshal(bz, &overrideFeesInfo)
	return overrideFeesInfo
}

func (k Keeper) SetAccountAllowed(ctx sdk.Context, overrider types.MsgAddOverride) {
	store := ctx.KVStore(k.storeKey)

	w := k.GetWhiteList(ctx)
	var totalOverrides uint64 = 0
	for _, accountRecord := range w {
		if overrider.Address == accountRecord.Account {
			totalOverrides = accountRecord.AllowedBypasses
		}
	}

	var whitelist types.WhitelistAccounts
	v := store.Get(types.ParamStoreKeyWhitelist)
	k.cdc.MustUnmarshal(v, &whitelist)

	// update whitelist with new amount for account
	for i, accountRecord := range whitelist.Records {
		if overrider.Address == accountRecord.Account {
			whitelist.Records[i].AllowedBypasses = overrider.Amount + totalOverrides
		}
	}

	store.Set(types.ParamStoreKeyWhitelist, k.cdc.MustMarshal(&whitelist))
}

func (k Keeper) RemoveAccount(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	// remove addr from the whitelist
	whitelist := k.GetWhiteList(ctx)
	for i, accountRecord := range whitelist {
		if addr.String() == accountRecord.Account {
			whitelist = append(whitelist[:i], whitelist[i+1:]...)
		}
	}

	var w types.WhitelistAccounts
	v := store.Get(types.ParamStoreKeyWhitelist)
	k.cdc.MustUnmarshal(v, &w)

	store.Set(types.ParamStoreKeyWhitelist, k.cdc.MustMarshal(&w))
}
