package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	types2 "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/furyahub/furya/testutil"
	testkeeper "github.com/furyahub/furya/testutil/interchaintxs/keeper"
	mock_types "github.com/furyahub/furya/testutil/mocks/interchaintxs/types"
	"github.com/furyahub/furya/x/interchaintxs/types"
)

func TestKeeper_InterchainAccountAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	icaKeeper := mock_types.NewMockICAControllerKeeper(ctrl)
	keeper, ctx, _ := testkeeper.InterchainTxsKeeper(t, nil, nil, icaKeeper, nil, nil)
	wctx := sdk.WrapSDKContext(ctx)

	resp, err := keeper.InterchainAccountAddress(wctx, nil)
	require.ErrorIs(t, err, sdkerrors.ErrInvalidRequest)
	require.Nil(t, resp)

	resp, err = keeper.InterchainAccountAddress(wctx, &types.QueryInterchainAccountAddressRequest{
		OwnerAddress:        "nonbetch32",
		InterchainAccountId: "test1",
		ConnectionId:        "connection-0",
	})
	require.ErrorContains(t, err, "failed to create ica owner")
	require.Nil(t, resp)

	portID := fmt.Sprintf("%s%s.%s", types2.PortPrefix, testutil.TestOwnerAddress, "test1")
	icaKeeper.EXPECT().GetInterchainAccountAddress(ctx, "connection-0", portID).Return("", false)
	resp, err = keeper.InterchainAccountAddress(wctx, &types.QueryInterchainAccountAddressRequest{
		OwnerAddress:        testutil.TestOwnerAddress,
		InterchainAccountId: "test1",
		ConnectionId:        "connection-0",
	})
	require.ErrorContains(t, err, "no interchain account found for portID")
	require.Nil(t, resp)

	portID = fmt.Sprintf("%s%s.%s", types2.PortPrefix, testutil.TestOwnerAddress, "test1")
	icaKeeper.EXPECT().GetInterchainAccountAddress(ctx, "connection-0", portID).Return("furya1interchainaccountaddress", true)
	resp, err = keeper.InterchainAccountAddress(wctx, &types.QueryInterchainAccountAddressRequest{
		OwnerAddress:        testutil.TestOwnerAddress,
		InterchainAccountId: "test1",
		ConnectionId:        "connection-0",
	})
	require.NoError(t, err)
	require.Equal(t, &types.QueryInterchainAccountAddressResponse{InterchainAccountAddress: "furya1interchainaccountaddress"}, resp)
}
