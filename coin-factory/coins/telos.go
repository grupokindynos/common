package coins

import (
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
)

var Telos = Coin{
	Info: CoinInfo{
		Icon:        "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",
		Tag:         "TELOS",
		Name:        "Telos (TELOS)",
		Trezor:      false,
		Ledger:      false,
		Segwit:      false,
		Masternodes: true,
		Token:       false,
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
				PubKeyHash: 38,
				ScriptHash: 138,
				Wif:        98,
			},
		},
	},
	Rates: RatesSource{
		Exchange:         "graviex",
		FallBackExchange: "southxchange",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "telos.polispay.com",
		BlockTime:        1,
		MinConfirmations: 6,
	},
	NetParams: &chaincfg.Params{
		Bech32HRPSegwit:   "",
		ScriptHashAddrID:  []byte{138},
		PubKeyHashAddrID:  []byte{38},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        424,
		Base58CksumHasher: base58.Sha256D,
		Net:               12, // Make sure doesn't collide with any other coin.
		AddressMagicLen:   1,
	},
}
