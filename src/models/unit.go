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
	Station        pgtype.Text `json:"station"`
	Unit_type      pgtype.Text `json:"unit_type"`
	Make           pgtype.Text `json:"make"`
	Unit_condition pgtype.Text `json:"unit_condition"`
	Vehicle_serial pgtype.Text `json:"vehicle_serial"`
	Motor_serial   pgtype.Text `json:"motor_serial"`
	Capacity       pgtype.Text `json:"capacity"`
	Details        []string    `json:"details"`
	Fuel_type      pgtype.Text `json:"fuel_type"`
	Water_capacity pgtype.Text `json:"water_capacity"`
	Observations   pgtype.Text `json:"observations"`
	Hurt_capacity  pgtype.Int4 `json:"hurt_capacity"`
	Doors          pgtype.Int4 `json:"doors"`
	Performance    pgtype.Text `json:"performance"`
	Load_capacity  pgtype.Int4 `json:"load_capacity"`
	Model          pgtype.Text `json:"model"`
	Alias          pgtype.Text `json:"alias"`
	Color          pgtype.Text `json:"color"`
	Year           pgtype.Text `json:"year"`
	Purpose        pgtype.Text `json:"purpose"`
	Init_kilometer pgtype.Int4 `json:"init_kilometer"`
}
