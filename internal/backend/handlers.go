package backend

import (
	"encoding/json"
	"net/http"
	"time"
)

type apiError struct {
	Error string `json:"error"`
}

type listWithdrawFlowsMeta struct {
	Count  int `json:"count"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type withdrawFlowItem struct {
	ChainID          int64  `json:"chain_id"`
	TxHash           string `json:"tx_hash"`
	LogIndex         int32  `json:"log_index"`
	BlockNumber      int64  `json:"block_number"`
	BlockTime        string `json:"block_time"`
	SenderAddress    string `json:"sender_address"`
	EventType        string `json:"event_type"`
	AmountRaw        string `json:"amount_raw"`
	AmountNormalized string `json:"amount_normalized"`
	UpdatedAt        string `json:"updated_at"`
}

type listWithdrawFlowsResponse struct {
	Items []withdrawFlowItem    `json:"items"`
	Meta  listWithdrawFlowsMeta `json:"meta"`
}

func (s *Server) handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *Server) handleListWithdrawFlows(w http.ResponseWriter, r *http.Request) {
	httpFilter, err := parseListWithdrawFlowsFilter(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		return
	}

	rows, err := s.repo.ListWithdrawFlows(r.Context(), httpFilter.toStorageFilter())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, apiError{Error: "query withdraw flows"})
		return
	}

	items := make([]withdrawFlowItem, 0, len(rows))
	for _, row := range rows {
		items = append(items, withdrawFlowItem{
			ChainID:          row.ChainID,
			TxHash:           row.TxHash,
			LogIndex:         row.LogIndex,
			BlockNumber:      row.BlockNumber,
			BlockTime:        row.BlockTime.UTC().Format(time.RFC3339),
			SenderAddress:    row.SenderAddress,
			EventType:        row.EventType,
			AmountRaw:        row.AmountRaw,
			AmountNormalized: row.AmountNormalized,
			UpdatedAt:        row.UpdatedAt.UTC().Format(time.RFC3339),
		})
	}

	writeJSON(w, http.StatusOK, listWithdrawFlowsResponse{
		Items: items,
		Meta: listWithdrawFlowsMeta{
			Count:  len(items),
			Limit:  httpFilter.Limit,
			Offset: httpFilter.Offset,
		},
	})
}

func writeJSON(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(body)
}

func badParam(message string) error {
	return &paramError{message: message}
}

type paramError struct {
	message string
}

func (e *paramError) Error() string {
	return e.message
}
