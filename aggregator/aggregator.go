package aggregator

import (
	"context"
	"database/sql"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
	"github.com/nsltharaka/newsWaveAggregator/database"
	"github.com/sym01/htmlsanitizer"
)

// StartAggregation is a async function starts the aggregation functionality. `limit` controls the number of feeds for aggregation in one batch. `timeDuration` is the interval between each batch for aggregation.
func StartAggregation(limit int, duration time.Duration) {

	// separate database connection for aggregator
	dbUrl := os.Getenv("DB_URL")
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("error with DSN: %v", err)
	}
	// check connection
	if err := conn.Ping(); err != nil {
		log.Fatalf("database connection failed in the aggregator: %v", err)
	}

	db := database.New(conn)
	ticker := time.NewTicker(duration)

	for ; ; <-ticker.C {
		log.Println("AGGREGATOR running...")
		// aggregation starts by getting all the feeds that are ready to get updated.
		// selected feed count is limited by the `limit`
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(limit))
		if err != nil {
			log.Println("error getting feeds from db ", err)
			log.Println("AGGREGATOR waiting for next tick...")
			continue
		}

		if len(feeds) == 0 {
			log.Printf("AGGREGATOR %d feeds to update\n", len(feeds))
			log.Println("AGGREGATOR waiting for next tick...")
			continue
		}

		//  for each feed, aggregate posts in parallel.
		// if error in fetching, continue
		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			// one routine per feed for efficiency
			wg.Add(1)
			go ScrapeFeeds(wg, db, feed)

		}

		wg.Wait()
		log.Println("AGGREGATOR scraping done")
		log.Println("AGGREGATOR waiting for next tick...")
	}

}

func ScrapeFeeds(wg *sync.WaitGroup, db *database.Queries, feed database.Feed) {
	defer wg.Done()

	// mark the feed as fetched
	if err := db.MarkFeedAsFetched(context.Background(), feed.ID); err != nil {
		log.Println("error marking feed as fetched ", err)
		return
	}

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	rssParser := gofeed.NewParser()
	rssFeed, err := rssParser.ParseURLWithContext(feed.Url, ctxWithTimeout)
	if err != nil {
		log.Println("error parsing url : ", feed.Url)
		return
	}

	log.Printf("AGGREGATOR found %d posts under %s\n", len(rssFeed.Items), feed.Url)

	for _, item := range rssFeed.Items {
		// save each in the db

		description := sql.NullString{}
		if item.Description != "" {
			description.String = sanitizeDescription(item.Description)
			description.Valid = true
		}

		authors := sql.NullString{}
		if len(item.Authors) > 0 {
			authors.String = item.Authors[0].Name
			authors.Valid = true
		}

		imgUrl := sql.NullString{}
		if item.Image != nil {
			imgUrl.String = item.Image.URL
			imgUrl.Valid = true
		}

		if _, err := db.CreatePost(context.Background(), database.CreatePostParams{
			PostID:      uuid.New(),
			Title:       item.Title,
			Description: description,
			Author:      authors,
			PubDate:     *item.PublishedParsed,
			FetchedAt:   time.Now().UTC(),
			PostImage:   imgUrl,
			Url:         item.Link,
			FeedID:      feed.ID,
		}); err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println(err)
		}
	}

}

func sanitizeDescription(description string) string {
	s := htmlsanitizer.NewHTMLSanitizer()
	// just set AllowList to nil to disable all tags
	s.AllowList = nil

	sanitizedHTML, _ := s.SanitizeString(description)
	return strings.TrimSpace(sanitizedHTML)
}
