-- name: CreateShipper :one
INSERT INTO shippers (name, phone, active)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetShipper :one
SELECT * FROM shippers WHERE id = $1;

-- name: ListShippers :many
SELECT * FROM shippers ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateShipper :one
UPDATE shippers
SET name = $2, phone = $3, active = $4
WHERE id = $1
RETURNING *;

-- name: DeleteShipper :exec
DELETE FROM shippers WHERE id = $1;