package coins

// Chainlink coinfactory information
var Chainlink = Coin{
	Info: CoinInfo{
		Icon:         "",
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
	},
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "",
	},
}
