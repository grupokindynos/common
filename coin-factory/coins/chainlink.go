package coins

// Chainlink coinfactory information
var Chainlink = Coin{
	Info: CoinInfo{
		Icon:         "data:image/png;base64,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",
		Tag:          "LINK",
		Name:         "Chainlink (LINK)",
		Trezor:       false,
		Ledger:       false,
		Segwit:       false,
		Masternodes:  false,
		Token:        true,
		StableCoin:   false,
		TokenNetwork: "ethereum",
		Contract:     "0x514910771af9ca656af840dff83e8264ecf986ca",
		Decimals:     18,
		Blockbook:    "https://eth2.trezor.io",
	},
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "",
	},
}
