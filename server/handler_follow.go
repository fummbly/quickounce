package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fummbly/quickounce/internal/auth"
	"github.com/fummbly/quickounce/internal/database"
	"github.com/google/uuid"
)

// defining follow json object
type Follow struct {
	FollowID   uuid.UUID `json:"follow_id"`
	FolloweeID uuid.UUID `json:"followee_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func (cfg *apiConfig) handleFollow(w http.ResponseWriter, r *http.Request) {

	// defining request parameters
	type parameters struct {
		FollowID uuid.UUID `json:"follow_id"`
	}

	params := parameters{}

	decoder := json.NewDecoder(r.Body)

	// checking if request parameters meet expected parameters
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "No follow id provided", err)
		return
	}

	// getting bearer token from request
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "No bearer token provided", err)
		return
	}

	// verifiying that token matches user in database
	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthoriezed request", err)
		return
	}

	// attemting to create follow
	follow, err := cfg.db.CreateFollow(r.Context(), database.CreateFollowParams{
		FollowID:   params.FollowID,
		FolloweeID: userID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not create follow", err)
		return
	}

	// responding with successful follow information
	respondWithJSON(w, http.StatusOK, Follow{
		FollowID:   follow.FollowID,
		FolloweeID: follow.FolloweeID,
		CreatedAt:  follow.CreatedAt,
	})

}
