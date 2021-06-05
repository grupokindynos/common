package coins

// PolisBSC coinfactory information
var Cake = Coin{
	Info: CoinInfo{
		Icon: "data:image/png;base64,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",

		Tag:          "CAKE",
		Name:         "Cake (CAKE)",
		Trezor:       true,
		Ledger:       true,
		Segwit:       false,
		Masternodes:  false,
		Token:        true,
		StableCoin:   false,
		TokenNetwork: "binance",
		Contract:     "0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82",
		Decimals:     18,
		Blockbook:    "https://bscscan.com",
	},
	Rates: RatesSource{
		Exchange:         "pancake_swap",
		FallBackExchange: "",
		CoinGeckoId:      "cake",
	},
}
