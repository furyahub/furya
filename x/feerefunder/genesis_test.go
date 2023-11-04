package feerefunder_test

import (
	"testing"

	"github.com/furyahub/furya/app/params"
	"github.com/furyahub/furya/testutil/feerefunder/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/furyahub/furya/testutil/interchainqueries/nullify"
	"github.com/furyahub/furya/x/feerefunder"
	"github.com/furyahub/furya/x/feerefunder/types"
)

const TestContractAddressFurya = "furya14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s5c2epq"

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		FeeInfos: []types.FeeInfo{{
			Payer:    TestContractAddressFurya,
			PacketId: types.NewPacketID("port", "channel-1", 64),
			Fee: types.Fee{
				RecvFee:    sdk.NewCoins(sdk.NewCoin(params.DefaultDenom, sdk.NewInt(0))),
				AckFee:     sdk.NewCoins(sdk.NewCoin(params.DefaultDenom, sdk.NewInt(types.DefaultFees.AckFee.AmountOf(params.DefaultDenom).Int64()+1))),
				TimeoutFee: sdk.NewCoins(sdk.NewCoin(params.DefaultDenom, sdk.NewInt(types.DefaultFees.TimeoutFee.AmountOf(params.DefaultDenom).Int64()+1))),
			},
		}},
	}

	require.EqualValues(t, genesisState.Params, types.DefaultParams())

	k, ctx := keeper.FeeKeeper(t, nil, nil)
	feerefunder.InitGenesis(ctx, *k, genesisState)
	got := feerefunder.ExportGenesis(ctx, *k)

	require.EqualValues(t, got.Params, types.DefaultParams())
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
