package models

import (
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrorMissionVehicleNotFound   = errors.New("vehículo no encontrado")
	ErrorMissionVehicleNotCreated = errors.New("vehículo no creado")
	ErrorMissionVehicleNotUpdated = errors.New("el vehículo no pudo ser actualizado")
	ErrorMissionVehicleNotDeleted = errors.New("el vehículo no pudo ser eliminado")
)

type MissionVehicle struct {
	Id               pgtype.Int4 `json:"id" db:"id"`
	ServiceId        pgtype.Int4 `json:"service_id" db:"service_id"`
	VehicleCondition pgtype.Text `json:"vehicle_condition" db:"vehicle_condition"`
	Make             pgtype.Text `json:"make" db:"make"`
	Model            pgtype.Text `json:"model" db:"model"`
	Year             pgtype.Text `json:"year" db:"year"`
	Plate            pgtype.Text `json:"plate" db:"plate"`
	Color            pgtype.Text `json:"color" db:"color"`
	VehicleType      pgtype.Text `json:"vehicle_type" db:"vehicle_type"`
	MotorSerial      pgtype.Text `json:"motor_serial" db:"motor_serial"`
	VehicleVerified  pgtype.Bool `json:"vehicle_verified" db:"vehicle_verified"`
}
