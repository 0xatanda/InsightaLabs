package main

import (
	"log"
	"net/http"
	"os"

	"github.com/0xatanda/InsightaLabs/internal/db"
	"github.com/0xatanda/InsightaLabs/internal/handler"
	"github.com/0xatanda/InsightaLabs/internal/repository"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	database := db.Connect()
	repo := &repository.Repo{DB: database}
	h := &handler.Handler{Repo: repo}

	http.HandleFunc("/api/profiles", h.GetProfiles)
	http.HandleFunc("/api/profiles/search", h.SearchProfiles)

	log.Println("Server running on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
