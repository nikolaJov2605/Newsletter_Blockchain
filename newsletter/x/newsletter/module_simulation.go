package newsletter

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"newsletter/testutil/sample"
	newslettersimulation "newsletter/x/newsletter/simulation"
	"newsletter/x/newsletter/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = newslettersimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateNewsletter = "op_weight_msg_create_newsletter"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateNewsletter int = 100

	opWeightMsgSubscribe = "op_weight_msg_subscribe"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubscribe int = 100

	opWeightMsgWithdraw = "op_weight_msg_withdraw"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdraw int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	newsletterGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&newsletterGenesis)
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

	var weightMsgCreateNewsletter int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateNewsletter, &weightMsgCreateNewsletter, nil,
		func(_ *rand.Rand) {
			weightMsgCreateNewsletter = defaultWeightMsgCreateNewsletter
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateNewsletter,
		newslettersimulation.SimulateMsgCreateNewsletter(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubscribe int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubscribe, &weightMsgSubscribe, nil,
		func(_ *rand.Rand) {
			weightMsgSubscribe = defaultWeightMsgSubscribe
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubscribe,
		newslettersimulation.SimulateMsgSubscribe(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdraw int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdraw, &weightMsgWithdraw, nil,
		func(_ *rand.Rand) {
			weightMsgWithdraw = defaultWeightMsgWithdraw
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdraw,
		newslettersimulation.SimulateMsgWithdraw(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
