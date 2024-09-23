package models

import (
	"database/sql"

	"github.com/jackc/pgx/v5"
)

type MissionAuthority struct {
	Id        int64          `db:"id"`
	MissionId int64          `db:"mission_id"`
	Type      sql.NullString `db:"type"`
	CreatedAt sql.NullTime   `db:"created_at"`
	Alias     sql.NullString `db:"alias"`
}

type MissionAuthoritySummary struct {
	Id        int64          `db:"id"`
	MissionId int64          `db:"mission_id"`
	CreatedAt sql.NullTime   `db:"created_at"`
	Alias     sql.NullString `db:"alias"`
	Vehicles  sql.NullInt64  `db:"vehicles"`
	People    sql.NullInt64  `db:"people"`
	Type      sql.NullString `db:"type"`
	Services  sql.NullInt64  `db:"services"`
}

func (s *MissionAuthority) SetId(id int64) {

	s.Id = id

}

func (s *MissionAuthority) GetNameArgs() pgx.NamedArgs {

	args := make(pgx.NamedArgs)

	args["id"] = s.Id
	args["mission_id"] = s.MissionId
	if s.Alias.Valid {
		args["alias"] = s.Alias.String
	}

	return args
}
