-- name: CreateFeed :one
INSERT INTO
    feeds (id, created_at, url)
VALUES ($1, $2, $3)
RETURNING
    *;
-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url = $1;

-- name: GetFeedsForUserTopic :many
SELECT f.id, f.url
FROM
    feeds f
    INNER JOIN topic_contains_feed tcf ON tcf.feed_id = f.id
WHERE
    tcf.topic_id = $1
    AND tcf.user_id = $2;

-- name: GetAllFeedsGroupedByTopic :many
SELECT t.name AS topic_name, f.url
FROM
    Topics t
    INNER JOIN topic_contains_feed tcf ON t.id = tcf.topic_id
    INNER JOIN feeds f ON tcf.feed_id = f.id
WHERE
    tcf.user_id = $1
GROUP BY
    t.name,
    f.url
ORDER BY t.name;

-- name: GetAllFeedsForUser :many
SELECT f.*
FROM
    feeds f
    INNER JOIN topic_contains_feed tcf ON f.id = tcf.feed_id
    INNER JOIN user_follows_topic uft ON uft.topic_id = tcf.topic_id
WHERE
    uft.user_id = $1;

-- name: GetNextFeedsToFetch :many
select * from feeds order by updated_at asc nulls first limit $1;

-- name: MarkFeedAsFetched :exec
UPDATE feeds SET updated_at = now() WHERE id = $1;