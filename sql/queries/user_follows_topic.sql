-- name: CreateUserFollowTopic :one
INSERT INTO
    user_follows_topic (user_id, topic_id)
VALUES ($1, $2)
RETURNING
    *;

-- name: DeleteUserFollowTopic :exec
DELETE FROM user_follows_topic WHERE topic_id = $1 AND user_id = $2;