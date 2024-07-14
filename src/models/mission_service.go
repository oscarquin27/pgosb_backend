package models

import (
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrorMissionServiceNotFound   = errors.New("servicio no encontrado")
	ErrorMissionServiceNotCreated = errors.New("servicio no creado")
	ErrorMissionServiceNotUpdated = errors.New("el servicio no pudo ser actualizado")
	ErrorMissionServiceNotDeleted = errors.New("el servicio no pudo ser eliminado")
)

type MissionService struct {
	Id		        	pgtype.Int4    	   `json:"id" binding:"required"`
	MissionId      	    pgtype.Int4    	   `json:"mission_id"`
	AntaresId			pgtype.Int4    	   `json:"antares_id"`
	Units               []pgtype.Int4      `json:"units"`
	Bombers           	[]pgtype.Int4      `json:"bombers"`
	Summary            	pgtype.Text    	   `json:"summary"`
	Description         pgtype.Text    	   `json:"description"`
}
