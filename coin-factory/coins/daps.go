package coins

import (
	"github.com/eabz/btcutil/base58"
	"github.com/eabz/btcutil/chaincfg"
)

var Daps = Coin{
	Info: CoinInfo{
		Icon:        "data:image/png;base64,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",
		Tag:         "DAPS",
		Name:        "Daps (DAPS)",
		Trezor:      false,
		Ledger:      false,
		Segwit:      false,
		Masternodes: false,
		Token:       false,
		StableCoin:  false,
		Blockbook:   "https://daps.polispay.com",
		Protocol:    "daps",
		TxVersion:   1,
		TxBuilder:   "bitcoinjs",
		HDIndex:     358,
		Networks: map[string]CoinNetworkInfo{
			"P2PKH": {
				MessagePrefix: "\x86Daps Signed Message:\n",
				Bech32:        "",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x488b21e,
					Private: 0x488ade4,
				},
				PubKeyHash: []int{30},
				ScriptHash: []int{13},
				Wif:        []int{212},
			},
		},
	},
	Rates: RatesSource{
		Exchange:         "stex",
		FallBackExchange: "",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "crw.polispay.com",
		BlockTime:        2,
		MinConfirmations: 3,
	},
	NetParams: &chaincfg.Params{
		AddressMagicLen:   1,
		Bech32HRPSegwit:   "",
		PubKeyHashAddrID:  []byte{30},
		ScriptHashAddrID:  []byte{13},
		PrivateKeyID:      []byte{212},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        358,
		Base58CksumHasher: base58.Sha256D,
	},
}
