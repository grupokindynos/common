package coins

var Telos = Coin{
	Tag:  "TELOS",
	Name: "teloscoin",
	Rates: RatesSource{
		Exchange:         "graviex",
		FallBackExchange: "southxchange",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "telos.polispay.com",
		BlockTime:        1,
		MinConfirmations: 6,
	},
	Token:         false,
	BlockExplorer: "https://telos.polispay.com",
}
