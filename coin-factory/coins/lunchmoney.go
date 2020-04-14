package coins

// LunchMoney coinfactory information
var LunchMoney = Coin{
	Info: CoinInfo{
		Icon:         "",
		Tag:          "LMY",
		Name:         "LunchMoney (LMY)",
		Trezor:       false,
		Ledger:       false,
		Segwit:       false,
		Masternodes:  false,
		Token:        true,
		StableCoin:   false,
		TokenNetwork: "ethereum",
		Contract:     "0x66fd97a78d8854fec445cd1c80a07896b0b4851f",
		Decimals:     18,
	},
	Rates: RatesSource{
		Exchange:         "",
		FallBackExchange: "",
	},
}
