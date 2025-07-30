package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/fummbly/quickounce/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	db        *database.Queries
	platform  string
	jwtSecret string
}

func main() {
	const filepathRoot = "."
	const port = "8080"

	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}
	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatal("PLATFORM must be set")
	}
	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Fatal("SECRET must be seet")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s\n", err)
	}
	dbQueries := database.New(dbConn)

	apiCfg := apiConfig{
		db:        dbQueries,
		platform:  platform,
		jwtSecret: secret,
	}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.HandleFunc("GET /", apiCfg.handlerIndex)
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.Handle("GET /uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	mux.HandleFunc("GET /api/users", apiCfg.handlerUsersGet)
	mux.HandleFunc("GET /api/users/{username}", apiCfg.handlerGetUser)

	mux.HandleFunc("POST /api/users", apiCfg.handlerUsersCreate)
	mux.HandleFunc("POST /api/posts", apiCfg.handlerUploadImage)
	mux.HandleFunc("DELETE /api/post/{postID}", apiCfg.handlerDeletePost)
	mux.HandleFunc("POST /api/login", apiCfg.handlerLogin)
	mux.HandleFunc("POST /api/comments", apiCfg.handlerCommentCreate)

	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
