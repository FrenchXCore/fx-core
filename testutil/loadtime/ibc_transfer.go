package loadtime

import (
	"fmt"
	"math"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	"github.com/gogo/protobuf/proto"
	"github.com/informalsystems/tm-load-test/pkg/loadtest"
	"github.com/pkg/errors"
	tmrand "github.com/tendermint/tendermint/libs/rand"

	"github.com/functionx/fx-core/v4/client"
	ibctypes "github.com/functionx/fx-core/v4/x/ibc/applications/transfer/types"
)

var _ loadtest.Client = (*MsgIbcTransferClient)(nil)

type MsgIbcTransferClient struct {
	*ChainInfo

	Denom               string
	SourcePort          string
	SourceChannel       string
	TargetAddressPrefix string
}

func NewMsgIbcTransferClient(denom, sourcePort, sourceChannel, targetAddressPrefix string) *MsgIbcTransferClient {
	return &MsgIbcTransferClient{Denom: denom, SourcePort: sourcePort, SourceChannel: sourceChannel, TargetAddressPrefix: targetAddressPrefix}
}

func (c *MsgIbcTransferClient) BuildMsgClient(chainInfo *ChainInfo) loadtest.Client {
	c.ChainInfo = chainInfo
	return c
}

func (c *MsgIbcTransferClient) GenerateTx() ([]byte, error) {
	account := c.Accounts[tmrand.Intn(len(c.Accounts)-1)]
	sourceAddress := sdk.AccAddress(account.PrivateKey.PubKey().Address())
	receiverAddress, err := bech32.ConvertAndEncode(c.TargetAddressPrefix, sourceAddress)
	if err != nil {
		panic(errors.WithStack(fmt.Errorf("source address to target address err! sourceAddress:[%v], targetPrefix:[%v], error:[%v]", sourceAddress.String(), c.TargetAddressPrefix, err)))
	}
	msgs := []sdk.Msg{&ibctypes.MsgTransfer{
		SourcePort:    c.SourcePort,
		SourceChannel: c.SourceChannel,
		Token:         sdk.NewCoin(c.Denom, sdkmath.NewInt(tmrand.Int63n(10)+1)),
		Sender:        sourceAddress.String(),
		Receiver:      receiverAddress,
		TimeoutHeight: clienttypes.Height{
			RevisionNumber: 0,
			RevisionHeight: math.MaxUint64,
		},
		TimeoutTimestamp: 0,
		Router:           "",
		Fee:              sdk.NewCoin(c.Denom, sdk.ZeroInt()),
	}}
	txRaw, err := client.BuildTxV1(c.ChainID, account.Sequence.Load(), account.AccountNumber, account.PrivateKey, msgs, c.GasPrice, c.GasLimit, "", 0)
	if err != nil {
		return nil, err
	}
	account.Sequence.Add(1)
	return proto.Marshal(txRaw)
}
