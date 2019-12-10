package coins

import "github.com/btcsuite/btcd/chaincfg"

// Ethereum coinfactory information
var Ethereum = Coin{
	Tag:  "ETH",
	Name: "ethereum",
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "",
	},
	BlockchainInfo: BlockchainInfo{
		BlockTime:        1,
		MinConfirmations: 8,
		ExternalSource:   "eth1.trezor.io",
	},
	NetParams: &chaincfg.Params{
		// TODO this is created different
		Bech32HRPSegwit: "",
		PubKeyHashAddrID: 0,
		HDPrivateKeyID:   [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:    [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:       60,
	},
	Token:         false,
	BlockExplorer: "https://eth1.trezor.io",
}
