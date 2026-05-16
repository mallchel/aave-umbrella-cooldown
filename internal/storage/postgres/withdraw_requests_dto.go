package postgres

import (
	"database/sql"
	"time"
)

// WithdrawRequest is a row model for raw_withdraw_requests.
type WithdrawRequest struct {
	ChainID           int64
	TxHash            string
	LogIndex          int32
	BlockNumber       int64
	BlockTime         time.Time
	UserAddress       string
	AssetSymbol       string
	AmountRaw         string
	AmountNormalized  string
	AmountUSD         string
	CooldownStartTime time.Time
	WithdrawableFrom  time.Time
	WithdrawableUntil time.Time
	Status            string
	WithdrawTxHash    sql.NullString
	UpdatedAt         time.Time
}

// TableName maps this DTO to the schema table.
func (WithdrawRequest) TableName() string {
	return "raw_withdraw_requests"
}
