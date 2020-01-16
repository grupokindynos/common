package coins

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"os"
)

// Tether coinfactory information
var Tether = Coin{
	Info: CoinInfo{
		Tag:          "USDT",
		Name:         "Tether (USDT)",
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

func NewTetherInfo() *Coin {
	f, err := os.Open("./icons/tether.png")
	if err != nil {
		return nil
	}
	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil
	}
	Tether.Info.Icon = base64.StdEncoding.EncodeToString(content)
	return &Tether
}
