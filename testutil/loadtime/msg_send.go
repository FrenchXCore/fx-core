package loadtime

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/gogo/protobuf/proto"
	"github.com/informalsystems/tm-load-test/pkg/loadtest"
	tmrand "github.com/tendermint/tendermint/libs/rand"

	"github.com/functionx/fx-core/v4/client"
	fxtypes "github.com/functionx/fx-core/v4/types"
)

var _ loadtest.Client = (*MsgSendClient)(nil)

type MsgSendClient struct {
	*ChainInfo
	toAddress string
}

func NewMsgSendClient(toAddress string) *MsgSendClient {
	return &MsgSendClient{toAddress: toAddress}
}

func (c *MsgSendClient) BuildMsgClient(chainInfo *ChainInfo) loadtest.Client {
	c.ChainInfo = chainInfo
	return c
}

func (c *MsgSendClient) GenerateTx() ([]byte, error) {
	account := c.Accounts[tmrand.Intn(len(c.Accounts)-1)]
	msgs := []sdk.Msg{&bank.MsgSend{
		FromAddress: sdk.AccAddress(account.PrivateKey.PubKey().Address()).String(),
		ToAddress:   c.toAddress,
		Amount:      sdk.NewCoins(sdk.NewCoin(fxtypes.DefaultDenom, sdkmath.NewInt(tmrand.Int63n(10)+1))),
	}}
	txRaw, err := client.BuildTxV1(c.ChainID, account.Sequence.Load(), account.AccountNumber, account.PrivateKey, msgs, c.GasPrice, c.GasLimit, "", 0)
	if err != nil {
		return nil, err
	}
	account.Sequence.Add(1)
	return proto.Marshal(txRaw)
}
