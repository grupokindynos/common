package coins

import (
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
)

var Snowgem = Coin{
	Tag:  "XSG",
	Name: "snowgem",
	Rates: RatesSource{
		Exchange:         "stex",
		FallBackExchange: "cryptobridge",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "xsg.polispay.com",
		BlockTime:        1,
		MinConfirmations: 12,
	},
	NetParams: &chaincfg.Params{
		// TODO this is created different
		Bech32HRPSegwit:   "",
		ScriptHashAddrID:  []byte{0},
		PubKeyHashAddrID:  []byte{0},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        410,
		Base58CksumHasher: base58.Sha256D,
		Net:               BtcNet,
	},
	Token:         false,
	BlockExplorer: "https://xsg.polispay.com",
}
