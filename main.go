package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Check if .env file exists
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Postgres connection string
	connStr := fmt.Sprintf("host=localhost port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USERNAME"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_PASSWORD"),
	)

	// Connect to Postgres
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Connect to Dragonfly
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%s", os.Getenv("DRAGONFLY_PORT")), // use default Addr
		Password: "",                                                       // no password set
		DB:       0,                                                        // use default DB
	})

	// Check Postgres connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	// Check Dragonfly connection
	if _, err := rdb.Ping().Result(); err != nil {
		log.Fatalf("failed to connect to Dragonfly: %v", err)
	}

	// Everything is working
	fmt.Println("🚀 Connected to Postgres and Redis")

}
