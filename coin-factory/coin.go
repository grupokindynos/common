package coinfactory

import (
	"errors"
	"strings"
)

// Coin is the basic coin structure to get the correct properties for each coin.
type Coin struct {
	Tag              string
	Name             string
	Exchange         string
	FallBackExchange string
}

var (
	bitcoin = Coin{
		Tag:              "BTC",
		Name:             "bitcoin",
		Exchange:         "bitso",
		FallBackExchange: "",
	}
	onion = Coin{
		Tag:              "ONION",
		Name:             "deeponion",
		Exchange:         "kucoin",
		FallBackExchange: "crex24",
	}
	colossus = Coin{
		Tag:              "COLX",
		Name:             "colossus",
		Exchange:         "cryptobridge",
		FallBackExchange: "novaexchange",
	}
	dash = Coin{
		Tag:              "DASH",
		Name:             "dash",
		Exchange:         "binance",
		FallBackExchange: "bittrex",
	}
	digibyte = Coin{
		Tag:              "DGB",
		Name:             "digibyte",
		Exchange:         "bittrex",
		FallBackExchange: "",
	}
	groestlcoin = Coin{
		Tag:              "GRS",
		Name:             "groestlcoin",
		Exchange:         "binance",
		FallBackExchange: "bittrex",
	}
	litecoin = Coin{
		Tag:              "LTC",
		Name:             "litecoin",
		Exchange:         "binance",
		FallBackExchange: "bittrex",
	}
	mnpcoin = Coin{
		Tag:              "MNP",
		Name:             "mnpcoin",
		Exchange:         "crex24",
		FallBackExchange: "stex",
	}
	polis = Coin{
		Tag:              "POLIS",
		Name:             "polis",
		Exchange:         "cryptobridge",
		FallBackExchange: "southxchange",
	}
	snowgem = Coin{
		Tag:              "XSG",
		Name:             "snowgem",
		Exchange:         "stex",
		FallBackExchange: "cryptobridge",
	}
	zcoin = Coin{
		Tag:              "XZC",
		Name:             "zcoin",
		Exchange:         "binance",
		FallBackExchange: "bittrex",
	}
)

// Coins is the main array where used coins are stored
type Coins []Coin

// CoinFactory refers to the coins that are being used on the API instance
var CoinFactory = Coins{
	bitcoin,
	polis,
	onion,
	dash,
	litecoin,
	digibyte,
	groestlcoin,
	zcoin,
	colossus,
	mnpcoin,
	snowgem,
}

// GetCoin is the safe way to check if a coin exists and retrieve the coin data
func GetCoin(tag string) (*Coin, error) {
	for _, v := range CoinFactory {
		if v.Tag == strings.ToUpper(tag) {
			return &v, nil
		}
	}
	return nil, errors.New("coin not available")
}
