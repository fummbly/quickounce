-- name: CreatePhoto :one
INSERT INTO photos(id, user_id, post_id, photo_path, created_at, updated_at)
VALUES(
	gen_random_uuid(),
	$1,
	$2,
	$3,
	NOW(),
	NOW()
	)
RETURNING *;
