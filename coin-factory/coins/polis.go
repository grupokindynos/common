package coins

import (
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
)

var Polis = Coin{
	Tag:  "POLIS",
	Name: "polis",
	Rates: RatesSource{
		Exchange:         "southxchange",
		FallBackExchange: "crex24",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "blockbook.polispay.org",
		BlockTime:        2,
		MinConfirmations: 2,
	},
	NetParams: &chaincfg.Params{
		Bech32HRPSegwit:   "",
		ScriptHashAddrID:  []byte{56},
		PubKeyHashAddrID:  []byte{55},
		HDPrivateKeyID:    [4]byte{0x03, 0xE2, 0x59, 0x45},
		HDPublicKeyID:     [4]byte{0x03, 0xE2, 0x5D, 0x7E},
		HDCoinType:        1997,
		Base58CksumHasher: base58.Sha256D,
		Net:               10, // Make sure doesn't collide with any other coin.
	},
	Token:         false,
	BlockExplorer: "https://blockbook.polispay.org",
}
