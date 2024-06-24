package unity_entity

import "github.com/jackc/pgx/v5/pgtype"


type Unity struct {
	Id pgtype.Int8 `json:"id"`
	Plate pgtype.Text `json:"plate"`
	Zone pgtype.Text `json:"zone"`
	Station pgtype.Text `json:"station"`
	Unity_type pgtype.Text `json:"unity_type"`
	Make pgtype.Text `json:"make"`
	Drivers pgtype.Int8 `json:"drivers"`
	Unity_condition pgtype.Text `json:"unity_condition"`
	Vehicle_serial pgtype.Text `json:"vehicle_serial"`
	Motor_serial pgtype.Text `json:"motor_serial"`
	Capacity pgtype.Text `json:"capacity"`
	Details pgtype.Text `json:"details"`
	Fuel_type pgtype.Text `json:"fuel_type"`
	Water_capacity pgtype.Text `json:"water_capacity"`
	Observations pgtype.Text `json:"observations"`
}