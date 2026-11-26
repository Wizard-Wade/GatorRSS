-- name: GetNextFeedToFetch :many
SELECT * FROM feeds ORDER BY last_fetch_at ASC NULLS FIRST;

-- name: UpdateFeed :exec
UPDATE feeds SET last_fetch_at = NOW(), updated_at = NOW() WHERE id = $1;
