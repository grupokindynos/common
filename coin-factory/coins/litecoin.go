package coins

import (
	coinfactory "github.com/grupokindynos/common/coin-factory"
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
		Bech32HRPSegwit:  "ltc",
		ScriptHashAddrID: []byte{50},
		PubKeyHashAddrID: []byte{48},
		HDPrivateKeyID:   [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:    [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:       2,
		Net:              coinfactory.BtcNet,
	},
	Token:         false,
	BlockExplorer: "https://ltc2.trezor.io",
}
