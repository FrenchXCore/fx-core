package types

// Evm module events
const (
	EventTypeEthereumTx = TypeMsgEthereumTx
	EventTypeBlockBloom = "block_bloom"
	EventTypeTxLog      = "tx_log"
	EventContractCode   = "contract_code"

	AttributeKeyRecipient        = "recipient"
	AttributeKeyTxHash           = "txHash"
	AttributeKeyEthereumTxHash   = "ethereumTxHash"
	AttributeKeyTxIndex          = "txIndex"
	AttributeKeyTxGasUsed        = "txGasUsed"
	AttributeKeyTxType           = "txType"
	AttributeKeyTxLog            = "txLog"
	AttributeKeyEthereumTxFailed = "ethereumTxFailed"
	AttributeValueCategory       = ModuleName
	AttributeKeyEthereumBloom    = "bloom"
	AttributeKeyContract         = "contract"
	AttributeKeyCodeHash         = "code_hash"
)
