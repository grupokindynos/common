package coins

import (
	"bufio"
	"encoding/base64"
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
	"io/ioutil"
	"os"
)

var Digibyte = Coin{
	Info: CoinInfo{
		Tag:         "DGB",
		Name:        "DigiByte (DGB)",
		Trezor:      true,
		Ledger:      true,
		Segwit:      true,
		Masternodes: false,
		Token:       false,
		Blockbook:   "https://dgb1.trezor.io",
		Protocol:    "dgb",
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
		Exchange:         "bittrex",
		FallBackExchange: "",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "dgb2.trezor.io",
		BlockTime:        0.25,
		MinConfirmations: 10,
	},
	NetParams: &chaincfg.Params{
		Bech32HRPSegwit:   "dgb",
		PubKeyHashAddrID:  []byte{30},
		ScriptHashAddrID:  []byte{63},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        20,
		Base58CksumHasher: base58.Sha256D,
		Net:               4, // Make sure doesn't collide with any other coin.
		AddressMagicLen:   1,
	},
}

func NewDigibyteInfo() *Coin {
	f, err := os.Open("../icons/digibyte.png")
	if err != nil {
		return nil
	}
	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil
	}
	Digibyte.Info.Icon = base64.StdEncoding.EncodeToString(content)
	return &Digibyte
}
