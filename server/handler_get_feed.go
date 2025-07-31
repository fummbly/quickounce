package main

import (
	"fmt"
	"net/http"

	"github.com/fummbly/quickounce/internal/auth"
)

func (cfg *apiConfig) handlerGetFeed(w http.ResponseWriter, r *http.Request) {

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't find bearer token", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Failed to validate token", err)
	}

	fmt.Println(userID)

}
