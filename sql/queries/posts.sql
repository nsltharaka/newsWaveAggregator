-- name: CreatePost :one
INSERT INTO
    posts (
        post_id,
        title,
        description,
        author,
        pub_date,
        fetched_at,
        post_image,
        url,
        feed_id
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9
    )
RETURNING
    *;

-- name: GetAllTopicsWithLimitAndOffset :many
SELECT p.*, f.url AS feed_url, t.name AS topic_name
FROM
    posts p
    INNER JOIN feeds f ON p.feed_id = f.id
    INNER JOIN topic_contains_feed tcf ON f.id = tcf.feed_id
    INNER JOIN user_follows_topic uft ON tcf.topic_id = uft.topic_id
    INNER JOIN topics t ON tcf.topic_id = t.id
WHERE
    uft.user_id = $1
ORDER BY p.pub_date DESC -- Order by latest fetched posts first (optional)
LIMIT $2
OFFSET
    $3;

-- name: GetAllPostsForTopic :many
SELECT p.*, f.url AS feed_url
FROM
    posts p
    INNER JOIN feeds f ON p.feed_id = f.id
    INNER JOIN topic_contains_feed tcf ON f.id = tcf.feed_id
    INNER JOIN user_follows_topic uft ON tcf.topic_id = uft.topic_id
WHERE
    uft.user_id = $1
    AND tcf.topic_id = $2
ORDER BY p.pub_date DESC -- Order by latest posts first (optional)
LIMIT $3
OFFSET
    $4;