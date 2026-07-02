-- name: ListDailyFlowPoints :many
SELECT
  date_trunc('day', block_time)::date AS day,
  COALESCE(SUM(CASE WHEN event_type = 'request' THEN amount_normalized::numeric ELSE 0 END), 0)::double precision AS requested,
  COALESCE(SUM(CASE WHEN event_type = 'request' AND amount_normalized > 0 THEN 1 ELSE 0 END), 0)::double precision AS request_count,
  COALESCE(SUM(CASE WHEN event_type = 'withdraw' THEN amount_normalized::numeric ELSE 0 END), 0)::double precision AS withdrawn,
  COALESCE(SUM(CASE WHEN event_type = 'request' AND NOW() < to_timestamp(cooldown_end_at) THEN amount_normalized::numeric ELSE 0 END), 0)::double precision AS active_cooldown
FROM raw_withdraw_flows
GROUP BY 1
ORDER BY 1 ASC;
