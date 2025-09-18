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
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

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
	if err != nil {
		log.Printf("error getting filedata from form: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, err := io.ReadAll(filedata)
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

	staticBucketName := os.Getenv("STATIC_BUCKET")
	if staticBucketName == "" {
		staticBucketName = "the-trinity-pallette-static-assets"
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	if err != nil {
		log.Printf("error creating AWS session: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s3Client := s3.New(sess)

	_, err3 := filedata.Seek(0, io.SeekStart)
	if err3 != nil {
		log.Printf("couldn't seek start of filedata")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	key := fmt.Sprintf("assets/%s%s", encodedFile, ctype)
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(staticBucketName),
		Key:         aws.String(key),
		Body:        filedata,
		ContentType: aws.String(mime.TypeByExtension(ctype)),
		ACL:         aws.String("public-read"),
	})
	if err != nil {
		log.Printf("error uploading to S3: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err2 := filedata.Close()
	if err2 != nil {
		log.Printf("error closing filedata: %v", err2)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	staticBaseURL := os.Getenv("STATIC_BASE_URL")
	if staticBaseURL == "" {
		staticBaseURL = fmt.Sprintf("https://%s.s3.amazonaws.com", staticBucketName)
	}
	dataURL := fmt.Sprintf("%s/%s", staticBaseURL, key)

	name := r.FormValue("product_name")
	description := r.FormValue("product_description")
	priceStr := r.FormValue("price")
	inStockStr := r.FormValue("in_stock")

	var price int32
	_, err5 := fmt.Sscanf(priceStr, "%d", &price)
	if err5 != nil {
		log.Printf("error parsing price: %v", err5)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var inStock bool
	if inStockStr == "true" || inStockStr == "1" {
		inStock = true
	}

	queryParams := database.CreateItemParams{
		ID:                 itemID,
		ProductName:        name,
		ProductDescription: description,
		Price:              price,
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
	_, err7 := w.Write(val)
	if err7 != nil {
		log.Printf("error writing response: %v", err7)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
