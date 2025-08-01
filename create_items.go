package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/tristenkelly/the-trinity-pallette/internal/database"
)

func (cfg *apiConfig) createItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Printf("error parsing multipart form: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	itemID := uuid.New()

	filedata, header, err := r.FormFile("image")
	data, err := io.ReadAll(filedata)
	filedata.Close()
	if err != nil {
		log.Printf("error reading filedata in image %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctype := header.Header.Get("Content-Type")
	image := productImage{
		data:      data,
		mediaType: ctype,
	}
	mimeType, _, _ := mime.ParseMediaType(ctype)
	if mimeType != "image/jpeg" && mimeType != "image/png" {
		log.Printf("unsupported file type")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch ctype {
	case "image/png":
		ctype = ".png"
	case "image/jpeg":
		ctype = ".jpeg"
	}
	randkey := make([]byte, 32)
	_, err4 := io.ReadFull(rand.Reader, randkey)
	if err4 != nil {
		log.Printf("error reading randkey")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	encodedFile := base64.RawURLEncoding.EncodeToString(randkey)
	dir := filepath.Join("static/assets/", encodedFile+ctype)
	file, err := os.Create(dir)
	if err != nil {
		log.Printf("coudn't create file")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()
	_, err3 := filedata.Seek(0, io.SeekStart)
	if err3 != nil {
		log.Printf("coudn't seek start of filedata")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	io.Copy(file, filedata)
	filedata.Close()
	dataURL := fmt.Sprintf("http://localhost:%v/static/assets/%v%v", cfg.port, encodedFile, ctype)

	name := r.FormValue("product_name")
	description := r.FormValue("product_description")
	priceStr := r.FormValue("price")
	inStockStr := r.FormValue("in_stock")

	var price int
	fmt.Sscanf(priceStr, "%d", &price)

	var inStock bool
	if inStockStr == "true" || inStockStr == "1" {
		inStock = true
	}

	queryParams := database.CreateItemParams{
		ID:                 itemID,
		ProductName:        name,
		ProductDescription: description,
		Price:              int32(price),
		InStock:            inStock,
		UpdatedAt:          time.Now(),
		ImageUrl:           dataURL,
	}
	itemImages[itemID] = image
	item, err := cfg.db.CreateItem(r.Context(), queryParams)
	if err != nil {
		log.Printf("error creating item in items table: %v", err)
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
