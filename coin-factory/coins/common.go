package coins

import (
	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

// This network params are just an override around the wire net information to make sure HD wallet prefixes are
// registered. For BTC defaults (all coins except Polis) just use the BtcNet param.
var (
	BtcNet   wire.BitcoinNet = 1
	PolisNet wire.BitcoinNet = 2
)

// Coin is the basic coin structure to get the correct properties for each coin.
type Coin struct {
	Tag            string
	Name           string
	Rates          RatesSource
	BlockchainInfo BlockchainInfo
	Mnemonic       string
	NetParams      *chaincfg.Params
	BlockExplorer  string
	Token          bool
	TokenNetwork   string
	Contract       string
}

// RatesSource is the prefered source of exchange for rates
type RatesSource struct {
	Exchange         string // The main exchange used to get rate information.
	FallBackExchange string // The fallback exchange to get if first fail.
}

// BlockchainInfo is a model to get information for a particular blockchain
type BlockchainInfo struct {
	BlockTime        float64 // Time between blocks in minutes.
	MinConfirmations int     // Minimum confirmations to mark a tx as usable.
	ExternalSource   string  // External source of information to double validate agains another trusted source
}
