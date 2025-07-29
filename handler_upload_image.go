package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

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

	file, fileHeader, err := r.FormFile("photo")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to get file from post", err)
		return
	}

	createdAt := fmt.Sprintf("%d", time.Now().UnixNano())

	uploadID := encrypting.Hash([]byte(createdAt))
	stringUploadID := base64.URLEncoding.EncodeToString(uploadID)

	userID, err := uuid.Parse(r.FormValue("user_id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to parse user id", err)
		return
	}

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

	respondWithJSON(w, http.StatusOK, Post{
		UserID:    post.UserID,
		ID:        post.ID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		ImageUrl:  post.ImageUrl,
	})

}
