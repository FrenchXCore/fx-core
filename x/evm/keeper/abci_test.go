package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/abci/types"

	evmtypes "github.com/functionx/fx-core/v5/x/evm/types"
)

func (suite *KeeperTestSuite) TestEndBlock() {
	ctx := suite.ctx.WithEventManager(sdk.NewEventManager())
	em := ctx.EventManager()
	suite.Require().Equal(0, len(em.Events()))

	res := suite.app.EvmKeeper.EndBlock(ctx, types.RequestEndBlock{})
	suite.Require().Equal([]types.ValidatorUpdate{}, res)

	// should emit 1 EventTypeBlockBloom event on EndBlock
	suite.Require().Equal(1, len(em.Events()))
	suite.Require().Equal(evmtypes.EventTypeBlockBloom, em.Events()[0].Type)
}
