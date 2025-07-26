package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/fummbly/quickounce/internal/database"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to open database %s\n", err)
	}

	dbQueries := database.New(db)

	user, err := dbQueries.CreateUser(context.Background(), database.CreateUserParams{
		Email:          "test@test.com",
		HashedPassword: "KJSDLKHJHkdfhaldshf",
	})

	if err != nil {
		log.Fatalf("failed to add user: %s\n", err)
	}

	_, err = dbQueries.CreatePost(context.Background(), uuid.NullUUID{UUID: user.ID, Valid: true})
	if err != nil {
		log.Fatalf("Failed to create post: %s\n", err)
	}

}
