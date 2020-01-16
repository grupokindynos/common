package coinfactory

import (
	"errors"
	"github.com/grupokindynos/common/coin-factory/coins"
	"os"
	"strings"
)

// Coins refers to the coins that are being used on the API instance
var Coins = map[string]*coins.Coin{
	"BTC":   coins.NewBitcoinInfo(),
	"COLX":  coins.NewColossusInfo(),
	"DASH":  coins.NewDashInfo(),
	"DGB":   coins.NewDigibyteInfo(),
	"DIVI":  coins.NewDiviInfo(),
	"ETH":   coins.NewEthereumInfo(),
	"GRS":   coins.NewGroestlCoinInfo(),
	"LTC":   coins.NewLitecoinInfo(),
	"ONION": coins.NewOnionInfo(),
	"POLIS": coins.NewPolisInfo(),
	"XSG":   coins.NewSnowGemInfo(),
	"TELOS": coins.NewTelosInfo(),
	"USDT":  coins.NewTetherInfo(),
	"TUSD":  coins.NewTrueUSDInfo(),
	"USDC":  coins.NewUSDCoinInfo(),
	"XZC":   coins.NewZcoinInfo(),
}

// GetCoin is the safe way to check if a coin exists and retrieve the coin data
func GetCoin(tag string) (*coins.Coin, error) {
	coin, ok := Coins[strings.ToUpper(tag)]
	if !ok {
		return nil, errors.New("coin not available")
	}
	coin = &coins.Coin{
		BlockchainInfo: coin.BlockchainInfo,
		Rates:          coin.Rates,
		NetParams:      coin.NetParams,
		Mnemonic:       os.Getenv("MNEMONIC_" + strings.ToUpper(tag)),
	}
	return coin, nil
}

func loadIcons() {

}
