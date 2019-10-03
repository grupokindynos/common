package plutus

//Plutus is the model for the response from the Plutus microservice
type Plutus struct {
	Data   interface{} `json:"data"`
	Status int64       `json:"status"`
}

type PlutusEncoded struct {
	Data   string `json:"data"`
	Status int64  `json:"status"`
}

//PlutusAddress is the model for the response from the Plutus address microservice
type PlutusAddress struct {
	Data string `json:"address"`
}

//Address is the model for the GET plutus address validation method
type ValidAddress struct {
	Valid bool `json:"valid"`
}

//Address is the model for the response from the hot-wallets address
type Address struct {
	Address string `json:"address"`
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

//PlutusBalance is the model for the response from the Plutus address microservice
type PlutusBalance struct {
	Data   Balance "json:data"
	Status int64   "json:status"
}

//Balance is the model for the hot-wallets balance
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
