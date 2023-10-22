-- name: CreateFeed :one
INSERT INTO
    feeds(id, createdAt, updatedAt, name, url, userId)
VALUES
    (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6
    ) RETURNING *;

-- name: GetFeeds :many
SELECT
    *
FROM
    feeds;

-- name: GetNextFeedsToFetch :many
SELECT
    *
FROM
    feeds
ORDER BY
    lastFetchedAt ASC NULLS FIRST
LIMIT
    $1;

-- name: MarkFeedAsFetched :one
UPDATE
    feeds
SET
    lastFetchedAt = NOW(), 
    updatedAt = NOW()
WHERE
    id = $1 RETURNING *;