package hestia

import "github.com/grupokindynos/olympus-utils/amount"

type Outputs struct {
	Address string
	Amount  amount.AmountType
}

type PolisTxInfo struct {
	Blockhash     string `json:"blockhash"`
	Blocktime     int    `json:"blocktime"`
	Confirmations int    `json:"confirmations"`
	Height        int    `json:"height"`
	Hex           string `json:"hex"`
	Instantlock   bool   `json:"instantlock"`
	Locktime      int    `json:"locktime"`
	Size          int    `json:"size"`
	Time          int    `json:"time"`
	Txid          string `json:"txid"`
	Type          int    `json:"type"`
	Version       int    `json:"version"`
	Vin           []struct {
		ScriptSig struct {
			Asm string `json:"asm"`
			Hex string `json:"hex"`
		} `json:"scriptSig"`
		Sequence int    `json:"sequence"`
		Txid     string `json:"txid"`
		Vout     int    `json:"vout"`
	} `json:"vin"`
	Vout []struct {
		N            int `json:"n"`
		ScriptPubKey struct {
			Addresses []string `json:"addresses"`
			Asm       string   `json:"asm"`
			Hex       string   `json:"hex"`
			ReqSigs   int      `json:"reqSigs"`
			Type      string   `json:"type"`
		} `json:"scriptPubKey"`
		Value    float64 `json:"value"`
		ValueSat int     `json:"valueSat"`
	} `json:"vout"`
}

type BitcoinTxInfo struct {
	Blockhash     string `json:"blockhash"`
	Blocktime     int    `json:"blocktime"`
	Confirmations int    `json:"confirmations"`
	Hash          string `json:"hash"`
	Hex           string `json:"hex"`
	Locktime      int    `json:"locktime"`
	Size          int    `json:"size"`
	Time          int    `json:"time"`
	Txid          string `json:"txid"`
	Version       int    `json:"version"`
	Vin           []struct {
		ScriptSig struct {
			Asm string `json:"asm"`
			Hex string `json:"hex"`
		} `json:"scriptSig"`
		Sequence    int      `json:"sequence"`
		Txid        string   `json:"txid"`
		Txinwitness []string `json:"txinwitness"`
		Vout        int      `json:"vout"`
	} `json:"vin"`
	Vout []struct {
		N            int `json:"n"`
		ScriptPubKey struct {
			Addresses []string `json:"addresses"`
			Asm       string   `json:"asm"`
			Hex       string   `json:"hex"`
			ReqSigs   int      `json:"reqSigs"`
			Type      string   `json:"type"`
		} `json:"scriptPubKey"`
		Value float64 `json:"value"`
	} `json:"vout"`
	Vsize  int `json:"vsize"`
	Weight int `json:"weight"`
}

type DashTxInfo struct {
	Blockhash           string `json:"blockhash"`
	Blocktime           int    `json:"blocktime"`
	Chainlock           bool   `json:"chainlock"`
	Confirmations       int    `json:"confirmations"`
	Height              int    `json:"height"`
	Hex                 string `json:"hex"`
	Instantlock         bool   `json:"instantlock"`
	InstantlockInternal bool   `json:"instantlock_internal"`
	Locktime            int    `json:"locktime"`
	Size                int    `json:"size"`
	Time                int    `json:"time"`
	Txid                string `json:"txid"`
	Type                int    `json:"type"`
	Version             int    `json:"version"`
	Vin                 []struct {
		ScriptSig struct {
			Asm string `json:"asm"`
			Hex string `json:"hex"`
		} `json:"scriptSig"`
		Sequence int    `json:"sequence"`
		Txid     string `json:"txid"`
		Vout     int    `json:"vout"`
	} `json:"vin"`
	Vout []struct {
		N            int `json:"n"`
		ScriptPubKey struct {
			Addresses []string `json:"addresses"`
			Asm       string   `json:"asm"`
			Hex       string   `json:"hex"`
			ReqSigs   int      `json:"reqSigs"`
			Type      string   `json:"type"`
		} `json:"scriptPubKey"`
		Value    float64 `json:"value"`
		ValueSat int     `json:"valueSat"`
	} `json:"vout"`
}
