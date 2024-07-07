package role_entity

import "errors"

var (
	ErrorRoleNotFound   = errors.New("rol no encontrado")
	ErrorRoleNotCreated = errors.New("rol no creado")
	ErrorRoleNotUpdated = errors.New("el rol no pudo ser actualizado")
	ErrorRoleNotDeleted = errors.New("el rol no pudo ser eliminado")
)
