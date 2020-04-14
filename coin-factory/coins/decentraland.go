package coins

// Decentraland coinfactory information
var Decentraland = Coin{
	Info: CoinInfo{
		Icon:         "",
		Tag:          "MANA",
		Name:         "Decentraland (MANA)",
		Trezor:       false,
		Ledger:       false,
		Segwit:       false,
		Masternodes:  false,
		Token:        true,
		StableCoin:   false,
		TokenNetwork: "ethereum",
		Contract:     "0x0f5d2fb29fb7d3cfee444a200298f468908cc942",
		Decimals:     18,
	},
	Rates: RatesSource{
		Exchange:         "",
		FallBackExchange: "",
	},
}
