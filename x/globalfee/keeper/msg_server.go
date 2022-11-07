package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	// govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/CosmosContracts/juno/v11/x/globalfee/types"
)

type msgServer struct {
	GlobalFeeKeeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the exp MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper GlobalFeeKeeper) types.MsgServer {
	return &msgServer{GlobalFeeKeeper: keeper}
}

func (k msgServer) TemporarilyOverrideFees(goCtx context.Context, msg *types.MsgTemporarilyOverrideFees) (*types.MsgTemporarilyOverrideFeesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	fromAddress, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil, err
	}

	// remove this
	if err := k.verifyAccountCanOverride(ctx, fromAddress); err != nil {
		return nil, err
	}

	params := k.GetParams(ctx)
	hours := params.GetOverridePeriodHours()

	current_time := ctx.BlockTime()
	end_time := current_time.Add(time.Duration(hours) * time.Hour)

	msg.EndTime = uint64(end_time.Second())
	k.SetOverrideFees(ctx, *msg)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeKeyOverrideFees),
		),
	)

	return &types.MsgTemporarilyOverrideFeesResponse{}, nil
}
