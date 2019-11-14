package coins

// Ethereum coinfactory information
var Ethereum = Coin{
	Tag:  "ETH",
	Name: "ethereum",
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
		GetWalletInfo:              "eth_getBalance",
		GetWalletAccounts: 			"eth_accounts",
		GetBlockchainInfo:          "",
		GetNetworkInfo:             "",
		GetNewAddress:              "",
		SendToAddress:              "",
		ValidateAddress:            "",
		GetRawTransaction:          "",
		DecodeRawTransaction:       "",
		GetRawTransactionVerbosity: "",
	},
	Token:         false,
	BlockExplorer: "https://eth1.trezor.io",
}
