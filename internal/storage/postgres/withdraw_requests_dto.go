package postgres

import "time"

const (
	FlowEventTypeRequest  = "request"
	FlowEventTypeWithdraw = "withdraw"
)

// WithdrawFlow is a row model for raw_withdraw_flows.
type WithdrawFlow struct {
	ChainID          int64
	TxHash           string
	LogIndex         int32
	BlockNumber      int64
	BlockTime        time.Time
	SenderAddress    string
	EventType        string
	AmountRaw        string
	AmountNormalized string
	UpdatedAt        time.Time
	CooldownEndAt    int64
}

// TableName maps this DTO to the schema table.
func (WithdrawFlow) TableName() string {
	return "raw_withdraw_flows"
}
