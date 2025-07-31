package main

import (
	"log"
	"net/http"

	"github.com/tristenkelly/the-trinity-pallette/internal/auth"
)

func (cfg *apiConfig) handleVerifyToken(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		log.Printf("error getting bearer token from header %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err2 := auth.ValidateJWT(token, cfg.jwtsecret)
	if err2 != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
