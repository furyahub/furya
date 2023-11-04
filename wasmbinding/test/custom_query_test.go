package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/furyahub/furya/app/params"
	feerefundertypes "github.com/furyahub/furya/x/feerefunder/types"
	tokenfactorytypes "github.com/furyahub/furya/x/tokenfactory/types"

	"github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/furyahub/furya/app"
	"github.com/furyahub/furya/testutil"
	"github.com/furyahub/furya/wasmbinding/bindings"
	icqtypes "github.com/furyahub/furya/x/interchainqueries/types"
	ictxtypes "github.com/furyahub/furya/x/interchaintxs/types"
)

type CustomQuerierTestSuite struct {
	testutil.IBCConnectionTestSuite
}

func (suite *CustomQuerierTestSuite) TestInterchainQueryResult() {
	var (
		furya = suite.GetFuryaZoneApp(suite.ChainA)
		ctx     = suite.ChainA.GetContext()
		owner   = keeper.RandomAccountAddress(suite.T()) // We don't care what this address is
	)

	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(ctx, owner, "../testdata/reflect.wasm")
	contractAddress := suite.InstantiateReflectContract(ctx, owner, codeID)
	suite.Require().NotEmpty(contractAddress)

	// Register and submit query result
	clientKey := host.FullClientStateKey(suite.Path.EndpointB.ClientID)
	lastID := furya.InterchainQueriesKeeper.GetLastRegisteredQueryKey(ctx) + 1
	furya.InterchainQueriesKeeper.SetLastRegisteredQueryKey(ctx, lastID)
	registeredQuery := &icqtypes.RegisteredQuery{
		Id: lastID,
		Keys: []*icqtypes.KVKey{
			{Path: host.StoreKey, Key: clientKey},
		},
		QueryType:    string(icqtypes.InterchainQueryTypeKV),
		UpdatePeriod: 1,
		ConnectionId: suite.Path.EndpointA.ConnectionID,
	}
	furya.InterchainQueriesKeeper.SetLastRegisteredQueryKey(ctx, lastID)
	err := furya.InterchainQueriesKeeper.SaveQuery(ctx, registeredQuery)
	suite.Require().NoError(err)

	chainBResp := suite.ChainB.App.Query(abci.RequestQuery{
		Path:   fmt.Sprintf("store/%s/key", host.StoreKey),
		Height: suite.ChainB.LastHeader.Header.Height - 1,
		Data:   clientKey,
		Prove:  true,
	})

	expectedQueryResult := &icqtypes.QueryResult{
		KvResults: []*icqtypes.StorageValue{{
			Key:           chainBResp.Key,
			Proof:         chainBResp.ProofOps,
			Value:         chainBResp.Value,
			StoragePrefix: host.StoreKey,
		}},
		// we don't have tests to test transactions proofs verification since it's a tendermint layer, and we don't have access to it here
		Block:    nil,
		Height:   uint64(chainBResp.Height),
		Revision: suite.ChainA.LastHeader.GetHeight().GetRevisionNumber(),
	}
	err = furya.InterchainQueriesKeeper.SaveKVQueryResult(ctx, lastID, expectedQueryResult)
	suite.Require().NoError(err)

	// Query interchain query result
	query := bindings.FuryaQuery{
		InterchainQueryResult: &bindings.QueryRegisteredQueryResultRequest{
			QueryID: lastID,
		},
	}
	resp := icqtypes.QueryRegisteredQueryResultResponse{}
	err = suite.queryCustom(ctx, contractAddress, query, &resp)
	suite.Require().NoError(err)

	suite.Require().Equal(uint64(chainBResp.Height), resp.Result.Height)
	suite.Require().Equal(suite.ChainA.LastHeader.GetHeight().GetRevisionNumber(), resp.Result.Revision)
	suite.Require().Empty(resp.Result.Block)
	suite.Require().NotEmpty(resp.Result.KvResults)
	suite.Require().Equal([]*icqtypes.StorageValue{{
		Key:           chainBResp.Key,
		Proof:         nil,
		Value:         chainBResp.Value,
		StoragePrefix: host.StoreKey,
	}}, resp.Result.KvResults)
}

