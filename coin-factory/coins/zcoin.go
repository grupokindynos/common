package coins

var Zcoin = Coin{
	Tag:  "XZC",
	Name: "zcoin",
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "bittrex",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "xzc.polispay.com",
		BlockTime:        5,
		MinConfirmations: 5,
	},
	Token:         false,
	BlockExplorer: "https://xzc.polispay.com",
}
