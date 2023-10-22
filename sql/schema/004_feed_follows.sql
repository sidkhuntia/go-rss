-- +goose Up
CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    userId UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feedId UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    UNIQUE (userId, feedId)
);

-- +goose Down
DROP TABLE feed_follows;