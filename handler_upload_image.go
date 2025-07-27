package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/fummbly/quickounce/internal/database"
	"github.com/fummbly/quickounce/internal/encrypting"
	"github.com/google/uuid"
)

type Post struct {
	UserID        uuid.UUID `json:"user_id"`
	ID            uuid.UUID `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	UploadID      []byte    `json:"upload_id"`
	TotalLikes    int       `json:"total_likes"`
	TotalComments int       `json:"total_comments"`
	PhotoID       uuid.UUID `json:"photo_id"`
	PhotoPath     string    `json:"photo_path"`
}

func (cfg *apiConfig) handlerUploadImage(w http.ResponseWriter, r *http.Request) {

	createdAt := fmt.Sprintf("%s", time.Now().UnixNano())

	uploadID := encrypting.Hash([]byte(createdAt))
	stringUploadID := base64.URLEncoding.EncodeToString(uploadID)

	userID, err := uuid.Parse(r.FormValue("user_id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to parse user id", err)
		return
	}

	post, err := cfg.db.CreatePost(r.Context(), database.CreatePostParams{
		UserID:   userID,
		UploadID: uploadID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create post", err)
	}

	file, fileHeader, err := r.FormFile("photo")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to open file", err)
		return
	}

	defer file.Close()

	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to read file", err)
		return
	}

	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" {
		respondWithError(w, http.StatusBadRequest, "Incorrect filetype: allowed files are jpeg and png", nil)
		return
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to go back to begining of file", err)
		return
	}

	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create uploads folder", err)
		return
	}

	dst, err := os.Create(fmt.Sprintf("./uploads/%s%s", stringUploadID, filepath.Ext(fileHeader.Filename)))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create file on server", err)
		return
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to copy data to new file", err)
		return
	}

	photo, err := cfg.db.CreatePhoto(r.Context(), database.CreatePhotoParams{
		UserID:    userID,
		PostID:    post.ID,
		PhotoPath: stringUploadID,
	})

	respondWithJSON(w, http.StatusOK, Post{
		UserID:        post.UserID,
		ID:            post.ID,
		CreatedAt:     post.CreatedAt,
		UpdatedAt:     post.UpdatedAt,
		UploadID:      post.UploadID,
		TotalLikes:    int(post.TotalLikes.Int32),
		TotalComments: int(post.TotalComments.Int32),
		PhotoID:       photo.ID,
		PhotoPath:     photo.PhotoPath,
	})

}
