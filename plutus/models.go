package plutus

//Address is the model for the GET plutus address validation method
type Address struct {
	Valid bool `json:"valid"`
}

//Status is the model for the GET plutus status method
type Status struct {
	Blocks          int  `json:"node_blocks"`
	Headers         int  `json:"node_headers"`
	ExternalBlocks  int  `json:"external_blocks"`
	ExternalHeaders int  `json:"external_headers"`
	SyncStatus      bool `json:"synced"`
}

//Info is the model for the GET plutus info method
type Info struct {
	Blocks      int    `json:"blocks"`
	Headers     int    `json:"headers"`
	Chain       string `json:"chain"`
	Protocol    int    `json:"protocol"`
	Version     int    `json:"version"`
	SubVersion  string `json:"subversion"`
	Connections int    `json:"connections"`
}

//Balance is the model for the GET plutus balance method
type Balance struct {
	Confirmed   float64 `json:"confirmed"`
	Unconfirmed float64 `json:"unconfirmed"`
}

//Transaction is the model for the GET plutus transaction method
type Transaction struct {
	Blockhash     string `json:"blockhash"`
	Blocktime     int    `json:"blocktime"`
	Confirmations int    `json:"confirmations"`
	Height        int    `json:"height"`
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
		Sequence int64  `json:"sequence"`
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
		ValueSat int64   `json:"valueSat"`
	} `json:"vout"`
}

//DecodedRawTX is the model for the verify RAW transaction
type DecodedRawTX struct {
	Locktime int    `json:"locktime"`
	Size     int    `json:"size"`
	Txid     string `json:"txid"`
	Version  int    `json:"version"`
	Vin      []struct {
		ScriptSig struct {
			Asm string `json:"asm"`
			Hex string `json:"hex"`
		} `json:"scriptSig"`
		Sequence int64  `json:"sequence"`
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

type SendAddressBodyReq struct {
	Address string  `json:"address"`
	Coin    string  `json:"coin"`
	Amount  float64 `json:"amount"`
}

type SendAddressInternalBodyReq struct {
	Coin   string  `json:"coin"`
	Amount float64 `json:"amount"`
}
