package coins

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"os"
	"path/filepath"
)

// TrueUSD coinfactory information
var TrueUSD = Coin{
	Info: CoinInfo{
		Tag:          "TUSD",
		Name:         "TrueUSD (TUSD)",
		Trezor:       true,
		Ledger:       true,
		Segwit:       false,
		Masternodes:  false,
		Token:        true,
		TokenNetwork: "ethereum",
		Contract:     "",
	},
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "",
	},
}

func NewTrueUSDInfo() *Coin {
	path, err := filepath.Abs("coins/icons/trueusd.png")
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
	TrueUSD.Info.Icon = base64.StdEncoding.EncodeToString(content)
	return &TrueUSD
}
