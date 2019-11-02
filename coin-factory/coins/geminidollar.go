package coins

// Geminidollar coinfactory information
var Geminidollar = Coin{
	Tag:  "GUSD",
	Name: "geminidollar",
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "",
	},
	TokenNetwork: "ethereum",
	Token:        true,
	Contract:     "0x056fd409e1d7a124bd7017459dfea2f387b6d5cd",
}
