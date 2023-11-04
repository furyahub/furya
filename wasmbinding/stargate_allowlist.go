package wasmbinding

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	ibcconnectiontypes "github.com/cosmos/ibc-go/v4/modules/core/03-connection/types"
	contractmanagertypes "github.com/furyahub/furya/x/contractmanager/types"
	feeburnertypes "github.com/furyahub/furya/x/feeburner/types"
	interchainqueriestypes "github.com/furyahub/furya/x/interchainqueries/types"
	interchaintxstypes "github.com/furyahub/furya/x/interchaintxs/types"
	tokenfactorytypes "github.com/furyahub/furya/x/tokenfactory/types"
)

func AcceptedStargateQueries() wasmkeeper.AcceptedStargateQueries {
	return wasmkeeper.AcceptedStargateQueries{
		// ibc
		"/ibc.core.client.v1.Query/ClientState":    &ibcclienttypes.QueryClientStateResponse{},
		"/ibc.core.client.v1.Query/ConsensusState": &ibcclienttypes.QueryConsensusStateResponse{},
		"/ibc.core.connection.v1.Query/Connection": &ibcconnectiontypes.QueryConnectionResponse{},

		// token factory
		"/osmosis.tokenfactory.v1beta1.Query/Params":                 &tokenfactorytypes.QueryParamsResponse{},
		"/osmosis.tokenfactory.v1beta1.Query/DenomAuthorityMetadata": &tokenfactorytypes.QueryDenomAuthorityMetadataResponse{},
		"/osmosis.tokenfactory.v1beta1.Query/DenomsFromCreator":      &tokenfactorytypes.QueryDenomsFromCreatorResponse{},

		// transfer
		"/ibc.applications.transfer.v1.Query/DenomTrace": &ibctransfertypes.QueryDenomTraceResponse{},

		// auth
		"/cosmos.auth.v1beta1.Query/Account": &authtypes.QueryAccountResponse{},
		"/cosmos.auth.v1beta1.Query/Params":  &authtypes.QueryParamsResponse{},

		// bank
		"/cosmos.bank.v1beta1.Query/Balance":       &banktypes.QueryBalanceResponse{},
		"/cosmos.bank.v1beta1.Query/DenomMetadata": &banktypes.QueryDenomsMetadataResponse{},
		"/cosmos.bank.v1beta1.Query/Params":        &banktypes.QueryParamsResponse{},
		"/cosmos.bank.v1beta1.Query/SupplyOf":      &banktypes.QuerySupplyOfResponse{},

		// contractmanager
		"/furya.contractmanager.Query/AddressFailures": &contractmanagertypes.QueryFailuresResponse{},
		"/furya.contractmanager.Query/Failures":        &contractmanagertypes.QueryFailuresResponse{},

		// interchaintxs
		"/furya.interchaintxs.Query/Params": &interchaintxstypes.QueryParamsResponse{},

		// interchainqueries
		"/furya.interchainqueries.Query/Params": &interchainqueriestypes.QueryParamsResponse{},

		// feeburner
		"/furya.feeburner.Query/Params": &feeburnertypes.QueryParamsResponse{},
	}
}
