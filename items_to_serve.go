package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (cfg *apiConfig) itemsToServe(w http.ResponseWriter, r *http.Request) {
	items, err := cfg.db.GetItems(r.Context())
	if err != nil {
		log.Printf("error getting items from database")
		w.WriteHeader(500)
		return
	}

	type ItemResponse struct {
		ProductName        string `json:"product_name"`
		ProductDescription string `json:"product_description"`
		Price              int32  `json:"price"`
		InStock            bool   `json:"in_stock"`
	}

	var responseItems []ItemResponse
	for _, item := range items {
		responseItems = append(responseItems, ItemResponse{
			ProductName:        item.ProductName,
			ProductDescription: item.ProductDescription,
			Price:              item.Price,
			InStock:            item.InStock,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseItems)
	log.Printf("items returned: %v", responseItems)

}
