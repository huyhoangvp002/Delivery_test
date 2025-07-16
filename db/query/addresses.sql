-- name: CreateAddress :one
INSERT INTO addresses (name, phone, address, status)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetAddressByID :one
SELECT * FROM addresses WHERE id = $1;

-- name: DeleteAddress :exec
DELETE FROM addresses WHERE id = $1;