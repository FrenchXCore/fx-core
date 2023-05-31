package loadtime

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client/flags"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/evmos/ethermint/crypto/hd"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/functionx/fx-core/v4/client"
	"github.com/functionx/fx-core/v4/client/jsonrpc"
	"github.com/functionx/fx-core/v4/testutil/helpers"
	"github.com/functionx/fx-core/v4/testutil/loadtime/logging"
	fxtypes "github.com/functionx/fx-core/v4/types"
)

const OutPath = "out"

type Account struct {
	PrivateKey    cryptotypes.PrivKey `json:"-"`
	Mnemonic      string
	AccountNumber uint64
	Sequence      atomic.Uint64 `json:"-"`
	Balance       sdk.Coin      `json:"-"`
}

type CreateAccount struct {
	*ChainInfo
	*jsonrpc.NodeRPC
	CreateNumber int
	InitAccount  *Account
	NewAccChan   chan *Account
	sync.Mutex

	Logger logging.Logger
}

func NewCreateAccount(url string, privateKey cryptotypes.PrivKey, number int) (*CreateAccount, error) {
	nodeRPC := jsonrpc.NewNodeRPC(jsonrpc.NewClient(url))
	accAddr := sdk.AccAddress(privateKey.PubKey().Address())
	balance, err := nodeRPC.QueryBalance(accAddr.String(), fxtypes.DefaultDenom)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	gasPrices, err := nodeRPC.GetGasPrices()
	if err != nil {
		return nil, err
	}
	chainId, err := nodeRPC.GetChainId()
	if err != nil {
		return nil, err
	}
	accounts := make([]*Account, 0, number)
	return &CreateAccount{
		ChainInfo: &ChainInfo{
			Accounts: accounts,
			ChainID:  chainId,
			GasPrice: sdk.NewCoin(fxtypes.DefaultDenom, gasPrices.AmountOf(fxtypes.DefaultDenom)),
			GasLimit: flags.DefaultGasLimit,
		},
		NodeRPC:      nodeRPC,
		CreateNumber: number,
		InitAccount: &Account{
			PrivateKey: privateKey,
			Balance:    balance,
		},
		NewAccChan: make(chan *Account, number),
		Logger:     logging.NewLogrusLogger("account"),
	}, nil
}

func (c *CreateAccount) SetAccounts(account []*Account) error {
	if err := c.syncSequence(account); err != nil {
		return err
	}
	c.Accounts = account
	return nil
}

func (c *CreateAccount) BatchCreateNewAccount() error {
	c.Logger.Info("start batch create accounts", "number", c.CreateNumber)
	if c.CreateNumber <= 1 {
		return errors.New("create number cannot be less than or equal to 1")
	}
	if err := c.CheckInitAccBalance(); err != nil {
		return err
	}
	count := 2
	index := 1
	for ; index <= 15; index++ {
		if count >= c.CreateNumber {
			break
		}
		count = count * 2
	}
	if index > 15 && count < c.CreateNumber {
		return fmt.Errorf("number of invalid created accounts: %d", c.CreateNumber)
	}
	acctAdd := make(map[string]bool, c.CreateNumber)
	c.NewAccChan <- c.InitAccount
	for i := 1; i <= count*2; i++ {
		newAcc := <-c.NewAccChan
		if newAcc == nil {
			continue
		}
		if !bytes.Equal(c.InitAccount.PrivateKey.Bytes(), newAcc.PrivateKey.Bytes()) && !acctAdd[newAcc.PrivateKey.String()] {
			acctAdd[newAcc.PrivateKey.String()] = true
			c.Accounts = append(c.Accounts, newAcc)
		}
		if len(c.Accounts) == c.CreateNumber {
			break
		}
		go c.CreateNewAccount(newAcc)
	}
	return c.Accounts.SaveAccountsToFile()
}

