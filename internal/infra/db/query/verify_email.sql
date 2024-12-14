-- name: CreateVerifyEmail :one
INSERT INTO verify_emails (
  id,
  name,
  email,
  password_hash,
  secret_code
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetVerifyEmail :one
SELECT * FROM verify_emails
WHERE id = $1 LIMIT 1;

-- name: Verify :one
SELECT * FROM verify_emails
WHERE
  id = @id
  AND secret_code = @secret_code
  AND expired_at > now()
  LIMIT 1;