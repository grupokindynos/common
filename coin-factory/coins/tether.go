package coins

// Tether coinfactory information
var Tether = Coin{
	Info: CoinInfo{
		Icon:         "data:image/png;base64,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",
		Tag:          "USDT",
		Name:         "Tether (USDT)",
		Trezor:       true,
		Ledger:       true,
		Segwit:       false,
		Masternodes:  false,
		Token:        true,
		StableCoin:   true,
		TokenNetwork: "ethereum",
		Contract:     "0xdac17f958d2ee523a2206206994597c13d831ec7",
		Decimals:     6,
	},
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "",
	},
}
