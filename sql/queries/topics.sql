-- name: CreateTopic :one
INSERT INTO topics (id, name)
VALUES ($1, $2)
RETURNING *;