-- name: CreateTopic :one
INSERT INTO topics (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetTopicByName :one
SELECT *
FROM topics
WHERE name = $1;

-- name: DeleteTopicByID :one
DELETE FROM topics WHERE id = $1 RETURNING *;