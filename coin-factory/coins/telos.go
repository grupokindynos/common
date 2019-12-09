package coins

import "github.com/btcsuite/btcd/chaincfg"

var Telos = Coin{
	Tag:  "TELOS",
	Name: "teloscoin",
	Rates: RatesSource{
		Exchange:         "graviex",
		FallBackExchange: "southxchange",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "telos.polispay.com",
		BlockTime:        1,
		MinConfirmations: 6,
	},
	NetParams: &chaincfg.Params{
		HDPrivateKeyID: [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:  [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:     424,
	},
	Token:         false,
	BlockExplorer: "https://telos.polispay.com",
}
