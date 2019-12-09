package coins

var Snowgem = Coin{
	Tag:  "XSG",
	Name: "snowgem",
	Rates: RatesSource{
		Exchange:         "stex",
		FallBackExchange: "cryptobridge",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "xsg.polispay.com",
		BlockTime:        1,
		MinConfirmations: 12,
	},
	Token:         false,
	BlockExplorer: "https://xsg.polispay.com",
}
