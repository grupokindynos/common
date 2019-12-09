package coins

import "github.com/btcsuite/btcd/chaincfg"

var DeepOnion = Coin{
	Tag:  "ONION",
	Name: "deeponion",
	Rates: RatesSource{
		Exchange:         "kucoin",
		FallBackExchange: "crex24",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "",
		BlockTime:        4,
		MinConfirmations: 20,
	},
	NetParams: &chaincfg.Params{
		HDPrivateKeyID: [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:  [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:     305,
	},
	Token:         false,
	BlockExplorer: "https://onion.polispay.com",
}
