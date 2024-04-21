-- name: CreateUserTopic :one
INSERT INTO user_topic(id, user_id, topic_id)
values ($1, $2, $3)
RETURNING *;
-- name: DeleteUserTopic :one
delete from user_topic
WHERE user_id = $1
    and topic_id = $2
RETURNING *;