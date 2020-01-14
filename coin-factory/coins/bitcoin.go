package coins

import (
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
)

// Bitcoin coinfactory information
var Bitcoin = Coin{
	Tag:  "BTC",
	Name: "bitcoin",
	Rates: RatesSource{
		Exchange:         "bitso",
		FallBackExchange: "",
	},
	BlockchainInfo: BlockchainInfo{
		BlockTime:        10,
		MinConfirmations: 1,
		ExternalSource:   "btc2.trezor.io",
	},
	NetParams: &chaincfg.Params{
		Bech32HRPSegwit:   "bc",
		PubKeyHashAddrID:  []byte{0},
		ScriptHashAddrID:  []byte{5},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        0,
		Base58CksumHasher: base58.Sha256D,
		Net:               1, // Make sure doesn't collide with any other coin.
		AddressMagicLen:   1,
	},
	Token:         false,
	BlockExplorer: "https://btc1.trezor.io",
}
