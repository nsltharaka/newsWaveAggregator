package topicImages

import (
	"context"
	"database/sql"
	"log"

	"github.com/nsltharaka/newsWaveAggregator/database"
)

type ImageFinder struct {
	imageFetcher func(string) ([]string, int, error)
}

func NewImageFinder(imageFetcher func(string) ([]string, int, error)) *ImageFinder {
	return &ImageFinder{
		imageFetcher,
	}
}

func (finder *ImageFinder) UpdateTopic(db *database.Queries, topic string) {

	log.Println("fetching images for: ", topic)

	images, imgIndex, err := finder.imageFetcher(topic)
	if err != nil {
		log.Println("image fetcher returned an error while fetching images for: ", topic)
		return
	}

	imageUrl := sql.NullString{}

	if len(images) > 0 {

		log.Printf("found %d images for: %s\n", len(images), topic)

		imageUrl.String = images[imgIndex]
		imageUrl.Valid = true

		_, _ = db.UpdateTopicImage(context.Background(), database.UpdateTopicImageParams{
			ImgUrl: imageUrl,
			Name:   topic,
		})

		log.Println("updated image url for : ", topic)

	} else {
		log.Println("no images found for : ", topic)
	}

}
