package coins

// USDCoin coinfactory information
var USDCoin = Coin{
	Tag:  "USDC",
	Name: "usdcoin",
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "",
	},
	TokenNetwork: "ethereum",
	Token:        true,
	Contract:     "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
}
