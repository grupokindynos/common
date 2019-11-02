package coins

// Tether coinfactory information
var Tether = Coin{
	Tag:  "USDT",
	Name: "tether",
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "",
	},
	TokenNetwork: "ethereum",
	Token:        true,
	Contract:     "0xdac17f958d2ee523a2206206994597c13d831ec7",
}
