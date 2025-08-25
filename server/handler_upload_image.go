package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/fummbly/quickounce/internal/auth"
	"github.com/fummbly/quickounce/internal/database"
	"github.com/fummbly/quickounce/internal/encrypting"
	"github.com/fummbly/quickounce/internal/photoproc"
	"github.com/google/uuid"
)

type Post struct {
	UserID    uuid.UUID `json:"user_id"`
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ImageUrl  string    `json:"image_url"`
}

func (cfg *apiConfig) handlerUploadImage(w http.ResponseWriter, r *http.Request) {

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

	file, fileHeader, err := r.FormFile("photo")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to get file from post", err)
		return
	}

	createdAt := fmt.Sprintf("%d", time.Now().UnixNano())

	uploadID := encrypting.Hash([]byte(createdAt))
	stringUploadID := base64.URLEncoding.EncodeToString(uploadID)

	err = photoproc.CopyPhoto(stringUploadID, file, fileHeader)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to copy photo", err)
		return
	}

	post, err := cfg.db.CreatePost(r.Context(), database.CreatePostParams{
		UserID:   userID,
		ImageUrl: stringUploadID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create post", err)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	respondWithJSON(w, http.StatusOK, Post{
		UserID:    post.UserID,
		ID:        post.ID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		ImageUrl:  post.ImageUrl,
	})
}

func (cfg *apiConfig) handlerPostsOption(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusNoContent)
}
