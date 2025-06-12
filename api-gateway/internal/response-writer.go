package internal

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, reponseBody interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(reponseBody); err != nil {
		http.Error(w, "Can't parse body", http.StatusInternalServerError)
		return
	}
}
