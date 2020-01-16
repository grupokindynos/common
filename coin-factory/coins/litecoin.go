package coins

import (
	"bufio"
	"encoding/base64"
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
	"io/ioutil"
	"os"
)

var Litecoin = Coin{
	Info: CoinInfo{
		Tag:         "LTC",
		Name:        "Litecoin (LTC)",
		Trezor:      true,
		Ledger:      true,
		Segwit:      true,
		Masternodes: false,
		Token:       false,
		Blockbook:   "https://ltc1.trezor.io",
		Protocol:    "ltc",
		TxVersion:   1,
		TxBuilder:   "bitcoinjs",

		HDIndex: 0,
		Networks: map[string]CoinNetworkInfo{
			"P2SHInP2WPKH": {
				MessagePrefix: "\x18Bitcoin Signed Message:\n",
				Bech32:        "bc",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x049d7cb2,
					Private: 0x049d7878,
				},
				PubKeyHash: 0x00,
				ScriptHash: 0x05,
				Wif:        0x80,
			},
			"P2WPKH": {
				MessagePrefix: "\x18Bitcoin Signed Message:\n",
				Bech32:        "bc",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x04b24746,
					Private: 0x04b2430c,
				},
				PubKeyHash: 0x00,
				ScriptHash: 0x05,
				Wif:        0x80,
			},
		},
	},
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "bittrex",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "ltc2.trezor.io",
		BlockTime:        2.5,
		MinConfirmations: 4,
	},
	NetParams: &chaincfg.Params{
		Bech32HRPSegwit:   "ltc",
		ScriptHashAddrID:  []byte{50},
		PubKeyHashAddrID:  []byte{48},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        2,
		Base58CksumHasher: base58.Sha256D,
		Net:               8, // Make sure doesn't collide with any other coin.
		AddressMagicLen:   1,
	},
}

func NewLitecoinInfo() *Coin {
	f, err := os.Open("coins/icons/litecoin.png")
	if err != nil {
		return nil
	}
	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil
	}
	Litecoin.Info.Icon = base64.StdEncoding.EncodeToString(content)
	return &Litecoin
}
