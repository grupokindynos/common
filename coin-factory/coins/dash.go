package coins

import (
	"github.com/eabz/blockbook/bchain/coins/dash"
)

var Dash = Coin{
	Info: CoinInfo{
		Icon:        "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",
		Tag:         "DASH",
		Name:        "Dash (DASH)",
		Trezor:      true,
		Ledger:      true,
		Segwit:      false,
		Masternodes: true,
		Token:       false,
		Blockbook:   "https://dash2.trezor.io",
		Protocol:    "dash",
		TxVersion:   2,
		TxBuilder:   "bitcoinjs",
		HDIndex:     5,
		Networks: map[string]CoinNetworkInfo{
			"P2PKH": {
				MessagePrefix: "\x18Dash Signed Message:\n",
				Bech32:        "",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x488b21e,
					Private: 0x488ade4,
				},
				PubKeyHash: 76,
				ScriptHash: 16,
				Wif:        204,
			},
		},
	},
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "bittrex",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "dash2.trezor.io",
		BlockTime:        2.5,
		MinConfirmations: 2,
	},
	NetParams: &dash.MainNetParams,
}
