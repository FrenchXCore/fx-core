package cli

import (
	"bufio"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/input"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/params"
	rpctypes "github.com/functionx/fx-core/rpc/ethereum/types"
	fxcoretypes "github.com/functionx/fx-core/types"
	feemarkettypes "github.com/functionx/fx-core/x/feemarket/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"

	"github.com/functionx/fx-core/x/evm/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	cmd.AddCommand(NewRawTxCmd())
	return cmd
}

// NewRawTxCmd command build cosmos transaction from raw ethereum transaction
func NewRawTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "raw [tx-hex]",
		Short: "Build cosmos transaction from raw ethereum transaction",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			data, err := hexutil.Decode(args[0])
			if err != nil {
				return errors.Wrap(err, "failed to decode ethereum tx hex bytes")
			}

			msg := &types.MsgEthereumTx{}
			if err := msg.UnmarshalBinary(data); err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			rsp, err := rpctypes.NewQueryClient(clientCtx).Params(cmd.Context(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			tx, err := msg.BuildTx(clientCtx.TxConfig.NewTxBuilder(), rsp.Params.EvmDenom)
			if err != nil {
				return err
			}

			if clientCtx.GenerateOnly {
				json, err := clientCtx.TxConfig.TxJSONEncoder()(tx)
				if err != nil {
					return err
				}

				return clientCtx.PrintString(fmt.Sprintf("%s\n", json))
			}

			if !clientCtx.SkipConfirm {
				out, err := clientCtx.TxConfig.TxJSONEncoder()(tx)
				if err != nil {
					return err
				}

				_, _ = fmt.Fprintf(os.Stderr, "%s\n\n", out)

				buf := bufio.NewReader(os.Stdin)
				ok, err := input.GetConfirmation("confirm transaction before signing and broadcasting", buf, os.Stderr)

				if err != nil || !ok {
					_, _ = fmt.Fprintf(os.Stderr, "%s\n", "canceled transaction")
					return err
				}
			}

			txBytes, err := clientCtx.TxConfig.TxEncoder()(tx)
			if err != nil {
				return err
			}

			// broadcast to a Tendermint node
			res, err := clientCtx.BroadcastTx(txBytes)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func getEvmParamsByFlags(cmd *cobra.Command) (*types.Params, error) {
	evmParamsEvmDenom, err := cmd.Flags().GetString(flagEvmParamsEvmDenom)
	if err != nil {
		return nil, err
	}
	homesteadBlock := sdk.ZeroInt()
	daoForkBlock := sdk.ZeroInt()
	eip150Block := sdk.ZeroInt()
	eip155Block := sdk.ZeroInt()
	eip158Block := sdk.ZeroInt()
	byzantiumBlock := sdk.ZeroInt()
	constantinopleBlock := sdk.ZeroInt()
	petersburgBlock := sdk.ZeroInt()
	istanbulBlock := sdk.ZeroInt()
	muirGlacierBlock := sdk.ZeroInt()
	berlinBlock := sdk.ZeroInt()
	londonBlock := sdk.ZeroInt()
	arrowGlacierBlock := sdk.ZeroInt()
	mergeForkBlock := sdk.ZeroInt()

	return &types.Params{
		EvmDenom:     evmParamsEvmDenom,
		EnableCreate: true,
		EnableCall:   true,
		ExtraEIPs:    nil,
		ChainConfig: types.ChainConfig{
			HomesteadBlock:      &homesteadBlock,
			DAOForkBlock:        &daoForkBlock,
			DAOForkSupport:      true,
			EIP150Block:         &eip150Block,
			EIP150Hash:          common.Hash{}.String(),
			EIP155Block:         &eip155Block,
			EIP158Block:         &eip158Block,
			ByzantiumBlock:      &byzantiumBlock,
			ConstantinopleBlock: &constantinopleBlock,
			PetersburgBlock:     &petersburgBlock,
			IstanbulBlock:       &istanbulBlock,
			MuirGlacierBlock:    &muirGlacierBlock,
			BerlinBlock:         &berlinBlock,
			LondonBlock:         &londonBlock,
			ArrowGlacierBlock:   &arrowGlacierBlock,
			MergeForkBlock:      &mergeForkBlock,
		},
	}, nil
}

func getFeeMarkerParamsByFlags(cmd *cobra.Command, noBaseFee bool, baseFee, minBaseFee, maxBaseFee, maxGas int64) (*feemarkettypes.Params, error) {
	var BaseFeeChangeDenominator uint32 = params.BaseFeeChangeDenominator
	var ElasticityMultiplier uint32 = params.ElasticityMultiplier
	var BaseFee = baseFee * params.InitialBaseFee //4000 gWei
	var EnableHeight = fxcoretypes.EvmSupportBlock()
	var MinBaseFee = sdk.NewInt(minBaseFee * params.InitialBaseFee)
	var MaxBaseFee = sdk.NewInt(maxBaseFee * params.InitialBaseFee)
	var MaxGas = sdk.NewInt(maxGas)
	return &feemarkettypes.Params{
		NoBaseFee:                noBaseFee,
		BaseFeeChangeDenominator: BaseFeeChangeDenominator,
		ElasticityMultiplier:     ElasticityMultiplier,
		BaseFee:                  sdk.NewInt(BaseFee),
		EnableHeight:             EnableHeight,
		MinBaseFee:               MinBaseFee,
		MaxBaseFee:               MaxBaseFee,
		MaxGas:                   MaxGas,
	}, nil
}

//func getErc20ParamsByFlags(cmd *cobra.Command) (*types.Erc20Params, error) {
//	var enableErc20 = true
//	var enableEVMHook = true
//	var ibcTransferTimeoutHeight = uint64(20000)
//	return &types.Erc20Params{
//		EnableErc20:              enableErc20,
//		EnableEVMHook:            enableEVMHook,
//		IbcTransferTimeoutHeight: ibcTransferTimeoutHeight,
//	}, nil
//}
//
//func InitEvmProposalCmd() *cobra.Command {
//	cmd := &cobra.Command{
//		Use:     "init-evm",
//		Short:   "Submit a init evm proposal",
//		Example: fmt.Sprintf(`$ %s tx gov submit-proposal init-evm --evm-denom=<denom> --metadata=<path/to/metadata> --from=<key_or_address>`, version.AppName),
//		RunE: func(cmd *cobra.Command, args []string) error {
//			cliCtx, err := client.GetClientTxContext(cmd)
//			if err != nil {
//				return err
//			}
//			initProposalAmount, err := sdk.ParseCoinsNormalized(viper.GetString(cli.FlagDeposit))
//			if err != nil {
//				return err
//			}
//			title, err := cmd.Flags().GetString(cli.FlagTitle)
//			if err != nil {
//				return err
//			}
//			description, err := cmd.Flags().GetString(cli.FlagDescription)
//			if err != nil {
//				return err
//			}
//
//			evmParams, err := getEvmParamsByFlags(cmd)
//			if err != nil {
//				return err
//			}
//			feeMarketParams, err := getFeeMarkerParamsByFlags(cmd, viper.GetBool(flagNoBaseFee),
//				viper.GetInt64(flagBaseFee), viper.GetInt64(flagMinBaseFee),
//				viper.GetInt64(flagMaxBaseFee), viper.GetInt64(flagMaxGas))
//			if err != nil {
//				return err
//			}
//
//			erc20Params, err := getErc20ParamsByFlags(cmd)
//			if err != nil {
//				return err
//			}
//
//			metadataPath := viper.GetString(flagMetadata)
//			var metadatas []banktypes.Metadata
//			if len(strings.TrimSpace(metadataPath)) > 0 {
//				metadatas, err = ReadMetadataFromPath(cliCtx.Codec, metadataPath)
//				if err != nil {
//					return err
//				}
//			}
//
//			proposal := &types.InitEvmProposal{
//				Title:           title,
//				Description:     description,
//				EvmParams:       evmParams,
//				FeemarketParams: feeMarketParams,
//				Erc20Params:     erc20Params,
//				Metadata:        metadatas,
//			}
//
//			fromAddress := cliCtx.GetFromAddress()
//			msg, err := govtypes.NewMsgSubmitProposal(proposal, initProposalAmount, fromAddress)
//			if err != nil {
//				return err
//			}
//			if err := msg.ValidateBasic(); err != nil {
//				return err
//			}
//			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
//		},
//	}
//
//	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
//	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
//	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")
//	cmd.Flags().String(flagEvmParamsEvmDenom, fxcoretypes.FX, "evm denom represents the token denomination used to run the EVM state transitions.")
//	cmd.Flags().String(flagMetadata, "", "path to metadata file/directory")
//	cmd.Flags().Bool(flagNoBaseFee, false, "no base fee")
//	cmd.Flags().Int64(flagBaseFee, 4000, "enable base fee(gwei)")
//	cmd.Flags().Int64(flagMinBaseFee, 4000, "min base fee(gwei)")
//	cmd.Flags().Int64(flagMaxBaseFee, 40000, "max base fee(gwei)")
//	cmd.Flags().Int64(flagMaxGas, 10000000, "max gas limit")
//
//	if err := cmd.MarkFlagRequired(cli.FlagTitle); err != nil {
//		panic(err)
//	}
//	if err := cmd.MarkFlagRequired(cli.FlagDescription); err != nil {
//		panic(err)
//	}
//	if err := cmd.MarkFlagRequired(cli.FlagDeposit); err != nil {
//		panic(err)
//	}
//	return cmd
//}

const (
	flagEvmParamsEvmDenom = "evm-denom"
	flagNoBaseFee         = "no-base-fee"
	flagBaseFee           = "base-fee"
	flagMinBaseFee        = "min-base-fee"
	flagMaxBaseFee        = "max-base-fee"
	flagMaxGas            = "max-gas"
	flagMetadata          = "metadata"
)
