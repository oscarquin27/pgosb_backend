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
	Id		        	pgtype.Int4    	   `json:"id" binding:"required"`
	ServiceId      	    pgtype.Int4    	   `json:"service_id"`
	VehicleCondition    pgtype.Text    	   `json:"vehicle_condition"`
	Make                pgtype.Text    	   `json:"make"`
	Model           	pgtype.Text    	   `json:"model"`
	Year            	pgtype.Text    	   `json:"year"`
	Plate            	pgtype.Text    	   `json:"plate"`
	Color            	pgtype.Text    	   `json:"color"`
	VehicleType         pgtype.Text    	   `json:"vehicle_type"`
	MotorSerial         pgtype.Text    	   `json:"motor_serial"`
	VehicleVerified     pgtype.Bool    	   `json:"vehicle_verified"`
}
