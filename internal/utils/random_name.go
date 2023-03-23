package utils

import (
	"github.com/google/uuid"
)

func RandomNameObjectUUID() string {
	objID, _ := uuid.NewRandom()
	return objID.String()
}

/* func RandomNameObjectUUIDUpdate(imageURL string) (string, error) {
	urlPath, err := url.Parse(imageURL)

	if err != nil {
		log.Printf("No se pudo extraer el path del url: %v\n", imageURL)
	}

	return path.Base(urlPath.Path), nil
} */
