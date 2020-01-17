package coins

import (
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
)

// Divi coinfactory information
var Divi = Coin{
	Info: CoinInfo{
		Icon:        "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",
		Tag:         "DIVI",
		Name:        "Divi (DIVI)",
		Trezor:      false,
		Ledger:      false,
		Segwit:      false,
		Masternodes: true,
		Token:       false,
		Blockbook:   "https://divi.polispay.com",
		Protocol:    "divi",
		TxVersion:   1,
		TxBuilder:   "bitcoinjs",
		HDIndex:     301,
		Networks: map[string]CoinNetworkInfo{
			"P2PKH": {
				MessagePrefix: "\x18Divi Signed Message:\n",
				Bech32:        "",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x488b21e,
					Private: 0x488ade4,
				},
				PubKeyHash: 30,
				ScriptHash: 13,
				Wif:        212,
			},
		},
	},
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
		PubKeyHashAddrID: []byte{30},
		ScriptHashAddrID: []byte{13},
		PrivateKeyID:     []byte{212},

		// Bitcoin defaults
		HDPrivateKeyID: [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:  [4]byte{0x04, 0x88, 0xB2, 0x1E},
		// HD Coin Slip44 https://github.com/satoshilabs/slips/blob/master/slip-0044.md
		HDCoinType:        301,
		Base58CksumHasher: base58.Sha256D,
	},
}
