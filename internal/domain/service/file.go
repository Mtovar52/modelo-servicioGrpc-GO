package service_file

import (
	"context"
)

type IFile interface {
	// Realiza la subida a Google Cloud. El primer parámetro corresponde al contexto.
	// Segundo parámetro es el valor actual del campo, debe estar codificado en Base64.
	// Tercer parámetro contiene la ruta con la especificación de se extensión.
	SetFile(ctx context.Context, field, routeSprintf string) (string, error)
	DeleteFile(ctx context.Context, routeSprintf, name string) error
}
