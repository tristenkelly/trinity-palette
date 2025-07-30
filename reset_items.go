package main

import (
	"log"
	"net/http"
)

func (cfg *apiConfig) resetItems(w http.ResponseWriter, r *http.Request) {
	err := cfg.db.ResetItems(r.Context())
	if err != nil {
		log.Printf("error resetting items table %v", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(http.StatusOK)
}
