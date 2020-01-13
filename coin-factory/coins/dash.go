package coins

import (
	coinfactory "github.com/grupokindynos/common/coin-factory"
	"github.com/martinboehm/btcutil/chaincfg"
)

var Dash = Coin{
	Tag:  "DASH",
	Name: "dash",
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "bittrex",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "dash2.trezor.io",
		BlockTime:        2.5,
		MinConfirmations: 2,
	},
	NetParams: &chaincfg.Params{
		Bech32HRPSegwit:  "",
		PubKeyHashAddrID: []byte{76},
		ScriptHashAddrID: []byte{16},
		HDPrivateKeyID:   [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:    [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:       5,
		Net:              coinfactory.BtcNet,
	},
	Token:         false,
	BlockExplorer: "https://dash2.trezor.io",
}
