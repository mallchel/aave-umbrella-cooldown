package backend

import (
	"fmt"
	"html/template"
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

type dailySeriesDataPoint struct {
	Day          string  `json:"day"`
	Requested    float64 `json:"requested"`
	Withdrawn    float64 `json:"withdrawn"`
	RequestCount float64 `json:"request_count"`
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

func (s *Server) handleRenderChart(w http.ResponseWriter, r *http.Request) {
	points, err := s.repo.ListDailyFlowPoints(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("query chart data: %v", err), http.StatusInternalServerError)
		return
	}

	requestedSvg, err := buildRequestedChartSVG(points)
	if err != nil {
		http.Error(w, fmt.Sprintf("render chart: %v", err), http.StatusInternalServerError)
		return
	}
	withdrawnSvg, err := buildWithdrawnChartSVG(points)
	if err != nil {
		http.Error(w, fmt.Sprintf("render chart: %v", err), http.StatusInternalServerError)
		return
	}
	requestCountSvg, err := buildRequestCountChartSVG(points)
	if err != nil {
		http.Error(w, fmt.Sprintf("render chart: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := chartPageTpl.Execute(w, chartPageData{RenderedAt: time.Now(), RequestedSVG: template.HTML(requestedSvg), WithdrawnSVG: template.HTML(withdrawnSvg), RequestCountSVG: template.HTML(requestCountSvg)}); err != nil {
		http.Error(w, fmt.Sprintf("render page: %v", err), http.StatusInternalServerError)
		return
	}
}

func (s *Server) handleDailySeriesData(w http.ResponseWriter, r *http.Request) {
	points, err := s.repo.ListDailyFlowPoints(r.Context())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, apiError{Error: "query daily series data"})
		return
	}

	items := make([]dailySeriesDataPoint, 0, len(points))
	for _, point := range points {
		items = append(items, dailySeriesDataPoint{
			Day:          point.Day.UTC().Format("2006-01-02"),
			Requested:    point.Requested,
			Withdrawn:    point.Withdrawn,
			RequestCount: point.RequestCount,
		})
	}

	writeJSON(w, http.StatusOK, items)
}
