package coins

var Polis = Coin{
	Tag:  "POLIS",
	Name: "polis",
	Rates: RatesSource{
		Exchange:         "southxchange",
		FallBackExchange: "cryptobridge",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "blockbook.polispay.org",
		BlockTime:        2,
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
	BlockExplorer: "https://blockbook.polispay.org",
}
