-- query/ip_log.sql
-- name: InsertIPLog :one
INSERT INTO ip_logs (
    ip_address, country, region, city, latitude, longitude, user_agent, referrer, request_time
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetIPLogByID :one
SELECT * FROM ip_logs WHERE id = $1;

-- name: ListIPLogs :many
SELECT * FROM ip_logs LIMIT $1 OFFSET $2;
