package coins

var Colossus = Coin{
	Tag:  "COLX",
	Name: "colossus",
	Rates: RatesSource{
		Exchange:         "cryptobridge",
		FallBackExchange: "novaexchange",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "colx.polispay.com",
		BlockTime:        1,
		MinConfirmations: 30,
	},
	Token:         false,
	BlockExplorer: "https://colx.polispay.com",
}
