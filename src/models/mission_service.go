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
	MissionId      	    pgtype.Int2    	   `json:"mission_id"`
	AntaresId			pgtype.Int2    	   `json:"antares_id,omitempty"`
	Units               []pgtype.Int2      `json:"units,omitempty"`
	Bombers           	[]pgtype.Int2      `json:"bombers,omitempty"`
	Summary            	pgtype.Text    	   `json:"summary,omitempty"`
	Description         pgtype.Text    	   `json:"description,omitempty"`
}
