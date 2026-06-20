package backend

import (
	"strings"
	"testing"
	"time"
)

func TestParseListWithdrawFlowsParamsNormalizesAndDefaults(t *testing.T) {
	senderAddress := " 0xabc "
	eventType := ListWithdrawFlowsParamsEventType(" Request ")
	fromTime := time.Date(2026, 6, 20, 10, 0, 0, 0, time.UTC)
	limit := 999
	offset := 2

	filter, err := parseListWithdrawFlowsParams(ListWithdrawFlowsParams{
		SenderAddress: &senderAddress,
		EventType:     &eventType,
		FromTime:      &fromTime,
		Limit:         &limit,
		Offset:        &offset,
	})
	if err != nil {
		t.Fatalf("parse params: %v", err)
	}

	if filter.SenderAddress == nil || *filter.SenderAddress != "0xabc" {
		t.Fatalf("sender address = %v, want trimmed 0xabc", filter.SenderAddress)
	}
	if filter.EventType == nil || *filter.EventType != withdrawFlowEventTypeRequest {
		t.Fatalf("event type = %v, want normalized request", filter.EventType)
	}
	if filter.FromTime == nil || !filter.FromTime.Equal(fromTime) {
		t.Fatalf("from time = %v, want %v", filter.FromTime, fromTime)
	}
	if filter.ToTime != nil {
		t.Fatalf("to time = %v, want nil", filter.ToTime)
	}
	if filter.Limit != 500 {
		t.Fatalf("limit = %d, want capped 500", filter.Limit)
	}
	if filter.Offset != offset {
		t.Fatalf("offset = %d, want %d", filter.Offset, offset)
	}
}

func TestParseListWithdrawFlowsParamsValidation(t *testing.T) {
	fromTime := time.Date(2026, 6, 21, 0, 0, 0, 0, time.UTC)
	toTime := time.Date(2026, 6, 20, 0, 0, 0, 0, time.UTC)
	zero := 0
	negative := -1

	tests := []struct {
		name    string
		params  ListWithdrawFlowsParams
		wantErr string
	}{
		{
			name:    "empty sender address",
			params:  ListWithdrawFlowsParams{SenderAddress: ptr("   ")},
			wantErr: "sender_address must not be empty",
		},
		{
			name:    "invalid event type",
			params:  ListWithdrawFlowsParams{EventType: ptr(ListWithdrawFlowsParamsEventType("invalid"))},
			wantErr: "event_type must be one of: request, withdraw",
		},
		{
			name:    "from time after to time",
			params:  ListWithdrawFlowsParams{FromTime: &fromTime, ToTime: &toTime},
			wantErr: "from_time must be before or equal to to_time",
		},
		{
			name:    "non-positive limit",
			params:  ListWithdrawFlowsParams{Limit: &zero},
			wantErr: "limit must be a positive integer",
		},
		{
			name:    "negative offset",
			params:  ListWithdrawFlowsParams{Offset: &negative},
			wantErr: "offset must be a non-negative integer",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := parseListWithdrawFlowsParams(tt.params)
			if err == nil {
				t.Fatal("expected validation error")
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("error = %q, want it to contain %q", err.Error(), tt.wantErr)
			}
		})
	}
}

func ptr[T any](value T) *T {
	return &value
}
