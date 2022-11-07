package types

import (
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTemporarilyOverrideFees{}

// Route Implements Msg.
func (m MsgTemporarilyOverrideFees) Route() string { return sdk.MsgTypeURL(&m) }

// Type Implements Msg.
func (m MsgTemporarilyOverrideFees) Type() string { return sdk.MsgTypeURL(&m) }

// GetSigners returns the expected signers for a MsgTemporarilyOverrideFees .
func (m MsgTemporarilyOverrideFees) GetSigners() []sdk.AccAddress {
	daoAccount, err := sdk.AccAddressFromBech32(m.FromAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{daoAccount}
}

// GetSignBytes Implements Msg.
func (m MsgTemporarilyOverrideFees) GetSignBytes() []byte {
	return sdk.MustSortJSON(legacy.Cdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m MsgTemporarilyOverrideFees) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.FromAddress)
	if err != nil {
		return sdkerrors.Wrap(err, "from address must be valid address")
	}

	if m.Coins.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "coins must be greater than 0")
	}

	return nil
}

func NewMsgTemporarilyOverrideFees(fromAddr sdk.AccAddress, endTimeSeconds uint64, coins sdk.Coins) *MsgTemporarilyOverrideFees {
	return &MsgTemporarilyOverrideFees{
		FromAddress: fromAddr.String(),
		EndTime:     endTimeSeconds,
		Coins:       coins,
	}
}
