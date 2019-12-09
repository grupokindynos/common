package coins

import "github.com/btcsuite/btcd/chaincfg"

var Zcoin = Coin{
	Tag:  "XZC",
	Name: "zcoin",
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "bittrex",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "xzc.polispay.com",
		BlockTime:        5,
		MinConfirmations: 5,
	},
	NetParams: &chaincfg.Params{
		PubKeyHashAddrID: 82,
		HDPrivateKeyID:   [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:    [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:       136,
	},
	Token:         false,
	BlockExplorer: "https://xzc.polispay.com",
}
