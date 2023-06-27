package types_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/functionx/fx-core/v5/x/evm/types"
)

func TestStorageValidate(t *testing.T) {
	testCases := []struct {
		name    string
		storage types.Storage
		expPass bool
	}{
		{
			"valid storage",
			types.Storage{
				types.NewState(common.BytesToHash([]byte{1, 2, 3}), common.BytesToHash([]byte{1, 2, 3})),
			},
			true,
		},
		{
			"empty storage key bytes",
			types.Storage{
				{Key: ""},
			},
			false,
		},
		{
			"duplicated storage key",
			types.Storage{
				{Key: common.BytesToHash([]byte{1, 2, 3}).String()},
				{Key: common.BytesToHash([]byte{1, 2, 3}).String()},
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		err := tc.storage.Validate()
		if tc.expPass {
			require.NoError(t, err, tc.name)
		} else {
			require.Error(t, err, tc.name)
		}
	}
}

func TestStorageCopy(t *testing.T) {
	testCases := []struct {
		name    string
		storage types.Storage
	}{
		{
			"single storage",
			types.Storage{
				types.NewState(common.BytesToHash([]byte{1, 2, 3}), common.BytesToHash([]byte{1, 2, 3})),
			},
		},
		{
			"empty storage key value bytes",
			types.Storage{
				{Key: common.Hash{}.String(), Value: common.Hash{}.String()},
			},
		},
		{
			"empty storage",
			types.Storage{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		require.Equal(t, tc.storage, tc.storage.Copy(), tc.name)
	}
}

func TestStorageString(t *testing.T) {
	storage := types.Storage{types.NewState(common.BytesToHash([]byte("key")), common.BytesToHash([]byte("value")))}
	str := "key:\"0x00000000000000000000000000000000000000000000000000000000006b6579\" value:\"0x00000000000000000000000000000000000000000000000000000076616c7565\" \n"
	require.Equal(t, str, storage.String())
}
