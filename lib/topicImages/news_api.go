package topicImages

import (
	"net/url"
	"os"
	"slices"
)

func FromNewsAPI(topic string) ([]string, error) {

	apiUrl, _ := url.Parse("https://newsapi.org/v2/everything")
	q := apiUrl.Query()
	q.Set("apiKey", os.Getenv("NEWS_API_KEY"))
	q.Set("q", topic)
	q.Set("pageSize", "10")

	apiUrl.RawQuery = q.Encode()

	newsApiResponse, err := fetchData[newsApiResponse](apiUrl.String())
	if err != nil {
		return nil, err
	}

	imageUrls := []string{}
	for _, article := range newsApiResponse.Articles {
		if slices.Index(imageUrls, article.ImageUrl) == -1 {
			imageUrls = append(imageUrls, article.ImageUrl)
		}
	}

	return imageUrls, nil
}

type newsApiResponse struct {
	Articles []struct {
		ImageUrl string `json:"urlToImage"`
	} `json:"articles"`
}
