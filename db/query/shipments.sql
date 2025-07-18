
-- name: CreateShipment :one
INSERT INTO shipments (client_id, from_address_id, to_address_id, shipper_id, shipment_code, fee, status)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetShipment :one
SELECT * FROM shipments WHERE id = $1;

-- name: ListShipments :many
SELECT * FROM shipments ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateShipment :one
UPDATE shipments
SET client_id = $2, from_address_id = $3, to_address_id = $4, shipper_id = $5, shipment_code = $6, fee = $7, status = $8, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteShipment :exec
DELETE FROM shipments WHERE id = $1;

-- name: ListShipmentsByClient :many
SELECT * FROM shipments WHERE client_id = $1 ORDER BY id LIMIT $2 OFFSET $3;

-- name: CheckShipmentCodeExists :one
SELECT EXISTS (
  SELECT 1 FROM shipments WHERE shipment_code = $1
);

-- name: UpdateShipmentStatus :one
UPDATE shipments SET status = $2, updated_at = $3 WHERE shipment_code = $1 RETURNING shipment_code,status,updated_at;