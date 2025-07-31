-- +goose Up
CREATE TABLE follows(
	follow_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	followee_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE follows;
