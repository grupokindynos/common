package coins

// Tether coinfactory information
var Gather = Coin{
	Info: CoinInfo{
		Icon:         "data:image/png;base64,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",
		Tag:          "GTH",
		Name:         "Gather (GTH)",
		Trezor:       true,
		Ledger:       true,
		Segwit:       false,
		Masternodes:  false,
		Token:        true,
		StableCoin:   false,
		TokenNetwork: "ethereum",
		Contract:     "0xc3771d47e2ab5a519e2917e61e23078d0c05ed7f",
		Decimals:     18,
		Blockbook:    "https://eth2.trezor.io",
	},
	Rates: RatesSource{
		Exchange:         "bithumb",
		FallBackExchange: "",
		CoinGeckoId:      "gather",
	},
	BlockchainInfo: BlockchainInfo{
		BlockTime:        0.25,
		MinConfirmations: 3,
	},
}
