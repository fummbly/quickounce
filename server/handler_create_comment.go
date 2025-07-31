package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fummbly/quickounce/internal/auth"
	"github.com/fummbly/quickounce/internal/database"
	"github.com/google/uuid"
)

type Comment struct {
	ID          uuid.UUID `json:"id"`
	PostID      uuid.UUID `json:"post_id"`
	UserID      uuid.UUID `json:"user_id"`
	CommentText string    `json:"comment_text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (cfg *apiConfig) handlerCommentCreate(w http.ResponseWriter, r *http.Request) {

	type parameter struct {
		PostID  string `json:"post_id"`
		Comment string `json:"comment"`
	}

	params := parameter{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode body", err)
		return
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't get access token", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate token", err)
		return
	}

	postID, err := uuid.Parse(params.PostID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not parse post id", err)
		return
	}

	_, err = cfg.db.GetPost(r.Context(), postID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't find post", err)
		return
	}

	comment, err := cfg.db.CreateComment(r.Context(), database.CreateCommentParams{
		UserID:      userID,
		PostID:      postID,
		CommentText: params.Comment,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create comment", err)
		return
	}

	respondWithJSON(w, http.StatusOK, Comment{
		ID:          comment.ID,
		PostID:      comment.PostID,
		UserID:      comment.UserID,
		CommentText: comment.CommentText,
		CreatedAt:   comment.CreatedAt,
		UpdatedAt:   comment.UpdatedAt,
	})
}
