package entity

import (
	"fmt"
	"log"
	"os"
)

func SetMediaUrl(s string) string {
	mediaUrl := os.Getenv("MEDIA_URL")
	log.Println(mediaUrl)
	url := fmt.Sprintf("%v/%v", mediaUrl, s)
	return url
}
