package municipality_entity

import "errors"

var (
	ErrorMunicipalityFound   = errors.New("Municipio no encontrado")
	ErrorMunicipalityNotCreated = errors.New("Municipio no creado")
	ErrorMunicipalityNotUpdated = errors.New("El Municipio no pudo ser actualizado")
	ErrorMunicipalityNotDeleted = errors.New("El Municipio no pudo ser eliminado")
)