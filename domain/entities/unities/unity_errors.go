package unity_entity

import "errors"

var (
	ErrorUnityNotFound   = errors.New("Unidad no encontrada")
	ErrorUnityNotCreated = errors.New("Unidad no creada")
	ErrorUnityNotUpdated = errors.New("La unidad no pudo ser actualizada")
	ErrorUnityNotDeleted = errors.New("La unidad no pudo ser eliminada")
)
