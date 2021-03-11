package coins

// PolisBSC coinfactory information
var PolisBSC = Coin{
	Info: CoinInfo{
		Icon: "data:image/png;base64,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",

		Tag:          "POLISBSC",
		Name:         "Polis (POLISBSC)",
		Trezor:       true,
		Ledger:       true,
		Segwit:       false,
		Masternodes:  false,
		Token:        true,
		StableCoin:   false,
		TokenNetwork: "binance",
		Contract:     "0x66fd97a78d8854fec445cd1c80a07896b0b4851f",
		Decimals:     18,
		Blockbook:    "https://bscscan.com",
	},
	Rates: RatesSource{
		Exchange:         "pancake_swap",
		FallBackExchange: "southxchange",
		CoinGeckoId:      "polis",
	},
}
