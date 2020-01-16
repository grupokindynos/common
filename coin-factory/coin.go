package coinfactory

import (
	"errors"
	"github.com/grupokindynos/common/coin-factory/coins"
	"os"
	"strings"
)

// Coins refers to the coins that are being used on the API instance
var Coins = map[string]*coins.Coin{
	"BTC": coins.NewBitcoinInfo(),

	"POLIS": &coins.Polis,
	"DGB":   &coins.Digibyte,
	"XZC":   &coins.Zcoin,
	"LTC":   &coins.Litecoin,
	"DASH":  &coins.Dash,
	"DIVI":  &coins.Divi,
	"GRS":   &coins.Groestlcoin,
	"COLX":  &coins.Colossus,
	"ONION": &coins.DeepOnion,
	"XSG":   &coins.Snowgem,
	"TELOS": &coins.Telos,
	"ETH":   &coins.Ethereum,
	"TUSD":  &coins.TrueUSD,
	"USDT":  &coins.Tether,
	"USDC":  &coins.USDCoin,
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
