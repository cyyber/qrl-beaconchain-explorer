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

// NOTE(rgeraldes24): unused for now
/*
type ZnsDomainResponse struct {
	Address string `json:"address"`
	Domain  string `json:"domain"`
}
*/
