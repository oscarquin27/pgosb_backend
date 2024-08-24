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
	Id        pgtype.Int4 `db:"id"`
	CreatedAt pgtype.Date `db:"created_at,omitempty"`
	Code      pgtype.Text `db:"code"`
	Alias     *string     `db:"alias"`
}
