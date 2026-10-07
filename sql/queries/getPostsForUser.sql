-- name: GetPostsForUser :many
SELECT posts.*, feed.name AS feed_name
FROM posts
JOIN feed_follows ON feed_follows.feed_id = posts.feed_id
JOIN feed ON posts.feed_id = feed.id
WHERE feed_follows.user_id = $1
ORDER BY published_at DESC
LIMIT $2;