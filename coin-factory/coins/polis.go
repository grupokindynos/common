package coins

import (
	"bufio"
	"encoding/base64"
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
	"io/ioutil"
	"os"
)

var icnPolis string

func init() {
	f, _ := os.Open("../icons/polis.png")
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	icnPolis = base64.StdEncoding.EncodeToString(content)
}

var Polis = Coin{
	Info: CoinInfo{
		Icon:        icnPolis,
		Tag:         "BTC",
		Name:        "Bitcoin (BTC)",
		Trezor:      true,
		Ledger:      true,
		Segwit:      true,
		Masternodes: false,
		Token:       false,
		Blockbook:   "https://btc2.trezor.io",
		Protocol:    "bitcoin",
		TxVersion:   1,
		TxBuilder:   "bitcoinjs",
		HDIndex:     0,
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
		Exchange:         "southxchange",
		FallBackExchange: "crex24",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "blockbook.polispay.org",
		BlockTime:        2,
		MinConfirmations: 2,
	},
	NetParams: &chaincfg.Params{
		Bech32HRPSegwit:   "",
		ScriptHashAddrID:  []byte{56},
		PubKeyHashAddrID:  []byte{55},
		HDPrivateKeyID:    [4]byte{0x03, 0xE2, 0x59, 0x45},
		HDPublicKeyID:     [4]byte{0x03, 0xE2, 0x5D, 0x7E},
		HDCoinType:        1997,
		Base58CksumHasher: base58.Sha256D,
		Net:               10, // Make sure doesn't collide with any other coin.
		AddressMagicLen:   1,
	},
}
