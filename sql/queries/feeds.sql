-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, url)
VALUES ($1, $2, $3, $4)
RETURNING *;