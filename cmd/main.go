package main

import (
	"log"
	"os"
	"spotify_mod/internal/db"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("file not found: %v", err)
	}

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL error!")
	}

	database, err := db.New(connStr)
	if err != nil {
		log.Fatal("DB error: ", err)
	}
	defer database.Close()

}
