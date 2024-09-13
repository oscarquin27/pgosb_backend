package models

import (
	"database/sql"

	"github.com/jackc/pgx/v5"
)

type MissionAuthority struct {
	Id              sql.NullInt64  `db:"id"`
	Type            sql.NullString `db:"type"`
	Name            sql.NullString `db:"name"`
	Lastname        sql.NullString `db:"lastname"`
	Rank            sql.NullString `db:"rank"`
	Identification  sql.NullString `db:"identification"`
	Phone           sql.NullString `db:"phone"`
	NumberVehicle   sql.NullString `db:"number_vehicle"`
	DetailOfVehicle sql.NullString `db:"detail_of_vehicle"`
}

func (s *MissionAuthority) SetId(id int64) {
	s.Id.Valid = true
	s.Id.Int64 = id
}

func (s *MissionAuthority) GetNameArgs() pgx.NamedArgs {

	args := make(pgx.NamedArgs)

	if s.Id.Valid {
		args["id"] = s.Id.Int64
	}

	if s.Type.Valid {
		args["type"] = s.Type.String
	}

	if s.Name.Valid {
		args["name"] = s.Name.String
	}

	if s.Lastname.Valid {
		args["lastname"] = s.Lastname.String
	}

	if s.Identification.Valid {
		args["identification"] = s.Identification.String
	}

	if s.Phone.Valid {
		args["phone"] = s.Phone.String
	}

	if s.NumberVehicle.Valid {
		args["number_vehicle"] = s.NumberVehicle.String
	}

	if s.DetailOfVehicle.Valid {
		args["detail_of_vehicle"] = s.DetailOfVehicle.String
	}

	return args
}
