-- name: CreatePost :one
INSERT INTO posts(id, user_id, created_at, updated_at, image_url)
VALUES(
	gen_random_uuid(),
	$1,
	NOW(),
	NOW(),
	$2
)
RETURNING *;
