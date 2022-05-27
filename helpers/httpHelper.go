package helpers

import (
	"encoding/json"
	"net/http"
)

func HttpJsonResponse(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
