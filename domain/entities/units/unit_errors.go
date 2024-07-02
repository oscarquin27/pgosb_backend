package unit_entity

import "errors"

var (
	ErrorUnitNotFound   = errors.New("Unidad no encontrada")
	ErrorUnitNotCreated = errors.New("Unidad no creada")
	ErrorUnitNotUpdated = errors.New("La unidad no pudo ser actualizada")
	ErrorUnitNotDeleted = errors.New("La unidad no pudo ser eliminada")
)
