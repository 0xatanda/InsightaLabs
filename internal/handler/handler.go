package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/0xatanda/InsightaLabs/internal/model"
	"github.com/0xatanda/InsightaLabs/internal/parser"
	"github.com/0xatanda/InsightaLabs/internal/repository"
)

type Handler struct {
	Repo *repository.Repo
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *Handler) GetProfiles(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()

	f := &model.Filters{
		Gender:    q.Get("gender"),
		AgeGroup:  q.Get("age_group"),
		CountryID: q.Get("country_id"),
		Page:      1,
		Limit:     10,
	}

	if v := q.Get("page"); v != "" {
		f.Page, _ = strconv.Atoi(v)
	}
	if v := q.Get("limit"); v != "" {
		f.Limit, _ = strconv.Atoi(v)
	}

	data, total, _ := h.Repo.FindAll(f)

	writeJSON(w, 200, map[string]any{
		"status": "success",
		"page":   f.Page,
		"limit":  f.Limit,
		"total":  total,
		"data":   data,
	})
}

func (h *Handler) SearchProfiles(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query().Get("q")

	f, err := parser.Parse(q)
	if err != nil {
		writeJSON(w, 400, map[string]string{
			"status":  "error",
			"message": "Unable to interpret query",
		})
		return
	}

	data, total, _ := h.Repo.FindAll(f)

	writeJSON(w, 200, map[string]any{
		"status": "success",
		"total":  total,
		"data":   data,
	})
}
