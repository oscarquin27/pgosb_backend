package models

import "database/sql"

type MissionServiceSummary struct {
	Id                int64          `db:"id"`
	MissionId         int64          `db:"mission_id"`
	AntaresId         sql.NullInt64  `db:"antares_id"`
	RegionName        sql.NullString `db:"region_name"`
	Description       sql.NullString `db:"description"`
	ServiceDate       sql.NullTime   `db:"service_date"`
	ManualServiceDate sql.NullTime   `db:"manual_service_date"`
	NumFirefighters   sql.NullInt64  `db:"num_firefighters"`
	NumVehicles       sql.NullInt64  `db:"num_vehicles"`
	NumUnits          sql.NullInt64  `db:"num_units"`
	StationName       sql.NullString `db:"station_name"`
	Unharmed          sql.NullInt64  `db:"unharmed"`
	Injured           sql.NullInt64  `db:"injured"`
	Transported       sql.NullInt64  `db:"transported"`
	Deceased          sql.NullInt64  `db:"deceased"`
	IsImportant       bool           `db:"is_important"`
	OperativesAreas   []string       `db:"operative_areas"`
	PendingForData    sql.NullBool   `db:"pending_for_data"`
	PeaceQuadrant     sql.NullString `db:"peace_quadrant"`
	Level             sql.NullString `db:"level"`
	State             sql.NullString `db:"state"`
	Municipality      sql.NullString `db:"municipality"`
	Parish            sql.NullString `db:"parish"`
	Commander         sql.NullString `db:"commander"`
	Alias             sql.NullString `db:"alias"`
	UnitsPlates       []string       `db:"units_plates"`
}
