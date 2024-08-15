package models

import (
	"github.com/jackc/pgx/v5"
)

type MissionLocation struct {
	Id        int64 `db:"id"`
	MissionId int64 `db:"mission_id"`

	Alias *string `db:"alias"`

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

	// Street *string `db:"street"`

	Address *string `db:"address"`
}

func (s *MissionLocation) SetId(id int64) {
	s.Id = id
}

func (s *MissionLocation) GetNameArgs() pgx.NamedArgs {
	return pgx.NamedArgs{
		"id":              s.Id,
		"mission_id":      s.MissionId,
		"alias":           s.Alias,
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
		"address":         s.Address,
	}
}
