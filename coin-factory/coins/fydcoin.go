package coins

import (
	"github.com/eabz/btcutil/base58"
	"github.com/eabz/btcutil/chaincfg"
)

// FYDCoin coinfactory information
var FYDCoin = Coin{
	Info: CoinInfo{
		Icon:        "data:image/png;base64,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",
		Tag:         "FYD",
		Name:        "FYDCoin (FYD)",
		Trezor:      false,
		Ledger:      false,
		Segwit:      false,
		Masternodes: true,
		Token:       false,
		StableCoin:  false,
		Blockbook:   "https://fyd.polispay.com",
		Protocol:    "fyd",
		TxVersion:   1,
		TxBuilder:   "bitcoinjs",
		HDIndex:     0,
		Networks: map[string]CoinNetworkInfo{
			"P2PKH": {
				MessagePrefix: "\x18FYDCoin Signed Message:\n",
				Bech32:        "",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x488b21e,
					Private: 0x488ade4,
				},
				PubKeyHash: []int{36},
				ScriptHash: []int{28},
				Wif:        []int{212},
			},
		},
	},
	Rates: RatesSource{
		Exchange:         "crex24",
		FallBackExchange: "",
	},
	BlockchainInfo: BlockchainInfo{
		BlockTime:        1,
		MinConfirmations: 6,
		ExternalSource:   "fyd.polispay.com",
	},
	NetParams: &chaincfg.Params{
		AddressMagicLen:   1,
		Bech32HRPSegwit:   "",
		PubKeyHashAddrID:  []byte{36},
		ScriptHashAddrID:  []byte{28},
		PrivateKeyID:      []byte{212},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        0,
		Base58CksumHasher: base58.Sha256D,
	},
}
