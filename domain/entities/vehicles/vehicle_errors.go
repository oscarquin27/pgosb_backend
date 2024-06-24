package vehicle_entity

import "errors"

var (
	ErrorVehicleNotFound   = errors.New("Vehículo no encontrado")
	ErrorVehicleNotCreated = errors.New("Vehículo no creado")
	ErrorVehicleNotUpdated = errors.New("El vehículo no pudo ser actualizado")
	ErrorVehicleNotDeleted = errors.New("El vehículo no pudo ser eliminado")
)
