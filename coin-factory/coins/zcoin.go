package coins

import (
	"github.com/eabz/btcutil/base58"
	"github.com/eabz/btcutil/chaincfg"
)

var Zcoin = Coin{
	Info: CoinInfo{
		Icon:        "data:image/png;base64,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",
		Tag:         "XZC",
		Name:        "Zcoin (XZC)",
		Trezor:      false,
		Ledger:      false,
		Segwit:      false,
		Masternodes: true,
		Token:       false,
		StableCoin:  false,
		Blockbook:   "https://xzc.polispay.com",
		Protocol:    "xzc",
		TxVersion:   1,
		TxBuilder:   "bitcoinjs",
		HDIndex:     136,
		Networks: map[string]CoinNetworkInfo{
			"P2PKH": {
				MessagePrefix: "\x18ZCoin Signed Message:\n",
				Bech32:        "",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x488b21e,
					Private: 0x488ade4,
				},
				PubKeyHash: 82,
				ScriptHash: 7,
				Wif:        210,
			},
		},
	},
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "bittrex",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "xzc.polispay.com",
		BlockTime:        5,
		MinConfirmations: 1,
	},
	NetParams: &chaincfg.Params{
		AddressMagicLen:   1,
		Bech32HRPSegwit:   "",
		ScriptHashAddrID:  []byte{7},
		PubKeyHashAddrID:  []byte{82},
		PrivateKeyID:      []byte{210},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        136,
		Base58CksumHasher: base58.Sha256D,
	},
}
