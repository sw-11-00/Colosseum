package colosseum

import (
	"math/rand"

	"Colosseum/testutil/sample"
	colosseumsimulation "Colosseum/x/colosseum/simulation"
	"Colosseum/x/colosseum/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = colosseumsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateCoinSymbol = "op_weight_msg_create_coin_symbol"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateCoinSymbol int = 100

	opWeightMsgDeleteCoinSymbol = "op_weight_msg_delete_coin_symbol"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteCoinSymbol int = 100

	opWeightMsgUpdateCoinSymbol = "op_weight_msg_update_coin_symbol"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateCoinSymbol int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	colosseumGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&colosseumGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateCoinSymbol int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateCoinSymbol, &weightMsgCreateCoinSymbol, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCoinSymbol = defaultWeightMsgCreateCoinSymbol
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCoinSymbol,
		colosseumsimulation.SimulateMsgCreateCoinSymbol(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteCoinSymbol int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteCoinSymbol, &weightMsgDeleteCoinSymbol, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteCoinSymbol = defaultWeightMsgDeleteCoinSymbol
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteCoinSymbol,
		colosseumsimulation.SimulateMsgDeleteCoinSymbol(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateCoinSymbol int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateCoinSymbol, &weightMsgUpdateCoinSymbol, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateCoinSymbol = defaultWeightMsgUpdateCoinSymbol
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateCoinSymbol,
		colosseumsimulation.SimulateMsgUpdateCoinSymbol(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
