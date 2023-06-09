package handlers

import (
	"encoding/json"
	"net/http"
)

type homeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(homeResponse{
		Message: "Welcome to my first web app with Go!!",
		Status:  true,
	})
}
