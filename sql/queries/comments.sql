-- name: CreateComment :one
INSERT INTO comments (id, post_id, user_id, comment_text, created_at, updated_at)
VALUES(
	gen_random_uuid(),
	$1,
	$2,
	$3,
	NOW(),
	NOW()
)RETURNING *;

-- name: GetCommentsByPost :many
SELECT * FROM comments
WHERE post_id = $1;
