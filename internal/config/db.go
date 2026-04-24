package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {

	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("DB open error:", err)
	}

	// verify connection
	if err := db.Ping(); err != nil {
		log.Fatal("DB ping error:", err)
	}

	// optional but recommended
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	return db
}
