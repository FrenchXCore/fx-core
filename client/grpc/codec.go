package grpc

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	ethcryptocodec "github.com/functionx/fx-core/v5/crypto/codec"
	fxtypes "github.com/functionx/fx-core/v5/types"
)

func newInterfaceRegistry() types.InterfaceRegistry {
	interfaceRegistry := types.NewInterfaceRegistry()
	fxtypes.RegisterInterfaces(interfaceRegistry)
	authtypes.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
	ethcryptocodec.RegisterInterfaces(interfaceRegistry)
	return interfaceRegistry
}
