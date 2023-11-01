package file

import (
	"bytes"
	"context"
	"encoding/base64"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
)

func UploadFile(file string) (string, error) {
	cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_CLOUD_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))
	uploadParams := uploader.UploadParams{
		PublicID: generateRandomString(10),
		Folder:   "FoodRecipe",
	}
	fileBytes, err := base64.StdEncoding.DecodeString(file)
	if err != nil {
		return "", err
	}
	uploadResult, err := cld.Upload.Upload(context.Background(), bytes.NewReader(fileBytes), uploadParams)

	if err != nil {
		return "", err
	}
	return uploadResult.SecureURL, nil
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	random := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		random[i] = charset[rand.Intn(len(charset))]
	}
	return string(random)
}
func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}
