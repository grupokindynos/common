package coins

import (
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
)

var Litecoin = Coin{
	Tag:  "LTC",
	Name: "litecoin",
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "bittrex",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "ltc2.trezor.io",
		BlockTime:        2.5,
		MinConfirmations: 4,
	},
	NetParams: &chaincfg.Params{
		Bech32HRPSegwit:   "ltc",
		ScriptHashAddrID:  []byte{50},
		PubKeyHashAddrID:  []byte{48},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        2,
		Base58CksumHasher: base58.Sha256D,
		Net:               8, // Make sure doesn't collide with any other coin.
		AddressMagicLen:   1,
	},
	Token:         false,
	BlockExplorer: "https://ltc2.trezor.io",
}
