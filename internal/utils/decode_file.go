package utils

import (
	"bytes"
	"encoding/base64"
	"errors"
)

func DecodeFile(strBase64 string) (*bytes.Buffer, error) {
	data, err := base64.StdEncoding.DecodeString(strBase64)
	if err != nil {
		return nil, errors.New("no es posible decodificar el archivo debido a errores existentes")
	}

	buf := bytes.NewBuffer(data)
	return buf, nil
}
