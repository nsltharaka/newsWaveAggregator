package topicImages

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"slices"
)

/*

finding an image for a topic.

current method:
	use the news api to search for articles about that topic and get one of the articles from there and get it's cover image.
	drawbacks:
		logos are there. specially in BBC - may be need to filter them out
		image might be irrelevant to the topic.

*/

func FromNewsAPI() ImageFinder {

	apiUrl, _ := url.Parse("https://newsapi.org/v2/everything")
	q := apiUrl.Query()
	q.Set("apiKey", os.Getenv("NEWS_API_KEY"))

	return func(topic string) (imageUrl []string, err error) {

		q.Set("q", topic)
		q.Set("pageSize", "10")
		apiUrl.RawQuery = q.Encode()

		newsApiResponse, err := fetchData(apiUrl.String())
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
}

func fetchData(apiUrl string) (*newsApiResponse, error) {
	response, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	apiResp, err := parseJSON(bytes)
	if err != nil {
		return nil, err
	}

	return apiResp, nil
}

func parseJSON(bytes []byte) (*newsApiResponse, error) {
	var apiResponse newsApiResponse
	if err := json.Unmarshal(bytes, &apiResponse); err != nil {
		return nil, err
	}
	return &apiResponse, nil
}

type newsApiResponse struct {
	Articles []struct {
		ImageUrl string `json:"urlToImage"`
	} `json:"articles"`
}
