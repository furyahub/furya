package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/furyahub/furya/app"
	"github.com/furyahub/furya/x/tokenfactory/types"
)

func TestDecomposeDenoms(t *testing.T) {
	app.GetDefaultConfig()
	for _, tc := range []struct {
		desc  string
		denom string
		valid bool
	}{
		{
			desc:  "empty is invalid",
			denom: "",
			valid: false,
		},
		{
			desc:  "normal",
			denom: "factory/furya1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2/bitcoin",
			valid: true,
		},
		{
			desc:  "multiple slashes in subdenom",
			denom: "factory/furya1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2/bitcoin/1",
			valid: true,
		},
		{
			desc:  "no subdenom",
			denom: "factory/furya1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2/",
			valid: true,
		},
		{
			desc:  "incorrect prefix",
			denom: "ibc/furya1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2/bitcoin",
			valid: false,
		},
		{
			desc:  "subdenom of only slashes",
			denom: "factory/furya1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2/////",
			valid: true,
		},
		{
			desc:  "too long name",
			denom: "factory/furya1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2/adsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsf",
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			_, _, err := types.DeconstructDenom(tc.denom)
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
