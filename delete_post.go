package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (cfg *apiConfig) deletePost(w http.ResponseWriter, r *http.Request) {
	postIDStr := chi.URLParam(r, "postID")
	var postID int32
	_, err := fmt.Sscanf(postIDStr, "%d", &postID)
	if err != nil {
		log.Printf("error parsing postID: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("deleting post: %v", postID)

	err2 := cfg.db.DeletePost(r.Context(), postID)
	if err2 != nil {
		log.Printf("error deleting item from items table: %v", err2)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("post deleted successfully")
	w.WriteHeader(http.StatusNoContent)

}
