package coins

// TrueUSD coinfactory information
var TrueUSD = Coin{
	Tag:  "TUSD",
	Name: "trueusd",
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "",
	},
	TokenNetwork: "ethereum",
	Token:        true,
	Contract:     "0x0000000000085d4780B73119b644AE5ecd22b376",
}
