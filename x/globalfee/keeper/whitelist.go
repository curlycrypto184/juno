package keeper

import (
	"github.com/CosmosContracts/juno/v11/x/globalfee/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) IterateWhitelist(ctx sdk.Context, cb func(record types.AccountRecord) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ParamStoreKeyWhitelist)

	for ; iterator.Valid(); iterator.Next() {
		var record types.AccountRecord
		err := k.cdc.Unmarshal(iterator.Value(), &record)
		if err != nil {
			panic(err)
		}
		if cb(record) {
			break
		}
	}
}

func (k Keeper) GetWhiteList(ctx sdk.Context) (records types.AccountRecords) {
	k.IterateWhitelist(ctx, func(record types.AccountRecord) bool {
		records = append(records, record)
		return false
	})
	return
}
