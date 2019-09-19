package coinfactory

import (
	"errors"
	"os"
	"strings"
)

// Coin is the basic coin structure to get the correct properties for each coin.
type Coin struct {
	Tag              string
	Name             string
	Exchange         string
	FallBackExchange string
	ExternalSource   string
	BlockchainInfo   BlockchainInfo
	RpcMethods       RPCMethods
	Address          string
	ColdAddress      string
	RpcUser          string
	RpcPass          string
	RpcPort          string
	Host             string
	Port             string
	User             string
	PrivKey          string
}

// Get information from Blockchain
type BlockchainInfo struct {
	MaxDiff        float64
	ValidationTime float64
	TxApiUrl       string
}

type RPCMethods struct {
	GetWalletInfo              string
	GetBlockchainInfo          string
	GetNetworkInfo             string
	GetNewAddress              string
	SendToAddress              string
	ValidateAddress            string
	GetRawTransaction          string
	GetRawTransactionVerbosity interface{}
}

var (
	bitcoin = Coin{
		Tag:              "BTC",
		Name:             "bitcoin",
		Exchange:         "bitso",
		FallBackExchange: "",
		ExternalSource:   "btc2.trezor.io",
		BlockchainInfo: BlockchainInfo{
			MaxDiff:        1000,
			ValidationTime: 10,
			TxApiUrl:       "https://insight.bitpay.com/api/tx/",
		},
		RpcMethods: RPCMethods{
			GetWalletInfo:              "getwalletinfo",
			GetBlockchainInfo:          "getblockchaininfo",
			GetNetworkInfo:             "getnetworkinfo",
			GetNewAddress:              "getnewaddress",
			SendToAddress:              "sendtoaddress",
			ValidateAddress:            "getaddressinfo",
			GetRawTransaction:          "getrawtransaction",
			GetRawTransactionVerbosity: "1",
		},
	}
	onion = Coin{
		Tag:              "ONION",
		Name:             "deeponion",
		Exchange:         "kucoin",
		FallBackExchange: "crex24",
		ExternalSource:   "",
		BlockchainInfo: BlockchainInfo{
			MaxDiff:        0,
			ValidationTime: 0,
			TxApiUrl:       "https://www.blockexperts.com/onion/tx/",
		},
		RpcMethods: RPCMethods{
			GetWalletInfo:              "getwalletinfo",
			GetBlockchainInfo:          "getblockchaininfo",
			GetNetworkInfo:             "getnetworkinfo",
			GetNewAddress:              "getnewaddress",
			SendToAddress:              "sendtoaddress",
			ValidateAddress:            "validateaddress",
			GetRawTransaction:          "getrawtransaction",
			GetRawTransactionVerbosity: true,
		},
	}
	colossus = Coin{
		Tag:              "COLX",
		Name:             "colossus",
		Exchange:         "cryptobridge",
		FallBackExchange: "novaexchange",
		ExternalSource:   "",
		BlockchainInfo: BlockchainInfo{
			MaxDiff:        2,
			ValidationTime: 1,
			TxApiUrl:       "https://delphi.polispay.com/api/v1/COLX/tx/", // TODO Update this as delphi will be deprecated
		},
		RpcMethods: RPCMethods{
			GetWalletInfo:              "getwalletinfo",
			GetBlockchainInfo:          "getblockchaininfo",
			GetNetworkInfo:             "getnetworkinfo",
			GetNewAddress:              "getnewaddress",
			SendToAddress:              "sendtoaddress",
			ValidateAddress:            "validateaddress",
			GetRawTransaction:          "getrawtransaction",
			GetRawTransactionVerbosity: true,
		},
	}
	dash = Coin{
		Tag:              "DASH",
		Name:             "dash",
		Exchange:         "binance",
		FallBackExchange: "bittrex",
		ExternalSource:   "dash2.trezor.io",
		BlockchainInfo: BlockchainInfo{
			MaxDiff:        10,
			ValidationTime: 2.5,
			TxApiUrl:       "https://insight.dash.org/api/tx/",
		},
		RpcMethods: RPCMethods{
			GetWalletInfo:              "getwalletinfo",
			GetBlockchainInfo:          "getblockchaininfo",
			GetNetworkInfo:             "getnetworkinfo",
			GetNewAddress:              "getnewaddress",
			SendToAddress:              "sendtoaddress",
			ValidateAddress:            "validateaddress",
			GetRawTransaction:          "getrawtransaction",
			GetRawTransactionVerbosity: true,
		},
	}
	digibyte = Coin{
		Tag:              "DGB",
		Name:             "digibyte",
		Exchange:         "bittrex",
		FallBackExchange: "",
		ExternalSource:   "dgb2.trezor.io",
		BlockchainInfo: BlockchainInfo{
			MaxDiff:        0, // TODO No info
			ValidationTime: 0.25,
			TxApiUrl:       "https://digiexplorer.info/api/tx/",
		},
		RpcMethods: RPCMethods{
			GetWalletInfo:              "getwalletinfo",
			GetBlockchainInfo:          "getblockchaininfo",
			GetNetworkInfo:             "getnetworkinfo",
			GetNewAddress:              "getnewaddress",
			SendToAddress:              "sendtoaddress",
			ValidateAddress:            "getaddressinfo",
			GetRawTransaction:          "getrawtransaction",
			GetRawTransactionVerbosity: "1",
		},
	}
	groestlcoin = Coin{
		Tag:              "GRS",
		Name:             "groestlcoin",
		Exchange:         "binance",
		FallBackExchange: "bittrex",
		ExternalSource:   "grs.polispay.com",
		BlockchainInfo: BlockchainInfo{
			MaxDiff:        0, // TODO No info
			ValidationTime: 0, // TODO No info
			TxApiUrl:       "https://groestlsight.groestlcoin.org/api/tx/",
		},
		RpcMethods: RPCMethods{
			GetWalletInfo:              "getwalletinfo",
			GetBlockchainInfo:          "getblockchaininfo",
			GetNetworkInfo:             "getnetworkinfo",
			GetNewAddress:              "getnewaddress",
			SendToAddress:              "sendtoaddress",
			ValidateAddress:            "getaddressinfo",
			GetRawTransaction:          "getrawtransaction",
			GetRawTransactionVerbosity: "1",
		},
	}
	litecoin = Coin{
		Tag:              "LTC",
		Name:             "litecoin",
		Exchange:         "binance",
		FallBackExchange: "bittrex",
		ExternalSource:   "ltc2.trezor.io",
		BlockchainInfo: BlockchainInfo{
			MaxDiff:        2,
			ValidationTime: 2.5,
			TxApiUrl:       "https://insight.litecore.io/api/tx/",
		},
		RpcMethods: RPCMethods{
			GetWalletInfo:              "getwalletinfo",
			GetBlockchainInfo:          "getblockchaininfo",
			GetNetworkInfo:             "getnetworkinfo",
			GetNewAddress:              "getnewaddress",
			SendToAddress:              "sendtoaddress",
			ValidateAddress:            "getaddressinfo",
			GetRawTransaction:          "getrawtransaction",
			GetRawTransactionVerbosity: "1",
		},
	}
	mnpcoin = Coin{
		Tag:              "MNP",
		Name:             "mnpcoin",
		Exchange:         "crex24",
		FallBackExchange: "stex",
		ExternalSource:   "",
		BlockchainInfo: BlockchainInfo{
			MaxDiff:        2,
			ValidationTime: 1,
			TxApiUrl:       "https://delphi.polispay.com/api/v1/MNP/tx/",
		},
		RpcMethods: RPCMethods{
			GetWalletInfo:              "getwalletinfo",
			GetBlockchainInfo:          "getblockchaininfo",
			GetNetworkInfo:             "getnetworkinfo",
			GetNewAddress:              "getnewaddress",
			SendToAddress:              "sendtoaddress",
			ValidateAddress:            "validateaddress",
			GetRawTransaction:          "getrawtransaction",
			GetRawTransactionVerbosity: true,
		},
	}
	polis = Coin{
		Tag:              "POLIS",
		Name:             "polis",
		Exchange:         "cryptobridge",
		FallBackExchange: "southxchange",
		ExternalSource:   "blockbook.polispay.org",
		BlockchainInfo: BlockchainInfo{
			MaxDiff:        2,
			ValidationTime: 2,
			TxApiUrl:       "https://blockbook.polispay.org/api/v2/tx/",
		},
		RpcMethods: RPCMethods{
			GetWalletInfo:              "getwalletinfo",
			GetBlockchainInfo:          "getblockchaininfo",
			GetNetworkInfo:             "getnetworkinfo",
			GetNewAddress:              "getnewaddress",
			SendToAddress:              "sendtoaddress",
			ValidateAddress:            "validateaddress",
			GetRawTransaction:          "getrawtransaction",
			GetRawTransactionVerbosity: true,
		},
	}
	snowgem = Coin{
		Tag:              "XSG",
		Name:             "snowgem",
		Exchange:         "stex",
		FallBackExchange: "cryptobridge",
		ExternalSource:   "",
		BlockchainInfo: BlockchainInfo{
			MaxDiff:        2, //TODO No info
			ValidationTime: 1,
			TxApiUrl:       "https://explorer.snowgem.org/tx/",
		},
		RpcMethods: RPCMethods{
			GetWalletInfo:              "getwalletinfo",
			GetBlockchainInfo:          "getblockchaininfo",
			GetNetworkInfo:             "getnetworkinfo",
			GetNewAddress:              "getnewaddress",
			SendToAddress:              "sendtoaddress",
			ValidateAddress:            "validateaddress",
			GetRawTransaction:          "getrawtransaction",
			GetRawTransactionVerbosity: true,
		},
	}
	zcoin = Coin{
		Tag:              "XZC",
		Name:             "zcoin",
		Exchange:         "binance",
		FallBackExchange: "bittrex",
		ExternalSource:   "xzc.polispay.com",
		BlockchainInfo: BlockchainInfo{
			MaxDiff:        0, //TODO No info
			ValidationTime: 5,
			TxApiUrl:       "https://insight.zcoin.io/api/tx/",
		},
		RpcMethods: RPCMethods{
			GetWalletInfo:              "getwalletinfo",
			GetBlockchainInfo:          "getblockchaininfo",
			GetNetworkInfo:             "getnetworkinfo",
			GetNewAddress:              "getnewaddress",
			SendToAddress:              "sendtoaddress",
			ValidateAddress:            "validateaddress",
			GetRawTransaction:          "getrawtransaction",
			GetRawTransactionVerbosity: true,
		},
	}
)

