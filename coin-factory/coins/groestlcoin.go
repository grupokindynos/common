package coins

import (
	"bufio"
	"encoding/base64"
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
	"io/ioutil"
	"os"
	"path/filepath"
)

var Groestlcoin = Coin{
	Info: CoinInfo{
		Tag:         "GRS",
		Name:        "GroestlCoin (GRS)",
		Trezor:      true,
		Ledger:      true,
		Segwit:      true,
		Masternodes: false,
		Token:       false,
		Blockbook:   "https://grs.polispay.com",
		Protocol:    "grs",
		TxVersion:   1,
		TxBuilder:   "groestlcoinjs",

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
		ExternalSource:   "grs.polispay.com",
		BlockTime:        10,
		MinConfirmations: 2,
	},
	NetParams: &chaincfg.Params{
		Bech32HRPSegwit:   "grs",
		ScriptHashAddrID:  []byte{5},
		PubKeyHashAddrID:  []byte{36},
		Base58CksumHasher: base58.Groestl512D,
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        17,
		Net:               7, // Make sure doesn't collide with any other coin.
		AddressMagicLen:   1,
	},
}

func NewGroestlCoinInfo() *Coin {
	path, err := filepath.Abs("coins/icons/groestlcoin.png")
	if err != nil {
		return nil
	}
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil
	}
	Groestlcoin.Info.Icon = base64.StdEncoding.EncodeToString(content)
	return &Groestlcoin
}
