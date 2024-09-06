package models

import "database/sql"

type MissionSummary struct {
	Id              int64          `db:"id"`
	Alias           sql.NullString `db:"alias"`
	CreatedAt       sql.NullTime   `db:"created_at"`
	NumServices     sql.NullInt64  `db:"num_services"`
	NumFireFighters sql.NullInt64  `db:"num_firefighters"`
	NumVehicles     sql.NullInt64  `db:"num_vehicles"`
}
