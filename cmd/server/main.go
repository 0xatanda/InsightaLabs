// cmd/server/main.go
package main

import (
	"log"
	"net/http"

	"github.com/0xatanda/InsightaLabs/internal/db"
	"github.com/0xatanda/InsightaLabs/internal/handler"

	"github.com/0xatanda/InsightaLabs/internal/repository"
)

func main() {

	database := db.Connect()
	repo := &repository.Repo{DB: database}
	h := &handler.Handler{Repo: repo}

	http.HandleFunc("/api/profiles", h.GetProfiles)
	http.HandleFunc("/api/profiles/search", h.SearchProfiles)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
