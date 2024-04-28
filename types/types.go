package types

import (
	"time"

	"github.com/google/uuid"
)

type CanValidated interface {
	LoginUserPayload |
		RegisterUserPayload |
		IncomingFollowTopicFeedPayload
}

/* types associated with user routes */

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=130"`
	ApiKey   string `json:"api_key"`
}

type OutgoingUserPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ApiKey   string `json:"api_key"`
}

/* types associated with followTopicFeed routes */

type IncomingFollowTopicFeedPayload struct {
	Topic    string   `json:"topic" validate:"required"`
	FeedURLs []string `json:"feeds" validate:"required,min=1,dive,required"`
}

/* types related to topic routes */
type OutgoingTopicPayload struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	ImgUrl      string    `json:"img_url"`
	UpdatedAt   time.Time `json:"updated_at"`
	SourceCount int       `json:"source_count"`
}
