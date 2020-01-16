package coins

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"os"
)

// USDCoin coinfactory information
var USDCoin = Coin{
	Info: CoinInfo{
		Tag:          "USDC",
		Name:         "USDCoin (USDC)",
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

func NewUSDCoinInfo() *Coin {
	f, err := os.Open("./icons/usdcoin.png")
	if err != nil {
		return nil
	}
	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil
	}
	USDCoin.Info.Icon = base64.StdEncoding.EncodeToString(content)
	return &USDCoin
}
