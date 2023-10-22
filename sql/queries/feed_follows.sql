-- name: CreateFeedFollow :one
INSERT INTO feed_follows(id, createdAt, updatedAt, userId, feedId )
VALUES
    (
        $1,
        $2,
        $3,
        $4,
        $5
    ) RETURNING *;

-- name: GetFeedFollows :many
SELECT * FROM feed_follows WHERE userId = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE userId = $1 AND id = $2;