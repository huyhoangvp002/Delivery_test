-- name: CreateAccount :one
INSERT INTO accounts (
  username,
  password,
  role,
  created_at
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetAccountByID :one
SELECT * FROM accounts WHERE id = $1;

-- name: GetAccountByUsername :one
SELECT * FROM accounts WHERE username = $1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts
SET
  username = $2,
  password = $3,
  role = $4
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;

-- name: GetPasswordByUsername :one

SELECT password FROM accounts WHERE username = $1;

-- name: GetAccountIDByUsername :one

SELECT id FROM accounts WHERE username = $1;