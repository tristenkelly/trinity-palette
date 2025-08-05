package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (cfg *apiConfig) postsToServe(w http.ResponseWriter, r *http.Request) {
	posts, err := cfg.db.GetPosts(r.Context())
	if err != nil {
		log.Printf("error getting posts from database")
		w.WriteHeader(500)
		return
	}

	type postResponse struct {
		Id         int       `json:"id"`
		Post_title string    `json:"title"`
		Post_body  string    `json:"body"`
		CreatedAt  time.Time `json:"created_at"`
		Username   string    `json:"username"`
	}

	var responsePosts []postResponse
	for _, post := range posts {
		responsePosts = append(responsePosts, postResponse{
			Id:         int(post.ID),
			Post_title: post.Title,
			Post_body:  post.Body,
			CreatedAt:  post.CreatedAt,
			Username:   post.Username,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	err2 := json.NewEncoder(w).Encode(responsePosts)
	if err2 != nil {
		log.Printf("error encoding struct: %v", err2)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
