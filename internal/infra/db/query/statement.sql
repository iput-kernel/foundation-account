-- name: CreateEntry :one
INSERT INTO statements (
  id,
  user_id,
  amount
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetStatements :one
SELECT * FROM statements
WHERE id = $1 LIMIT 1;

-- name: ListStatements :many
SELECT * FROM statements
WHERE user_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;