package usecase

import (
	sf "GRPC-AUTH/internal/domain/service"
	"GRPC-AUTH/internal/utils"
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	"cloud.google.com/go/storage"
)

type gcImageRepository struct {
	Storage    *storage.Client
	BucketName string
}

func FileService(gcClient *storage.Client, bucketName string) sf.IFile {
	return &gcImageRepository{
		Storage:    gcClient,
		BucketName: bucketName,
	}
}

func (r *gcImageRepository) SetFile(ctx context.Context, field, routeSprintf string) (string, error) {
	buf, err := utils.DecodeFile(field)
	if err != nil {
		return "", errors.New("no es posible decodificar el archivo especificado")
	}
	randomName := fmt.Sprintf(routeSprintf, utils.RandomNameObjectUUID())

	bckt := r.Storage.Bucket(r.BucketName)
	object := bckt.Object(randomName)
	wc := object.NewWriter(ctx)
	wc.ObjectAttrs.CacheControl = "Cache-Control:no-cache, max-age=0"

	if _, err := io.Copy(wc, buf); err != nil {
		log.Printf("no se puede escribir el archivo en Google Cloud Storage: %v\n", err)
		return "", errors.New("no se puede escribir el archivo en Google Cloud Storage")
	}

	if err := wc.Close(); err != nil {
		fmt.Println(err)
		return "", errors.New("problemas de acceso al sistema de almacenamiento")
	}

	imageURL := fmt.Sprintf(
		"https://storage.googleapis.com/%s/%s",
		r.BucketName,
		randomName,
	)

	return imageURL, nil
}

func (r *gcImageRepository) DeleteFile(ctx context.Context, routeSprintf, name string) error {
	bckt := r.Storage.Bucket(r.BucketName)
	objectName := fmt.Sprintf(routeSprintf, name)
	object := bckt.Object(objectName)

	if err := object.Delete(ctx); err != nil {
		return errors.New("no fue posible eliminar el archivo con el ID especificado")
	}

	return nil
}
