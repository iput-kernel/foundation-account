-- name: CreateFollow :one
INSERT INTO follows (
  following_user_id,
  followed_user_id
) VALUES (
  $1, $2
) RETURNING *;


-- name: ListFollows :many
SELECT * FROM follows
WHERE 
    following_user_id = $1 OR
    followed_user_id = $2
ORDER BY created_at
LIMIT $3
OFFSET $4;