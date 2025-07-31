package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/tristenkelly/the-trinity-pallette/internal/auth"
	"github.com/tristenkelly/the-trinity-pallette/internal/database"
)

func (cfg *apiConfig) handleLogin(w http.ResponseWriter, r *http.Request) {
	type paramaters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	params := paramaters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("error decoding json for login %v", err)
		w.WriteHeader(500)
		return
	}

	user, err := cfg.db.GetPassHash(r.Context(), params.Email)
	if err != nil {
		log.Printf("error getting user in login query %v", err)
	}
	token, err := auth.MakeJWT(user.ID, cfg.jwtsecret)
	if err != nil {
		log.Printf("error creating JWT %v", err)
		w.WriteHeader(500)
		return
	}
	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		log.Printf("error creating refresh token %v", err)
		w.WriteHeader(500)
		return
	}
	revokedAt := sql.NullTime{
		Time:  time.Time{},
		Valid: false,
	}

	rtParams := database.CreateRefreshTokenParams{
		Token:     refreshToken,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(60 * 24 * time.Hour),
		RevokedAt: revokedAt,
	}

	type returnVals struct {
		Id            uuid.UUID `json:"id"`
		Created_at    time.Time `json:"created_at"`
		Updated_at    time.Time `json:"updated_at"`
		Email         string    `json:"email"`
		Token         string    `json:"token"`
		Refresh_token string    `json:"refresh_token"`
	}

	returnuserVals := returnVals{
		Id:            user.ID,
		Created_at:    user.CreatedAt,
		Updated_at:    user.UpdatedAt,
		Email:         user.Email,
		Token:         token,
		Refresh_token: refreshToken,
	}

	err2 := auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err2 != nil {
		log.Println("incorrect password")
		w.WriteHeader(401)
		return
	}

	_, err3 := cfg.db.CreateRefreshToken(r.Context(), rtParams)
	if err3 != nil {
		log.Printf("error creating refresh token in table: %v", err)
		w.WriteHeader(500)
		return
	}

	val, err := json.Marshal(returnuserVals)
	if err != nil {
		log.Println("error marshalling login data")
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(val)
}
