package coins

import "github.com/btcsuite/btcd/chaincfg"

// Bitcoin coinfactory information
var Bitcoin = Coin{
	Tag:  "BTC",
	Name: "bitcoin",
	Rates: RatesSource{
		Exchange:         "bitso",
		FallBackExchange: "",
	},
	BlockchainInfo: BlockchainInfo{
		BlockTime:        10,
		MinConfirmations: 1,
		ExternalSource:   "btc2.trezor.io",
	},
	NetParams: &chaincfg.Params{
		Bech32HRPSegwit:  "bc",
		PubKeyHashAddrID: 0,
		ScriptHashAddrID: 5,
		HDPrivateKeyID:   [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:    [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:       0,
	},
	Token:         false,
	BlockExplorer: "https://btc1.trezor.io",
}
