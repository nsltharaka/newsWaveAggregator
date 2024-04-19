package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload" // loads and setup the env variables
	_ "github.com/lib/pq"                 // postgres database driver
	"github.com/nsltharaka/newsWaveAggregator/cmd/api"
	"github.com/nsltharaka/newsWaveAggregator/database"
)

func main() {

	// database connection
	dbUrl := os.Getenv("DB_URL")
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("error with DSN")
	}
	// check connection
	if err := conn.Ping(); err != nil {
		log.Fatal("database connection failed")
	}

	db := database.New(conn)

	// server setup
	PORT := os.Getenv("PORT")
	HOST := os.Getenv("HOST")
	apiServer := api.NewAPIServer(fmt.Sprintf("%s:%s", HOST, PORT), db)

	fmt.Printf("server started on http://%s:%s\n", HOST, PORT)
	if err := apiServer.Run(); err != nil {
		log.Fatal(err)
	}
}
