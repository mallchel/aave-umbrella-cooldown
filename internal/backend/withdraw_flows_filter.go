package backend

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"1-task/internal/storage/postgres"
)

const (
	withdrawFlowEventTypeRequest  = "request"
	withdrawFlowEventTypeWithdraw = "withdraw"
)

// listWithdrawFlowsFilter is the HTTP-layer filter model for listing withdraw flows.
type listWithdrawFlowsFilter struct {
	SenderAddress *string
	EventType     *string
	FromTime      *time.Time
	ToTime        *time.Time
	Limit         int
	Offset        int
}

func parseListWithdrawFlowsFilter(r *http.Request) (listWithdrawFlowsFilter, error) {
	q := r.URL.Query()
	filter := listWithdrawFlowsFilter{}

	if sender := strings.TrimSpace(q.Get("sender_address")); sender != "" {
		filter.SenderAddress = &sender
	}

	if eventType := strings.ToLower(strings.TrimSpace(q.Get("event_type"))); eventType != "" {
		switch eventType {
		case withdrawFlowEventTypeRequest, withdrawFlowEventTypeWithdraw:
			filter.EventType = &eventType
		default:
			return filter, badParam("event_type must be one of: request, withdraw")
		}
	}

	if fromTimeText := strings.TrimSpace(q.Get("from_time")); fromTimeText != "" {
		fromTime, err := time.Parse(time.RFC3339, fromTimeText)
		if err != nil {
			return filter, badParam("from_time must be RFC3339")
		}
		filter.FromTime = &fromTime
	}

	if toTimeText := strings.TrimSpace(q.Get("to_time")); toTimeText != "" {
		toTime, err := time.Parse(time.RFC3339, toTimeText)
		if err != nil {
			return filter, badParam("to_time must be RFC3339")
		}
		filter.ToTime = &toTime
	}

	if filter.FromTime != nil && filter.ToTime != nil && filter.FromTime.After(*filter.ToTime) {
		return filter, badParam("from_time must be before or equal to to_time")
	}

	limit := 100
	if limitText := strings.TrimSpace(q.Get("limit")); limitText != "" {
		parsedLimit, err := strconv.Atoi(limitText)
		if err != nil || parsedLimit <= 0 {
			return filter, badParam("limit must be a positive integer")
		}
		limit = min(parsedLimit, 500)
	}

	offset := 0
	if offsetText := strings.TrimSpace(q.Get("offset")); offsetText != "" {
		parsedOffset, err := strconv.Atoi(offsetText)
		if err != nil || parsedOffset < 0 {
			return filter, badParam("offset must be a non-negative integer")
		}
		offset = parsedOffset
	}

	filter.Limit = limit
	filter.Offset = offset

	return filter, nil
}

func (f listWithdrawFlowsFilter) toStorageFilter() postgres.ListWithdrawFlowsFilter {
	return postgres.ListWithdrawFlowsFilter{
		SenderAddress: f.SenderAddress,
		EventType:     f.EventType,
		FromTime:      f.FromTime,
		ToTime:        f.ToTime,
		Limit:         f.Limit,
		Offset:        f.Offset,
	}
}
