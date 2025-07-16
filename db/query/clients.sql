-- name: CreateClient :one
INSERT INTO clients (name, contact_email,account_id)
VALUES ($1, $2,$3)
RETURNING *;

-- name: GetClient :one
SELECT * FROM clients WHERE id = $1;

-- name: ListClients :many
SELECT * FROM clients ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateClient :one
UPDATE clients
SET name = $2, contact_email = $3
WHERE id = $1
RETURNING *;

-- name: DeleteClient :exec
DELETE FROM clients WHERE id = $1;

-- name: GetClientIDByAccountID :one 
SELECT id FROM clients WHERE account_id = $1;
