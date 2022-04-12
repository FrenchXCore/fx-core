package evm_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/functionx/fx-core/app/fxcore"
	_ "github.com/functionx/fx-core/app/fxcore"
	"github.com/functionx/fx-core/contracts"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/rpc/client/http"
	"strings"
	"testing"
)

func TestQueryBalance(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	client, err := ethclient.Dial("http://0.0.0.0:8545")
	require.NoError(t, err)

	addressBytes, err := sdk.AccAddressFromBech32("fx17ykqect7ee5e9r4l2end78d8gmp6mauzj87cwz")
	require.NoError(t, err)

	address := common.BytesToAddress(addressBytes)
	println(address.Hex())
	balanceRes, err := client.BalanceAt(context.Background(), address, nil)
	require.NoError(t, err)
	println(balanceRes.String())
}

func TestQueryTransaction(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	client, err := ethclient.Dial("http://0.0.0.0:8545")
	require.NoError(t, err)

	transactionReceipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0x74a90ed91f42baa375804c22e2fa17087a6060bbca4ffb8f1e0fc1446883a0f7"))
	require.NoError(t, err)
	t.Logf("transactionReceipt:%+#v", transactionReceipt)
}

func TestQueryFxTxByEvmHash(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	client, err := ethclient.Dial("http://0.0.0.0:8545")
	require.NoError(t, err)

	transactionReceipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0x74a90ed91f42baa375804c22e2fa17087a6060bbca4ffb8f1e0fc1446883a0f7"))
	require.NoError(t, err)
	t.Logf("transactionReceipt:%+#v", transactionReceipt)

	fxClient, err := http.New("http://0.0.0.0:26657", "/websocket")
	require.NoError(t, err)
	evmHashBlockNumber := transactionReceipt.BlockNumber.Int64()
	blockData, err := fxClient.Block(context.Background(), &evmHashBlockNumber)
	require.NoError(t, err)
	require.True(t, uint(len(blockData.Block.Txs)) > transactionReceipt.TransactionIndex)
	fxTx := blockData.Block.Txs[transactionReceipt.TransactionIndex]
	encodingConfig := fxcore.MakeEncodingConfig()
	tx, err := encodingConfig.TxConfig.TxDecoder()(fxTx)
	require.NoError(t, err)
	txJsonStr, err := encodingConfig.TxConfig.TxJSONEncoder()(tx)
	//marshalIndent, err := json.MarshalIndent(string(txJsonStr), "", "  ")
	//require.NoError(t, err)
	t.Logf("\nTxHash:%x\nData:\n%v", fxTx.Hash(), string(txJsonStr))

}
func TestMnemonicToFxPrivate(t *testing.T) {
	privKey, err := mnemonicToFxPrivKey("december slow blue fury silly bread friend unknown render resource dry buyer brand final abstract gallery slow since hood shadow neglect travel convince foil")
	require.NoError(t, err)
	t.Logf("%x", privKey.Key)
}

func TestEthPrivateKeyToAddress(t *testing.T) {
	//privateKey, err := crypto.GenerateKey()
	//require.NoError(t, err)
	//fromECDSA := crypto.FromECDSA(privateKey)
	//t.Logf("fromEc:%x", fromECDSA)

	// 1ce31354ff0a3f057c9b70ebbbdacb68ace4bf9c008ac722f2b996328ab3ca08
	hexPrivKey := "86b87f127b6e0901f7f00aa77b6c82624847f2628a901bf1833b2d48883b73d3"
	recoverPrivKey, err := crypto.HexToECDSA(hexPrivKey)
	require.NoError(t, err)
	address := crypto.PubkeyToAddress(recoverPrivKey.PublicKey)
	t.Logf("Eth address:%v, FxAddress:%v", address.Hex(), sdk.AccAddress(address.Bytes()).String())
}

func TestEthAddressToFxAddress(t *testing.T) {
	ethAddress := common.HexToAddress("0xf12C0Ce17eCE69928ebf5666Df1Da746c3adf782")
	t.Logf("%o", ethAddress.Bytes())
	t.Logf("EthAddress:%v, FxAddress:%v", ethAddress.Hex(), sdk.AccAddress(ethAddress.Bytes()).String())
}

func TestFxAddressToEthAddress(t *testing.T) {
	fxAddress, err := sdk.AccAddressFromBech32("fx10kg059hhxc2pevxssszunvgc70jpmxsjal4xf6")
	require.NoError(t, err)
	ethAddress := common.BytesToAddress(fxAddress)
	t.Logf("EthAddress:%v, FxAddress:%v", ethAddress.Hex(), sdk.AccAddress(ethAddress.Bytes()).String())
}

func TestTraverseBlockERC20(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	fxClient, err := http.New("http://127.0.0.1:26657", "/websocket")
	require.NoError(t, err)

	ctx := context.Background()
	info, err := fxClient.Status(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for i := int64(1); i < info.SyncInfo.LatestBlockHeight; i++ {
		block, err := fxClient.BlockResults(ctx, &i)
		if err != nil {
			t.Fatal(err)
		}
		for _, v := range block.EndBlockEvents {
			for _, vv := range v.Attributes {
				if strings.EqualFold("fip20_symbol", string(vv.Key)) {
					fmt.Println(i, "fip20 symbol:", string(vv.Value))
				}
				if strings.EqualFold("fip20_token", string(vv.Key)) {
					fmt.Println(i, "fip20 address:", string(vv.Value))
				}
			}
		}
	}
}

func mnemonicToFxPrivKey(mnemonic string) (*secp256k1.PrivKey, error) {
	algo := hd.Secp256k1
	bytes, err := algo.Derive()(mnemonic, "", "m/44'/118'/0'/0/0")
	if err != nil {
		return nil, err
	}
	privKey := algo.Generate()(bytes)
	priv, ok := privKey.(*secp256k1.PrivKey)
	if !ok {
		return nil, fmt.Errorf("not secp256k1.PrivKey")
	}
	return priv, nil
}

func TestFIP20Code(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	codeAddr := "0x5f123738067a8BAA3E9bb8Cd7e4A8827474a2F53"
	codeBase64 := "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"

	code, codeNew := replayCodeAddress(t, codeBase64, codeAddr, contracts.FIP20LogicAddress)
	t.Log("code", hex.EncodeToString(code))
	t.Log("new-code", hex.EncodeToString(codeNew))
}

func TestWFXCode(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	codeAddr := "0x5882566ad042E16F14633a0D77b705E9a912e94d"
	codeBase64 := "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"

	code, codeNew := replayCodeAddress(t, codeBase64, codeAddr, contracts.WFXLogicAddress)
	t.Log("code", hex.EncodeToString(code))
	t.Log("new-code", hex.EncodeToString(codeNew))
}

func replayCodeAddress(t *testing.T, codeBase64, addr, addrNew string) (code, codeNew []byte) {
	bz, err := base64.StdEncoding.DecodeString(codeBase64)
	require.NoError(t, err)

	addr1 := common.HexToAddress(addr)
	addr2 := common.HexToAddress(addrNew)

	bzZero := bytes.ReplaceAll(bz, addr1.Bytes(), common.HexToAddress(contracts.EmptyEvmAddress).Bytes())
	bzNew := bytes.ReplaceAll(bz, addr1.Bytes(), addr2.Bytes())

	return bzZero, bzNew
}
