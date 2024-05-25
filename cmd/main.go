package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // postgres database driver
	"github.com/nsltharaka/newsWaveAggregator/cmd/api"
	"github.com/nsltharaka/newsWaveAggregator/database"
)

func main() {

	// environment variables
	if err := godotenv.Load("example.env"); err != nil {
		log.Fatalf("error loading environmental variables : %v", err)
	}

	// database connection
	dbUrl := os.Getenv("DB_URL")
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("error with DSN: %v", err)
	}
	// check connection
	if err := conn.Ping(); err != nil {
		log.Fatalf("database connection failed: %v", err)
	}

	db := database.New(conn)

	// server setup
	// default port
	port, exists := os.LookupEnv("API_PORT")
	if !exists {
		port = "3030"
	}

	// default host
	hostname, exists := os.LookupEnv("API_HOST")
	if !exists {
		var err error
		hostname, err = os.Hostname()
		if err != nil {
			hostname = "localhost"
		}
	}

	apiServer := api.NewAPIServer(fmt.Sprintf("%s:%s", hostname, port), db)

	apiBaseUrl := fmt.Sprintf("http://%s:%s", hostname, port)
	if err := os.Setenv("API_BASE_URL", apiBaseUrl); err != nil {
		slog.Warn("setting env variable API_BASE_URL failed")
	}

	fmt.Printf("server started at %s\n", apiBaseUrl)
	if err := apiServer.Run(); err != nil {
		log.Fatal(err)
	}
}
