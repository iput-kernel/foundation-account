-- name: CreateUser :one
INSERT INTO "user" (
  name,
  email,
  password_hash,
  role
)
VALUES (
  $1, $2, $3, $4
) RETURNING *;
