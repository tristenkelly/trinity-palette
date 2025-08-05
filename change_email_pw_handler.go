package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tristenkelly/the-trinity-pallette/internal/auth"
	"github.com/tristenkelly/the-trinity-pallette/internal/database"
)

func (cfg *apiConfig) changePassword(w http.ResponseWriter, r *http.Request) {
	type passParams struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := passParams{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("error decoding password change params: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPass, err := auth.HashPassword(params.Password)
	if err != nil {
		log.Printf("error hashing password (changepass): %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	queryParams := database.ChangePassParams{
		Email:          params.Email,
		HashedPassword: hashedPass,
	}

	err3 := cfg.db.ChangePass(r.Context(), queryParams)
	if err3 != nil {
		log.Printf("error changing password: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (cfg *apiConfig) changeEmail(w http.ResponseWriter, r *http.Request) {
	type passParams struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	params := passParams{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("error decoding email change params: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		log.Printf("error getting jwt: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtsecret)
	if err != nil {
		log.Printf("couldn't validate jwt: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, err := cfg.db.GetUserByName(r.Context(), params.Username)

	if userID != user.ID {
		log.Printf("can't change email, account not associated: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	queryParams := database.ChangeEmailParams{
		Username: params.Username,
		Email:    params.Email,
	}

	err3 := cfg.db.ChangeEmail(r.Context(), queryParams)
	if err3 != nil {
		log.Printf("error changing email: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
