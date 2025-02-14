package models

import "database/sql"

type MissionServiceSummary struct {
	Id                   int64          `db:"id"`
	Alias                sql.NullString `db:"alias"`
	Code                 sql.NullString `db:"code"`
	CreatedAt            sql.NullTime   `db:"created_at"`
	NumServices          sql.NullInt64  `db:"num_services"`
	Unharmed             sql.NullInt64  `db:"unharmed"`
	Injured              sql.NullInt64  `db:"injured"`
	Transported          sql.NullInt64  `db:"transported"`
	Deceased             sql.NullInt64  `db:"deceased"`
	NumVehicles          sql.NullInt64  `db:"num_vehicles"`
	NumFirefighters      sql.NullInt64  `db:"num_firefighters"`
	OperativesAreas      []string       `db:"operative_areas"`
	NumAuthorities       sql.NullInt64  `db:"num_authorities"`
	NumAuthorityServices sql.NullInt64  `db:"num_authority_services"`
	NumAuthorityPerson   sql.NullInt64  `db:"num_authority_person"`
	NumAuthorityVehicle  sql.NullInt64  `db:"num_authority_vehicle"`
	IsImportant          bool           `db:"is_important"`
	ManualMissionDate    sql.NullTime   `db:"manual_mission_date"`
	StationName          sql.NullString `db:"station_name"`
	AntaresId            sql.NullInt64  `db:"antares_id"`
	AntaresDescription   sql.NullString `db:"antares_description"`
}
