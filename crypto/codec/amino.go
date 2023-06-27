package codec

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/functionx/fx-core/v5/crypto/ethsecp256k1"
)

// RegisterLegacyAminoCodec registers all crypto dependency types with the provided Amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&ethsecp256k1.PubKey{}, ethsecp256k1.PubKeyName, nil)
	cdc.RegisterConcrete(&ethsecp256k1.PrivKey{}, ethsecp256k1.PrivKeyName, nil)
}
