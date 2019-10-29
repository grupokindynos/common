package coins

// Bitcoin coinfactory information
var Bitcoin = Coin{
	Tag:  "BTC",
	Name: "bitcoin",
	Rates: RatesSource{
		Exchange:         "bitso",
		FallBackExchange: "",
	},
	BlockchainInfo: BlockchainInfo{
		BlockTime:        10,
		MinConfirmations: 1,
		ExternalSource:   "btc2.trezor.io",
	},
	RpcMethods: RPCMethods{
		GetWalletInfo:              "getwalletinfo",
		GetBlockchainInfo:          "getblockchaininfo",
		GetNetworkInfo:             "getnetworkinfo",
		GetNewAddress:              "getnewaddress",
		SendToAddress:              "sendtoaddress",
		ValidateAddress:            "getaddressinfo",
		GetRawTransaction:          "getrawtransaction",
		DecodeRawTransaction:       "decoderawtransaction",
		GetRawTransactionVerbosity: "1",
	},
	BlockExplorer: "https://btc1.trezor.io",
}
