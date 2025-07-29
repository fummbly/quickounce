package main

import (
	"net/http"
	"text/template"
)

func (cfg *apiConfig) handlerIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to get index file", err)
		return
	}

	t.Execute(w, nil)

}
