package coins

import (
	"github.com/eabz/btcutil/base58"
	"github.com/eabz/btcutil/chaincfg"
)

// Bitcoin coinfactory information
var Bitcoin = Coin{
	Info: CoinInfo{
		Icon:        "data:image/png;base64,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",
		Tag:         "BTC",
		Name:        "Bitcoin (BTC)",
		Trezor:      true,
		Ledger:      true,
		Segwit:      true,
		Masternodes: false,
		Token:       false,
		StableCoin:  false,
		Blockbook:   "https://btc2.trezor.io",
		Protocol:    "bitcoin",
		TxVersion:   1,
		TxBuilder:   "bitcoinjs",
		HDIndex:     0,
		Networks: map[string]CoinNetworkInfo{
			"P2SHInP2WPKH": {
				MessagePrefix: "\x18Bitcoin Signed Message:\n",
				Bech32:        "bc",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x049d7cb2,
					Private: 0x049d7878,
				},
				PubKeyHash: 0x00,
				ScriptHash: 0x05,
				Wif:        0x80,
			},
			"P2WPKH": {
				MessagePrefix: "\x18Bitcoin Signed Message:\n",
				Bech32:        "bc",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x04b24746,
					Private: 0x04b2430c,
				},
				PubKeyHash: 0x00,
				ScriptHash: 0x05,
				Wif:        0x80,
			},
		},
	},
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "bitso",
	},
	BlockchainInfo: BlockchainInfo{
		BlockTime:        10,
		MinConfirmations: 1,
		ExternalSource:   "btc2.trezor.io",
	},
	NetParams: &chaincfg.Params{
		AddressMagicLen:   1,
		Bech32HRPSegwit:   "bc",
		PubKeyHashAddrID:  []byte{0x00},
		ScriptHashAddrID:  []byte{0x05},
		PrivateKeyID:      []byte{0x80},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        0,
		Base58CksumHasher: base58.Sha256D,
	},
}
