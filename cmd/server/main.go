package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/0xatanda/InsightaLabs/internal/handler"
	"github.com/0xatanda/InsightaLabs/internal/repository"
)

func main() {

	db, err := sql.Open("postgres", "postgres://appuser:password@localhost:5432/profiles_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	repo := &repository.Repo{DB: db}
	h := &handler.Handler{Repo: repo}

	http.HandleFunc("/api/profiles", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			h.CreateProfile(w, r)
			return
		}
		if r.Method == "GET" {
			h.GetProfiles(w, r)
			return
		}
		w.WriteHeader(405)
	})

	http.HandleFunc("/api/profiles/search", h.SearchProfiles)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
