package utils

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func ValidateError(err error) string {
	var mysqlErr *mysql.MySQLError
	switch {
	case errors.As(err, &mysqlErr) && mysqlErr.Number == 1062:
		return "Datos duplicados, asignados a otro usuario "
	case errors.As(err, &mysqlErr) && mysqlErr.Number == 1452:
		return "Datos inexistentes "
	case errors.Is(err, gorm.ErrRecordNotFound):
		return "No se pudo encontrar el dato "
	default:
		return "Registro no creado "
	}
}
