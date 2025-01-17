-- name: CreateUser :one
INSERT INTO "users" (
  id,
  name,
  email,
  password_hash,
  role,
  credit
)
VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByName :one
SELECT * FROM users
WHERE name = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET 
  name = COALESCE(sqlc.narg(name), name),
  email = COALESCE(sqlc.narg(email), email),
  password_hash = COALESCE(sqlc.narg(password_hash), password_hash),
  credit = COALESCE(sqlc.narg(credit), credit)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: AddUserCredit :one
UPDATE users
SET 
  credit = credit + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;