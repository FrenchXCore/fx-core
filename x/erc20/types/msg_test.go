package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"

	"github.com/functionx/fx-core/v5/testutil/helpers"
	"github.com/functionx/fx-core/v5/x/erc20/types"
)

type MsgsTestSuite struct {
	suite.Suite
}

func TestMsgsTestSuite(t *testing.T) {
	suite.Run(t, new(MsgsTestSuite))
}

func (suite *MsgsTestSuite) TestMsgConvertCoinGetters() {
	msgInvalid := types.MsgConvertCoin{}
	msg := types.NewMsgConvertCoin(
		sdk.NewCoin("test", sdkmath.NewInt(100)),
		helpers.GenerateAddress(),
		helpers.GenerateAddress().Bytes(),
	)
	suite.Require().Equal(types.RouterKey, msg.Route())
	suite.Require().Equal(types.TypeMsgConvertCoin, msg.Type())
	suite.Require().NotNil(msgInvalid.GetSignBytes())
	suite.Require().NotNil(msg.GetSigners())
}

func (suite *MsgsTestSuite) TestMsgConvertCoinNew() {
	testCases := []struct {
		msg        string
		coin       sdk.Coin
		receiver   common.Address
		sender     sdk.AccAddress
		expectPass bool
	}{
		{
			"msg convert coin - pass",
			sdk.NewCoin("test", sdkmath.NewInt(100)),
			helpers.GenerateAddress(),
			sdk.AccAddress(helpers.GenerateAddress().Bytes()),
			true,
		},
	}

	for i, tc := range testCases {
		tx := types.NewMsgConvertCoin(tc.coin, tc.receiver, tc.sender)
		err := tx.ValidateBasic()

		if tc.expectPass {
			suite.Require().NoError(err, "valid test %d failed: %s, %v", i, tc.msg)
		} else {
			suite.Require().Error(err, "invalid test %d passed: %s, %v", i, tc.msg)
		}
	}
}

func (suite *MsgsTestSuite) TestMsgConvertCoin() {
	testCases := []struct {
		msg        string
		coin       sdk.Coin
		receiver   string
		sender     string
		expectPass bool
	}{
		{
			"invalid denom",
			sdk.Coin{
				Denom:  "",
				Amount: sdkmath.NewInt(100),
			},
			"0x0000",
			helpers.GenerateAddress().String(),
			false,
		},
		{
			"negative coin amount",
			sdk.Coin{
				Denom:  "coin",
				Amount: sdkmath.NewInt(-100),
			},
			"0x0000",
			helpers.GenerateAddress().String(),
			false,
		},
		{
			"msg convert coin - invalid sender",
			sdk.NewCoin("coin", sdkmath.NewInt(100)),
			helpers.GenerateAddress().String(),
			"evmosinvalid",
			false,
		},
		{
			"msg convert coin - invalid receiver",
			sdk.NewCoin("coin", sdkmath.NewInt(100)),
			"0x0000",
			sdk.AccAddress(helpers.GenerateAddress().Bytes()).String(),
			false,
		},
		{
			"msg convert coin - pass",
			sdk.NewCoin("coin", sdkmath.NewInt(100)),
			helpers.GenerateAddress().String(),
			sdk.AccAddress(helpers.GenerateAddress().Bytes()).String(),
			true,
		},
		{
			"msg convert coin - pass with `erc20/` denom",
			sdk.NewCoin("erc20/0xdac17f958d2ee523a2206206994597c13d831ec7", sdkmath.NewInt(100)),
			helpers.GenerateAddress().String(),
			sdk.AccAddress(helpers.GenerateAddress().Bytes()).String(),
			true,
		},
		{
			"msg convert coin - pass with `ibc/{hash}` denom",
			sdk.NewCoin("ibc/7F1D3FCF4AE79E1554D670D1AD949A9BA4E4A3C76C63093E17E446A46061A7A2", sdkmath.NewInt(100)),
			helpers.GenerateAddress().String(),
			sdk.AccAddress(helpers.GenerateAddress().Bytes()).String(),
			true,
		},
	}

	for i, tc := range testCases {
		tx := types.MsgConvertCoin{tc.coin, tc.receiver, tc.sender}
		err := tx.ValidateBasic()

		if tc.expectPass {
			suite.Require().NoError(err, "valid test %d failed: %s, %v", i, tc.msg)
		} else {
			suite.Require().Error(err, "invalid test %d passed: %s, %v", i, tc.msg)
		}
	}
}

