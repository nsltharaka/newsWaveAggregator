package topicImages

import (
	"net/url"
	"os"
)

func FromPixabay(topic string) ([]string, error) {

	apiUrl, _ := url.Parse("https://pixabay.com/api/")
	q := apiUrl.Query()
	q.Set("key", os.Getenv("PIXABAY_API_KEY"))
	q.Set("image_type", "photo")
	q.Set("q", topic)

	apiUrl.RawQuery = q.Encode()

	apiResponse, err := fetchData[pixabayResponse](apiUrl.String())
	if err != nil {
		return nil, err
	}

	imageUrls := []string{}
	for _, hit := range apiResponse.Hits {
		imageUrls = append(imageUrls, hit.ImageUrl)
	}

	return imageUrls, nil

}

type pixabayResponse struct {
	Hits []struct {
		ImageUrl string `json:"largeImageURL"`
	} `json:"hits"`
}
