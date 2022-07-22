package cli

import (
	"encoding/json"
	"fmt"

	"github.com/functionx/fx-core/types"

	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/functionx/fx-core/app"
)

func Network() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "network",
		Args:    cobra.NoArgs,
		Short:   "Show fxcored network and upgrade info",
		Example: "fxcored network",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			outputBytes, err := json.Marshal(map[string]interface{}{
				"ChainId":                                app.ChainID,
				"Network":                                types.Network(),
				"GravityPruneValsetsAndAttestationBlock": fmt.Sprintf("%d", types.GravityPruneValsetsAndAttestationBlock()),
				"GravityValsetSlashBlock":                fmt.Sprintf("%d", types.GravityValsetSlashBlock()),
				"CrossChainSupportBscBlock":              fmt.Sprintf("%d", types.CrossChainSupportBscBlock()),
				"CrossChainSupportTronBlock":             fmt.Sprintf("%d", types.CrossChainSupportTronBlock()),
				"CrossChainSupportPolygonBlock":          fmt.Sprintf("%d", types.CrossChainSupportPolygonBlock()),
			})
			if err != nil {
				return err
			}
			return PrintOutput(clientCtx, outputBytes)
		},
	}
	cmd.Flags().StringP(tmcli.OutputFlag, "o", "text", "Output format (text|json)")
	return cmd
}
