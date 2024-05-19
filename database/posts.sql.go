// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: posts.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
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
    post_id, title, description, author, pub_date, post_image, url, feed_id
`

type CreatePostParams struct {
	PostID      uuid.UUID      `json:"post_id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	Author      sql.NullString `json:"author"`
	PubDate     time.Time      `json:"pub_date"`
	PostImage   sql.NullString `json:"post_image"`
	Url         string         `json:"url"`
	FeedID      uuid.UUID      `json:"feed_id"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.PostID,
		arg.Title,
		arg.Description,
		arg.Author,
		arg.PubDate,
		arg.PostImage,
		arg.Url,
		arg.FeedID,
	)
	var i Post
	err := row.Scan(
		&i.PostID,
		&i.Title,
		&i.Description,
		&i.Author,
		&i.PubDate,
		&i.PostImage,
		&i.Url,
		&i.FeedID,
	)
	return i, err
}
