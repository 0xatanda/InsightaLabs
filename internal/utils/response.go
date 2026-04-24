package utils

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, code int, payload any) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
