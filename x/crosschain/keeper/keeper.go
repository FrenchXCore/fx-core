package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/functionx/fx-core/v2/x/crosschain/types"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	moduleName string
	cdc        codec.BinaryCodec // The wire codec for binary encoding/decoding.
	storeKey   sdk.StoreKey      // Unexposed key to access store from sdk.Context
	paramSpace paramtypes.Subspace

	stakingKeeper      types.StakingKeeper
	distributionKeeper types.DistributionKeeper
	bankKeeper         types.BankKeeper
	ibcTransferKeeper  types.IBCTransferKeeper
	ibcChannelKeeper   types.IBCChannelKeeper
	erc20Keeper        types.Erc20Keeper
}

// NewKeeper returns a new instance of the gravity keeper
func NewKeeper(cdc codec.BinaryCodec, moduleName string, storeKey sdk.StoreKey, paramSpace paramtypes.Subspace,
	stakingKeeper types.StakingKeeper, distributionKeeper types.DistributionKeeper, bankKeeper types.BankKeeper,
	ibcTransferKeeper types.IBCTransferKeeper, channelKeeper types.IBCChannelKeeper, erc20Keeper types.Erc20Keeper) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}
	// set KeyTable if it has not already been set
	return Keeper{
		moduleName: moduleName,
		cdc:        cdc,
		storeKey:   storeKey,
		paramSpace: paramSpace,

		stakingKeeper:      stakingKeeper,
		distributionKeeper: distributionKeeper,
		bankKeeper:         bankKeeper,
		ibcTransferKeeper:  ibcTransferKeeper,
		ibcChannelKeeper:   channelKeeper,
		erc20Keeper:        erc20Keeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+k.moduleName)
}

/////////////////////////////
//       PARAMETERS        //
/////////////////////////////

// GetParams returns the parameters from the store
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return
}

// SetParams sets the parameters in the store
func (k Keeper) SetParams(ctx sdk.Context, ps *types.Params) {
	k.paramSpace.SetParamSet(ctx, ps)
}

// GetGravityID returns the GravityID is essentially a salt value
// for bridge signatures, provided each chain running Gravity has a unique ID
// it won't be possible to play back signatures from one bridge onto another
// even if they share a oracle set.
//
// The lifecycle of the GravityID is that it is set in the Genesis file
// read from the live chain for the contract deployment, once a Gravity contract
// is deployed the GravityID CAN NOT BE CHANGED. Meaning that it can't just be the
// same as the chain id since the chain id may be changed many times with each
// successive chain in charge of the same bridge
func (k Keeper) GetGravityID(ctx sdk.Context) string {
	var gravityId string
	k.paramSpace.Get(ctx, types.ParamsStoreKeyGravityID, &gravityId)
	return gravityId
}

func (k Keeper) GetOracleDelegateThreshold(ctx sdk.Context) sdk.Coin {
	var threshold sdk.Coin
	k.paramSpace.Get(ctx, types.ParamOracleDelegateThreshold, &threshold)
	return threshold
}

func (k Keeper) GetOracleDelegateMultiple(ctx sdk.Context) int64 {
	var multiple int64
	k.paramSpace.Get(ctx, types.ParamOracleDelegateMultiple, &multiple)
	return multiple
}

func (k Keeper) GetSlashFraction(ctx sdk.Context) sdk.Dec {
	var dec sdk.Dec
	k.paramSpace.Get(ctx, types.ParamsStoreSlashFraction, &dec)
	return dec
}

func (k Keeper) GetSignedWindow(ctx sdk.Context) uint64 {
	var i uint64
	k.paramSpace.Get(ctx, types.ParamsStoreKeySignedWindow, &i)
	return i
}

func (k Keeper) GetIbcTransferTimeoutHeight(ctx sdk.Context) uint64 {
	var i uint64
	k.paramSpace.Get(ctx, types.ParamStoreIbcTransferTimeoutHeight, &i)
	return i
}

func (k Keeper) GetOracleSetUpdatePowerChangePercent(ctx sdk.Context) sdk.Dec {
	var dec sdk.Dec
	k.paramSpace.Get(ctx, types.ParamStoreOracleSetUpdatePowerChangePercent, &dec)
	return dec
}

// SetLastOracleSlashBlockHeight sets the last proposal block height
func (k Keeper) SetLastOracleSlashBlockHeight(ctx sdk.Context, blockHeight uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LastOracleSlashBlockHeight, sdk.Uint64ToBigEndian(blockHeight))
}

// GetLastOracleSlashBlockHeight returns the last proposal block height
func (k Keeper) GetLastOracleSlashBlockHeight(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	data := store.Get(types.LastOracleSlashBlockHeight)
	if len(data) == 0 {
		return 0
	}
	return sdk.BigEndianToUint64(data)
}

// setLastEventBlockHeightByOracle set the latest event blockHeight for a give oracle
func (k Keeper) setLastEventBlockHeightByOracle(ctx sdk.Context, oracleAddr sdk.AccAddress, blockHeight uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetLastEventBlockHeightByOracleKey(oracleAddr), sdk.Uint64ToBigEndian(blockHeight))
}

// getLastEventBlockHeightByOracle get the latest event blockHeight for a give oracle
func (k Keeper) getLastEventBlockHeightByOracle(ctx sdk.Context, oracleAddr sdk.AccAddress) uint64 {
	store := ctx.KVStore(k.storeKey)
	key := types.GetLastEventBlockHeightByOracleKey(oracleAddr)
	if !store.Has(key) {
		return 0
	}
	data := store.Get(key)
	return sdk.BigEndianToUint64(data)
}

// prefixRange turns a prefix into a (start, end) range. The start is the given prefix value and
// the end is calculated by adding 1 bit to the start value. Nil is not allowed as prefix.
//
//	Example: []byte{1, 3, 4} becomes []byte{1, 3, 5}
//			 []byte{15, 42, 255, 255} becomes []byte{15, 43, 0, 0}
//
// In case of an overflow the end is set to nil.
//
//	Example: []byte{255, 255, 255, 255} becomes nil
//
// MARK finish-batches: this is where some crazy shit happens
func prefixRange(prefix []byte) ([]byte, []byte) {
	if prefix == nil {
		panic("nil key not allowed")
	}
	// special case: no prefix is whole range
	if len(prefix) == 0 {
		return nil, nil
	}

	// copy the prefix and update last byte
	end := make([]byte, len(prefix))
	copy(end, prefix)
	l := len(end) - 1
	end[l]++

	// wait, what if that overflowed?....
	for end[l] == 0 && l > 0 {
		l--
		end[l]++
	}

	// okay, funny guy, you gave us FFF, no end to this range...
	if l == 0 && end[0] == 0 {
		end = nil
	}
	return prefix, end
}
