package entity

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func SetMediaUrl(s string) string {
	_ = godotenv.Load("/usr/src/app/.env")
	url := fmt.Sprintf("%v/%v", os.Getenv("MEDIA_URL"), s)
	return url
}
