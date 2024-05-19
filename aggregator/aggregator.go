package aggregator

import (
	"context"
	"database/sql"
	"log"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/nsltharaka/newsWaveAggregator/database"
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
		// aggregation starts by getting all the feeds that are ready to get updated.
		// selected feed count is limited by the `limit`
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(limit))
		if err != nil {
			log.Println("error getting feeds from db ", err)
			continue
		}

		//  for each feed, aggregate posts in parallel.
		// if error in fetching, continue
		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)
			// one routine per feed for efficiency
			go ScrapeFeeds(wg, db, feed)

		}

		println("receiving posts...")
		wg.Wait()
	}

}

func ScrapeFeeds(wg *sync.WaitGroup, db *database.Queries, feed database.Feed) {
	defer wg.Done()

	// mark the feed as fetched
	if err := db.MarkFeedAsFetched(context.Background(), feed.ID); err != nil {
		log.Println("error marking feed as fetched ", err)
	}

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	rssParser := gofeed.NewParser()
	rssFeed, err := rssParser.ParseURLWithContext(feed.Url, ctxWithTimeout)
	if err != nil {
		log.Println("error parsing url : ", feed.Url)
	}

	for _, item := range rssFeed.Items {
		// save each in the db
		_ = item
		slog.Info("FEED : ", "url", feed.Url)
		slog.Info("FEED ITEM : ", "Title", item.Title, "Published", item.PublishedParsed, "link", item.Link)
	}

}
