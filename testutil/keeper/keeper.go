package keeper

import (
	tokenfactorykeeper "noble/x/tokenfactory/keeper"
	tokenfactorytypes "noble/x/tokenfactory/types"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	// ExampleTimestamp is a timestamp used as the current time for the context of the keepers returned from the package
	ExampleTimestamp = time.Date(2020, time.January, 1, 12, 0, 0, 0, time.UTC)

	// ExampleHeight is a block height used as the current block height for the context of test keeper
	ExampleHeight = int64(1111)
)

// TestKeepers holds all keepers used during keeper tests for all modules
type TestKeepers struct {
	T                  testing.TB
	AccountKeeper      authkeeper.AccountKeeper
	BankKeeper         bankkeeper.Keeper
	TokenfactoryKeeper *tokenfactorykeeper.Keeper
}

// TestMsgServers holds all message servers used during keeper tests for all modules
type TestMsgServers struct {
	T               testing.TB
	TokenfactorySrv tokenfactorytypes.MsgServer
}

func NewTestSetup(t testing.TB) (sdk.Context, TestKeepers, TestMsgServers) {
	initializer := newInitializer()

	paramKeeper := initializer.Param()
	authKeeper := initializer.Auth(paramKeeper)
	bankKeeper := initializer.Bank(paramKeeper, authKeeper)
	tokenfactoryKeeper := initializer.Tokenfactory(bankKeeper, paramKeeper, authKeeper)
	initializer.StateStore.LoadLatestVersion()
	ctx := sdk.NewContext(initializer.StateStore, tmproto.Header{
		Time:   ExampleTimestamp,
		Height: ExampleHeight,
	}, false, log.NewNopLogger())

	tokenfactoryKeeper.SetParams(ctx, tokenfactorytypes.DefaultParams())

	tokenfactorySrv := tokenfactorykeeper.NewMsgServerImpl(*tokenfactoryKeeper)

	return ctx, TestKeepers{
			T:                  t,
			TokenfactoryKeeper: tokenfactoryKeeper,
			BankKeeper:         bankKeeper,
			AccountKeeper:      authKeeper,
		}, TestMsgServers{
			T:               t,
			TokenfactorySrv: tokenfactorySrv,
		}
}
