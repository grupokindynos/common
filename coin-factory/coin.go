package coinfactory

import (
	"errors"
	"os"
	"strings"

	"github.com/eabz/btcutil/chaincfg"
	"github.com/grupokindynos/common/coin-factory/coins"
	"github.com/martinboehm/btcd/wire"
)

// For bitcoin-like coins we need to asing a network magic to register addresses
// and prevent prefix collitions.
var nets = map[string]wire.BitcoinNet{
	"BTC":   wire.BitcoinNet(1),
	"BITG":  wire.BitcoinNet(2),
	"COLX":  wire.BitcoinNet(3),
	"DASH":  wire.BitcoinNet(4),
	"DGB":   wire.BitcoinNet(5),
	"DIVI":  wire.BitcoinNet(6),
	"GRS":   wire.BitcoinNet(7),
	"LTC":   wire.BitcoinNet(8),
	"POLIS": wire.BitcoinNet(10),
	"TELOS": wire.BitcoinNet(11),
	"XZC":   wire.BitcoinNet(12),
	"FYD":   wire.BitcoinNet(13),
	"AYA":   wire.BitcoinNet(14),
	"RPD":   wire.BitcoinNet(15),
	"MW":    wire.BitcoinNet(16),
	"XSG":   wire.BitcoinNet(17),
	"CRW":   wire.BitcoinNet(18),
	"DAPS":  wire.BitcoinNet(19),
}

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
	"POLIS": &coins.Polis,
	"TELOS": &coins.Telos,
	"USDT":  &coins.Tether,
	"TUSD":  &coins.TrueUSD,
	"USDC":  &coins.USDCoin,
	"XZC":   &coins.Zcoin,
	"FYD":   &coins.FYDCoin,
	"AYA":   &coins.AryaCoin,
	"RPD":   &coins.Rapids,
	"MW":    &coins.MasterWin,
	"LMY":   &coins.LunchMoney,
	"BAT":   &coins.BAT,
	"LINK":  &coins.Chainlink,
	"MANA":  &coins.Decentraland,
	"XSG":   &coins.SnowGem,
	"CRW":   &coins.Crown,
	"DAPS":  &coins.Daps,
	"NULS":  &coins.Nuls,
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
	if !coin.Info.Token && coin.Info.Tag != "NULS" {
		coin.NetParams.Net = nets[coin.Info.Tag]
		if !chaincfg.IsRegistered(coin.NetParams) {
			_ = chaincfg.Register(coin.NetParams)
		}
	}
	return coin, nil
}
