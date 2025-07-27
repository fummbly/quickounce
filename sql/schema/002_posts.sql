-- +goose Up
CREATE TABLE posts(
	id UUID PRIMARY KEY,
	total_likes INT,
	total_comments INT,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	upload_id BYTEA,
	user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;
