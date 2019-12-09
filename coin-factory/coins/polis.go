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
		MinConfirmations: 2,
	},
	Token:         false,
	BlockExplorer: "https://blockbook.polispay.org",
}
