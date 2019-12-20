package coins

import "github.com/btcsuite/btcd/chaincfg"

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
		ScriptHashAddrID: 50,
		PubKeyHashAddrID: 48,
		HDPrivateKeyID:   [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:    [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:       2,
	},
	Token:         false,
	BlockExplorer: "https://ltc2.trezor.io",
}
