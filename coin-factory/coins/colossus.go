package coins

import (
	coinfactory "github.com/grupokindynos/common/coin-factory"
	"github.com/martinboehm/btcutil/chaincfg"
)

var Colossus = Coin{
	Tag:  "COLX",
	Name: "colossus",
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
		Bech32HRPSegwit:  "",
		PubKeyHashAddrID: []byte{0x1e},
		ScriptHashAddrID: []byte{0xd},
		HDPrivateKeyID:   [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:    [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:       1999,
		Net:              coinfactory.BtcNet,
	},
	Token:         false,
	BlockExplorer: "https://colx.polispay.com",
}
