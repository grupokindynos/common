package coins

import (
	"github.com/eabz/btcutil/base58"
	"github.com/eabz/btcutil/chaincfg"
)

// Ethereum coinfactory information
var Ethereum = Coin{
	Info: CoinInfo{
		Icon:         "data:image/png;base64,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",
		Tag:          "ETH",
		Name:         "Ethereum (ETH)",
		Trezor:       true,
		Ledger:       true,
		Segwit:       false,
		Masternodes:  false,
		Token:        true,
		TokenNetwork: "ethereum",
		Blockbook:    "https://eth2.trezor.io",
		Protocol:     "ethereum",
		TxVersion:    1,
		TxBuilder:    "ethereum",
		HDIndex:      60,
		Networks: map[string]CoinNetworkInfo{
			"ETHEREUM": {
				MessagePrefix: "\x18Ethereum Signed Message:\n",
				Bech32:        "",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x488b21e,
					Private: 0x488ade4,
				},
				PubKeyHash: 0x00,
				ScriptHash: 0x05,
				Wif:        0x80,
			},
		},
	},
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "",
	},
	BlockchainInfo: BlockchainInfo{
		BlockTime:        1,
		MinConfirmations: 1,
		ExternalSource:   "eth1.trezor.io",
	},
	NetParams: &chaincfg.Params{
		Bech32HRPSegwit:   "",
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        60,
		Base58CksumHasher: base58.Sha256D,
	},
}
