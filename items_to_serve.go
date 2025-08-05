package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) itemsToServe(w http.ResponseWriter, r *http.Request) {
	items, err := cfg.db.GetItems(r.Context())
	if err != nil {
		log.Printf("error getting items from database")
		w.WriteHeader(500)
		return
	}

	type ItemResponse struct {
		Id                 uuid.UUID `json:"id"`
		ProductName        string    `json:"product_name"`
		ProductDescription string    `json:"product_description"`
		Price              int32     `json:"price"`
		InStock            bool      `json:"in_stock"`
		Image_url          string    `json:"image_url"`
	}

	var responseItems []ItemResponse
	for _, item := range items {
		responseItems = append(responseItems, ItemResponse{
			Id:                 item.ID,
			ProductName:        item.ProductName,
			ProductDescription: item.ProductDescription,
			Price:              item.Price,
			InStock:            item.InStock,
			Image_url:          item.ImageUrl,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	err2 := json.NewEncoder(w).Encode(responseItems)
	if err2 != nil {
		log.Printf("error encoding struct: %v", err2)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
