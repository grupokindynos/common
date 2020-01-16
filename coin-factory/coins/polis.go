package coins

import (
	"bufio"
	"encoding/base64"
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
	"io/ioutil"
	"os"
)

var Polis = Coin{
	Info: CoinInfo{
		Tag:         "POLIS",
		Name:        "Polis (POLIS)",
		Trezor:      true,
		Ledger:      false,
		Segwit:      false,
		Masternodes: true,
		Token:       false,
		Blockbook:   "https://blockbook.polispay.orgr",
		Protocol:    "polis",
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

func NewPolisInfo() *Coin {
	f, err := os.Open("./icons/polis.png")
	if err != nil {
		return nil
	}
	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil
	}
	Polis.Info.Icon = base64.StdEncoding.EncodeToString(content)
	return &Polis
}
