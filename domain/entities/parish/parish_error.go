package parish_entity

import "errors"

var (
	ErrorParishFound   = errors.New("Parroquia no encontrada")
	ErrorParishNotCreated = errors.New("Parroquia no creada")
	ErrorParishNotUpdated = errors.New("La Parroquia no pudo ser actualizada")
	ErrorParishNotDeleted = errors.New("La Parroquia no pudo ser eliminada")
)