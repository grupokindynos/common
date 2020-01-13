package coinfactory

import (
	"errors"
	"github.com/grupokindynos/common/coin-factory/coins"
	"os"
	"strings"
)

// Coins refers to the coins that are being used on the API instance
var Coins = map[string]*coins.Coin{
	"POLIS": &coins.Polis,
	"DGB":   &coins.Digibyte,
	"XZC":   &coins.Zcoin,
	"LTC":   &coins.Litecoin,
	"BTC":   &coins.Bitcoin,
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
		Tag:            coin.Tag,
		Name:           coin.Name,
		BlockchainInfo: coin.BlockchainInfo,
		Rates:          coin.Rates,
		BlockExplorer:  coin.BlockExplorer,
		Token:          coin.Token,
		TokenNetwork:   coin.TokenNetwork,
		Contract:       coin.Contract,
		NetParams:      coin.NetParams,
		Mnemonic:       os.Getenv("MNEMONIC_" + strings.ToUpper(tag)),
	}
	return coin, nil
}
