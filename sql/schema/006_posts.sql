-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    url TEXT UNIQUE NOT NULL,
    feedId UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    publishedAt TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE posts;