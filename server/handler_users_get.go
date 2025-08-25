package main

import "net/http"

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request) {

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

	type response struct {
		User  User
		Posts []Post
	}

	username := r.PathValue("username")
	dbUser, err := cfg.db.GetUserByUsername(r.Context(), username)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve user", err)
		return
	}

	user := User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Email:     dbUser.Email,
		Username:  dbUser.Username,
	}

	dbPosts, err := cfg.db.GetPostsByUserID(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to get posts by user", err)
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

	respondWithJSON(w, http.StatusOK, response{
		User:  user,
		Posts: posts,
	})

}
