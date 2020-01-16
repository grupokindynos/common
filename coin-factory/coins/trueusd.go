package coins

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"os"
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
	f, err := os.Open("../icons/trueusd.png")
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
