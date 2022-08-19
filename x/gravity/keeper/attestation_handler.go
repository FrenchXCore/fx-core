package keeper

import (
	"fmt"

	fxtypes "github.com/functionx/fx-core/v2/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/functionx/fx-core/v2/x/gravity/types"
)

// AttestationHandler processes `observed` Attestations
type AttestationHandler struct {
	keeper     *Keeper
	bankKeeper types.BankKeeper
}

// Handle is the entry point for Attestation processing.
func (a AttestationHandler) Handle(ctx sdk.Context, att types.Attestation, ethereumClaim types.EthereumClaim) error {
	switch claim := ethereumClaim.(type) {
	case *types.MsgDepositClaim:
		// Check if coin is fx-originated asset and get denom
		isCosmosOriginated, denom := a.keeper.ERC20ToDenomLookup(ctx, claim.TokenContract)
		coin := sdk.NewCoin(denom, claim.Amount)
		coins := sdk.Coins{coin}
		receiveAddr, err := sdk.AccAddressFromBech32(claim.FxReceiver)
		if err != nil {
			return sdkerrors.Wrap(err, "invalid receiver address")
		}
		if !isCosmosOriginated {
			// If it is not cosmos originated, mint the coins (aka vouchers)
			if err := a.bankKeeper.MintCoins(ctx, types.ModuleName, coins); err != nil {
				return sdkerrors.Wrapf(err, "mint vouchers coins: %s", coins)
			}
		}
		if err = a.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiveAddr, coins); err != nil {
			return sdkerrors.Wrap(err, "transfer vouchers")
		}
		a.handlerRelayTransfer(ctx, claim, receiveAddr, coin)
		return nil
	case *types.MsgWithdrawClaim:
		err := a.keeper.OutgoingTxBatchExecuted(ctx, claim.TokenContract, claim.BatchNonce)
		if err != nil {
			return err
		}
	case *types.MsgFxOriginatedTokenClaim:
		// Check if it already exists
		existingERC20, exists := a.keeper.GetFxOriginatedERC20(ctx, claim.Symbol)
		if exists {
			return sdkerrors.Wrap(
				types.ErrInvalid,
				fmt.Sprintf("ERC20 %s already exists for denom %s", existingERC20, claim.Symbol))
		}

		// Check if denom exists
		baseDenom := claim.Symbol
		metadata, found := a.keeper.bankKeeper.GetDenomMetaData(ctx, baseDenom)
		if !found {
			return sdkerrors.Wrap(types.ErrUnknown, fmt.Sprintf("denom not found %s", claim.Symbol))
		}

		// Check if attributes of ERC20 match fx denom
		if claim.Name != metadata.Name {
			return sdkerrors.Wrap(
				types.ErrInvalid,
				fmt.Sprintf("ERC20 name %s does not match denom display %s", claim.Name, metadata.Description))
		}

		if claim.Symbol != metadata.Symbol {
			return sdkerrors.Wrap(
				types.ErrInvalid,
				fmt.Sprintf("ERC20 symbol %s does not match denom display %s", claim.Symbol, metadata.Display))
		}

		if fxtypes.DenomUnit != uint32(claim.Decimals) {
			return sdkerrors.Wrap(
				types.ErrInvalid,
				fmt.Sprintf("ERC20 decimals %d does not match denom decimals %d", claim.Decimals, fxtypes.DenomUnit))
		}

		// Add to denom-erc20 mapping
		a.keeper.SetFxOriginatedDenomToERC20(ctx, claim.Symbol, claim.TokenContract)
		a.keeper.Logger(ctx).Debug("set fx originated denom to erc20 success", "denom", claim.Symbol, "token", claim.TokenContract)
	case *types.MsgValsetUpdatedClaim:
		a.keeper.SetLastObservedValset(ctx, types.Valset{
			Nonce:   claim.ValsetNonce,
			Members: claim.Members,
		})
	default:
		return sdkerrors.Wrapf(types.ErrInvalid, "event type: %s", claim.GetType())
	}
	return nil
}
