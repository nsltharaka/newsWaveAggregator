package topicImages

import (
	"encoding/json"
	"io"
	"net/http"
)

func fetchData[T any](apiUrl string) (*T, error) {
	response, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse T
	if err := json.Unmarshal(bytes, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}
