package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/tristenkelly/the-trinity-pallette/internal/auth"
	"github.com/tristenkelly/the-trinity-pallette/internal/database"
)

func (cfg *apiConfig) signUpHandler(w http.ResponseWriter, r *http.Request) {
	type userParams struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := userParams{}
	log.Println(params)
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("error decoding user data %v", err)
		w.WriteHeader(500)
		return
	}

	hashedPass, err := auth.HashPassword(params.Password)
	if err != nil {
		log.Printf("couldn't hash password: %v", err)
		w.WriteHeader(500)
		return
	}

	userID := uuid.New()

	queryParams := database.CreateUserParams{
		ID:             userID,
		Username:       params.Username,
		Email:          params.Email,
		HashedPassword: hashedPass,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		IsAdmin:        false,
	}

	user, err := cfg.db.CreateUser(r.Context(), queryParams)
	if err != nil {
		log.Printf("error creating user in db: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, err := auth.MakeJWT(user.ID, cfg.jwtsecret)
	if err != nil {
		log.Printf("error making jwt: %v", err)
		w.WriteHeader(500)
		return
	}

	type returnUser struct {
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		Token     string    `json:"token"`
	}

	returnUserInfo := returnUser{
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		Token:     token,
	}

	val, err := json.Marshal(returnUserInfo)
	if err != nil {
		log.Printf("couldn't return user data: %v", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	_, err2 := w.Write(val)
	if err2 != nil {
		log.Printf("error writing response: %v", err2)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
