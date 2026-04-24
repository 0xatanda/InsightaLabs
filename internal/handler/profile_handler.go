package handler

import (
	"net/http"

	"github.com/0xatanda/InsightaLabs/internal/service"
	"github.com/0xatanda/InsightaLabs/internal/utils"
)

type Handler struct {
	ProfileService *service.ProfileService
	SearchService  *service.SearchService
}

func NewHandler(p *service.ProfileService, s *service.SearchService) *Handler {
	return &Handler{
		ProfileService: p,
		SearchService:  s,
	}
}

func (h *Handler) Profiles(w http.ResponseWriter, r *http.Request) {

	q := utils.ParseRequest(r)

	data, total, err := h.ProfileService.Get(q)
	if err != nil {
		utils.JSON(w, 500, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	utils.JSON(w, 200, map[string]any{
		"status": "success",
		"page":   q.Page,
		"limit":  q.Limit,
		"total":  total,
		"data":   data,
	})
}

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query().Get("q")

	page := 1
	limit := 10

	data, total, ok, err := h.SearchService.Search(q, page, limit)

	if err != nil {
		utils.JSON(w, 500, map[string]string{"status": "error"})
		return
	}

	if !ok {
		utils.JSON(w, 200, map[string]string{
			"status":  "error",
			"message": "Unable to interpret query",
		})
		return
	}

	utils.JSON(w, 200, map[string]any{
		"status": "success",
		"data":   data,
		"total":  total,
	})
}
