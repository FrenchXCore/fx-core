package keeper

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/functionx/fx-core/contracts"
	"github.com/functionx/fx-core/x/erc20/types"
)

func (k Keeper) UpgradeSystemContract(ctx sdk.Context) error {
	ctx.Logger().Info("upgrade system contract", "height", ctx.BlockHeight())
	for _, contract := range contracts.GetUpgradeContracts(ctx.BlockHeight()) {
		if len(contract.Code) <= 0 || contract.Address == common.HexToAddress(contracts.EmptyEvmAddress) {
			return errors.New("invalid contract")
		}
		err := k.evmKeeper.CreateContractWithCode(ctx, contract.Address, contract.Code)
		if err != nil {
			return err
		}
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeUpgradeSystemContract,
			sdk.NewAttribute(types.AttributeKeyContractAddress, contract.Address.String()),
		))
	}
	return nil
}
