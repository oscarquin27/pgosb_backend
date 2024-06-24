package city_entity

import "errors"

var (
	ErrorCityFound   = errors.New("Ciudad no encontrada")
	ErrorCityNotCreated = errors.New("Ciudad no creada")
	ErrorCityNotUpdated = errors.New("La Ciudad no pudo ser actualizada")
	ErrorCityNotDeleted = errors.New("La Ciudad no pudo ser eliminada")
)