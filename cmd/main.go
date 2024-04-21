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
	host, err := os.Hostname()
	if err != nil {
		fmt.Println("could not retrieve hostname, using localhost")
		host = "localhost"
	}

	apiServer := api.NewAPIServer(fmt.Sprintf("%s:%s", host, PORT), db)

	fmt.Printf("server started at http://%s:%s\n", host, PORT)
	if err := apiServer.Run(); err != nil {
		log.Fatal(err)
	}
}
