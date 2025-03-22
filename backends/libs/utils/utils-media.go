package utils

import (
	"log"
	"regexp"

	"github.com/adrisongomez/pti-ecommerce-site/backends/databases/db"
)

const Image db.MediaType = "IMAGE"
const Video db.MediaType = "VIDEO"
const Unknown db.MediaType = "UNKWNON"

func GetMediaTypeByMimeType(mimeType string) db.MediaType {
	if IsImage(mimeType) {
		return Image
	}
	if IsVideo(mimeType) {
		return Video
	}
	return Unknown
}

func IsImage(mimeType string) bool {
	re, err := regexp.Compile("^images/.*")
	if err != nil {
		log.Println("Error in matching function for images")
		return false
	}
	return re.MatchString(mimeType)

}

func IsVideo(mimeType string) bool {
	re, err := regexp.Compile("^video/.*")
	if err != nil {
		log.Println("Error in matching function for video")
		return false
	}
	return re.MatchString(mimeType)
}