// Coins refers to the coins that are being used on the API instance
var Coins = map[string]*Coin{
	"POLIS": &polis,
	"DGB":   &digibyte,
	"XZC":   &zcoin,
	"LTC":   &litecoin,
	"BTC":   &bitcoin,
	"DASH":  &dash,
	"GRS":   &groestlcoin,
	"COLX":  &colossus,
	"ONION": &onion,
	"MNP":   &mnpcoin,
	"XSG":   &snowgem,
}

// GetCoin is the safe way to check if a coin exists and retrieve the coin data
func GetCoin(tag string) (*Coin, error) {
	coin, ok := Coins[strings.ToUpper(tag)]
	if !ok {
		return nil, errors.New("coin not available")
	}
	coin = &Coin{
		Tag:              coin.Tag,
		Name:             coin.Name,
		Exchange:         coin.Exchange,
		FallBackExchange: coin.FallBackExchange,
		ExternalSource:   coin.ExternalSource,
		RpcMethods:       coin.RpcMethods,
		ColdAddress:      os.Getenv(strings.ToUpper(tag) + "_COLD_ADDRESS"),
		RpcUser:          os.Getenv(strings.ToUpper(tag) + "_RPC_USER"),
		RpcPass:          os.Getenv(strings.ToUpper(tag) + "_RPC_PASS"),
		RpcPort:          os.Getenv(strings.ToUpper(tag) + "_RPC_PORT"),
		Host:             os.Getenv(strings.ToUpper(tag) + "_IP"),
		Port:             os.Getenv(strings.ToUpper(tag) + "_SSH_PORT"),
		User:             os.Getenv(strings.ToUpper(tag) + "_SSH_USER"),
		PrivKey:          os.Getenv(strings.ToUpper(tag) + "_SSH_PRIVKEY"),
	}
	return coin, nil
}

func CheckCoinConfigs(coin *Coin) error {
	if coin.RpcUser == "" {
		return errors.New("missing rpc username")
	}
	if coin.RpcPass == "" {
		return errors.New("missing rpc password")
	}
	if coin.RpcPort == "" {
		return errors.New("missing rpc port")
	}
	if coin.Host == "" {
		return errors.New("missing host ip")
	}
	if coin.Port == "" {
		return errors.New("missing host port")
	}
	if coin.User == "" {
		return errors.New("missing host user")
	}
	if coin.PrivKey == "" {
		return errors.New("missing authorization token")
	}
	if coin.ColdAddress == "" {
		return errors.New("missing cold address to send")
	}

	return nil
}
