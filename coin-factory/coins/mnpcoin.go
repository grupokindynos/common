package coins

var MNPCoin = Coin{
	Tag:  "MNP",
	Name: "mnpcoin",
	Rates: RatesSource{
		Exchange:         "crex24",
		FallBackExchange: "stex",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "",
		BlockTime:        1,
		MinConfirmations: 6,
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
	BlockExplorer: "https://mnp.polispay.com",
}
