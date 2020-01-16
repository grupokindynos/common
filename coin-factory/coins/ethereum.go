package coins

import (
	"bufio"
	"encoding/base64"
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
	"io/ioutil"
	"os"
)

// Ethereum coinfactory information
var Ethereum = Coin{
	Info: CoinInfo{
		Tag:          "ETH",
		Name:         "Ethereum (ETH)",
		Trezor:       true,
		Ledger:       true,
		Segwit:       false,
		Masternodes:  false,
		Token:        true,
		TokenNetwork: "ethereum",
		Blockbook:    "https://eth2.trezor.io",
		Protocol:     "eth",
		TxVersion:    1,
		TxBuilder:    "ethereum",

		HDIndex: 0,
		Networks: map[string]CoinNetworkInfo{
			"ETHEREUM": {
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
		FallBackExchange: "",
	},
	BlockchainInfo: BlockchainInfo{
		BlockTime:        1,
		MinConfirmations: 8,
		ExternalSource:   "eth1.trezor.io",
	},
	NetParams: &chaincfg.Params{
		Bech32HRPSegwit:   "",
		PubKeyHashAddrID:  []byte{0},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        60,
		Base58CksumHasher: base58.Sha256D,
		Net:               6, // Make sure doesn't collide with any other coin.
		AddressMagicLen:   1,
	},
}

func NewEthereumInfo() *Coin {
	f, err := os.Open("./icons/ethereum.png")
	if err != nil {
		return nil
	}
	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil
	}
	Ethereum.Info.Icon = base64.StdEncoding.EncodeToString(content)
	return &Ethereum
}
