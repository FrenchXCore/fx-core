package gov

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/functionx/fx-core/v3/x/gov/keeper"
)

var (
	_ module.AppModule         = AppModule{}
	_ module.EndBlockAppModule = AppModule{}
)

// AppModule implements an application module for the gov module.
type AppModule struct {
	gov.AppModule
	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, keeper keeper.Keeper, ak govtypes.AccountKeeper, bk govtypes.BankKeeper) AppModule {
	return AppModule{
		AppModule: gov.NewAppModule(cdc, keeper.Keeper, ak, bk),
		keeper:    keeper,
	}
}

// Deprecated: Route returns the message routing key
func (am AppModule) Route() sdk.Route {
	return sdk.Route{}
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	govtypes.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(govkeeper.NewMsgServerImpl(am.keeper.Keeper), am.keeper))
	govtypes.RegisterQueryServer(cfg.QueryServer(), am.keeper)

	m := govkeeper.NewMigrator(am.keeper.Keeper)
	err := cfg.RegisterMigration(govtypes.ModuleName, 1, m.Migrate1to2)
	if err != nil {
		panic(err)
	}
}

// EndBlock returns the end blocker for the gov module. It returns no validator
// updates.
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	am.keeper.EndBlocker(ctx)
	return []abci.ValidatorUpdate{}
}
