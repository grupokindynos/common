package coins

import (
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
)

// BinanceCoin coinfactory information
var BinanceCoin = Coin{
	Info: CoinInfo{
		Icon:         "data:image/png;base64,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",
		Tag:          "BNB",
		Name:         "Binance Coin (BNB)",
		Trezor:       true,
		Ledger:       true,
		Segwit:       false,
		Masternodes:  false,
		Token:        true,
		StableCoin:   false,
		TokenNetwork: "bsc",
		Blockbook:    "https://eth2.trezor.io",
		Protocol:     "bsc",
		TxVersion:    1,
		TxBuilder:    "bsc",
		HDIndex:      714,
		Networks: map[string]CoinNetworkInfo{
			"BINANCE": {
				MessagePrefix: "\x18Ethereum Signed Message:\n",
				Bech32:        "",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x488b21e,
					Private: 0x488ade4,
				},
				PubKeyHash: []int{0x00},
				ScriptHash: []int{0x05},
				Wif:        []int{0x80},
			},
		},
	},
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "",
		CoinGeckoId:      "binancecoin",
	},
	BlockchainInfo: BlockchainInfo{
		BlockTime:        1,
		MinConfirmations: 1,
		ExternalSource:   "eth1.trezor.io",
	},
	NetParams: &chaincfg.Params{
		AddressMagicLen:   1,
		Bech32HRPSegwit:   "",
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        714,
		Base58CksumHasher: base58.Sha256D,
	},
}
