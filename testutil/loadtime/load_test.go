package loadtime_test

import (
	"testing"

	"github.com/informalsystems/tm-load-test/pkg/loadtest"
	"github.com/sirupsen/logrus"

	"github.com/functionx/fx-core/v4/app"
	"github.com/functionx/fx-core/v4/testutil/helpers"
	"github.com/functionx/fx-core/v4/testutil/loadtime"
	fxtypes "github.com/functionx/fx-core/v4/types"
)

func TestBankTest(t *testing.T) {
	var (
		mnemonic = "test test test test test test test test test test test junk"
		rpc      = "http://127.0.0.1:26657"
	)
	fxtypes.SetConfig(true)
	toAddress := helpers.CreateRandomAccounts(1)[0].String()
	if err := loadtest.RegisterClientFactory("msgSend", loadtime.NewBaseFactory(mnemonic, rpc, loadtime.NewMsgSendClient(toAddress).BuildMsgClient)); err != nil {
		panic(err)
	}
	if err := loadtest.RegisterClientFactory("ibc_transfer", loadtime.NewBaseFactory(mnemonic, rpc, loadtime.NewMsgIbcTransferClient("", "transfer", "channel-0", "").BuildMsgClient)); err != nil {
		panic(err)
	}

	if err := loadtest.RegisterClientFactory("ethereum_transfer", loadtime.NewBaseFactory(mnemonic, rpc, loadtime.NewMsgEthereumTxClient(530, fxtypes.DefaultDenom, fxtypes.EmptyEvmAddress, loadtime.EthereumTxTransferFx, app.MakeEncodingConfig().TxConfig.NewTxBuilder()).BuildMsgClient)); err != nil {
		panic(err)
	}

	logrus.SetLevel(logrus.DebugLevel)

	cfg := loadtest.Config{
		ClientFactory:        "ethereum_transfer",
		Connections:          1,
		Time:                 60,
		SendPeriod:           1,
		Rate:                 1000,
		Count:                60000,
		PeerConnectTimeout:   600,
		Endpoints:            []string{"ws://127.0.0.1:26657/websocket"},
		BroadcastTxMethod:    "async",
		NoTrapInterrupts:     true,
		EndpointSelectMethod: loadtest.SelectSuppliedEndpoints,
	}
	if err := loadtest.ExecuteStandalone(cfg); err != nil {
		panic(err)
	}
}
