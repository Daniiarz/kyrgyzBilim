package entity

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func SetMediaUrl(s string) string {
	_ = godotenv.Load(".env")
	url := fmt.Sprintf("%v/%v", os.Getenv("MEDIA_URL"), s)
	return url
}
