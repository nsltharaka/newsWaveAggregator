-- name: CreateTopic :one
INSERT INTO topics (id, name, created_by, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetTopicByName :one
SELECT *
FROM topics
WHERE name = $1;