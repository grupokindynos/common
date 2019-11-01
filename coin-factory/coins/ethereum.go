package coins

// Ethereum coinfactory information
var Ethereum = Coin{
	Tag:  "ETH",
	Name: "etherem",
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "",
	},
	BlockchainInfo: BlockchainInfo{
		BlockTime:        1,
		MinConfirmations: 8,
		ExternalSource:   "eth1.trezor.io",
	},
	RpcMethods: RPCMethods{
		GetWalletInfo:              "",
		GetBlockchainInfo:          "",
		GetNetworkInfo:             "",
		GetNewAddress:              "",
		SendToAddress:              "",
		ValidateAddress:            "",
		GetRawTransaction:          "",
		DecodeRawTransaction:       "",
		GetRawTransactionVerbosity: "",
	},
	BlockExplorer: "https://eth1.trezor.io",
}
