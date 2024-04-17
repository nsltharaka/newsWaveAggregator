package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/nsltharaka/newsWaveAggregator/cmd/api"
	"github.com/nsltharaka/newsWaveAggregator/database"
)

func main() {

	dbUrl := os.Getenv("DB_URL")
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("error with DSN")
	}

	if err := conn.Ping(); err != nil {
		log.Fatal("database connection failed")
	}

	db := database.New(conn)

	apiServer := api.NewAPIServer("localhost:3000", db)

	fmt.Println("server running on port 3000...")
	if err := apiServer.Run(); err != nil {
		log.Fatal(err)
	}
}
