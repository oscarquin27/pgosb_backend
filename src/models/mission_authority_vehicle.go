package models

import (
	"database/sql"

	"github.com/jackc/pgx/v5"
)

type MissionAuthorityVehicle struct {
	Id          int64          `db:"id"`
	MissionId   int64          `db:"mission_id"`
	AuthorityId int64          `db:"authority_id"`
	Type        sql.NullString `db:"type"`
	Make        sql.NullString `db:"make"`
	Model       sql.NullString `db:"model"`
	Plate       sql.NullString `db:"plate"`
	Year        sql.NullString `db:"year"`
	Color       sql.NullString `db:"color"`
	Description sql.NullString `db:"description"`
	CreatedAt   sql.NullTime   `db:"created_at"`
}

func (s *MissionAuthorityVehicle) SetId(id int64) {
	s.Id = id
}

func (s *MissionAuthorityVehicle) GetNameArgs() pgx.NamedArgs {

	args := make(pgx.NamedArgs)

	args["id"] = s.Id
	args["mission_id"] = s.MissionId

	args["authority_id"] = s.AuthorityId

	if s.Type.Valid {
		args["type"] = s.Type.String
	}

	if s.Make.Valid {
		args["make"] = s.Make.String
	}

	if s.Model.Valid {
		args["model"] = s.Model.String
	}

	if s.Plate.Valid {
		args["plate"] = s.Plate.String
	}

	if s.Year.Valid {
		args["year"] = s.Year.String
	}

	if s.Color.Valid {
		args["color"] = s.Color.String
	}

	if s.Description.Valid {
		args["description"] = s.Description.String
	}

	if s.CreatedAt.Valid {
		args["created_at"] = s.CreatedAt.Time
	}

	return args

}
