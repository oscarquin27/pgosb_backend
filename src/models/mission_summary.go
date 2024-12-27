package models

import "database/sql"

type MissionSummary struct {
	Id                   int64            `db:"id"`
	Code                 sql.NullString   `db:"code"`
	Alias                sql.NullString   `db:"alias"`
	AntaresId            sql.NullInt64    `db:"antares_id"`
	AntaresDescription   sql.NullString   `db:"antares_description"`
	NumServices          sql.NullInt64    `db:"num_services"`
	IsImportant          bool             `db:"is_important"`
	StationName          sql.NullString   `db:"station_name"`
	Commander            sql.NullString   `db:"commander"`
	NumFirefighters      sql.NullInt64    `db:"num_firefighters"`
	FirefighterArray     []sql.NullString `db:"firefighter_array"`
	NumVehicles          sql.NullInt64    `db:"num_vehicles"`
	UnitArray            []sql.NullString `db:"unit_array"`
	OperativesAreas      []sql.NullString `db:"operative_areas"`
	Description          sql.NullString   `db:"description"`
	Level                sql.NullString   `db:"level"`
	PeaceQuadrant        sql.NullString   `db:"peace_quadrant"`
	PendingForData       bool             `db:"pending_for_data"`
	CancelReason         sql.NullString   `db:"cancel_reason"`
	CreatedAt            sql.NullTime     `db:"created_at"`
	ManualMissionDate    sql.NullTime     `db:"manual_mission_date"`
	Unharmed             sql.NullInt64    `db:"unharmed"`
	Injured              sql.NullInt64    `db:"injured"`
	Transported          sql.NullInt64    `db:"transported"`
	Deceased             sql.NullInt64    `db:"deceased"`
	LocationId           sql.NullInt64    `db:"location_id"`
	LocationDestinyId    sql.NullInt64    `db:"location_destiny_id"`
	NumAuthorities       sql.NullInt64    `db:"num_authorities"`
	NumAuthorityServices sql.NullInt64    `db:"num_authority_services"`
	NumAuthorityPerson   sql.NullInt64    `db:"num_authority_person"`
	NumAuthorityVehicle  sql.NullInt64    `db:"num_authority_vehicle"`
	StationId            sql.NullInt64    `db:"station_id"`
}
