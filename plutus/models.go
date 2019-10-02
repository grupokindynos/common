package plutus

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
