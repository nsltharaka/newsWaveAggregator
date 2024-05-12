package topic

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nsltharaka/newsWaveAggregator/types"
)

func TestEditTopic(t *testing.T) {

	// request body
	reqBody := types.IncomingFollowTopicFeedPayload{
		Topic:    "test_topic",
		FeedURLs: []string{"testuri", "testuri2"},
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf(err.Error())
	}

	// setup
	req := httptest.NewRequest(http.MethodPut, "/topics", bytes.NewReader(reqBodyBytes))
	req.Header.Set("Authorization", "Bearer 437c4058f8c10d8b589aaea669a18cb7a97cea6f2115cba0484b04e9e07d20ec")

	// making the request
	// handle
	handler := NewHandler(nil)

	// router
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.handleUpdateTopic)

	// response

}
