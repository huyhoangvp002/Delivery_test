-- name: CreateAddress :one
INSERT INTO addresses (name, phone, street, ward, district, city, country)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetAddress :one
SELECT * FROM addresses WHERE id = $1;

-- name: ListAddresses :many
SELECT * FROM addresses ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateAddress :one
UPDATE addresses
SET name = $2, phone = $3, street = $4, ward = $5, district = $6, city = $7, country = $8
WHERE id = $1
RETURNING *;

-- name: DeleteAddress :exec
DELETE FROM addresses WHERE id = $1;