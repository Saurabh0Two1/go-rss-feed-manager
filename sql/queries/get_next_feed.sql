-- name: GetNextFeed :one

SELECT * FROM feeds
    ORDER BY last_fetched_at DESC 
    NULLS FIRST
    LIMIT 1;

