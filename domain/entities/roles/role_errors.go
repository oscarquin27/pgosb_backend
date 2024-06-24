package role_entity

import "errors"

var (
	ErrorRoleNotFound   = errors.New("Rol no encontrado")
	ErrorRoleNotCreated = errors.New("Rol no creado")
	ErrorRoleNotUpdated = errors.New("El rol no pudo ser actualizado")
	ErrorRoleNotDeleted = errors.New("El rol no pudo ser eliminado")
)
