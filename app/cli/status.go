package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/tendermint/tendermint/libs/bytes"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/p2p"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

// ValidatorInfo is info about the node's validator, same as Tendermint,
// except that we use our own PubKey.
type validatorInfo struct {
	Address     bytes.HexBytes
	PubKey      cryptotypes.PubKey
	VotingPower int64
}

// ResultStatus is node's info, same as Tendermint, except that we use our own
// PubKey.
type resultStatus struct {
	NodeInfo      p2p.DefaultNodeInfo
	SyncInfo      ctypes.SyncInfo
	ValidatorInfo validatorInfo
}

// StatusCommand returns the command to return the status of the network.
func StatusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Query remote node for status",
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			status, err := getNodeStatus(clientCtx)
			if err != nil {
				return err
			}

			// `status` has TM pubkeys, we need to convert them to our pubkeys.
			pk, err := cryptocodec.FromTmPubKeyInterface(status.ValidatorInfo.PubKey)
			if err != nil {
				return err
			}
			statusWithPk := resultStatus{
				NodeInfo: status.NodeInfo,
				SyncInfo: status.SyncInfo,
				ValidatorInfo: validatorInfo{
					Address:     status.ValidatorInfo.Address,
					PubKey:      pk,
					VotingPower: status.ValidatorInfo.VotingPower,
				},
			}
			//nolint
			return clientCtx.PrintObjectLegacy(statusWithPk)
		},
	}

	cmd.Flags().StringP(tmcli.OutputFlag, "o", "json", "Output format (text|json)")
	cmd.Flags().String(flags.FlagNode, "tcp://localhost:26657", "<host>:<port> to Tendermint RPC interface for this chain")

	return cmd
}

func getNodeStatus(clientCtx client.Context) (*ctypes.ResultStatus, error) {
	node, err := clientCtx.GetNode()
	if err != nil {
		return &ctypes.ResultStatus{}, err
	}

	return node.Status(context.Background())
}