func (c *CreateAccount) CreateNewAccount(newAcc *Account) {
	retryCount := 0
	for {
		accountI, _ := c.QueryAccount(newAcc.Address().String())
		if (accountI != nil && newAcc.AccountNumber == 0 && newAcc.Sequence.Load() == 0) || (accountI != nil && accountI.GetSequence() == newAcc.Sequence.Load()+1) {
			c.Logger.Info("query account success", "address", newAcc.Address().String(), "accountNumber", newAcc.AccountNumber, "sequence", newAcc.Sequence.Load())
			newAcc.Sequence.Store(accountI.GetSequence())
			newAcc.AccountNumber = accountI.GetAccountNumber()
			break
		} else {
			if retryCount > 5 {
				c.deleteElement(newAcc)
				c.NewAccChan <- c.InitAccount
				return
			}
		}
		c.Lock()
		retryCount++
		c.Unlock()
		time.Sleep(5 * time.Second)
	}
	transferCoins := sdk.NewCoin(newAcc.Balance.Denom, newAcc.Balance.Amount.QuoRaw(2))
	if transferCoins.Amount.LTE(sdk.NewInt(1e17).MulRaw(5)) {
		c.NewAccChan <- nil
		return
	}
	mnemonic := helpers.NewMnemonic()
	privKey, err := helpers.PrivKeyFromMnemonic(mnemonic, hd.EthSecp256k1Type, 0, 0)
	if err != nil {
		panic(err)
	}
	newNextAcc := &Account{
		PrivateKey: privKey,
		Mnemonic:   mnemonic,
		Balance:    transferCoins,
	}
	transferMsg := banktypes.NewMsgSend(newAcc.Address(), newNextAcc.PrivateKey.PubKey().Address().Bytes(), sdk.NewCoins(transferCoins))
	txRaw, err := client.BuildTxV1(c.ChainID, newAcc.Sequence.Load(), newAcc.AccountNumber, newAcc.PrivateKey, []sdk.Msg{transferMsg}, c.GasPrice, c.GasLimit, "", 0)
	if err != nil {
		c.NewAccChan <- newAcc
		return
	}
	txResponse, err := c.BroadcastTx(txRaw, tx.BroadcastMode_BROADCAST_MODE_SYNC)
	if err != nil {
		c.NewAccChan <- newAcc
		return
	}
	c.Logger.Info("success send tx", "txHash", txResponse.TxHash, "newAccount", newNextAcc.Address().String())
	newAcc.Balance.Amount = transferCoins.Amount.Sub(sdk.NewIntFromBigInt(big.NewInt(txResponse.GasUsed)))
	c.NewAccChan <- newNextAcc
	c.NewAccChan <- newAcc
}

func (c *CreateAccount) CheckInitAccBalance() error {
	var totalAmount sdkmath.Int
	PerAccountFee := sdk.NewInt(1e17).MulRaw(5)
	totalAmount = sdk.NewIntFromBigInt(PerAccountFee.BigInt()).Mul(sdk.NewInt(int64(c.CreateNumber)))
	if c.InitAccount.Balance.Amount.LTE(totalAmount) {
		return errors.New(fmt.Sprintf("account balance is insufficient, need: %s, but balance: %s", totalAmount, c.InitAccount.Balance.Amount))
	}
	return nil
}

func (c *CreateAccount) deleteElement(account *Account) {
	result := make([]*Account, 0, c.CreateNumber)
	for _, acc := range c.Accounts {
		if acc.PrivateKey != account.PrivateKey {
			result = append(result, acc)
		}
	}
	c.Accounts = result
}

func (c *CreateAccount) syncSequence(acc []*Account) error {
	c.Logger.Info("start sync accounts sequence~")
	g, _ := errgroup.WithContext(context.Background())
	for _, acct := range acc {
		g.Go(func(acc *Account) func() error {
			return func() error {
				account, err := c.QueryAccount(sdk.AccAddress(acc.PrivateKey.PubKey().Address()).String())
				if err != nil {
					return err
				}
				acc.Sequence.Store(account.GetSequence())
				acc.AccountNumber = account.GetAccountNumber()
				return nil
			}
		}(acct))
	}
	return g.Wait()
}

type Accounts []*Account

func (acc Accounts) SaveAccountsToFile() error {
	root, err := GetRootPath()
	if err != nil {
		return err
	}
	newAccSaveFileName := path.Join(root, OutPath, fmt.Sprintf("%s.json", "fx-load-account"))
	filePath, _ := path.Split(newAccSaveFileName)
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		return err
	}
	data, err := json.MarshalIndent(acc, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(newAccSaveFileName, data, os.ModePerm)
}

func (acc *Account) Address() sdk.AccAddress {
	return sdk.AccAddress(acc.PrivateKey.PubKey().Address())
}

func ReadAccountsFromFile() ([]*Account, error) {
	root, err := GetRootPath()
	if err != nil {
		return nil, err
	}
	fileName := path.Join(root, OutPath, fmt.Sprintf("%s.json", "fx-load-account"))
	accounts := make([]*Account, 0)
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return accounts, nil
		}
		return nil, err
	}
	if err = json.Unmarshal(data, &accounts); err != nil {
		return nil, err
	}
	for _, account := range accounts {
		privKey, err := helpers.PrivKeyFromMnemonic(account.Mnemonic, hd.EthSecp256k1Type, 0, 0)
		if err != nil {
			panic(err)
		}
		account.PrivateKey = privKey
	}
	return accounts, nil
}

func GetRootPath() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		return path.Dir(path.Dir(path.Dir(filename))), nil
	}
	return "", fmt.Errorf("can not get root path")
}
