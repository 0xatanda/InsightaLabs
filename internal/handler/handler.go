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

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func errorResp(msg string) map[string]string {
	return map[string]string{
		"status":  "error",
		"message": msg,
	}
}

// ✅ CREATE PROFILE (THIS WAS MISSING — ROOT ISSUE)
func (h *Handler) CreateProfile(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		writeJSON(w, 405, errorResp("method not allowed"))
		return
	}

	var req struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, 400, errorResp("invalid request body"))
		return
	}

	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		writeJSON(w, 400, errorResp("Missing or empty name"))
		return
	}

	// 🔥 CALL SERVICE LAYER (you must already have this logic in repo/service)
	profile, isExisting, err := h.Repo.CreateOrGetProfile(req.Name)
	if err != nil {
		writeJSON(w, 502, errorResp("external API or database error"))
		return
	}

	if isExisting {
		writeJSON(w, 200, map[string]any{
			"status":  "success",
			"message": "Profile already exists",
			"data":    profile,
		})
		return
	}

	writeJSON(w, 201, map[string]any{
		"status": "success",
		"data":   profile,
	})
}

// ✅ GET ALL PROFILES
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
		if p, err := strconv.Atoi(v); err == nil && p > 0 {
			f.Page = p
		}
	}

	if v := q.Get("limit"); v != "" {
		if l, err := strconv.Atoi(v); err == nil && l > 0 && l <= 50 {
			f.Limit = l
		}
	}

	data, total, err := h.Repo.FindAll(f)
	if err != nil {
		writeJSON(w, 500, errorResp("failed to fetch profiles"))
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

// ✅ NATURAL LANGUAGE SEARCH
func (h *Handler) SearchProfiles(w http.ResponseWriter, r *http.Request) {

	q := strings.TrimSpace(r.URL.Query().Get("q"))

	if q == "" {
		writeJSON(w, 400, errorResp("Missing query"))
		return
	}

	f, err := parser.Parse(q)
	if err != nil {
		writeJSON(w, 400, errorResp("Unable to interpret query"))
		return
	}

	page := 1
	limit := 10

	if v := r.URL.Query().Get("page"); v != "" {
		if p, err := strconv.Atoi(v); err == nil {
			page = p
		}
	}

	if v := r.URL.Query().Get("limit"); v != "" {
		if l, err := strconv.Atoi(v); err == nil {
			limit = l
		}
	}

	f.Page = page
	f.Limit = limit

	data, total, err := h.Repo.FindAll(f)
	if err != nil {
		writeJSON(w, 500, errorResp("failed to fetch profiles"))
		return
	}

	writeJSON(w, 200, map[string]any{
		"status": "success",
		"page":   page,
		"limit":  limit,
		"total":  total,
		"data":   data,
	})
}
