-- name: CreateApiKey :one
INSERT INTO api_keys (client_id, api_key, revoked, expires_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetApiKey :one
SELECT * FROM api_keys WHERE id = $1;

-- name: ListApiKeys :many
SELECT * FROM api_keys ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateApiKey :one
UPDATE api_keys
SET api_key = $2, revoked = $3, expires_at = $4, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteApiKey :exec
DELETE FROM api_keys WHERE id = $1;

-- name: ListApiKeysByClientID :many
SELECT * FROM api_keys WHERE client_id = $1 ORDER BY id LIMIT $2 OFFSET $3;