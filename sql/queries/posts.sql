-- name: CreatePost :one
INSERT INTO posts(id, user_id, total_likes, total_comments, created_at, updated_at)
VALUES(
	gen_random_uuid(),
	$1,
	0,
	0,
	NOW(),
	NOW()
)
RETURNING *;
