-- name: CreateUserTopic :one
INSERT INTO user_topics (
        id,
        created_at,
        updated_at,
        user_id,
        topic_id
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;