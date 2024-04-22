-- name: CreateTopic :one
INSERT INTO topics (id, name, created_by)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetTopicByName :one
SELECT *
FROM topics
WHERE name = $1;