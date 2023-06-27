package types_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"

	"github.com/functionx/fx-core/v5/x/evm/types"
)

func TestParamsValidate(t *testing.T) {
	extraEips := []int64{2929, 1884, 1344}
	testCases := []struct {
		name     string
		params   types.Params
		expError bool
	}{
		{"default", types.DefaultParams(), false},
		{
			"valid",
			types.NewParams("ara", false, true, true, types.DefaultChainConfig(), extraEips),
			false,
		},
		{
			"empty",
			types.Params{},
			true,
		},
		{
			"invalid evm denom",
			types.Params{
				EvmDenom: "@!#!@$!@5^32",
			},
			true,
		},
		{
			"invalid eip",
			types.Params{
				EvmDenom:  "stake",
				ExtraEIPs: []int64{1},
			},
			true,
		},
	}

	for _, tc := range testCases {
		err := tc.params.Validate()

		if tc.expError {
			require.Error(t, err, tc.name)
		} else {
			require.NoError(t, err, tc.name)
		}
	}
}

func TestParamsEIPs(t *testing.T) {
	extraEips := []int64{2929, 1884, 1344}
	params := types.NewParams("ara", false, true, true, types.DefaultChainConfig(), extraEips)
	actual := params.EIPs()

	require.Equal(t, []int{2929, 1884, 1344}, actual)
}

func TestParamsValidatePriv(t *testing.T) {
	require.Error(t, types.ValidateEVMDenom(false))
	require.NoError(t, types.ValidateEVMDenom("inj"))
	require.Error(t, types.ValidateBool(""))
	require.NoError(t, types.ValidateBool(true))
	require.Error(t, types.ValidateEIPs(""))
	require.NoError(t, types.ValidateEIPs([]int64{1884}))
}

func TestValidateChainConfig(t *testing.T) {
	testCases := []struct {
		name     string
		i        interface{}
		expError bool
	}{
		{
			"invalid chain config type",
			"string",
			true,
		},
		{
			"valid chain config type",
			types.DefaultChainConfig(),
			false,
		},
	}
	for _, tc := range testCases {
		err := types.ValidateChainConfig(tc.i)

		if tc.expError {
			require.Error(t, err, tc.name)
		} else {
			require.NoError(t, err, tc.name)
		}
	}
}

func TestIsLondon(t *testing.T) {
	testCases := []struct {
		name   string
		height int64
		result bool
	}{
		{
			"Before london block",
			5,
			false,
		},
		{
			"After london block",
			12_965_001,
			true,
		},
		{
			"london block",
			12_965_000,
			true,
		},
	}

	for _, tc := range testCases {
		ethConfig := params.MainnetChainConfig
		require.Equal(t, types.IsLondon(ethConfig, tc.height), tc.result)
	}
}
