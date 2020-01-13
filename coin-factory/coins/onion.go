package coins

import (
	coinfactory "github.com/grupokindynos/common/coin-factory"
	"github.com/martinboehm/btcutil/chaincfg"
)

var DeepOnion = Coin{
	Tag:  "ONION",
	Name: "deeponion",
	Rates: RatesSource{
		Exchange:         "kucoin",
		FallBackExchange: "crex24",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "",
		BlockTime:        4,
		MinConfirmations: 20,
	},
	NetParams: &chaincfg.Params{
		Bech32HRPSegwit:  "",
		ScriptHashAddrID: []byte{0x4e},
		PubKeyHashAddrID: []byte{48},
		HDPrivateKeyID:   [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:    [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:       305,
		Net:              coinfactory.BtcNet,
	},
	Token:         false,
	BlockExplorer: "https://onion.polispay.com",
}
