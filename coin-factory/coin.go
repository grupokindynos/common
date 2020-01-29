package coinfactory

import (
	"errors"
	"github.com/grupokindynos/common/coin-factory/coins"
	"os"
	"strings"
)

// Coins refers to the coins that are being used on the API instance
var Coins = map[string]*coins.Coin{
	"BTC":   &coins.Bitcoin,
	"BITG":  &coins.Bitgreen,
	"COLX":  &coins.Colossus,
	"DASH":  &coins.Dash,
	"DGB":   &coins.Digibyte,
	"DIVI":  &coins.Divi,
	"ETH":   &coins.Ethereum,
	"GRS":   &coins.Groestlcoin,
	"LTC":   &coins.Litecoin,
	"ONION": &coins.DeepOnion,
	"POLIS": &coins.Polis,
	"TELOS": &coins.Telos,
	"USDT":  &coins.Tether,
	"TUSD":  &coins.TrueUSD,
	"USDC":  &coins.USDCoin,
	"XZC":   &coins.Zcoin,
}

// GetCoin is the safe way to check if a coin exists and retrieve the coin data
func GetCoin(tag string) (*coins.Coin, error) {
	coin, ok := Coins[strings.ToUpper(tag)]
	if !ok {
		return nil, errors.New("coin not available")
	}
	coin = &coins.Coin{
		Info:           coin.Info,
		BlockchainInfo: coin.BlockchainInfo,
		Rates:          coin.Rates,
		NetParams:      coin.NetParams,
		Mnemonic:       os.Getenv("MNEMONIC_" + strings.ToUpper(tag)),
	}
	return coin, nil
}

func loadIcons() {

}
