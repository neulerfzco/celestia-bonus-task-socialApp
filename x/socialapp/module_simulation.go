package socialapp

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"socialapp/testutil/sample"
	socialappsimulation "socialapp/x/socialapp/simulation"
	"socialapp/x/socialapp/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = socialappsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateProfile = "op_weight_msg_create_profile"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateProfile int = 100

	opWeightMsgUpdateProfile = "op_weight_msg_update_profile"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateProfile int = 100

	opWeightMsgDeleteProfile = "op_weight_msg_delete_profile"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteProfile int = 100

	opWeightMsgCreatePost = "op_weight_msg_create_post"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePost int = 100

	opWeightMsgDeletePost = "op_weight_msg_delete_post"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeletePost int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	socialappGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&socialappGenesis)
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

	var weightMsgCreateProfile int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateProfile, &weightMsgCreateProfile, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProfile = defaultWeightMsgCreateProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateProfile,
		socialappsimulation.SimulateMsgCreateProfile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateProfile int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateProfile, &weightMsgUpdateProfile, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProfile = defaultWeightMsgUpdateProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateProfile,
		socialappsimulation.SimulateMsgUpdateProfile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteProfile int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteProfile, &weightMsgDeleteProfile, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteProfile = defaultWeightMsgDeleteProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteProfile,
		socialappsimulation.SimulateMsgDeleteProfile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreatePost int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreatePost, &weightMsgCreatePost, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePost = defaultWeightMsgCreatePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePost,
		socialappsimulation.SimulateMsgCreatePost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeletePost int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeletePost, &weightMsgDeletePost, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePost = defaultWeightMsgDeletePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePost,
		socialappsimulation.SimulateMsgDeletePost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
