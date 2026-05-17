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
	query := `
		SELECT
			date_trunc('day', block_time)::date AS day,
			COALESCE(SUM(CASE WHEN event_type = 'request' THEN amount_normalized::numeric ELSE 0 END), 0)::double precision AS requested,
			COALESCE(SUM(CASE WHEN event_type = 'request' AND amount_normalized > 0 THEN 1 ELSE 0 END), 0)::double precision AS request_count,
			COALESCE(SUM(CASE WHEN event_type = 'withdraw' THEN amount_normalized::numeric ELSE 0 END), 0)::double precision AS withdrawn
		FROM raw_withdraw_flows
		GROUP BY 1
		ORDER BY 1 ASC
	`

	var rows []dailyFlowAggRow
	if err := r.db.WithContext(ctx).Raw(query).Scan(&rows).Error; err != nil {
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
