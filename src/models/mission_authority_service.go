package models

import (
	"database/sql"

	"github.com/jackc/pgx/v5"
)

type MissionAuthorityService struct {
	Id          int64 `db:"id"`
	MissionId   int64 `db:"mission_id"`
	ServiceId   int64 `db:"service_id"`
	AuthorityId int64 `db:"authority_id"`
}

type MissionAuthorityServiceSummary struct {
	Id          int64 `db:"main_id"`
	AuthorityId int64 `db:"id"`
	MissionId   int64 `db:"mission_id"`

	CreatedAt sql.NullTime   `db:"created_at"`
	Alias     sql.NullString `db:"alias"`
	Vehicles  sql.NullInt64  `db:"vehicles"`
	People    sql.NullInt64  `db:"people"`
	Type      sql.NullString `db:"type"`
	Services  sql.NullInt64  `db:"services"`
}

func (s *MissionAuthorityService) SetId(id int64) {
	s.Id = id
}

func (s *MissionAuthorityService) GetNameArgs() pgx.NamedArgs {

	args := make(pgx.NamedArgs)

	args["id"] = s.Id
	args["mission_id"] = s.MissionId
	args["service_id"] = s.ServiceId
	args["authority_id"] = s.AuthorityId

	return args

}

func (s *MissionAuthorityServiceSummary) GetNameArgs() pgx.NamedArgs {

	args := make(pgx.NamedArgs)

	args["id"] = s.Id
	args["mission_id"] = s.MissionId
	args["authority_id"] = s.AuthorityId

	if s.CreatedAt.Valid {
		args["created_at"] = s.CreatedAt.Time
	}

	if s.Alias.Valid {
		args["alias"] = s.Alias.String
	}

	if s.Vehicles.Valid {
		args["vehicles"] = s.Vehicles.Int64
	}

	if s.People.Valid {
		args["people"] = s.People.Int64
	}

	if s.Type.Valid {
		args["type"] = s.Type.String
	}

	if s.Services.Valid {
		args["services"] = s.Services.Int64
	}

	return args

}
