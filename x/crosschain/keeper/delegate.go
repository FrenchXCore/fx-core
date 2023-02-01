package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/functionx/fx-core/v3/x/crosschain/types"
)

func (k Keeper) GetOracleDelegateToken(ctx sdk.Context, delegateAddr sdk.AccAddress, valAddr sdk.ValAddress) (sdk.Int, error) {
	delegation, found := k.stakingKeeper.GetDelegation(ctx, delegateAddr, valAddr)
	if !found {
		return sdk.ZeroInt(), sdkerrors.Wrap(types.ErrInvalid, "no delegation for (address, validator) tuple")
	}
	validator, found := k.stakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return sdk.ZeroInt(), stakingtypes.ErrNoValidatorFound
	}

	delegateToken := validator.TokensFromSharesTruncated(delegation.GetShares()).TruncateInt()
	sharesTruncated, err := validator.SharesFromTokensTruncated(delegateToken)
	if err != nil {
		return sdk.ZeroInt(), sdkerrors.Wrapf(types.ErrInvalid, "shares from tokens:%v", delegateToken)
	}
	delShares := delegation.GetShares()
	if sharesTruncated.GT(delShares) {
		delegateToken = validator.TokensFromSharesTruncated(sharesTruncated).TruncateInt()
	}
	return delegateToken, nil
}