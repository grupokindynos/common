package coins

var Dash = Coin{
	Tag:  "DASH",
	Name: "dash",
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "bittrex",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "dash2.trezor.io",
		BlockTime:        2.5,
		MinConfirmations: 2,
	},
	Token:         false,
	BlockExplorer: "https://dash2.trezor.io",
}
