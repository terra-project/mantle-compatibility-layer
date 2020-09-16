package app

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
	terra "github.com/terra-project/core/app"
	wasmconfig "github.com/terra-project/core/x/wasm/config"
	"io/ioutil"
)

func NewTerraApp(db dbm.DB) *terra.TerraApp {
	return terra.NewTerraApp(
		log.NewTMLogger(ioutil.Discard),
		db,
		nil,
		true, // need this so KVStores are set
		0,
		make(map[int64]bool),
		&wasmconfig.Config{BaseConfig: wasmconfig.BaseConfig{
			ContractQueryGasLimit: viper.GetUint64(wasmconfig.FlagContractQueryGasLimit),
			CacheSize:             viper.GetUint64(wasmconfig.FlagCacheSize),
		}},
		fauxMerkleModeOpt, // error
		setPruningOptions(),
	)
}

// Pass this in as an option to use a dbStoreAdapter instead of an IAVLStore for simulation speed.
func fauxMerkleModeOpt(bapp *baseapp.BaseApp) {
	bapp.SetFauxMerkleMode()
}

func setPruningOptions() func(*baseapp.BaseApp) {
	// prune nothing
	pruningOptions := sdk.PruningOptions{
		KeepRecent: 0,
		KeepEvery:  0,
		Interval:   10,
	}
	return baseapp.SetPruning(pruningOptions)
}