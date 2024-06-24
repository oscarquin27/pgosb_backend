package state_entity

import "errors"

var (
	ErrorStateFound   = errors.New("Estado no encontrado")
	ErrorStateNotCreated = errors.New("Estado no creado")
	ErrorStateNotUpdated = errors.New("El Estado no pudo ser actualizado")
	ErrorStateNotDeleted = errors.New("El Estado no pudo ser eliminado")
)