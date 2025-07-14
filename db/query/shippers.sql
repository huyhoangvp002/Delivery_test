-- name: CreateShipper :one
INSERT INTO shippers (name, phone, carrier, active)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetShipper :one
SELECT * FROM shippers WHERE id = $1;

-- name: ListShippers :many
SELECT * FROM shippers ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateShipper :one
UPDATE shippers
SET name = $2, phone = $3, carrier = $4, active = $5
WHERE id = $1
RETURNING *;

-- name: DeleteShipper :exec
DELETE FROM shippers WHERE id = $1;