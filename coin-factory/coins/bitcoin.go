package coins

import (
	"bufio"
	"encoding/base64"
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
	"io/ioutil"
	"os"
)

// Bitcoin coinfactory information
var Bitcoin = Coin{
	Info: CoinInfo{
		Icon:        "",
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
		Exchange:         "bitso",
		FallBackExchange: "",
	},
	BlockchainInfo: BlockchainInfo{
		BlockTime:        10,
		MinConfirmations: 1,
		ExternalSource:   "btc2.trezor.io",
	},
	NetParams: &chaincfg.Params{
		Bech32HRPSegwit:   "bc",
		PubKeyHashAddrID:  []byte{0},
		ScriptHashAddrID:  []byte{5},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        0,
		Base58CksumHasher: base58.Sha256D,
		Net:               1, // Make sure doesn't collide with any other coin.
		AddressMagicLen:   1,
	},
}

func NewBitcoinInfo() *Coin {
	f, _ := os.Open("../icons/bitcoin.png")
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	Bitcoin.Info.Icon = base64.StdEncoding.EncodeToString(content)
	return &Bitcoin
}
