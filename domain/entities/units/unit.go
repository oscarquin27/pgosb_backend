package unit_entity

import "github.com/jackc/pgx/v5/pgtype"


type Unit struct {
	Id pgtype.Int4 `json:"id"`
	Plate pgtype.Text `json:"plate"`
	Zone pgtype.Text `json:"zone"`
	Station pgtype.Text `json:"station"`
	Unit_type pgtype.Text `json:"unit_type"`
	Make pgtype.Text `json:"make"`
	Drivers pgtype.Int4 `json:"drivers"`
	Unit_condition pgtype.Text `json:"unit_condition"`
	Vehicle_serial pgtype.Text `json:"vehicle_serial"`
	Motor_serial pgtype.Text `json:"motor_serial"`
	Capacity pgtype.Text `json:"capacity"`
	Details pgtype.Text `json:"details"`
	Fuel_type pgtype.Text `json:"fuel_type"`
	Water_capacity pgtype.Text `json:"water_capacity"`
	Observations pgtype.Text `json:"observations"`
}

type UnitDto struct {
	Id string `json:"id"`
	Plate string `json:"plate"`
	Zone string `json:"zone"`
	Station string `json:"station"`
	Unit_type string `json:"unit_type"`
	Make string `json:"make"`
	Drivers string `json:"drivers"`
	Unit_condition string `json:"unit_condition"`
	Vehicle_serial string `json:"vehicle_serial"`
	Motor_serial string `json:"motor_serial"`
	Capacity string `json:"capacity"`
	Details string `json:"details"`
	Fuel_type string `json:"fuel_type"`
	Water_capacity string `json:"water_capacity"`
	Observations string `json:"observations"`
}