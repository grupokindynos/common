package coins

import (
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
)

var Litecoin = Coin{
	Info: CoinInfo{
		Icon:        "data:image/png;base64,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",
		Tag:         "LTC",
		Name:        "Litecoin (LTC)",
		Trezor:      true,
		Ledger:      true,
		Segwit:      true,
		Masternodes: false,
		Token:       false,
		StableCoin:  false,
		Blockbook:   "https://ltc1.trezor.io",
		Protocol:    "ltc",
		TxVersion:   1,
		TxBuilder:   "bitcoinjs",
		HDIndex:     2,
		Networks: map[string]CoinNetworkInfo{
			"P2SHInP2WPKH": {
				MessagePrefix: "\x18Litecoin Signed Message:\n",
				Bech32:        "ltc",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x1b26ef6,
					Private: 0x19d9cfe,
				},
				PubKeyHash: []int{48},
				ScriptHash: []int{50},
				Wif:        []int{176},
			},
			"P2WPKH": {
				MessagePrefix: "\x18Litecoin Signed Message:\n",
				Bech32:        "ltc",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x4b24746,
					Private: 0x19d9cfe,
				},
				PubKeyHash: []int{48},
				ScriptHash: []int{50},
				Wif:        []int{176},
			},
		},
	},
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "southxchange",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "ltc2.trezor.io",
		BlockTime:        2.5,
		MinConfirmations: 1,
	},
	NetParams: &chaincfg.Params{
		AddressMagicLen:   1,
		Bech32HRPSegwit:   "ltc",
		ScriptHashAddrID:  []byte{50},
		PubKeyHashAddrID:  []byte{48},
		PrivateKeyID:      []byte{176},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        2,
		Base58CksumHasher: base58.Sha256D,
	},
}
