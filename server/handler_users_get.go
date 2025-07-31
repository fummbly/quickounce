package main

import "net/http"

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request) {

	enableCors(w)

	dbUsers, err := cfg.db.GetUsers(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch users", err)
		return
	}

	users := []User{}

	for _, dbUser := range dbUsers {
		users = append(users, User{
			ID:        dbUser.ID,
			CreatedAt: dbUser.CreatedAt,
			UpdatedAt: dbUser.UpdatedAt,
			Email:     dbUser.Email,
			Username:  dbUser.Username,
		})
	}

	respondWithJSON(w, http.StatusOK, users)
}

func (cfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {

	username := r.PathValue("username")
	user, err := cfg.db.GetUserByUsername(r.Context(), username)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve user", err)
		return
	}

	respondWithJSON(w, http.StatusOK, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
		Username:  user.Username,
	})

}
