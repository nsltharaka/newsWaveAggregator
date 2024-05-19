-- name: CreatePost :one
INSERT INTO
    posts (
        post_id,
        title,
        description,
        author,
        pub_date,
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
        $8
    )
RETURNING
    *;