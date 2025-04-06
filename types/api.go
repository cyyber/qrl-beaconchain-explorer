package types

type ExecBlockProposer struct {
	ExecBlock uint64 `db:"exec_block_number" json:"executionBlockNumber"`
	Proposer  uint64 `db:"proposer" json:"proposerIndex"`
	Slot      uint64 `db:"slot" json:"slot"`
	Epoch     uint64 `db:"epoch" json:"epoch"`
	Finalized bool   `json:"finalized"`
}

type ApiResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// TODO(now.youtrack.cloud/issue/TZB-1)
/*
type ZnsDomainResponse struct {
	Address string `json:"address"`
	Domain  string `json:"domain"`
}
*/
