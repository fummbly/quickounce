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

-- name: GetPostsByUserID :many
SELECT * FROM posts
WHERE user_id = $1
ORDER BY created_at ASC;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;


-- name: GetFollowPosts :many
SELECT posts.* FROM posts
JOIN follows ON posts.user_id = follows.follow_id
WHERE follows.follow_id = $1;

