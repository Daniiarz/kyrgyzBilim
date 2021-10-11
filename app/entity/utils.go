package entity

import (
	"fmt"
	"os"
)

func SetMediaUrl(s string) string {
	url := fmt.Sprintf("%v/%v", os.Getenv("MEDIA_URL"), s)
	return url
}
