package coins

import "github.com/btcsuite/btcd/chaincfg"

var Polis = Coin{
	Tag:  "POLIS",
	Name: "polis",
	Rates: RatesSource{
		Exchange:         "southxchange",
		FallBackExchange: "cryptobridge",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "blockbook.polispay.org",
		BlockTime:        2,
		MinConfirmations: 2,
	},
	NetParams: &chaincfg.Params{
		HDPrivateKeyID: [4]byte{0x03, 0xE2, 0x59, 0x45},
		HDPublicKeyID:  [4]byte{0x03, 0xE2, 0x5D, 0x7E},
		HDCoinType:     1997,
	},
	Token:         false,
	BlockExplorer: "https://blockbook.polispay.org",
}
