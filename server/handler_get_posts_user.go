package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetPostByUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		UserID uuid.UUID `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to decode user id", err)
		return
	}

	dbPosts, err := cfg.db.GetPostsByUserID(r.Context(), params.UserID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "User not found", err)
		return
	}

	posts := []Post{}

	for _, post := range dbPosts {
		posts = append(posts, Post{
			UserID:    post.UserID,
			ID:        post.ID,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
			ImageUrl:  post.ImageUrl,
		})
	}

	respondWithJSON(w, http.StatusOK, posts)

}
