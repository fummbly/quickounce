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

// server state object for holding
// information needed between functions
type apiConfig struct {
	db        *database.Queries
	platform  string
	jwtSecret string
}

func main() {
	const filepathRoot = "."
	// port for server to run on
	const port = "8080"

	// getting enviroment variables
	godotenv.Load()
	// checking that each is setup
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

	// making connection to database
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

	// starting mutex server
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// redirect routes for content
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.Handle("GET /uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	// user api routes
	mux.Handle("GET /api/users", apiCfg.middlewareCors(apiCfg.handlerUsersGet))
	mux.Handle("POST /api/users", apiCfg.middlewareCors(apiCfg.handlerUsersCreate))
	mux.Handle("GET /api/users/{username}", apiCfg.middlewareCors(apiCfg.handlerGetUser))

	// user login route
	mux.Handle("POST /api/login", apiCfg.middlewareCors(apiCfg.handlerLogin))
	mux.Handle("GET /api/refresh", apiCfg.middlewareCors(apiCfg.handlerRefresh))

	// api post routes
	mux.Handle("POST /api/posts", apiCfg.middlewareCors(apiCfg.handlerUploadImage))
	mux.HandleFunc("OPTIONS /api/posts", apiCfg.handlerPostsOption)
	mux.Handle("DELETE /api/post/{postID}", apiCfg.middlewareCors(apiCfg.handlerDeletePost))

	// api follow routes
	mux.Handle("POST /api/follows", apiCfg.middlewareCors(apiCfg.handleFollow))

	// api comment routes
	mux.Handle("POST /api/comments", apiCfg.middlewareCors(apiCfg.handlerCommentCreate))

	// admin route to reset databases
	mux.Handle("POST /admin/reset", apiCfg.middlewareCors(apiCfg.handlerReset))

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}

// CORS enabler for development setup
func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")
}

// CORS middleware to pass access control to handlers
func (cfg *apiConfig) middlewareCors(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check if the server is in a dev enviroment
		if cfg.platform == "dev" {
			enableCors(w)
		}
		next.ServeHTTP(w, r)
	})
}
