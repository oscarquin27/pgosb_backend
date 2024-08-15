package models

import (
	"errors"

	"github.com/jackc/pgx/v5"
)

var (
	ErrorStationFound      = errors.New("estación no encontrada")
	ErrorStationtCreated   = errors.New("estación no creada")
	ErrorStationNotUpdated = errors.New("la Estación no pudo ser actualizada")
	ErrorStationNotDeleted = errors.New("la Estación no pudo ser eliminada")
)

type Station struct {
	Id           int64    `db:"id"`
	Name         *string  `db:"name"`
	Description  *string  `db:"description"`
	Abbreviation *string  `db:"abbreviation"`
	Phones       []string `db:"phones"`
	RegionId     *int64   `db:"region_id"`

	StateId *int64  `db:"state_id"`
	State   *string `db:"state"`

	MunicipalityId *int64  `db:"municipality_id"`
	Municipality   *string `db:"municipality"`

	ParishId *int64  `db:"parish_id"`
	Parish   *string `db:"parish"`

	SectorId *int64  `db:"sector_id"`
	Sector   *string `db:"sector"`

	UrbId *int64  `db:"urb_id"`
	Urb   *string `db:"urb"`

	Street *string `db:"street"`

	Address *string `db:"address"`
}

func (s *Station) SetId(id int64) {
	s.Id = id
}

func (s *Station) GetNameArgs() pgx.NamedArgs {
	return pgx.NamedArgs{
		"id":              s.Id,
		"name":            s.Name,
		"description":     s.Description,
		"abbreviation":    s.Abbreviation,
		"phones":          s.Phones,
		"region_id":       s.RegionId,
		"state_id":        s.StateId,
		"state":           s.State,
		"municipality_id": s.MunicipalityId,
		"municipality":    s.Municipality,
		"parish_id":       s.ParishId,
		"parish":          s.Parish,
		"sector_id":       s.SectorId,
		"sector":          s.Sector,
		"urb_id":          s.UrbId,
		"urb":             s.Urb,
		"street":          s.Street,
		"address":         s.Address,
	}
}
