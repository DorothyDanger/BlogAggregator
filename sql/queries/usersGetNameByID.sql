-- name: GetNameByID :one
SELECT name
FROM users
WHERE ID = $1;