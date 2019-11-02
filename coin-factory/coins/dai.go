package coins

// Dai coinfactory information
var Dai = Coin{
	Tag:  "DAI",
	Name: "dai",
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "",
	},
	TokenNetwork: "ethereum",
	Token:        true,
	Contract:     "0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359",
}
