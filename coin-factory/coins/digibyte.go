package coins

var Digibyte = Coin{
	Tag:  "DGB",
	Name: "digibyte",
	Rates: RatesSource{
		Exchange:         "bittrex",
		FallBackExchange: "",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "dgb2.trezor.io",
		BlockTime:        0.25,
		MinConfirmations: 25,
	},
	RpcMethods: RPCMethods{
		GetWalletInfo:              "getwalletinfo",
		GetBlockchainInfo:          "getblockchaininfo",
		GetNetworkInfo:             "getnetworkinfo",
		GetNewAddress:              "getnewaddress",
		SendToAddress:              "sendtoaddress",
		ValidateAddress:            "getaddressinfo",
		GetRawTransaction:          "getrawtransaction",
		GetRawTransactionVerbosity: "1",
	},
}
