-- name: CreateShipmentStatusLog :one
INSERT INTO shipment_status_logs (shipment_id, status, note)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetShipmentStatusLog :one
SELECT * FROM shipment_status_logs WHERE id = $1;

-- name: ListShipmentStatusLogs :many
SELECT * FROM shipment_status_logs ORDER BY id LIMIT $1 OFFSET $2;

-- name: ListLogsByShipmentID :many
SELECT * FROM shipment_status_logs WHERE shipment_id = $1 ORDER BY created_at DESC;

-- name: DeleteShipmentStatusLog :exec
DELETE FROM shipment_status_logs WHERE id = $1;