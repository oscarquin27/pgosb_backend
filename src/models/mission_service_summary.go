package models

import "database/sql"

type MissionServiceSummary struct {
	Id                int64          `db:"id"`
	Alias             sql.NullString `db:"alias"`
	CreatedAt         sql.NullTime   `db:"created_at"`
	ServiceId         sql.NullInt64  `db:"service_id"`
	AntaresId         sql.NullInt64  `db:"antares_id"`
	Description       sql.NullString `db:"description"`
	ServiceDate       sql.NullTime   `db:"service_date"`
	ManualServiceDate sql.NullTime   `db:"manual_service_date"`
	NumFirefighters   sql.NullInt64  `db:"num_firefighters"`
	NumVehicles       sql.NullInt64  `db:"num_vehicles"`
	StationName       sql.NullString `db:"station_name"`
}
