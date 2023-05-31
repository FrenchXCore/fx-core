package loadtime

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/crypto/hd"
	"github.com/informalsystems/tm-load-test/pkg/loadtest"

	"github.com/functionx/fx-core/v4/testutil/helpers"
)

var _ loadtest.ClientFactory = (*BaseFactory)(nil)

type BuildMsgClient func(chainInfo *ChainInfo) loadtest.Client

type ChainInfo struct {
	Accounts
	ChainID  string
	GasPrice sdk.Coin
	GasLimit int64
}

type BaseFactory struct {
	Mnemonic string
	RPC      string

	msgClient BuildMsgClient
}

func NewBaseFactory(mnemonic, rpc string, msgClient BuildMsgClient) *BaseFactory {
	return &BaseFactory{Mnemonic: mnemonic, RPC: rpc, msgClient: msgClient}
}

func (f *BaseFactory) ValidateConfig(_ loadtest.Config) error { return nil }

func (f *BaseFactory) NewClient(_ loadtest.Config) (loadtest.Client, error) {
	privateKey, err := helpers.PrivKeyFromMnemonic(f.Mnemonic, hd.EthSecp256k1Type, 0, 0)
	if err != nil {
		return nil, err
	}
	newCreateAccount, err := NewCreateAccount(f.RPC, privateKey, 1000)
	if err != nil {
		return nil, err
	}
	accounts, err := ReadAccountsFromFile()
	if err != nil {
		return nil, err
	}
	if len(accounts) != 0 {
		if err = newCreateAccount.SetAccounts(accounts); err != nil {
			accounts = nil
		}
	}
	if len(accounts) == 0 {
		if err = newCreateAccount.BatchCreateNewAccount(); err != nil {
			return nil, err
		}
	}
	chainInfo := newCreateAccount.ChainInfo
	return f.msgClient(chainInfo), nil
}
