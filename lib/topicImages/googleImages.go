package topicImages

import (
	"net/url"
	"os"
)

func FromGoogleImages(topic string) ([]string, int, error) {

	apiUrl, _ := url.Parse("https://www.googleapis.com/customsearch/v1")
	q := apiUrl.Query()
	q.Set("key", os.Getenv("GOOGLE_API_KEY"))
	q.Set("cx", os.Getenv("GOOGLE_SEARCH_ENGINE_ID"))
	q.Set("searchType", "image")
	q.Set("imgType", "photo")
	q.Set("q", topic)

	apiUrl.RawQuery = q.Encode()

	googleImageResponse, err := fetchData[googleImageResponse](apiUrl.String())
	if err != nil {
		return nil, -1, err
	}

	imageUrls := []string{}
	for _, item := range googleImageResponse.Items {
		imageUrls = append(imageUrls, item.Link)
	}

	return imageUrls, 0, nil

}

type googleImageResponse struct {
	Items []struct {
		Link string `json:"link"`
	} `json:"items"`
}
