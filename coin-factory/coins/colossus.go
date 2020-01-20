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
				PubKeyHash: 0x1e,
				ScriptHash: 0xd,
				Wif:        0xd4,
			},
		},
	},
	Rates: RatesSource{
		Exchange:         "crex24",
		FallBackExchange: "",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "colx.polispay.com",
		BlockTime:        1,
		MinConfirmations: 30,
	},
	NetParams: &chaincfg.Params{
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
