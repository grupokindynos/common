package coins

import (
	"github.com/eabz/btcutil/base58"
	"github.com/eabz/btcutil/chaincfg"
)

var Telos = Coin{
	Info: CoinInfo{
		Icon:        "data:image/png;base64,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",
		Tag:         "TELOS",
		Name:        "Telos (TELOS)",
		Trezor:      false,
		Ledger:      false,
		Segwit:      false,
		Masternodes: true,
		Token:       false,
		StableCoin:  false,
		Blockbook:   "https://telos.polispay.com",
		Protocol:    "telos",
		TxVersion:   2,
		TxBuilder:   "bitcoinjs",
		HDIndex:     424,
		Networks: map[string]CoinNetworkInfo{
			"P2PKH": {
				MessagePrefix: "\x18Telos Signed Message:\n",
				Bech32:        "bc",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x488b21e,
					Private: 0x488ade4,
				},
				PubKeyHash: []int{38},
				ScriptHash: []int{138},
				Wif:        []int{98},
			},
		},
	},
	Rates: RatesSource{
		Exchange:         "crex24",
		FallBackExchange: "southxchange",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "telos.polispay.com",
		BlockTime:        1,
		MinConfirmations: 1,
	},
	NetParams: &chaincfg.Params{
		AddressMagicLen:   1,
		Bech32HRPSegwit:   "",
		ScriptHashAddrID:  []byte{138},
		PubKeyHashAddrID:  []byte{38},
		PrivateKeyID:      []byte{98},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        424,
		Base58CksumHasher: base58.Sha256D,
	},
}
