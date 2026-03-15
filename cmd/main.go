package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=spotify_user password=spotify123 dbname=spotify_stats sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error Open DB:", err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Error ping DB:", err)
	}

	var count int

	err = db.QueryRow("SELECT COUNT(*) FROM Users").Scan(&count)
	if err != nil {
		log.Fatal("Error Query:", err)
	}

	log.Println("DB connect!")
	log.Printf("Table Users: %d count string", count)
}
