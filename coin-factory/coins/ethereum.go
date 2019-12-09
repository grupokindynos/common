package coins

// Ethereum coinfactory information
var Ethereum = Coin{
	Tag:  "ETH",
	Name: "ethereum",
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "",
	},
	BlockchainInfo: BlockchainInfo{
		BlockTime:        1,
		MinConfirmations: 8,
		ExternalSource:   "eth1.trezor.io",
	},
	Token:         false,
	BlockExplorer: "https://eth1.trezor.io",
}
