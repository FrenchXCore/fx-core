package types_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"

	"github.com/functionx/fx-core/v5/crypto/ethsecp256k1"
	"github.com/functionx/fx-core/v5/x/evm/types"
)

type GenesisTestSuite struct {
	suite.Suite

	address string
	hash    common.Hash
	code    string
}

func (suite *GenesisTestSuite) SetupTest() {
	priv, err := ethsecp256k1.GenerateKey()
	suite.Require().NoError(err)

	suite.address = common.BytesToAddress(priv.PubKey().Address().Bytes()).String()
	suite.hash = common.BytesToHash([]byte("hash"))
	suite.code = common.Bytes2Hex([]byte{1, 2, 3})
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}

func (suite *GenesisTestSuite) TestValidateGenesisAccount() {
	testCases := []struct {
		name           string
		genesisAccount types.GenesisAccount
		expPass        bool
	}{
		{
			"valid genesis account",
			types.GenesisAccount{
				Address: suite.address,
				Code:    suite.code,
				Storage: types.Storage{
					types.NewState(suite.hash, suite.hash),
				},
			},
			true,
		},
		{
			"empty account address bytes",
			types.GenesisAccount{
				Address: "",
				Code:    suite.code,
				Storage: types.Storage{
					types.NewState(suite.hash, suite.hash),
				},
			},
			false,
		},
		{
			"empty code bytes",
			types.GenesisAccount{
				Address: suite.address,
				Code:    "",
				Storage: types.Storage{
					types.NewState(suite.hash, suite.hash),
				},
			},
			true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		err := tc.genesisAccount.Validate()
		if tc.expPass {
			suite.Require().NoError(err, tc.name)
		} else {
			suite.Require().Error(err, tc.name)
		}
	}
}

func (suite *GenesisTestSuite) TestValidateGenesis() {
	testCases := []struct {
		name     string
		genState *types.GenesisState
		expPass  bool
	}{
		{
			name:     "default",
			genState: types.DefaultGenesisState(),
			expPass:  true,
		},
		{
			name: "valid genesis",
			genState: &types.GenesisState{
				Accounts: []types.GenesisAccount{
					{
						Address: suite.address,

						Code: suite.code,
						Storage: types.Storage{
							{Key: suite.hash.String()},
						},
					},
				},
				Params: types.DefaultParams(),
			},
			expPass: true,
		},
		{
			name:     "empty genesis",
			genState: &types.GenesisState{},
			expPass:  false,
		},
		{
			name:     "copied genesis",
			genState: types.NewGenesisState(types.DefaultGenesisState().Params, types.DefaultGenesisState().Accounts),
			expPass:  true,
		},
		{
			name: "invalid genesis",
			genState: &types.GenesisState{
				Accounts: []types.GenesisAccount{
					{
						Address: common.Address{}.String(),
					},
				},
			},
			expPass: false,
		},
		{
			name: "invalid genesis account",
			genState: &types.GenesisState{
				Accounts: []types.GenesisAccount{
					{
						Address: "123456",

						Code: suite.code,
						Storage: types.Storage{
							{Key: suite.hash.String()},
						},
					},
				},
				Params: types.DefaultParams(),
			},
			expPass: false,
		},
		{
			name: "duplicated genesis account",
			genState: &types.GenesisState{
				Accounts: []types.GenesisAccount{
					{
						Address: suite.address,

						Code: suite.code,
						Storage: types.Storage{
							types.NewState(suite.hash, suite.hash),
						},
					},
					{
						Address: suite.address,

						Code: suite.code,
						Storage: types.Storage{
							types.NewState(suite.hash, suite.hash),
						},
					},
				},
			},
			expPass: false,
		},
		{
			name: "duplicated tx log",
			genState: &types.GenesisState{
				Accounts: []types.GenesisAccount{
					{
						Address: suite.address,

						Code: suite.code,
						Storage: types.Storage{
							{Key: suite.hash.String()},
						},
					},
				},
			},
			expPass: false,
		},
		{
			name: "invalid tx log",
			genState: &types.GenesisState{
				Accounts: []types.GenesisAccount{
					{
						Address: suite.address,

						Code: suite.code,
						Storage: types.Storage{
							{Key: suite.hash.String()},
						},
					},
				},
			},
			expPass: false,
		},
		{
			name: "invalid params",
			genState: &types.GenesisState{
				Params: types.Params{},
			},
			expPass: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		err := tc.genState.Validate()
		if tc.expPass {
			suite.Require().NoError(err, tc.name)
		} else {
			suite.Require().Error(err, tc.name)
		}
	}
}
