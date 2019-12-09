package coins

var Groestlcoin = Coin{
	Tag:  "GRS",
	Name: "groestlcoin",
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "bittrex",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "grs.polispay.com",
		BlockTime:        10,
		MinConfirmations: 2,
	},
	Token:         false,
	BlockExplorer: "https://grs.polispay.com",
}
