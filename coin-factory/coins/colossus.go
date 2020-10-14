package coins

import (
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
)

var Colossus = Coin{
	Info: CoinInfo{
		Icon:        "data:image/png;base64,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",
		Tag:         "COLX",
		Name:        "ColossusXT (COLX)",
		Trezor:      false,
		Ledger:      false,
		Segwit:      false,
		Masternodes: true,
		Token:       false,
		StableCoin:  false,
		Blockbook:   "https://colx.polispay.com",
		Protocol:    "colx",
		TxVersion:   1,
		TxBuilder:   "bitcoinjs",
		HDIndex:     1999,
		Networks: map[string]CoinNetworkInfo{
			"P2PKH": {
				MessagePrefix: "\x18Colx Signed Message:\n",
				Bech32:        "",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x488b21e,
					Private: 0x488ade4,
				},
				PubKeyHash: []int{0x1e},
				ScriptHash: []int{0xd},
				Wif:        []int{0xd4},
			},
		},
	},
	Rates: RatesSource{
		Exchange:         "crex24",
		FallBackExchange: "",
		CoinGeckoId:      "colossuscoinxt",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "colx.polispay.com",
		BlockTime:        1,
		MinConfirmations: 1,
	},
	NetParams: &chaincfg.Params{
		AddressMagicLen:   1,
		Bech32HRPSegwit:   "",
		PubKeyHashAddrID:  []byte{0x1e},
		ScriptHashAddrID:  []byte{0xd},
		PrivateKeyID:      []byte{0xd4},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        1999,
		Base58CksumHasher: base58.Sha256D,
	},
}
