package types_test

import (
	"math/big"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	"github.com/functionx/fx-core/v5/x/evm/types"
)

func TestTxData_chainID(t *testing.T) {
	chainID := sdkmath.NewInt(1)

	testCases := []struct {
		msg        string
		data       types.TxData
		expChainID *big.Int
	}{
		{
			"access list tx", &types.AccessListTx{Accesses: types.AccessList{}, ChainID: &chainID}, big.NewInt(1),
		},
		{
			"access list tx, nil chain ID", &types.AccessListTx{Accesses: types.AccessList{}}, nil,
		},
		{
			"legacy tx, derived", &types.LegacyTx{}, nil,
		},
	}

	for _, tc := range testCases {
		chainID := tc.data.GetChainID()
		require.Equal(t, chainID, tc.expChainID, tc.msg)
	}
}

func TestTxData_DeriveChainID(t *testing.T) {
	bitLen64, ok := new(big.Int).SetString("0x8000000000000000", 0)
	require.True(t, ok)

	bitLen80, ok := new(big.Int).SetString("0x80000000000000000000", 0)
	require.True(t, ok)

	expBitLen80, ok := new(big.Int).SetString("302231454903657293676526", 0)
	require.True(t, ok)

	testCases := []struct {
		msg        string
		data       types.TxData
		expChainID *big.Int
	}{
		{
			"v = -1", &types.LegacyTx{V: big.NewInt(-1).Bytes()}, nil,
		},
		{
			"v = 0", &types.LegacyTx{V: big.NewInt(0).Bytes()}, nil,
		},
		{
			"v = 1", &types.LegacyTx{V: big.NewInt(1).Bytes()}, nil,
		},
		{
			"v = 27", &types.LegacyTx{V: big.NewInt(27).Bytes()}, new(big.Int),
		},
		{
			"v = 28", &types.LegacyTx{V: big.NewInt(28).Bytes()}, new(big.Int),
		},
		{
			"Ethereum mainnet", &types.LegacyTx{V: big.NewInt(37).Bytes()}, big.NewInt(1),
		},
		{
			"chain ID 9000", &types.LegacyTx{V: big.NewInt(18035).Bytes()}, big.NewInt(9000),
		},
		{
			"bit len 64", &types.LegacyTx{V: bitLen64.Bytes()}, big.NewInt(4611686018427387886),
		},
		{
			"bit len 80", &types.LegacyTx{V: bitLen80.Bytes()}, expBitLen80,
		},
		{
			"v = nil ", &types.LegacyTx{V: nil}, nil,
		},
	}

	for _, tc := range testCases {
		v, _, _ := tc.data.GetRawSignatureValues()

		chainID := types.DeriveChainID(v)
		require.Equal(t, tc.expChainID, chainID, tc.msg)
	}
}
