package coins

// Bitcoin coinfactory information
var Bitcoin = Coin{
	Tag:  "BTC",
	Name: "bitcoin",
	Rates: RatesSource{
		Exchange:         "bitso",
		FallBackExchange: "",
	},
	BlockchainInfo: BlockchainInfo{
		BlockTime:        10,
		MinConfirmations: 1,
		ExternalSource:   "btc2.trezor.io",
	},
	Token:         false,
	BlockExplorer: "https://btc1.trezor.io",
}
