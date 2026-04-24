package main

import (
	"log"
	"net/http"

	"github.com/0xatanda/InsightaLabs/internal/config"
	"github.com/0xatanda/InsightaLabs/internal/handler"
	"github.com/0xatanda/InsightaLabs/internal/parser"
	"github.com/0xatanda/InsightaLabs/internal/query"
	"github.com/0xatanda/InsightaLabs/internal/repository"
	"github.com/0xatanda/InsightaLabs/internal/service"
	_ "github.com/lib/pq"
)

func main() {
	db := config.ConnectDB()

	repo := repository.NewProfileRepository(db)
	builder := query.NewProfileQueryBuilder()
	parser := parser.NewParser()
	profileService := service.NewProfileService(repo, builder)
	searchService := service.NewSearchService(repo, builder, parser)

	h := handler.NewHandler(profileService, searchService)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/profiles", h.Profiles)
	mux.HandleFunc("/api/profiles/search", h.Search)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", cors(mux)))
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
