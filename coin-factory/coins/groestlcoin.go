package coins

import "github.com/martinboehm/btcutil/chaincfg"

var Groestlcoin = Coin{
	Tag:  "GRS",
	Name: "groestlcoin",
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "bittrex",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "grs.polispay.com",
		BlockTime:        10,
		MinConfirmations: 2,
	},
	NetParams: &chaincfg.Params{
		// TODO this is created different
		Bech32HRPSegwit:  "grs",
		ScriptHashAddrID: []byte{5},
		PubKeyHashAddrID: []byte{24},
		HDPrivateKeyID:   [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:    [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:       17,
	},
	Token:         false,
	BlockExplorer: "https://grs.polispay.com",
}
