package loadtime

import (
	"math/big"
	"math/rand"
	"sync"

	"github.com/cosmos/cosmos-sdk/client"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"github.com/informalsystems/tm-load-test/pkg/loadtest"
	tmrand "github.com/tendermint/tendermint/libs/rand"
)

var _ loadtest.Client = (*MsgEthereumTxClient)(nil)

type RandomRawFn func(priKey cryptotypes.PrivKey, chainId int64, nonce uint64, contract common.Address) ([]byte, error)

type MsgEthereumTxClient struct {
	*ChainInfo

	EvmChainId int64
	Denom      string
	Contract   common.Address
	RandomRaw  RandomRawFn
	Builder    client.TxBuilder
	Mux        sync.Mutex
}

func NewMsgEthereumTxClient(evmChainId int64, denom, contract string, randomRaw RandomRawFn, builder client.TxBuilder,
) *MsgEthereumTxClient {
	return &MsgEthereumTxClient{
		EvmChainId: evmChainId,
		Denom:      denom,
		Contract:   common.HexToAddress(contract),
		RandomRaw:  randomRaw,
		Builder:    builder,
	}
}

func (c *MsgEthereumTxClient) BuildMsgClient(chainInfo *ChainInfo) loadtest.Client {
	c.ChainInfo = chainInfo
	return c
}

func (c *MsgEthereumTxClient) GenerateTx() ([]byte, error) {
	account := c.Accounts[tmrand.Intn(len(c.Accounts)-1)]
	raw, err := c.RandomRaw(account.PrivateKey, c.EvmChainId, account.Sequence.Load(), c.Contract)
	if err != nil {
		return nil, err
	}
	account.Sequence.Add(1)
	return c.wrapEthereumTx(raw)
}

func (c *MsgEthereumTxClient) wrapEthereumTx(txRaw []byte) ([]byte, error) {
	msg := &evmtypes.MsgEthereumTx{}
	if err := msg.UnmarshalBinary(txRaw); err != nil {
		return nil, err
	}
	c.Mux.Lock()
	defer c.Mux.Unlock()
	tx, err := msg.BuildTx(c.Builder, c.Denom)
	if err != nil {
		return nil, err
	}
	encoder := authtx.DefaultTxEncoder()
	txBytes, err := encoder(tx)
	if err != nil {
		return nil, err
	}
	return txBytes, nil
}

var GasFeeCap, _ = big.NewInt(0).SetString("500000000000", 10)

func RandomAddress() common.Address {
	bz := make([]byte, 20)
	rand.Read(bz)
	return common.BytesToAddress(bz)
}

func WrapDynamicFeeTx(priKey cryptotypes.PrivKey, chainId int64, nonce uint64, to common.Address, value *big.Int, data []byte) ([]byte, error) {
	if value == nil {
		value = big.NewInt(0)
	}
	baseTx := &types.DynamicFeeTx{
		ChainID:   big.NewInt(chainId),
		Nonce:     nonce,
		GasFeeCap: GasFeeCap,
		GasTipCap: big.NewInt(0),
		Gas:       25000000,
		To:        &to,
		Value:     value,
		Data:      data,
	}
	rawTx := types.NewTx(baseTx)
	signer := types.NewLondonSigner(big.NewInt(chainId))
	signature, err := priKey.Sign(signer.Hash(rawTx).Bytes())
	if err != nil {
		return nil, err
	}
	rawTx, err = rawTx.WithSignature(signer, signature)
	if err != nil {
		return nil, err
	}
	return rawTx.MarshalBinary()
}

func EthereumTxTransferFx(priKey cryptotypes.PrivKey, chainId int64, nonce uint64, _ common.Address) ([]byte, error) {
	address := RandomAddress()
	value := big.NewInt(tmrand.Int63n(10) + 1)

	return WrapDynamicFeeTx(priKey, chainId, nonce, address, value, nil)
}
