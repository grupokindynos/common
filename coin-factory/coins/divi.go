package coins

import (
	"bufio"
	"encoding/base64"
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
	"io/ioutil"
	"os"
)

var icnDivi string

func init() {
	f, _ := os.Open("../icons/divi.png")
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	icnDivi = base64.StdEncoding.EncodeToString(content)
}

// Divi coinfactory information
var Divi = Coin{
	Info: CoinInfo{
		Icon:        icnDivi,
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
		Exchange:         "bitrue",
		FallBackExchange: "crex24",
	},
	BlockchainInfo: BlockchainInfo{
		BlockTime:        2,
		MinConfirmations: 6,
		ExternalSource:   "divi.polispay.com",
	},
	NetParams: &chaincfg.Params{
		// chainparams https://github.com/DiviProject/Divi/blob/master0/divi/src/chainparams.cpp#L236-L238
		Bech32HRPSegwit:  "",
		PubKeyHashAddrID: []byte{30},
		ScriptHashAddrID: []byte{13},
		// Bitcoin defaults
		HDPrivateKeyID: [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:  [4]byte{0x04, 0x88, 0xB2, 0x1E},
		// HD Coin Slip44 https://github.com/satoshilabs/slips/blob/master/slip-0044.md
		HDCoinType:        301,
		Base58CksumHasher: base58.Sha256D,
		Net:               5, // Make sure doesn't collide with any other coin.
		AddressMagicLen:   1,
	},
}
