package models

import (
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrorStationFound      = errors.New("estación no encontrada")
	ErrorStationtCreated   = errors.New("estación no creada")
	ErrorStationNotUpdated = errors.New("la Estación no pudo ser actualizada")
	ErrorStationNotDeleted = errors.New("la Estación no pudo ser eliminada")
)

type Phones struct {
	AreaCode string `json:"area_code"`
	Number   string `json:"number"`
}

type Station struct {
	Id              pgtype.Int4 `json:"id" db:"station_id"`
	Municipality_id pgtype.Int4 `json:"municipality_id"`
	Name            pgtype.Text `json:"name"`
	Coordinates     pgtype.Text `json:"coordinates"`
	Description     pgtype.Text `json:"description"`
	Code            pgtype.Text `json:"code"`
	Abbreviation    pgtype.Text `json:"abbreviation"`
	Phones          []Phones    `json:"phones"`
	State_id        pgtype.Int4 `json:"state_id"`
	Parish_id       pgtype.Int4 `json:"parish_id"`
	Sector          pgtype.Text `json:"sector"`
	Community       pgtype.Text `json:"community"`
	Street          pgtype.Text `json:"street"`
	Address         pgtype.Text `json:"address"`
}

type State struct {
	Id          pgtype.Int4 `json:"id" db:"state_id" mk:"state_id"`
	Name        pgtype.Text `json:"name" mk:"name"`
	Coordinates pgtype.Text `json:"coordinates" mk:"coordinates"`
}
