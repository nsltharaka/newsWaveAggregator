package main

import (
	"database/sql"
	"fmt"
	"log"
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
	PORT := os.Getenv("PORT")
	host, err := os.Hostname()
	if err != nil {
		fmt.Println("could not retrieve hostname\n\tusing default value 'localhost'")
		host = "localhost"
	}

	apiServer := api.NewAPIServer(fmt.Sprintf("%s:%s", host, PORT), db)

	fmt.Printf("server started at http://%s:%s\n", host, PORT)
	if err := apiServer.Run(); err != nil {
		log.Fatal(err)
	}
}
