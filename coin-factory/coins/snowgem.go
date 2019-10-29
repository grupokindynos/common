package coins

var Snowgem = Coin{
	Tag:  "XSG",
	Name: "snowgem",
	Rates: RatesSource{
		Exchange:         "stex",
		FallBackExchange: "cryptobridge",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "xsg.polispay.com",
		BlockTime:        1,
		MinConfirmations: 12,
	},
	RpcMethods: RPCMethods{
		GetWalletInfo:              "getwalletinfo",
		GetBlockchainInfo:          "getblockchaininfo",
		GetNetworkInfo:             "getnetworkinfo",
		GetNewAddress:              "getnewaddress",
		SendToAddress:              "sendtoaddress",
		ValidateAddress:            "validateaddress",
		GetRawTransaction:          "getrawtransaction",
		DecodeRawTransaction:       "decoderawtransaction",
		GetRawTransactionVerbosity: true,
	},
	BlockExplorer: "https://xsg.polispay.com",
}
