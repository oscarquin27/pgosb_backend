package models

import (
	"database/sql"

	"github.com/jackc/pgx/v5"
)

type MissionUnit struct {
	Id        int64        `db:"id"`
	CreatedAt sql.NullTime `db:"created_at"`
	MissionId int64        `db:"mission_id"`
	UnitId    int64        `db:"unit_id"`
}

func (s *MissionUnit) GetNameArgs() pgx.NamedArgs {
	return pgx.NamedArgs{
		"id":         s.Id,
		"mission_id": s.MissionId,
		"unit_id":    s.UnitId,
		"created_at": s.CreatedAt,
	}
}

func (s *MissionUnit) SetId(id int64) {
	s.Id = id
}

type MissionUnitSummary struct {
	Id        int64          `db:"id"`
	MissionId int64          `db:"mission_id"`
	UnitId    int64          `db:"unit_id"`
	Plate     sql.NullString `db:"plate"`
	Alias     sql.NullString `db:"alias"`
	UnitType  sql.NullString `db:"unit_type"`
}

func (s *MissionUnitSummary) GetNameArgs() pgx.NamedArgs {
	return pgx.NamedArgs{
		"id":         s.Id,
		"mission_id": s.MissionId,
		"unit_id":    s.UnitId,
		"plate":      s.Plate,
		"alias":      s.Alias,
		"unit_type":  s.UnitType,
	}
}

func (s *MissionUnitSummary) SetId(id int64) {
	s.Id = id
}
