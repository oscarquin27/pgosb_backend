package models

import (
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrorMissionNotFound   = errors.New("misión no encontrada")
	ErrorMissionNotCreated = errors.New("misión no creada")
	ErrorMissionNotUpdated = errors.New("la misión pudo ser actualizado")
	ErrorMissionNotDeleted = errors.New("la misión no pudo ser eliminada")
)

type Mission struct {
	Id		        	pgtype.Int4    	   `json:"id" binding:"required"`
	CreatedAt      	    pgtype.Date    	   `json:"created_at,omitempty"`
	Code				pgtype.Text    	   `json:"code"`
}
