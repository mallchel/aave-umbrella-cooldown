package postgres

import "time"

// IndexerState stores indexer checkpoint data.
type IndexerState struct {
	Name               string
	LastProcessedBlock uint64
	LastProcessedTime  time.Time
	UpdatedAt          time.Time
}

// TableName maps this DTO to the schema table.
func (IndexerState) TableName() string {
	return "indexer_state"
}
