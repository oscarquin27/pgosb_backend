package models

import (
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrorMissionPersonNotFound   = errors.New("persona no encontrado")
	ErrorMissionPersonNotCreated = errors.New("persona no creado")
	ErrorMissionPersonNotUpdated = errors.New("la persona no pudo ser actualizada")
	ErrorMissionPersonNotDeleted = errors.New("la persona no pudo ser eliminada")
)

type MissionPerson struct {
	Id		        	pgtype.Int4    	   `json:"id" binding:"required"`
	ServiceId      	    pgtype.Int4    	   `json:"service_id"`
	UnitId       		pgtype.Int4    	   `json:"unit_id"`
	InfrastructureId    pgtype.Int4    	   `json:"infrastructure_id"`
	VehicleId           pgtype.Int4    	   `json:"vehicle_id"`
	FirstName           pgtype.Text    	   `json:"first_name"`
	LastName            pgtype.Text    	   `json:"last_name"`
	Age                 pgtype.Int2    	   `json:"age"`
	Gender              pgtype.Text    	   `json:"gender"`
	LegalId             pgtype.Text    	   `json:"legal_id"`
	Phone               pgtype.Text    	   `json:"phone"`
	Employment          pgtype.Text    	   `json:"employment"`
	State               pgtype.Text    	   `json:"state"`
	Municipality        pgtype.Text    	   `json:"municipality"`
	Parish              pgtype.Text 	   `json:"parish"`
	Address             pgtype.Text 	   `json:"address"`
	Pathology           pgtype.Text    	   `json:"pathology"`
	Observations        pgtype.Text    	   `json:"observations"`
	Condition           pgtype.Text 	   `json:"condition"`
}
