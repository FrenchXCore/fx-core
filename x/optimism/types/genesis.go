package types

import (
	crosschaintypes "github.com/functionx/fx-core/v5/x/crosschain/types"
)

func DefaultGenesisState() *crosschaintypes.GenesisState {
	params := crosschaintypes.DefaultParams()
	params.GravityId = "fx-optimism-bridge"
	params.AverageExternalBlockTime = 500
	return &crosschaintypes.GenesisState{
		Params: params,
	}
}
