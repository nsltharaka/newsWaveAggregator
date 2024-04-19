package feed

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nsltharaka/newsWaveAggregator/database"
)

func performTopicTransaction(r *http.Request, db *database.Queries, userId int32, topic string) {
	topicId := uuid.New()

	_, _ = db.CreateTopic(r.Context(), database.CreateTopicParams{
		ID:   topicId,
		Name: topic,
	})

	_, _ = db.CreateUserTopic(r.Context(), database.CreateUserTopicParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    userId,
		TopicID:   topicId,
	})
}

func performFeedTransaction(r *http.Request, db *database.Queries, userId int32, urls []string) {
	feedId := uuid.New()
	for _, url := range urls {

		_, _ = db.CreateFeed(r.Context(), database.CreateFeedParams{
			ID:        feedId,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Url:       url,
		})

		_, _ = db.CreateUserFeed(r.Context(), database.CreateUserFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			UserID:    userId,
			FeedID:    feedId,
		})

	}
}
