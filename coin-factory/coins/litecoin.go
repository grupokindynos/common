package coins

var Litecoin = Coin{
	Tag:  "LTC",
	Name: "litecoin",
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "bittrex",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "ltc2.trezor.io",
		BlockTime:        2.5,
		MinConfirmations: 20,
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
	Token:         false,
	BlockExplorer: "https://ltc2.trezor.io",
}
