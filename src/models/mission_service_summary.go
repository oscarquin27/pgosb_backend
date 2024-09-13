package models

import "database/sql"

type MissionServiceSummary struct {
	Id                int64          `db:"id"`
	MissionId         int64          `db:"mission_id"`
	AntaresId         sql.NullInt64  `db:"antares_id"`
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
	//Code              sql.NullString `db:"code"`

	Alias sql.NullString `db:"alias"`
}
