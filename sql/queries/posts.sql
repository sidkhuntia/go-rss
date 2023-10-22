-- name: CreatePost :one
INSERT INTO
    posts(
        id,
        createdAt,
        updatedAt,
        title,
        description,
        url,
        feedId,
        publishedAt
    )
VALUES
    (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8
    ) RETURNING *;

-- name: GetPostsForUser :many
SELECT
    p.*
FROM
    posts p
WHERE
    p.feedId IN (
        SELECT
            f.feedId
        FROM
            feed_follows f
        WHERE
            f.userId = $1
    )
ORDER BY p.publishedAt DESC
LIMIT $2;