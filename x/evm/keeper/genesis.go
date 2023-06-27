package keeper

import (
	"bytes"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	abci "github.com/tendermint/tendermint/abci/types"

	fxtypes "github.com/functionx/fx-core/v5/types"
	"github.com/functionx/fx-core/v5/x/evm/types"
)

// InitGenesis initializes genesis state based on exported genesis
func (k *Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) []abci.ValidatorUpdate {
	if err := k.SetParams(ctx, data.Params); err != nil {
		panic(fmt.Errorf("error setting params %s", err))
	}
	// ensure evm module account is set
	if acc := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName); acc == nil {
		panic("the EVM module account has not been set")
	}

	for _, account := range data.Accounts {
		address := common.HexToAddress(account.Address)
		accAddress := sdk.AccAddress(address.Bytes())
		// check that the EVM balance the matches the account balance
		acc := k.accountKeeper.GetAccount(ctx, accAddress)
		if acc == nil {
			panic(fmt.Errorf("account not found for address %s", account.Address))
		}

		ethAcct, ok := acc.(fxtypes.EthAccountI)
		if !ok {
			panic(fmt.Errorf("account %s must be an EthAccount interface, got %T", account.Address, acc))
		}

		code := common.Hex2Bytes(account.Code)
		codeHash := crypto.Keccak256Hash(code)

		// we ignore the empty Code hash checking, see ethermint PR#1234
		if len(account.Code) != 0 && !bytes.Equal(ethAcct.GetCodeHash().Bytes(), codeHash.Bytes()) {
			s := "the evm state code doesn't match with the codehash\n"
			panic(fmt.Sprintf("%s account: %s , evm state codehash: %v, ethAccount codehash: %v, evm state code: %s\n",
				s, account.Address, codeHash, ethAcct.GetCodeHash(), account.Code))
		}

		k.SetCode(ctx, codeHash.Bytes(), code)

		for _, storage := range account.Storage {
			k.SetState(ctx, address, common.HexToHash(storage.Key), common.HexToHash(storage.Value).Bytes())
		}
	}

	// init logic contract
	initContract := []fxtypes.Contract{fxtypes.GetFIP20(), fxtypes.GetWFX()}
	for _, contract := range initContract {
		if len(contract.Code) <= 0 || contract.Address == common.HexToAddress(fxtypes.EmptyEvmAddress) {
			panic(fmt.Sprintf("invalid contract: %s", contract.Address.String()))
		}
		if err := k.CreateContractWithCode(ctx, contract.Address, contract.Code); err != nil {
			panic(fmt.Sprintf("create contract %s with code error %s", contract.Address.String(), err.Error()))
		}
	}

	return []abci.ValidatorUpdate{}
}

// ExportGenesis exports genesis state of the EVM module
func (k *Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	var ethGenAccounts []types.GenesisAccount
	k.accountKeeper.IterateAccounts(ctx, func(account authtypes.AccountI) bool {
		ethAccount, ok := account.(fxtypes.EthAccountI)
		if !ok {
			// ignore non EthAccounts
			return false
		}

		addr := ethAccount.EthAddress()

		storage := k.GetAccountStorage(ctx, addr)

		genAccount := types.GenesisAccount{
			Address: addr.String(),
			Code:    common.Bytes2Hex(k.GetCode(ctx, ethAccount.GetCodeHash())),
			Storage: storage,
		}

		ethGenAccounts = append(ethGenAccounts, genAccount)
		return false
	})

	return &types.GenesisState{
		Accounts: ethGenAccounts,
		Params:   k.GetParams(ctx),
	}
}
