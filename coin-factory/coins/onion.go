package coins

var DeepOnion = Coin{
	Tag:  "ONION",
	Name: "deeponion",
	Rates: RatesSource{
		Exchange:         "kucoin",
		FallBackExchange: "crex24",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "",
		BlockTime:        4,
		MinConfirmations: 20,
	},
	Token:         false,
	BlockExplorer: "https://onion.polispay.com",
}
