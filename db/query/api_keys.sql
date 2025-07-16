-- name: CreateApiKey :one
INSERT INTO api_keys (client_id, api_key)
VALUES ($1, $2)
RETURNING *;

-- name: GetApiKey :one
SELECT * FROM api_keys WHERE id = $1;

-- name: ListApiKeys :many
SELECT * FROM api_keys ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateApiKey :one
UPDATE api_keys
SET api_key = $2, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteApiKey :exec
DELETE FROM api_keys WHERE id = $1;

-- name: ListApiKeysByClientID :many
SELECT * FROM api_keys WHERE client_id = $1 ORDER BY id LIMIT $2 OFFSET $3;

-- name: GetAPIKeyByValue :one
SELECT * FROM api_keys WHERE api_key = $1 LIMIT 1;

-- name: CheckAPIKeyExists :one
SELECT EXISTS (
  SELECT 1 FROM api_keys WHERE api_key = $1
);