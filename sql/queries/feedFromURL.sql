-- name: GetFeedFromURL :one
SELECT *
FROM feed
WHERE url = $1;