package feed

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nsltharaka/newsWaveAggregator/database"
	"github.com/nsltharaka/newsWaveAggregator/types"
)

func performFeedTransaction(r *http.Request, db *database.Queries, userId int32, payload *types.CreateFeedPayload) {

	topicId := uuid.New()

	existingTopic, err := db.GetTopicByName(r.Context(), payload.Topic)
	if err == nil {
		topicId = existingTopic.ID
	}

	db.CreateTopic(r.Context(), database.CreateTopicParams{
		ID:   topicId,
		Name: payload.Topic,
	})

	db.CreateUserTopic(r.Context(), database.CreateUserTopicParams{
		ID:      uuid.New(),
		UserID:  userId,
		TopicID: topicId,
	})

	for _, u := range payload.FeedURLs {

		existingFeed, err := db.GetFeedByURL(r.Context(), u)
		if err == nil && existingFeed.UserID == userId {
			continue
		}

		feedId := uuid.New()

		db.CreateFeed(r.Context(), database.CreateFeedParams{
			ID:        feedId,
			Url:       u,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			TopicID:   topicId,
			UserID:    userId,
		})

	}

}
