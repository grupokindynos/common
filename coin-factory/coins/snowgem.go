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

var Snowgem = Coin{
	Info: CoinInfo{
		Tag:         "XSG",
		Name:        "SnowGem (XSG)",
		Trezor:      false,
		Ledger:      false,
		Segwit:      false,
		Masternodes: true,
		Token:       false,
		Blockbook:   "https://xsg.polispay.com",
		Protocol:    "xsg",
		TxVersion:   1,

		TxBuilder: "bitcoinjs",
		HDIndex:   0,
		Networks: map[string]CoinNetworkInfo{
			"P2PKH": {
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
		Exchange:         "stex",
		FallBackExchange: "cryptobridge",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "xsg.polispay.com",
		BlockTime:        1,
		MinConfirmations: 12,
	},
	NetParams: &chaincfg.Params{
		// TODO this is created different
		Bech32HRPSegwit:   "",
		ScriptHashAddrID:  []byte{0},
		PubKeyHashAddrID:  []byte{0},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        410,
		Base58CksumHasher: base58.Sha256D,
		Net:               11, // Make sure doesn't collide with any other coin.
		AddressMagicLen:   2,
	},
}

func NewSnowGemInfo() *Coin {
	path, err := filepath.Abs("coins/icons/snowgem.png")
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
	Snowgem.Info.Icon = base64.StdEncoding.EncodeToString(content)
	return &Snowgem
}
