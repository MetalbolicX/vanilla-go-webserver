package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MetalbolicX/vanilla-go-webserver/pkg/repository"
)

func GetExercisesHandler(w http.ResponseWriter, r *http.Request) {

	exercises, err := repository.Get(r.Context(), `
		SELECT
			exercise_name_en
		FROM exercises
		WHERE movement_id = $1
		`, 1)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exercises)
}
