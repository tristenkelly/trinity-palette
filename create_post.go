package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/tristenkelly/the-trinity-pallette/internal/auth"
	"github.com/tristenkelly/the-trinity-pallette/internal/database"
)

func (cfg *apiConfig) createPost(w http.ResponseWriter, r *http.Request) {
	type postParams struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := postParams{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("error creating post")
		w.WriteHeader(500)
		return
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		log.Printf("error getting jwt: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtsecret)
	if err != nil {
		log.Printf("error validating jwt: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	queryParams := database.CreatePostParams{
		Title:     params.Title,
		Body:      params.Body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userID,
	}

	post, err := cfg.db.CreatePost(r.Context(), queryParams)
	if err != nil {
		log.Printf("error creating post in db: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	type returnPost struct {
		Title     string    `json:"title"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"created_at"`
		UserID    uuid.UUID `json:"user_id"`
	}

	returnPostResults := returnPost{
		Title:     post.Title,
		Body:      post.Body,
		CreatedAt: post.CreatedAt,
		UserID:    post.UserID,
	}

	val, err := json.Marshal(returnPostResults)
	if err != nil {
		log.Printf("error marshalling post data %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err2 := w.Write(val)
	if err2 != nil {
		log.Printf("error writing response: %v", err2)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
