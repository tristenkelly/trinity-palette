package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/tristenkelly/the-trinity-pallette/internal/auth"
)

func (cfg *apiConfig) userInfo(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		log.Printf("error getting bearer token in userinfo %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	userID, err := auth.ValidateJWT(token, cfg.jwtsecret)
	if err != nil {
		log.Printf("jwt not valid in userinfo %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, err2 := cfg.db.GetUser(r.Context(), userID)
	if err2 != nil {
		log.Printf("user not found: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	type returnVals struct {
		Id       uuid.UUID `json:"id"`
		Email    string    `json:"email"`
		Username string    `json:"username"`
		IsAdmin  bool      `json:"is_admin"`
	}

	returnUserVals := returnVals{
		Id:       user.ID,
		Email:    user.Email,
		Username: user.Email,
		IsAdmin:  user.IsAdmin,
	}

	val, err := json.Marshal(returnUserVals)
	if err != nil {
		log.Printf("error marshalling user info %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(val)
}
