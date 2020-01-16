package coins

import (
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
)

var DeepOnion = Coin{
	Info: CoinInfo{
		Icon:        "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",
		Tag:         "ONION",
		Name:        "DeepOnion (ONION)",
		Trezor:      false,
		Ledger:      false,
		Segwit:      false,
		Masternodes: false,
		Token:       false,
		Blockbook:   "https://onion.polispay.com",
		Protocol:    "onion",
		TxVersion:   2,
		TxBuilder:   "bitcoinjs",
		HDIndex:     305,
		Networks: map[string]CoinNetworkInfo{
			"P2PKH": {
				MessagePrefix: "\x18DeepOnion Signed Message:\n",
				Bech32:        "",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x03e25d7e,
					Private: 0x03e25945,
				},
				PubKeyHash: 0x1f,
				ScriptHash: 0x4e,
				Wif:        0x9f,
			},
		},
	},
	Rates: RatesSource{
		Exchange:         "kucoin",
		FallBackExchange: "crex24",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "",
		BlockTime:        4,
		MinConfirmations: 20,
	},
	NetParams: &chaincfg.Params{
		Bech32HRPSegwit:   "",
		ScriptHashAddrID:  []byte{0x4e},
		PubKeyHashAddrID:  []byte{48},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        305,
		Base58CksumHasher: base58.Sha256D,
		Net:               9, // Make sure doesn't collide with any other coin.
		AddressMagicLen:   1,
	},
}
