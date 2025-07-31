package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/tristenkelly/the-trinity-pallette/internal/auth"
	"github.com/tristenkelly/the-trinity-pallette/internal/database"
)

func (cfg *apiConfig) getRefreshToken(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		log.Printf("error getting bearer token: %v", err)
		w.WriteHeader(500)
		return
	}
	type jwebToken struct {
		Token string `json:"token"`
	}
	rfToken, err := cfg.db.GetResponseToken(r.Context(), token)
	if err != nil {
		log.Printf("error getting refresh token from table: %v", err)
		w.WriteHeader(500)
		return
	}

	if rfToken.ExpiresAt.Before(time.Now()) {
		w.WriteHeader(401)
		return
	}

	if rfToken.RevokedAt.Valid {
		w.WriteHeader(401)
		return
	}

	jwt, err := auth.MakeJWT(rfToken.UserID, cfg.jwtsecret)
	if err != nil {
		log.Printf("error making new jwt: %v", err)
		w.WriteHeader(500)
		return
	}
	validToken := jwebToken{
		Token: jwt,
	}
	val, err := json.Marshal(validToken)
	if err != nil {
		log.Printf("error marshalling json data for token: %v", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(val)

}

func (cfg *apiConfig) revokeRefreshToken(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		log.Printf("error getting bearer token: %v", err)
		w.WriteHeader(500)
		return
	}

	revokedAt := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	tokenParams := database.RevokeTokenParams{
		Token:     token,
		RevokedAt: revokedAt,
		UpdatedAt: time.Now(),
	}

	_, err2 := cfg.db.RevokeToken(r.Context(), tokenParams)
	if err2 != nil {
		log.Printf("error updating revoke token field %v", err)
		w.WriteHeader(500)
	}
	w.WriteHeader(204)
}
