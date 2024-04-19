package feed

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nsltharaka/newsWaveAggregator/service/auth"
)

func TestCreateFeedRoute(t *testing.T) {

	handler := NewHandler(nil)

	t.Run("create feed input validation", func(t *testing.T) {

		payload := map[string]any{
			"topic": "fsafd",
			"feeds": []string{"fsdfsd", "dfdsf"},
		}

		requestBody, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Fatal(err)
		}

		ctx := context.WithValue(req.Context(), auth.ContextKey("authUser"), 10)

		rr := httptest.NewRecorder()

		handler.handleCreateTopicWithFeeds(rr, req.WithContext(ctx))

		if rr.Result().StatusCode != http.StatusBadRequest {
			t.Errorf("expected %v\nactual %v", http.StatusBadRequest, rr.Result().StatusCode)
		}

		bytes, _ := io.ReadAll(rr.Result().Body)
		fmt.Printf("%v\n", string(bytes))

	})

}
