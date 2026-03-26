-- name: CreatePosts :one
INSERT INTO posts (id, title, url, description, feed_id, created_at, updated_at, published_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;