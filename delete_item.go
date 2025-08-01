package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) deleteItem(w http.ResponseWriter, r *http.Request) {
	log.Printf("deleting item")
	ItemIDStr := chi.URLParam(r, "itemID")
	itemID, err := uuid.Parse(ItemIDStr)
	if err != nil {
		log.Printf("error parsing itemid: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Deleting item with ID: %s", itemID)

	err2 := cfg.db.DeleteItem(r.Context(), itemID)
	if err2 != nil {
		log.Printf("error deleting item from items table: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("item deleted successfully")
	w.WriteHeader(http.StatusNoContent)

}
