package coins

import (
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
)

// Ethereum coinfactory information
var Ethereum = Coin{
	Tag:  "ETH",
	Name: "ethereum",
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "",
	},
	BlockchainInfo: BlockchainInfo{
		BlockTime:        1,
		MinConfirmations: 8,
		ExternalSource:   "eth1.trezor.io",
	},
	NetParams: &chaincfg.Params{
		// TODO this is created different
		Bech32HRPSegwit:   "",
		PubKeyHashAddrID:  []byte{0},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        60,
		Base58CksumHasher: base58.Sha256D,
		Net:               6, // Make sure doesn't collide with any other coin.
		AddressMagicLen:   1,
	},
	Token:         false,
	BlockExplorer: "https://eth1.trezor.io",
}
