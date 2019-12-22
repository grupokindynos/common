package coins

import "github.com/btcsuite/btcd/chaincfg"

// Divi coinfactory information
var Divi = Coin{
	Tag:  "DIVI",
	Name: "divi",
	Rates: RatesSource{
		Exchange:         "bitrue",
		FallBackExchange: "crex24",
	},
	BlockchainInfo: BlockchainInfo{
		BlockTime:        2,
		MinConfirmations: 6,
		ExternalSource:   "divi.polispay.com",
	},
	NetParams: &chaincfg.Params{
		// chainparams https://github.com/DiviProject/Divi/blob/master0/divi/src/chainparams.cpp#L236-L238
		Bech32HRPSegwit:  "",
		PubKeyHashAddrID: 30,
		ScriptHashAddrID: 13,
		// Bitcoin defaults
		HDPrivateKeyID:   [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:    [4]byte{0x04, 0x88, 0xB2, 0x1E},
		// HD Coin Slip44 https://github.com/satoshilabs/slips/blob/master/slip-0044.md
		HDCoinType:       301,
	},
	Token:         false,
	BlockExplorer: "https://divi.polispay.com",
}