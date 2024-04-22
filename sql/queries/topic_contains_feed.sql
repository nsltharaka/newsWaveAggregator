-- name: CreateTopicContainsFeed :one
INSERT INTO topic_contains_feed(topic_id, feed_id, user_id)
VALUES($1, $2, $3)
RETURNING *;
-- name: GetTopicContainsFeed :one
SELECT *
FROM topic_contains_feed
WHERE feed_id = $1
    AND topic_id = $2;