package topicImages

import (
	"fmt"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func TestGoogleImages(t *testing.T) {

	godotenv.Load("../../example.env")

	response, _, err := FromGoogleImages("Go Programming")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(strings.Join(response, "\n"))

}