func (suite *MsgsTestSuite) TestMsgConvertERC20Getters() {
	msgInvalid := types.MsgConvertERC20{}
	msg := types.NewMsgConvertERC20(
		sdkmath.NewInt(100),
		helpers.GenerateAddress().Bytes(),
		helpers.GenerateAddress(),
		helpers.GenerateAddress(),
	)
	suite.Require().Equal(types.RouterKey, msg.Route())
	suite.Require().Equal(types.TypeMsgConvertERC20, msg.Type())
	suite.Require().NotNil(msgInvalid.GetSignBytes())
	suite.Require().NotNil(msg.GetSigners())
}

func (suite *MsgsTestSuite) TestMsgConvertERC20New() {
	testCases := []struct {
		msg        string
		amount     sdkmath.Int
		receiver   sdk.AccAddress
		contract   common.Address
		sender     common.Address
		expectPass bool
	}{
		{
			"msg convert erc20 - pass",
			sdkmath.NewInt(100),
			sdk.AccAddress(helpers.GenerateAddress().Bytes()),
			helpers.GenerateAddress(),
			helpers.GenerateAddress(),
			true,
		},
	}

	for i, tc := range testCases {
		tx := types.NewMsgConvertERC20(tc.amount, tc.receiver, tc.contract, tc.sender)
		err := tx.ValidateBasic()

		if tc.expectPass {
			suite.Require().NoError(err, "valid test %d failed: %s, %v", i, tc.msg)
		} else {
			suite.Require().Error(err, "invalid test %d passed: %s, %v", i, tc.msg)
		}
	}
}

func (suite *MsgsTestSuite) TestMsgConvertERC20() {
	testCases := []struct {
		msg        string
		amount     sdkmath.Int
		receiver   string
		contract   string
		sender     string
		expectPass bool
	}{
		{
			"invalid contract hex address",
			sdkmath.NewInt(100),
			sdk.AccAddress(helpers.GenerateAddress().Bytes()).String(),
			sdk.AccAddress{}.String(),
			helpers.GenerateAddress().String(),
			false,
		},
		{
			"negative coin amount",
			sdkmath.NewInt(-100),
			sdk.AccAddress(helpers.GenerateAddress().Bytes()).String(),
			helpers.GenerateAddress().String(),
			helpers.GenerateAddress().String(),
			false,
		},
		{
			"invalid receiver address",
			sdkmath.NewInt(100),
			sdk.AccAddress{}.String(),
			helpers.GenerateAddress().String(),
			helpers.GenerateAddress().String(),
			false,
		},
		{
			"invalid sender address",
			sdkmath.NewInt(100),
			sdk.AccAddress(helpers.GenerateAddress().Bytes()).String(),
			helpers.GenerateAddress().String(),
			sdk.AccAddress{}.String(),
			false,
		},
		{
			"msg convert erc20 - pass",
			sdkmath.NewInt(100),
			sdk.AccAddress(helpers.GenerateAddress().Bytes()).String(),
			helpers.GenerateAddress().String(),
			helpers.GenerateAddress().String(),
			true,
		},
	}

	for i, tc := range testCases {
		tx := types.MsgConvertERC20{ContractAddress: tc.contract, Amount: tc.amount, Receiver: tc.receiver, Sender: tc.sender}
		err := tx.ValidateBasic()

		if tc.expectPass {
			suite.Require().NoError(err, "valid test %d failed: %s, %v", i, tc.msg)
		} else {
			suite.Require().Error(err, "invalid test %d passed: %s, %v", i, tc.msg)
		}
	}
}
