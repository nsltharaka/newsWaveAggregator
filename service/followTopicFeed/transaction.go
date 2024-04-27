package followTopicFeed

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nsltharaka/newsWaveAggregator/database"
	"github.com/nsltharaka/newsWaveAggregator/types"
)

func (h *Handler) performTransaction(r *http.Request, payload *types.IncomingFollowTopicFeedPayload, userID int) {

	// if the topic is new, add the topic, or get the existing id
	topicID := h.getOrInsertTopic(r, payload.Topic, userID)

	// make user follow the created or existing topic
	h.insertUserFollowTopicOrFail(r, topicID, userID)

	for _, url := range payload.FeedURLs {

		// if the feed doesn't exist, create one and make topic contains the feed
		existingFeed, err := h.db.GetFeedByURL(r.Context(), url)
		if err != nil {
			feedID := h.insertFeed(r, url)
			h.insertTopicContainsFeedOrFail(r, topicID, feedID, userID)
			continue
		}

		// if the feed exists, make the feed associated with the topic
		h.insertTopicContainsFeedOrFail(r, topicID, existingFeed.ID, userID)
	}
}

func (h *Handler) getOrInsertTopic(r *http.Request, name string, userID int) uuid.UUID {

	topicID := uuid.New()

	topic, err := h.db.GetTopicByName(r.Context(), name)
	if err != nil {

		_, _ = h.db.CreateTopic(r.Context(), database.CreateTopicParams{
			ID:        topicID,
			Name:      name,
			ImgUrl:    sql.NullString{},
			CreatedBy: int32(userID),
			UpdatedAt: time.Now().UTC(),
		})

	} else {
		topicID = topic.ID
	}

	return topicID

}

func (h *Handler) insertUserFollowTopicOrFail(r *http.Request, topicID uuid.UUID, userID int) {
	_, _ = h.db.CreateUserFollowTopic(r.Context(), database.CreateUserFollowTopicParams{
		UserID:  int32(userID),
		TopicID: topicID,
	})
}

func (h *Handler) insertFeed(r *http.Request, url string) uuid.UUID {
	feedID := uuid.New()
	_, _ = h.db.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        feedID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       url,
	})
	return feedID
}

func (h *Handler) insertTopicContainsFeedOrFail(r *http.Request, topicID, feedID uuid.UUID, user_id int) {
	_, _ = h.db.CreateTopicContainsFeed(r.Context(), database.CreateTopicContainsFeedParams{
		TopicID: topicID,
		FeedID:  feedID,
		UserID:  int32(user_id),
	})
}
