-- name: CreateTopic :one
INSERT INTO topics (id, name, img_url, created_by, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: GetTopicByName :one
SELECT *
FROM topics
WHERE name = $1;

-- name: UpdateTopicImage :one
UPDATE topics SET img_url = $1 WHERE name = $2 RETURNING *;