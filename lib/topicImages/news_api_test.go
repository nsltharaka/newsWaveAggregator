package topicImages

import (
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func TestBuidURL(t *testing.T) {

	godotenv.Load("../../.env")

	t.Run("url extract", func(t *testing.T) {

		extractor := FromNewsAPI()

		imgUrls, err := extractor("stock prices")
		if err != nil {
			t.Error(err)
		}

		t.Log(strings.Join(imgUrls, "\n"))

	})

}
