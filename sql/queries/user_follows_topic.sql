-- name: CreateUserFollowTopic :one
INSERT INTO user_follows_topic (user_id, topic_id)
VALUES ($1, $2)
RETURNING *;