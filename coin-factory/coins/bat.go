package coins

// BAT coinfactory information
var BAT = Coin{
	Info: CoinInfo{
		Icon:         "",
		Tag:          "BAT",
		Name:         "BAT (BAT)",
		Trezor:       false,
		Ledger:       false,
		Segwit:       false,
		Masternodes:  false,
		Token:        true,
		StableCoin:   false,
		TokenNetwork: "ethereum",
		Contract:     "0x0d8775f648430679a709e98d2b0cb6250d2887ef",
		Decimals:     18,
	},
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "",
	},
}
