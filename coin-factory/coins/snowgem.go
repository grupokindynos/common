package coins

import "github.com/btcsuite/btcd/chaincfg"

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
		Bech32HRPSegwit: "",
		ScriptHashAddrID: 0,
		PubKeyHashAddrID: 0,
		HDPrivateKeyID:   [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:    [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:       410,
	},
	Token:         false,
	BlockExplorer: "https://xsg.polispay.com",
}
