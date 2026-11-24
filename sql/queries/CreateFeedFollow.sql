-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS 
( 
	INSERT INTO feed_follows (id, 
		created_at, 
		updated_at, 
		user_id, 
		feed_id) 
	Values(
		$1,
		$2,
		$3,
		$4,
		$5
	)
	RETURNING *
)

SELECT inserted_feed_follow.*,
users.name AS user_name,
feeds.name AS feed_name
FROM inserted_feed_follow
INNER JOIN users ON inserted_feed_follow.user_id = users.id
INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
WITH found_follows AS 
(
	SELECT * 
	FROM feed_follows 
	WHERE feed_follows.user_id = $1
)

SELECT found_follows.*,
users.name AS user_name,
feeds.name AS feed_name
FROM found_follows
INNER JOIN users ON found_follows.user_id = users.id
INNER JOIN feeds ON found_follows.feed_id = feeds.id;

-- name: DeleteFollow :exec
DELETE FROM feed_follows 
WHERE feed_follows.user_id == $1 
AND feed_follows.feed_id == $2;