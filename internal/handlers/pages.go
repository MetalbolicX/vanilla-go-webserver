package handlers

import (
	"html/template"
	"net/http"
	"path"

	"github.com/MetalbolicX/vanilla-go-webserver/pkg/repository"
)

const tmpFolder string = "tmp"

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	indexFilePath := path.Join(tmpFolder, "index.html")
	htmlTemplate, err := template.ParseFiles(indexFilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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

	if err := htmlTemplate.Execute(w, exercises); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html")
}
