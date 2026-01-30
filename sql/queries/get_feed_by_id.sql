-- name: GetFeedById :one
SELECT * FROM feeds 
WHERE id = $1;


