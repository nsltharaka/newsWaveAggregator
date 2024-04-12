package main

import (
	"fmt"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/nsltharaka/newsWaveAggregator/cmd/api"
)

func main() {

	apiServer := api.NewAPIServer("localhost:3000", nil)

	fmt.Println("server running on port 3000...")
	if err := apiServer.Run(); err != nil {
		log.Fatal(err)
	}
}
