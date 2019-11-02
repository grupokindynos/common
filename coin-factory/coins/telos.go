package coins

var Telos = Coin{
	Tag:  "TELOS",
	Name: "teloscoin",
	Rates: RatesSource{
		Exchange:         "graviex",
		FallBackExchange: "southxchange",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "telos.polispay.com",
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
	Token:         false,
	BlockExplorer: "https://telos.polispay.com",
}
