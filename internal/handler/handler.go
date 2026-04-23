package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/0xatanda/InsightaLabs/internal/model"
	"github.com/0xatanda/InsightaLabs/internal/parser"
	"github.com/0xatanda/InsightaLabs/internal/repository"
)

type Handler struct {
	Repo *repository.Repo
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func errResp(msg string) map[string]string {
	return map[string]string{"status": "error", "message": msg}
}

func (h *Handler) GetProfiles(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()

	f := &model.Filters{
		Gender:    strings.ToLower(q.Get("gender")),
		AgeGroup:  strings.ToLower(q.Get("age_group")),
		CountryID: strings.ToUpper(q.Get("country_id")),
		SortBy:    q.Get("sort_by"),
		Order:     strings.ToLower(q.Get("order")),
		Page:      1,
		Limit:     10,
	}

	if v := q.Get("page"); v != "" {
		val, err := strconv.Atoi(v)
		if err != nil || val < 1 {
			writeJSON(w, 422, errResp("Invalid query parameters"))
			return
		}
		f.Page = val
	}

	if v := q.Get("limit"); v != "" {
		val, err := strconv.Atoi(v)
		if err != nil || val < 1 {
			writeJSON(w, 422, errResp("Invalid query parameters"))
			return
		}
		if val > 50 {
			val = 50
		}
		f.Limit = val
	}

	if v := q.Get("min_age"); v != "" {
		val, err := strconv.Atoi(v)
		if err != nil {
			writeJSON(w, 422, errResp("Invalid query parameters"))
			return
		}
		f.MinAge = &val
	}

	if v := q.Get("max_age"); v != "" {
		val, err := strconv.Atoi(v)
		if err != nil {
			writeJSON(w, 422, errResp("Invalid query parameters"))
			return
		}
		f.MaxAge = &val
	}

	data, total, err := h.Repo.FindAll(f)
	if err != nil {
		writeJSON(w, 500, errResp("failed to fetch profiles"))
		return
	}

	writeJSON(w, 200, map[string]any{
		"status": "success",
		"page":   f.Page,
		"limit":  f.Limit,
		"total":  total,
		"data":   data,
	})
}

func (h *Handler) SearchProfiles(w http.ResponseWriter, r *http.Request) {

	q := strings.TrimSpace(r.URL.Query().Get("q"))

	if q == "" {
		writeJSON(w, 400, errResp("Missing query"))
		return
	}

	f, err := parser.Parse(q)
	if err != nil {
		writeJSON(w, 400, errResp("Unable to interpret query"))
		return
	}

	f.Page = 1
	f.Limit = 10

	data, total, err := h.Repo.FindAll(f)
	if err != nil {
		writeJSON(w, 500, errResp("failed to fetch profiles"))
		return
	}

	writeJSON(w, 200, map[string]any{
		"status": "success",
		"page":   f.Page,
		"limit":  f.Limit,
		"total":  total,
		"data":   data,
	})
}
