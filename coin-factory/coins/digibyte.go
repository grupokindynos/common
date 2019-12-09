package coins

import "github.com/btcsuite/btcd/chaincfg"

var Digibyte = Coin{
	Tag:  "DGB",
	Name: "digibyte",
	Rates: RatesSource{
		Exchange:         "bittrex",
		FallBackExchange: "",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "dgb2.trezor.io",
		BlockTime:        0.25,
		MinConfirmations: 25,
	},
	NetParams: &chaincfg.Params{
		PubKeyHashAddrID: 30,
		HDPrivateKeyID:   [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:    [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:       20,
	},
	Token:         false,
	BlockExplorer: "https://dgb2.trezor.io",
}
