package coins

import (
	"github.com/eabz/btcutil/base58"
	"github.com/eabz/btcutil/chaincfg"
)

var Idx = Coin{
	Info: CoinInfo{
		Icon:        "data:image/png;base64,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",
		Tag:         "IDX",
		Name:        "Index Chain (IDX)",
		Trezor:      false,
		Ledger:      false,
		Segwit:      false,
		Masternodes: true,
		Token:       false,
		StableCoin:  false,
		Blockbook:   "",
		Protocol:    "idx",
		TxVersion:   1,
		TxBuilder:   "bitcoinjs",
		HDIndex:     860,
		Networks: map[string]CoinNetworkInfo{
			"P2PKH": {
				MessagePrefix: "\x18IDX Signed Message:\n",
				Bech32:        "",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x488b21e,
					Private: 0x488ade4,
				},
				PubKeyHash: []int{82},
				ScriptHash: []int{7},
				Wif:        []int{210},
			},
		},
	},
	Rates: RatesSource{
		Exchange:         "",
		FallBackExchange: "",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "",
		BlockTime:        1,
		MinConfirmations: 1,
	},
	NetParams: &chaincfg.Params{
		AddressMagicLen:   1,
		Bech32HRPSegwit:   "",
		ScriptHashAddrID:  []byte{7},
		PubKeyHashAddrID:  []byte{82},
		PrivateKeyID:      []byte{210},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        860,
		Base58CksumHasher: base58.Sha256D,
	},
}
