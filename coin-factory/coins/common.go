package coins

import "github.com/eabz/btcutil/chaincfg"

type CoinNetWorkBip32Info struct {
	Public  int `json:"public"`
	Private int `json:"private"`
}

type CoinNetworkInfo struct {
	MessagePrefix string               `json:"messagePrefix"`
	Bech32        string               `json:"bech32"`
	Bip32         CoinNetWorkBip32Info `json:"bip32"`
	PubKeyHash    int                  `json:"pubKeyHash"`
	ScriptHash    int                  `json:"scriptHash"`
	Wif           int                  `json:"wif"`
}

type CoinInfo struct {
	Icon         string                     `json:"icon"`
	Tag          string                     `json:"tag"`
	Name         string                     `json:"name"`
	Trezor       bool                       `json:"trezor"`
	Ledger       bool                       `json:"ledger"`
	Segwit       bool                       `json:"segwit"`
	Masternodes  bool                       `json:"masternodes"`
	Token        bool                       `json:"token"`
	TokenNetwork string                     `json:"token_network,omitempty"`
	Contract     string                     `json:"contract,omitempty"`
	Blockbook    string                     `json:"blockbook"`
	Protocol     string                     `json:"protocol"`
	TxVersion    int                        `json:"tx_version"`
	TxBuilder    string                     `json:"tx_builder"`
	HDIndex      int                        `json:"hd_index"`
	Networks     map[string]CoinNetworkInfo `json:"networks"`
}

// Coin is the basic coin structure to get the correct properties for each coin.
type Coin struct {
	Info           CoinInfo `json:"info"`
	Rates          RatesSource
	BlockchainInfo BlockchainInfo
	Mnemonic       string
	NetParams      *chaincfg.Params
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
