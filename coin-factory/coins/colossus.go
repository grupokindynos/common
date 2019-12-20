package coins

import "github.com/btcsuite/btcd/chaincfg"

var Colossus = Coin{
	Tag:  "COLX",
	Name: "colossus",
	Rates: RatesSource{
		Exchange:         "cryptobridge",
		FallBackExchange: "novaexchange",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "colx.polispay.com",
		BlockTime:        1,
		MinConfirmations: 30,
	},
	NetParams: &chaincfg.Params{
		Bech32HRPSegwit:  "",
		PubKeyHashAddrID: 0x1e,
		ScriptHashAddrID: 0xd,
		HDPrivateKeyID:   [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:    [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:       1999,
	},
	Token:         false,
	BlockExplorer: "https://colx.polispay.com",
}
