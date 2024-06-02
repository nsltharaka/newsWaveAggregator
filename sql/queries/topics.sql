-- name: CreateTopic :one
INSERT INTO
    topics (
        id,
        name,
        img_url,
        created_by,
        updated_at
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING
    *;
-- name: GetTopicByName :one
SELECT * FROM topics WHERE name = $1;

-- name: GetTopic :one
SELECT * FROM topics WHERE id = $1;

-- name: UpdateTopicImage :one
UPDATE topics SET img_url = $1 WHERE name = $2 RETURNING *;

-- name: GetAllTopicsForUserWithSourceCount :many
SELECT t.*, COUNT(DISTINCT tcf.feed_id) AS feed_count
FROM
    Topics t
    INNER JOIN User_Follows_Topic uft ON t.id = uft.topic_id
    LEFT JOIN Topic_Contains_Feed tcf ON t.id = tcf.topic_id
    AND uft.user_id = tcf.user_id -- Use LEFT JOIN for optional matching
WHERE
    uft.user_id = $1
GROUP BY
    t.id
ORDER BY t.updated_at DESC;

-- name: GetTopicsCount :one
SELECT COUNT(*) AS total_topics
FROM
    user_follows_topic uft
    INNER JOIN topics t ON uft.topic_id = t.id
WHERE
    uft.user_id = $1;

-- name: GetTopicsLike :many
SELECT t.*
FROM
    topics t
    INNER JOIN user_follows_topic uft on uft.topic_id = t.id
WHERE
    uft.user_id = $1
    AND LOWER(t.name) LIKE LOWER('%' || $2 || '%');