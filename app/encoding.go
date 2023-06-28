package app

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	cryptocodec "github.com/functionx/fx-core/v5/crypto/codec"
	fxtypes "github.com/functionx/fx-core/v5/types"
	crosschaintypes "github.com/functionx/fx-core/v5/x/crosschain/types"
	gravitytypes "github.com/functionx/fx-core/v5/x/gravity/types"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

// MakeEncodingConfig creates an EncodingConfig for testing
func MakeEncodingConfig() EncodingConfig {
	encodingConfig := makeEncodingConfig()
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	cryptocodec.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	cryptocodec.RegisterLegacyAminoCodec(encodingConfig.Amino)
	fxtypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	crosschaintypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	crosschaintypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	gravitytypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	gravitytypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	govv1beta1.ModuleCdc = codec.NewAminoCodec(encodingConfig.Amino)
	govv1.ModuleCdc = codec.NewAminoCodec(encodingConfig.Amino)

	// NOTE: update SDK's amino codec to include the ethsecp256k1 keys.
	// DO NOT REMOVE unless deprecated on the SDK.
	legacy.Cdc = encodingConfig.Amino
	keys.KeysCdc = encodingConfig.Amino
	return encodingConfig
}

// MakeEncodingConfig creates an EncodingConfig for an amino based test configuration.
func makeEncodingConfig() EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(cdc, tx.DefaultSignModes)

	encodingConfig := EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             cdc,
		TxConfig:          txCfg,
		Amino:             amino,
	}
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	keyring.RegisterLegacyAminoCodec(amino)
	return encodingConfig
}
