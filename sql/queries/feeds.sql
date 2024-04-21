-- name: CreateFeed :one
INSERT INTO feeds (
        id,
        created_at,
        updated_at,
        url,
        topic_id,
        user_id
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: GetFeedByURL :one
SELECT *
FROM feeds
WHERE url = $1;
-- name: GetAllFeedsForTopic :many
SELECT DISTINCT f.url,
    t.name
FROM feeds f
    JOIN user_topic ut ON f.user_id = ut.user_id
    JOIN topics t ON ut.topic_id = t.id
WHERE t.name = $1;
-- name: GetAllFeedsForTopicID :many
SELECT DISTINCT f.url
FROM feeds f
WHERE f.topic_id = $1;