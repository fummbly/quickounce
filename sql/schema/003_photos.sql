-- +goose Up 
CREATE TABLE photos(
	id UUID PRIMARY KEY,
	photo_path TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	post_id UUID NOT NULL REFERENCES posts(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;
