package station_entity

import "errors"

var (
	ErrorStationFound   = errors.New("Estación no encontrada")
	ErrorStationtCreated = errors.New("Estación no creada")
	ErrorStationNotUpdated = errors.New("La Estación no pudo ser actualizada")
	ErrorStationNotDeleted = errors.New("La Estación no pudo ser eliminada")
)