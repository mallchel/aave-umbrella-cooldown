package postgres

import (
	"context"
	"fmt"
	"time"
)

// DailyFlowPoint stores daily token and request flows for chart bars.
type DailyFlowPoint struct {
	Day          time.Time
	Requested    float64
	RequestCount float64
	Withdrawn    float64
}

type dailyFlowAggRow struct {
	Day          time.Time
	Requested    float64
	RequestCount float64
	Withdrawn    float64
}

// ListDailyFlowPoints returns per-day requested/withdrawn token volume and request counts.
func (r *Repository) ListDailyFlowPoints(ctx context.Context) ([]DailyFlowPoint, error) {
	rows, err := r.q.ListDailyFlowPoints(ctx)
	if err != nil {
		return nil, fmt.Errorf("query daily flow points: %w", err)
	}

	points := make([]DailyFlowPoint, 0, len(rows))
	for _, row := range rows {
		points = append(points, DailyFlowPoint{
			Day:          row.Day.UTC(),
			Requested:    row.Requested,
			RequestCount: row.RequestCount,
			Withdrawn:    row.Withdrawn,
		})
	}

	return points, nil
}
