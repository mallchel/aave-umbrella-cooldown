package backend

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (s *Server) GetHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, HealthResponse{Status: "ok"})
}

func (s *Server) ListWithdrawFlows(w http.ResponseWriter, r *http.Request, params ListWithdrawFlowsParams) {
	httpFilter, err := parseListWithdrawFlowsParams(params)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		return
	}

	rows, err := s.repo.ListWithdrawFlows(r.Context(), httpFilter.toStorageFilter())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, ApiError{Error: "query withdraw flows"})
		return
	}

	items := make([]WithdrawFlowItem, 0, len(rows))
	for _, row := range rows {
		items = append(items, WithdrawFlowItem{
			ChainId:          row.ChainID,
			TxHash:           row.TxHash,
			LogIndex:         row.LogIndex,
			BlockNumber:      row.BlockNumber,
			BlockTime:        row.BlockTime.UTC(),
			SenderAddress:    row.SenderAddress,
			EventType:        WithdrawFlowItemEventType(row.EventType),
			AmountRaw:        row.AmountRaw,
			AmountNormalized: row.AmountNormalized,
			UpdatedAt:        row.UpdatedAt.UTC(),
		})
	}

	writeJSON(w, http.StatusOK, ListWithdrawFlowsResponse{
		Items: items,
		Meta: ListWithdrawFlowsMeta{
			Count:  len(items),
			Limit:  httpFilter.Limit,
			Offset: httpFilter.Offset,
		},
	})
}

func (s *Server) RenderChartPage(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) ListDailySeriesData(w http.ResponseWriter, r *http.Request) {
	points, err := s.repo.ListDailyFlowPoints(r.Context())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, ApiError{Error: "query daily series data"})
		return
	}

	items := make([]DailySeriesDataPoint, 0, len(points))
	for _, point := range points {
		items = append(items, DailySeriesDataPoint{
			Day:          openapi_types.Date{Time: point.Day.UTC()},
			Requested:    point.Requested,
			Withdrawn:    point.Withdrawn,
			RequestCount: point.RequestCount,
		})
	}

	writeJSON(w, http.StatusOK, items)
}
