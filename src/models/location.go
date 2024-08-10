package models

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrorStateFound      = errors.New("estado no encontrado")
	ErrorStateNotCreated = errors.New("estado no creado")
	ErrorStateNotUpdated = errors.New("el Estado no pudo ser actualizado")
	ErrorStateNotDeleted = errors.New("el Estado no pudo ser eliminado")

	ErrorMunicipalityFound      = errors.New("municipio no encontrado")
	ErrorMunicipalityNotCreated = errors.New("municipio no creado")
	ErrorMunicipalityNotUpdated = errors.New("el Municipio no pudo ser actualizado")
	ErrorMunicipalityNotDeleted = errors.New("el Municipio no pudo ser eliminado")

	ErrorParishFound      = errors.New("parroquia no encontrada")
	ErrorParishNotCreated = errors.New("parroquia no creada")
	ErrorParishNotUpdated = errors.New("la Parroquia no pudo ser actualizada")
	ErrorParishNotDeleted = errors.New("la Parroquia no pudo ser eliminada")

	ErrorCityFound      = errors.New("ciudad no encontrada")
	ErrorCityNotCreated = errors.New("ciudad no creada")
	ErrorCityNotUpdated = errors.New("la Ciudad no pudo ser actualizada")
	ErrorCityNotDeleted = errors.New("la Ciudad no pudo ser eliminada")
)

type State struct {
	Id          int64   `db:"id"`
	Name        *string `db:"name"`
	Capital     *string `db:"capital"`
	Coordinates *string `db:"coordinates"`
}

func (s *State) SetId(id int64) {
	s.Id = id
}

func (s *State) GetNameArgs() pgx.NamedArgs {
	return pgx.NamedArgs{
		"id":          s.Id,
		"name":        s.Name,
		"coordinates": s.Coordinates,
		"capital":     s.Capital,
	}
}

type Municipality struct {
	Id          int64   `db:"id"`
	StateId     int64   `db:"state_id"`
	Name        *string `db:"name"`
	Capital     *string `db:"capital"`
	Coordinates *string `db:"coordinates"`
}

func (s *Municipality) SetId(id int64) {
	s.Id = id
}

func (s *Municipality) GetNameArgs() pgx.NamedArgs {
	return pgx.NamedArgs{
		"id":          s.Id,
		"state_id":    s.StateId,
		"name":        s.Name,
		"coordinates": s.Coordinates,
		"capital":     s.Capital,
	}
}

type Parish struct {
	Id             int64   `db:"id"`
	StateId        int64   `db:"state_id"`
	MunicipalityId int64   `db:"municipality_id"`
	Name           *string `db:"name"`
	Capital        *string `db:"capital"`
	Coordinates    *string `db:"coordinates"`
}

func (s *Parish) SetId(id int64) {
	s.Id = id
}

func (s *Parish) GetNameArgs() pgx.NamedArgs {
	return pgx.NamedArgs{
		"id":          s.Id,
		"state_id":    s.StateId,
		"name":        s.Name,
		"coordinates": s.Coordinates,
		"capital":     s.Capital,
	}
}

type City struct {
	Id          pgtype.Int8 `json:"id" db:"city_id"`
	State_Id    pgtype.Int8 `json:"state_id"`
	Name        pgtype.Text `json:"name"`
	Area_Code   pgtype.Text `json:"area_code"`
	Zip_Code    pgtype.Text `json:"zip_code"`
	Coordinates pgtype.Text `json:"Coordinates"`
}