func (suite *CustomQuerierTestSuite) TestInterchainQueryResultNotFound() {
	var (
		ctx   = suite.ChainA.GetContext()
		owner = keeper.RandomAccountAddress(suite.T()) // We don't care what this address is
	)

	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(ctx, owner, "../testdata/reflect.wasm")
	contractAddress := suite.InstantiateReflectContract(ctx, owner, codeID)
	suite.Require().NotEmpty(contractAddress)

	// Query interchain query result
	query := bindings.FuryaQuery{
		InterchainQueryResult: &bindings.QueryRegisteredQueryResultRequest{
			QueryID: 1,
		},
	}
	resp := icqtypes.QueryRegisteredQueryResultResponse{}
	err := suite.queryCustom(ctx, contractAddress, query, &resp)
	expectedErrMsg := fmt.Sprintf("Generic error: Querier contract error: codespace: interchainqueries, code: %d: query wasm contract failed", icqtypes.ErrNoQueryResult.ABCICode())
	suite.Require().ErrorContains(err, expectedErrMsg)
}

func (suite *CustomQuerierTestSuite) TestInterchainAccountAddress() {
	var (
		ctx   = suite.ChainA.GetContext()
		owner = keeper.RandomAccountAddress(suite.T()) // We don't care what this address is
	)

	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(ctx, owner, "../testdata/reflect.wasm")
	contractAddress := suite.InstantiateReflectContract(ctx, owner, codeID)
	suite.Require().NotEmpty(contractAddress)

	err := testutil.SetupICAPath(suite.Path, contractAddress.String())
	suite.Require().NoError(err)

	query := bindings.FuryaQuery{
		InterchainAccountAddress: &bindings.QueryInterchainAccountAddressRequest{
			OwnerAddress:        contractAddress.String(),
			InterchainAccountID: testutil.TestInterchainID,
			ConnectionID:        suite.Path.EndpointA.ConnectionID,
		},
	}
	resp := ictxtypes.QueryInterchainAccountAddressResponse{}
	err = suite.queryCustom(ctx, contractAddress, query, &resp)
	suite.Require().NoError(err)

	hostFuryaApp, ok := suite.ChainB.App.(*app.App)
	suite.Require().True(ok)

	expected := hostFuryaApp.ICAHostKeeper.GetAllInterchainAccounts(suite.ChainB.GetContext())[0].AccountAddress // we expect only one registered ICA
	suite.Require().Equal(expected, resp.InterchainAccountAddress)
}

func (suite *CustomQuerierTestSuite) TestUnknownInterchainAcc() {
	var (
		ctx   = suite.ChainA.GetContext()
		owner = keeper.RandomAccountAddress(suite.T()) // We don't care what this address is
	)

	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(ctx, owner, "../testdata/reflect.wasm")
	contractAddress := suite.InstantiateReflectContract(ctx, owner, codeID)
	suite.Require().NotEmpty(contractAddress)

	err := testutil.SetupICAPath(suite.Path, contractAddress.String())
	suite.Require().NoError(err)

	query := bindings.FuryaQuery{
		InterchainAccountAddress: &bindings.QueryInterchainAccountAddressRequest{
			OwnerAddress:        testutil.TestOwnerAddress,
			InterchainAccountID: "wrong_account_id",
			ConnectionID:        suite.Path.EndpointA.ConnectionID,
		},
	}
	resp := ictxtypes.QueryInterchainAccountAddressResponse{}
	expectedErrorMsg := "Generic error: Querier contract error: codespace: interchaintxs, code: 1102: query wasm contract failed"

	err = suite.queryCustom(ctx, contractAddress, query, &resp)
	suite.Require().ErrorContains(err, expectedErrorMsg)
}

