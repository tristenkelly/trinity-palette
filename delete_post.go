package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (cfg *apiConfig) deletePost(w http.ResponseWriter, r *http.Request) {
	postIDStr := chi.URLParam(r, "postID")
	var postID int
	fmt.Sscanf(postIDStr, "%d", &postID)

	log.Printf("deleting post: %v", postID)

	err := cfg.db.DeletePost(r.Context(), int32(postID))
	if err != nil {
		log.Printf("error deleting item from items table: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("post deleted successfully")
	w.WriteHeader(http.StatusNoContent)

}
