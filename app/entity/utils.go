package entity

import (
	"fmt"
	"os"
)

func SetMediaUrl(s string) string {
	mediaUrl := os.Getenv("MEDIA_URL")
	fmt.Println(mediaUrl)
	url := fmt.Sprintf("%v/%v", mediaUrl, s)
	return url
}
