package topic

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/nsltharaka/newsWaveAggregator/database"
	"github.com/nsltharaka/newsWaveAggregator/lib/topicImages"
	"github.com/nsltharaka/newsWaveAggregator/types"
)

// updating the topic goes three ways.

// either new topic is same as editing topic
//     - in this treat as existing topic.
// or new topic already in the database
//     - in this case user wants to follow the topic.
// or new topic is totally new
//     - in this case create the new topic and make user follow the topic.
//     - id of newly created topic
// make user unfollow the old topic.

func (h *Handler) handleUpdateTopicLogic(
	ctx context.Context,
	userId int,
	topicId uuid.UUID,
	payload *types.IncomingFollowTopicFeedPayload,
) error {

	h.unfollowTopic(ctx, userId, topicId)

	h.db.DeleteTopicContainsFeed(ctx, database.DeleteTopicContainsFeedParams{
		TopicID: topicId,
		UserID:  int32(userId),
	})

	// create topic if not exist
	createdTopicId := uuid.New()
	existingTopic, err := h.db.GetTopicByName(ctx, payload.Topic)
	if err != nil {
		h.db.CreateTopic(ctx, database.CreateTopicParams{
			ID:        createdTopicId,
			Name:      payload.Topic,
			ImgUrl:    sql.NullString{},
			CreatedBy: int32(userId),
			UpdatedAt: time.Now().UTC(),
		})

		go topicImages.NewImageFinder(topicImages.FromGoogleImages).UpdateTopic(h.db, payload.Topic)

	} else {
		createdTopicId = existingTopic.ID
	}

	h.followTopic(ctx, userId, createdTopicId)

	for _, feed := range payload.FeedURLs {

		newFeedId := uuid.New()
		existingFeedId, err := h.db.GetFeedByURL(ctx, feed)
		if err != nil {
			h.db.CreateFeed(ctx, database.CreateFeedParams{
				ID:        newFeedId,
				CreatedAt: time.Now().UTC(),
				Url:       feed,
			})
		} else {
			newFeedId = existingFeedId.ID
		}

		h.db.CreateTopicContainsFeed(ctx, database.CreateTopicContainsFeedParams{
			TopicID: createdTopicId,
			FeedID:  newFeedId,
			UserID:  int32(userId),
		})

	}

	return nil

}

func (h *Handler) unfollowTopic(ctx context.Context, userId int, topicId uuid.UUID) error {
	return h.db.DeleteUserFollowTopic(ctx, database.DeleteUserFollowTopicParams{
		TopicID: topicId,
		UserID:  int32(userId),
	})
}

func (h *Handler) followTopic(ctx context.Context, userId int, topicId uuid.UUID) error {
	_, err := h.db.CreateUserFollowTopic(ctx, database.CreateUserFollowTopicParams{
		UserID:  int32(userId),
		TopicID: topicId,
	})

	return err
}
