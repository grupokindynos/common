package plutus

type PlutusInfo struct {
	Blocks          int  `json:"node_blocks"`
	Headers         int  `json:"node_headers"`
	ExternalBlocks  int  `json:"external_blocks"`
	ExternalHeaders int  `json:"external_headers"`
	SyncStatus      bool `json:"synced"`
}
