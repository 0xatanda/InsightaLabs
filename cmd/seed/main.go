package main

import (
	"log"

	"github.com/0xatanda/InsightaLabs/internal/config"
)

func main() {
	db := config.ConnectDB()

	if err := Seed(db); err != nil {
		log.Fatal(err)
	}

	log.Println("Seeding completed")
}
