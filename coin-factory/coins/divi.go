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

// Divi coinfactory information
var Divi = Coin{
	Info: CoinInfo{
		Tag:         "DIVI",
		Name:        "Divi (DIVI)",
		Trezor:      false,
		Ledger:      false,
		Segwit:      false,
		Masternodes: true,
		Token:       false,
		Blockbook:   "https://divi.polispay.com",
		Protocol:    "divi",
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

func NewDiviInfo() *Coin {
	path, err := filepath.Abs("coins/icons/divi.png")
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
	Divi.Info.Icon = base64.StdEncoding.EncodeToString(content)
	return &Divi
}
