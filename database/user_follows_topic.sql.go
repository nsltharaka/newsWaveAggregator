// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user_follows_topic.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createUserFollowTopic = `-- name: CreateUserFollowTopic :one
INSERT INTO
    user_follows_topic (user_id, topic_id)
VALUES ($1, $2)
RETURNING
    user_id, topic_id
`

type CreateUserFollowTopicParams struct {
	UserID  int32     `json:"user_id"`
	TopicID uuid.UUID `json:"topic_id"`
}

func (q *Queries) CreateUserFollowTopic(ctx context.Context, arg CreateUserFollowTopicParams) (UserFollowsTopic, error) {
	row := q.db.QueryRowContext(ctx, createUserFollowTopic, arg.UserID, arg.TopicID)
	var i UserFollowsTopic
	err := row.Scan(&i.UserID, &i.TopicID)
	return i, err
}

const deleteUserFollowTopic = `-- name: DeleteUserFollowTopic :exec
DELETE FROM user_follows_topic WHERE topic_id = $1 AND user_id = $2
`

type DeleteUserFollowTopicParams struct {
	TopicID uuid.UUID `json:"topic_id"`
	UserID  int32     `json:"user_id"`
}

func (q *Queries) DeleteUserFollowTopic(ctx context.Context, arg DeleteUserFollowTopicParams) error {
	_, err := q.db.ExecContext(ctx, deleteUserFollowTopic, arg.TopicID, arg.UserID)
	return err
}
