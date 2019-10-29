package coins

var Colossus = Coin{
	Tag:  "COLX",
	Name: "colossus",
	Rates: RatesSource{
		Exchange:         "cryptobridge",
		FallBackExchange: "novaexchange",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "colx.polispay.com",
		BlockTime:        1,
		MinConfirmations: 30,
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
	BlockExplorer: "https://colx.polispay.com",
}
