package main

import (
	"net/http"

	"github.com/fummbly/quickounce/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerDeletePost(w http.ResponseWriter, r *http.Request) {

	postIDString := r.PathValue("postID")
	postID, err := uuid.Parse(postIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid post id", err)
		return
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}

	dbPost, err := cfg.db.GetPost(r.Context(), postID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get post", err)
		return
	}

	if dbPost.UserID != userID {
		respondWithError(w, http.StatusForbidden, "You can't delete this post", err)
		return
	}

	err = cfg.db.DeletePost(r.Context(), postID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete post", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
