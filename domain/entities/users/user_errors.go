package user_entity

import "errors"

var (
	ErrorUserNotFound   = errors.New("Usuario no encontrado")
	ErrorUserNotCreated = errors.New("Usuario no creado")
	ErrorUserNotUpdated = errors.New("El usuario no pudo ser actualizado")
	ErrorUserNotDeleted = errors.New("El usuario no pudo ser eliminado")
)