func (suite *CustomQuerierTestSuite) TestMinIbcFee() {
	var (
		ctx   = suite.ChainA.GetContext()
		owner = keeper.RandomAccountAddress(suite.T()) // We don't care what this address is
	)

	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(ctx, owner, "../testdata/reflect.wasm")
	contractAddress := suite.InstantiateReflectContract(ctx, owner, codeID)
	suite.Require().NotEmpty(contractAddress)

	query := bindings.FuryaQuery{
		MinIbcFee: &bindings.QueryMinIbcFeeRequest{},
	}
	resp := bindings.QueryMinIbcFeeResponse{}

	err := suite.queryCustom(ctx, contractAddress, query, &resp)
	suite.Require().NoError(err)
	suite.Require().Equal(
		feerefundertypes.Fee{
			RecvFee: sdk.Coins{},
			AckFee: sdk.Coins{
				sdk.Coin{Denom: "ufury", Amount: sdk.NewInt(1000)},
			},
			TimeoutFee: sdk.Coins{
				sdk.Coin{Denom: "ufury", Amount: sdk.NewInt(1000)},
			},
		},
		resp.MinFee,
	)
}

func (suite *CustomQuerierTestSuite) TestFullDenom() {
	var (
		ctx   = suite.ChainA.GetContext()
		owner = keeper.RandomAccountAddress(suite.T()) // We don't care what this address is
	)

	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(ctx, owner, "../testdata/reflect.wasm")
	contractAddress := suite.InstantiateReflectContract(ctx, owner, codeID)
	suite.Require().NotEmpty(contractAddress)

	query := bindings.FuryaQuery{
		FullDenom: &bindings.FullDenom{
			CreatorAddr: contractAddress.String(),
			Subdenom:    "test",
		},
	}
	resp := bindings.FullDenomResponse{}
	err := suite.queryCustom(ctx, contractAddress, query, &resp)
	suite.Require().NoError(err)

	expected := fmt.Sprintf("factory/%s/test", contractAddress.String())
	suite.Require().Equal(expected, resp.Denom)
}

func (suite *CustomQuerierTestSuite) TestDenomAdmin() {
	var (
		furya = suite.GetFuryaZoneApp(suite.ChainA)
		ctx     = suite.ChainA.GetContext()
		owner   = keeper.RandomAccountAddress(suite.T()) // We don't care what this address is
	)

	furya.TokenFactoryKeeper.SetParams(ctx, tokenfactorytypes.NewParams(
		sdk.NewCoins(sdk.NewInt64Coin(tokenfactorytypes.DefaultFuryaDenom, 10_000_000)),
		FeeCollectorAddress,
	))

	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(ctx, owner, "../testdata/reflect.wasm")
	contractAddress := suite.InstantiateReflectContract(ctx, owner, codeID)
	suite.Require().NotEmpty(contractAddress)

	senderAddress := suite.ChainA.SenderAccounts[0].SenderAccount.GetAddress()
	coinsAmnt := sdk.NewCoins(sdk.NewCoin(params.DefaultDenom, sdk.NewInt(int64(10_000_000))))
	bankKeeper := furya.BankKeeper
	err := bankKeeper.SendCoins(ctx, senderAddress, contractAddress, coinsAmnt)
	suite.NoError(err)

	denom, _ := furya.TokenFactoryKeeper.CreateDenom(ctx, contractAddress.String(), "test")

	query := bindings.FuryaQuery{
		DenomAdmin: &bindings.DenomAdmin{
			Subdenom: denom,
		},
	}
	resp := bindings.DenomAdminResponse{}
	err = suite.queryCustom(ctx, contractAddress, query, &resp)
	suite.Require().NoError(err)

	suite.Require().Equal(contractAddress.String(), resp.Admin)
}

type ChainRequest struct {
	Reflect wasmvmtypes.QueryRequest `json:"reflect"`
}

type ChainResponse struct {
	Data []byte `json:"data"`
}

func (suite *CustomQuerierTestSuite) queryCustom(ctx sdk.Context, contract sdk.AccAddress, request, response interface{}) error {
	msgBz, err := json.Marshal(request)
	suite.Require().NoError(err)

	query := ChainRequest{
		Reflect: wasmvmtypes.QueryRequest{Custom: msgBz},
	}

	queryBz, err := json.Marshal(query)
	if err != nil {
		return err
	}

	resBz, err := suite.GetFuryaZoneApp(suite.ChainA).WasmKeeper.QuerySmart(ctx, contract, queryBz)
	if err != nil {
		return err
	}

	var resp ChainResponse
	err = json.Unmarshal(resBz, &resp)
	if err != nil {
		return err
	}

	return json.Unmarshal(resp.Data, response)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(CustomQuerierTestSuite))
}
