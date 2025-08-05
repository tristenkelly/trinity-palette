package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/lpernett/godotenv"
	"github.com/tristenkelly/the-trinity-pallette/internal/database"
)

type apiConfig struct {
	s3bucket     string
	filepath_dir string
	db           *database.Queries
	jwtsecret    string
	port         string
}

type productImage struct {
	data      []byte
	mediaType string
}

var itemImages = map[uuid.UUID]productImage{}

func main() {
	r := chi.NewRouter()
	godotenv.Load(".env")
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)
	jwt_secret := os.Getenv("SECRET")
	port := os.Getenv("PORT")

	if err != nil {
		log.Printf("error getting database url: %v", err)
	}

	cfg := apiConfig{
		db:        dbQueries,
		jwtsecret: jwt_secret,
		port:      port,
	}

	server := http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index.html")
	})

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "login.html")
	})

	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "register.html")
	})

	r.HandleFunc("/forgotpassword", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "changepass.html")
	})

	r.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "admin.html")
	})

	r.Get("/shop", shopHandler)
	r.Post("/admin/item/create", cfg.createItem)
	r.Post("/admin/reset", cfg.resetItems)
	r.Get("/blog", cfg.blogHandler)
	r.Post("/api/post", cfg.createPost)
	r.HandleFunc("/api/posts", cfg.postsToServe)
	r.HandleFunc("/api/items", cfg.itemsToServe)
	r.HandleFunc("/api/login", cfg.handleLogin)
	r.HandleFunc("/api/register", cfg.signUpHandler)
	r.HandleFunc("/api/getrt", cfg.getRefreshToken)
	r.Post("/admin/revoketoken", cfg.revokeRefreshToken)
	r.HandleFunc("/api/changepassword", cfg.changePassword)
	r.HandleFunc("/api/changeemail", cfg.changeEmail)
	r.HandleFunc("/api/verify", cfg.handleVerifyToken)
	r.HandleFunc("/api/userInfo", cfg.userInfo)
	r.Delete("/api/item/{itemID}", cfg.deleteItem)
	r.Delete("/api/post/{postID}", cfg.deletePost)
	log.Println("Server running on http://localhost:8080")
	err2 := server.ListenAndServe()
	if err2 != nil {
		log.Fatal("Error starting server:", err2)
	}
}

func renderTemplate(w http.ResponseWriter, filename string) {
	tmpl, err := template.ParseFiles("templates/"+filename, "templates/navbar.html")
	if err != nil {
		http.Error(w, "Page not found", 500)
		log.Println("Template error:", err)
		return
	}
	err2 := tmpl.Execute(w, nil)
	if err2 != nil {
		http.Error(w, "Error rendering template", 500)
		log.Println("Template execution error:", err2)
		return
	}
}
