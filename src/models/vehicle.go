package models

import (
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrorVehicleNotFound   = errors.New("vehículo no encontrado")
	ErrorVehicleNotCreated = errors.New("vehículo no creado")
	ErrorVehicleNotUpdated = errors.New("el vehículo no pudo ser actualizado")
	ErrorVehicleNotDeleted = errors.New("el vehículo no pudo ser eliminado")
)

type Vehicle struct {
	Id                  pgtype.Int8 `json:"id"`
	Make                pgtype.Text `json:"make"`
	Model               pgtype.Text `json:"model"`
	Year                pgtype.Text `json:"year"`
	Drive               pgtype.Text `json:"drive"`
	Cylinders           pgtype.Text `json:"cylinders"`
	Engine_displacement pgtype.Text `json:"engine_displacement"`
	Fuel_type           pgtype.Text `json:"fuel_type"`
	Transmission        pgtype.Text `json:"transmission"`
	Vehicle_size_class  pgtype.Text `json:"vehicle_size_class"`
	Base_model          pgtype.Text `json:"base_model"`
}

type VehicleSimple struct {
	Id                  pgtype.Int8 `json:"id"`
	Make                pgtype.Text `json:"make"`
	Model               pgtype.Text `json:"model"`
}
