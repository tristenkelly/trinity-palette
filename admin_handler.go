package main

import (
	"log"
	"net/http"

	"github.com/tristenkelly/the-trinity-pallette/internal/auth"
)

func (cfg *apiConfig) adminPageHandler(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		log.Printf("couldn't get bearer token on admin page: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtsecret)
	if err != nil {
		log.Printf("couldn't validate jwt in admin portal %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, err := cfg.db.GetUser(r.Context(), userID)
	if err != nil {
		log.Printf("couldn't get user from db (admin page): %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.IsAdmin == false {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Not authorized to access this page"))
		return
	} else {
		renderTemplate(w, "admin.html")
		w.WriteHeader(http.StatusOK)
	}
}
