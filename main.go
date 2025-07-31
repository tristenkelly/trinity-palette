package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/lpernett/godotenv"
	"github.com/tristenkelly/the-trinity-pallette/internal/database"
)

type apiConfig struct {
	s3bucket     string
	filepath_dir string
	db           *database.Queries
	jwtsecret    string
}

func main() {
	godotenv.Load(".env")
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)
	jwt_secret := os.Getenv("SECRET")

	if err != nil {
		log.Printf("error getting database url: %v", err)
	}

	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	cfg := apiConfig{
		db:        dbQueries,
		jwtsecret: jwt_secret,
	}

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index.html")
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "login.html")
	})

	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "register.html")
	})

	mux.HandleFunc("/forgotpassword", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "changepass.html")
	})

	mux.HandleFunc("GET /shop", shopHandler)
	mux.HandleFunc("POST /admin/item/create", cfg.createItem)
	mux.HandleFunc("POST /admin/reset", cfg.resetItems)
	mux.HandleFunc("GET /blog", cfg.blogHandler)
	mux.HandleFunc("POST /admin/blog/create", cfg.createPost)
	mux.HandleFunc("GET /admin", cfg.adminPageHandler)
	mux.HandleFunc("/api/posts", cfg.postsToServe)
	mux.HandleFunc("/api/items", cfg.itemsToServe)
	mux.HandleFunc("/api/login", cfg.handleLogin)
	mux.HandleFunc("/api/register", cfg.signUpHandler)
	mux.HandleFunc("api/getrt", cfg.getRefreshToken)
	mux.HandleFunc("POST admin/revoketoken", cfg.revokeRefreshToken)
	mux.HandleFunc("/api/changepassword", cfg.changePassword)
	mux.HandleFunc("/api/changeemail", cfg.changeEmail)
	mux.HandleFunc("/api/verify", cfg.handleVerifyToken)
	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(server.Addr, server.Handler)
}

func renderTemplate(w http.ResponseWriter, filename string) {
	tmpl, err := template.ParseFiles("templates/"+filename, "templates/navbar.html")
	if err != nil {
		http.Error(w, "Page not found", 500)
		log.Println("Template error:", err)
		return
	}
	tmpl.Execute(w, nil)
}
