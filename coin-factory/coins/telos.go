package coins

import (
	"bufio"
	"encoding/base64"
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
	"io/ioutil"
	"os"
)

var Telos = Coin{
	Info: CoinInfo{
		Tag:         "TELOS",
		Name:        "Telos (TELOS)",
		Trezor:      false,
		Ledger:      false,
		Segwit:      false,
		Masternodes: true,
		Token:       false,
		Blockbook:   "https://telos.polispay.com",
		Protocol:    "telos",
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
		Exchange:         "graviex",
		FallBackExchange: "southxchange",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "telos.polispay.com",
		BlockTime:        1,
		MinConfirmations: 6,
	},
	NetParams: &chaincfg.Params{
		Bech32HRPSegwit:   "",
		ScriptHashAddrID:  []byte{138},
		PubKeyHashAddrID:  []byte{38},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        424,
		Base58CksumHasher: base58.Sha256D,
		Net:               12, // Make sure doesn't collide with any other coin.
		AddressMagicLen:   1,
	},
}

func NewTelosInfo() *Coin {
	f, err := os.Open("./icons/telos.png")
	if err != nil {
		return nil
	}
	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil
	}
	Telos.Info.Icon = base64.StdEncoding.EncodeToString(content)
	return &Telos
}
