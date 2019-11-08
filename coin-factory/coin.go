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
	coinKeys := coins.Keys{
		RpcUser: os.Getenv(strings.ToUpper(tag) + "_RPC_USER"),
		RpcPass: os.Getenv(strings.ToUpper(tag) + "_RPC_PASS"),
		RpcPort: os.Getenv(strings.ToUpper(tag) + "_RPC_PORT"),
		Host:    os.Getenv(strings.ToUpper(tag) + "_IP"),
		Port:    os.Getenv(strings.ToUpper(tag) + "_SSH_PORT"),
		User:    os.Getenv(strings.ToUpper(tag) + "_SSH_USER"),
		PrivKey: os.Getenv(strings.ToUpper(tag) + "_SSH_PRIVKEY"),
	}
	coin = &coins.Coin{
		Tag:            coin.Tag,
		Name:           coin.Name,
		BlockchainInfo: coin.BlockchainInfo,
		Rates:          coin.Rates,
		RpcMethods:     coin.RpcMethods,
		ColdAddress:    os.Getenv(strings.ToUpper(tag) + "_COLD_ADDRESS"),
		Keys:           coinKeys,
		BlockExplorer:  coin.BlockExplorer,
		Token:          coin.Token,
		TokenNetwork:   coin.TokenNetwork,
		Contract:       coin.Contract,
	}
	return coin, nil
}

func CheckCoinKeys(coin *coins.Coin) error {
	keys := coin.Keys
	if keys.RpcUser == "" {
		return errors.New("missing rpc username")
	}
	if keys.RpcPass == "" {
		return errors.New("missing rpc password")
	}
	if keys.RpcPort == "" {
		return errors.New("missing rpc port")
	}
	if keys.Host == "" {
		return errors.New("missing host ip")
	}
	if keys.Port == "" {
		return errors.New("missing host port")
	}
	if keys.User == "" {
		return errors.New("missing host user")
	}
	if keys.PrivKey == "" {
		return errors.New("missing authorization token")
	}
	if coin.ColdAddress == "" {
		return errors.New("missing cold address to send")
	}

	return nil
}
