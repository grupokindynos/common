package coins

import (
	"bufio"
	"encoding/base64"
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
	"io/ioutil"
	"os"
)

var DeepOnion = Coin{
	Info: CoinInfo{
		Tag:         "ONION",
		Name:        "DeepOnion (ONION)",
		Trezor:      false,
		Ledger:      false,
		Segwit:      false,
		Masternodes: false,
		Token:       false,
		Blockbook:   "https://onion.polispay.com",
		Protocol:    "onion",
		TxVersion:   1,
		TxBuilder:   "bitcoinjs",

		HDIndex: 0,
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
		Exchange:         "kucoin",
		FallBackExchange: "crex24",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "",
		BlockTime:        4,
		MinConfirmations: 20,
	},
	NetParams: &chaincfg.Params{
		Bech32HRPSegwit:   "",
		ScriptHashAddrID:  []byte{0x4e},
		PubKeyHashAddrID:  []byte{48},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        305,
		Base58CksumHasher: base58.Sha256D,
		Net:               9, // Make sure doesn't collide with any other coin.
		AddressMagicLen:   1,
	},
}

func NewOnionInfo() *Coin {
	f, err := os.Open("coins/icons/deeponion.png")
	if err != nil {
		return nil
	}
	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil
	}
	DeepOnion.Info.Icon = base64.StdEncoding.EncodeToString(content)
	return &DeepOnion
}
