package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/tristenkelly/the-trinity-pallette/internal/database"
)

func (cfg *apiConfig) createItem(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Product_name        string `json:"product_name"`
		Product_description string `json:"product_description"`
		Price               int    `json:"price"`
		In_stock            bool   `json:"in_stock"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("error decoding createitem response body")
		w.WriteHeader(500)
		return
	}

	itemID := uuid.New()

	queryParams := database.CreateItemParams{
		ID:                 itemID,
		ProductName:        params.Product_name,
		ProductDescription: params.Product_description,
		Price:              int32(params.Price),
		InStock:            params.In_stock,
		UpdatedAt:          time.Now(),
	}

	item, err := cfg.db.CreateItem(r.Context(), queryParams)
	if err != nil {
		log.Printf("error creating item in items table")
		w.WriteHeader(500)
		return
	}

	val, err := json.Marshal(item)
	if err != nil {
		log.Printf("error marshalling createitem val: %v", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(val)
}
