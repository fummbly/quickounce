-- name: CreateFollow :one
INSERT INTO follows(follow_id, followee_id, created_at)
VALUES (
	$1,
	$2,
	NOW()
)RETURNING *;

-- name: DeleteFollow :exec
DELETE FROM follows
WHERE $1 = follow_id AND $2 = followee_id;



