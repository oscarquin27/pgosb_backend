package models

import (
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrorUnitNotFound   = errors.New("uninidad no encontrada")
	ErrorUnitNotCreated = errors.New("unidad no creada")
	ErrorUnitNotUpdated = errors.New("la unidad no pudo ser actualizada")
	ErrorUnitNotDeleted = errors.New("la unidad no pudo ser eliminada")
)

type Unit struct {
	Id             pgtype.Int4 `json:"id"`
	Plate          pgtype.Text `json:"plate"`
	Zone           pgtype.Text `json:"zone"`
	Station        pgtype.Text `json:"station"`
	Unit_type      pgtype.Text `json:"unit_type"`
	Make           pgtype.Text `json:"make"`
	Drivers        pgtype.Int4 `json:"drivers"`
	Unit_condition pgtype.Text `json:"unit_condition"`
	Vehicle_serial pgtype.Text `json:"vehicle_serial"`
	Motor_serial   pgtype.Text `json:"motor_serial"`
	Capacity       pgtype.Text `json:"capacity"`
	Details        []string    `json:"details"`
	Fuel_type      pgtype.Text `json:"fuel_type"`
	Water_capacity pgtype.Text `json:"water_capacity"`
	Observations   pgtype.Text `json:"observations"`
}
