package config

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
)

var Storage *storage.Client

func setupStorageClient(configuration *Configuration) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	storage, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println("error creating cloud storage client: $w", err)
	}
	Storage = storage
}

func GetStorageClient() *storage.Client {
	return Storage
}
