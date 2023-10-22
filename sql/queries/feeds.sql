-- name: CreateFeed :one
INSERT INTO feeds(id, createdAt, updatedAt, name, url, userId)
VALUES
    (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6
    ) RETURNING *;