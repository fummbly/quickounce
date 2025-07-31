-- +goose Up
CREATE TABLE comments (
	id UUID PRIMARY KEY,
	comment_text TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	post_id UUID NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
	user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE comments;
